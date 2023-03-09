package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/golang/protobuf/ptypes/empty"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"

	pb "github.com/shin5ok/sample-grpc-app-with-spanner/pb"

	"github.com/google/uuid"
	"github.com/pereslava/grpc_zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var port string = os.Getenv("PORT")
var appPort = "8080"
var promPort = "18080"

var dbString = os.Getenv("SPANNER_STRING")
var domain = os.Getenv("DOMAIN")

var tracer trace.Tracer

type healthCheck struct{}

func init() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano

}

type newServerImplement struct {
	Client GameUserOperation
}

func (s *newServerImplement) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	log.
		Info().
		Str("method", "CreateUser").
		Str("arg", fmt.Sprintf("%+v", fmt.Sprintf("%+v", user))).
		Send()

	userName := user.GetName()

	if userName == "" {
		return nil, status.Error(codes.InvalidArgument, "user name is empty")
	}

	userId, _ := uuid.NewRandom()
	w := io.Discard

	ctx, span := tracer.Start(ctx, "create user into Cloud Spanner")
	defer span.End()

	err := s.Client.createUser(ctx, w, userParams{userID: userId.String(), userName: userName})
	span.End()

	return &pb.User{Name: userName, Id: userId.String()}, err
}

func (s *newServerImplement) AddItemUser(ctx context.Context, userItem *pb.UserItem) (*empty.Empty, error) {
	log.
		Info().
		Str("method", "AddItemUser").
		Str("arg", fmt.Sprintf("%+v", fmt.Sprintf("%+v", userItem))).
		Send()

	if userItem.User.Id == "" || userItem.Item.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "user id or/and itemid is/are empty")
	}

	w := io.Discard
	err := s.Client.addItemToUser(ctx, w, userParams{userID: userItem.User.Id}, itemParams{itemID: userItem.Item.Id})
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *newServerImplement) GetUserItems(user *pb.User, stream pb.Game_GetUserItemsServer) error {
	w := io.Discard
	ctx := context.Background()
	txn, iter, err := s.Client.userItems(ctx, w, user.GetId())
	if err != nil {
		log.Err(err).Send()
		return err
	}
	defer iter.Stop()
	defer txn.Close()

	var records_number int
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Err(err).Send()
			return err
		}
		var userName string
		var itemNames string
		var itemIds string
		if err := row.Columns(&userName, &itemNames, &itemIds); err != nil {
			log.Err(err).Send()
			return status.Error(codes.Unavailable, "columns bind error")
		}

		data := &pb.Item{Id: itemIds, Name: itemNames}

		if err := stream.Send(data); err != nil {
			return status.Error(codes.Unavailable, err.Error())
		}
		records_number++
	}

	if records_number == 0 {
		message := fmt.Sprintf("row count %d: not found", iter.RowCount)
		return status.Error(codes.NotFound, message)
	}

	return nil

}

func (n *newServerImplement) PingPong(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}

func main() {
	serverLogger := log.Level(zerolog.TraceLevel)
	grpc_zerolog.ReplaceGrpcLogger(zerolog.New(os.Stderr).Level(zerolog.ErrorLevel))

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	tp := tpExporter(projectID, "sample")
	ctx := context.Background()
	defer tp.ForceFlush(ctx)
	otel.SetTracerProvider(tp)

	tracer = otel.GetTracerProvider().Tracer(domain)

	interceptorOpt := otelgrpc.WithTracerProvider(otel.GetTracerProvider())

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_zerolog.NewPayloadUnaryServerInterceptor(serverLogger),
			grpc_prometheus.UnaryServerInterceptor,
			otelgrpc.UnaryServerInterceptor(interceptorOpt),
		),
		grpc.ChainStreamInterceptor(
			grpc_zerolog.NewStreamServerInterceptor(serverLogger),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zerolog.NewPayloadStreamServerInterceptor(serverLogger),
			otelgrpc.StreamServerInterceptor(interceptorOpt),
		),
	)

	if port == "" {
		port = appPort
	}

	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		serverLogger.Fatal().Msg(err.Error())
	}

	newServer := newServerImplement{}
	spannerClient, err := newClient(ctx, dbString)
	if err != nil {
		panic(err)
	}
	newServer.Client = spannerClient

	pb.RegisterGameServer(server, &newServer)

	var h = &healthCheck{}
	health.RegisterHealthServer(server, h)

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(server)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(":"+promPort, nil); err != nil {
			panic(err)
		}
		serverLogger.Info().Msgf("prometheus listening on :%s\n", promPort)
	}()

	reflection.Register(server)
	serverLogger.Info().Msgf("Listening on %s\n", port)
	server.Serve(listenPort)

}

func (h *healthCheck) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (h *healthCheck) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "No implementation for Watch")
}

func tpExporter(projectID, serviceName string) *sdktrace.TracerProvider {
	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		log.Err(err).Send()
		return nil
	}

	ctx := context.Background()
	res, err := resource.New(ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	if err != nil {
		log.Err(err).Send()
		return nil
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	return tp
}

func (s *newServerImplement) ListItems(ctx context.Context, _ *empty.Empty) (*pb.Items, error) {
	data, err := s.Client.listItems(ctx)
	if err != nil {
		log.Err(err).Send()
		return &pb.Items{}, err
	}
	// Just for temporary
	// TODO: use def of protobuf directly
	items := []*pb.Item{}
	for _, v := range data {
		items = append(items, &pb.Item{Id: v.itemID})
	}
	return &pb.Items{Items: items}, err
}
