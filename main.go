package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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

	userId, _ := uuid.NewRandom()
	userName := user.GetName()

	w := os.Stdout
	s.Client.createUser(ctx, w, userParams{userID: userId.String(), userName: userName})

	return &pb.User{Name: userName, Id: userId.String()}, nil
}

func (s *newServerImplement) AddItemUser(ctx context.Context, userItem *pb.UserItem) (*empty.Empty, error) {
	log.
		Info().
		Str("method", "AddItemUser").
		Str("arg", fmt.Sprintf("%+v", fmt.Sprintf("%+v", userItem))).
		Send()

	w := os.Stdout
	s.Client.addItemToUser(ctx, w, userParams{userID: userItem.User.Id}, itemParams{itemID: userItem.Item.Id})

	return &emptypb.Empty{}, nil
}

func (s *newServerImplement) GetUserItems(user *pb.User, stream pb.Game_GetUserItemsServer) error {
	w := os.Stdout
	ctx := context.Background()
	txn, iter, err := s.Client.userItems(ctx, w, user.GetId())
	if err != nil {
		log.Err(err).Send()
		return err
	}
	defer iter.Stop()
	defer txn.Close()

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
			return err
		}
		log.Info().Str("itemIds", itemIds).Send()

		data := &pb.Item{Id: itemIds}
		stream.Send(data)
	}
	return nil

}

func (n *newServerImplement) PingPong(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &emptypb.Empty{}, nil
}

func main() {
	serverLogger := log.Level(zerolog.TraceLevel)
	grpc_zerolog.ReplaceGrpcLogger(zerolog.New(os.Stderr).Level(zerolog.ErrorLevel))

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_zerolog.NewPayloadUnaryServerInterceptor(serverLogger),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc.ChainStreamInterceptor(
			grpc_zerolog.NewStreamServerInterceptor(serverLogger),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zerolog.NewPayloadStreamServerInterceptor(serverLogger),
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
	ctx := context.Background()
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
