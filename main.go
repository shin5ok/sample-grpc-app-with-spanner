package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var appName = "myapp"

var spannerString = os.Getenv("SPANNER_STRING")
var servicePort = os.Getenv("PORT")

type Serving struct {
	Client GameUserOperation
}

func main() {

	ctx := context.Background()
	db, err := newClient(ctx, spannerString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.sc.Close()

	s := Serving{
		Client: db,
	}

	run(s)

}

func run(s Serving) {

	oplog := httplog.LogEntry(context.Background())
	/* jsonify logging */
	httpLogger := httplog.NewLogger(appName, httplog.Options{JSON: true, LevelFieldName: "severity", Concise: true})

	/* exporter for prometheus */
	m := chiprometheus.NewMiddleware(appName)

	r := chi.NewRouter()
	// r.Use(middleware.Throttle(8))
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httplog.RequestLogger(httpLogger))
	r.Use(m)

	r.Handle("/metrics", promhttp.Handler())

	errorRender := func(w http.ResponseWriter, r *http.Request, httpCode int, err error) {
		render.Status(r, httpCode)
		render.JSON(w, r, map[string]interface{}{"ERROR": err.Error()})
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"Ping": "Pong"})
	})

	r.Get("/api/users/{user:[a-z0-9-.]+}", func(w http.ResponseWriter, r *http.Request) {
		userName := chi.URLParam(r, "user")
		ctx := r.Context()
		results, err := s.Client.ListUsers(ctx, w, userName)
		if err != nil {
			errorRender(w, r, http.StatusInternalServerError, err)
			return
		}
		render.JSON(w, r, results)
	})

	r.Post("/api/users/{user:[a-z0-9-.]+}", func(w http.ResponseWriter, r *http.Request) {
		userId, _ := uuid.NewUUID()
		userName := chi.URLParam(r, "user")
		ctx := r.Context()
		err := s.Client.createUser(ctx, w, userParams{userID: userId.String(), userName: userName})
		if err != nil {
			errorRender(w, r, http.StatusInternalServerError, err)
			return
		}
		render.JSON(w, r, map[string]string{})
	})

	r.Put("/api/users/{user:[a-z0-9-.]+}", func(w http.ResponseWriter, r *http.Request) {
		type bodyParams struct {
			Score int `json:"score"`
		}
		params := bodyParams{}
		jsonDecorder := json.NewDecoder(r.Body)
		if err := jsonDecorder.Decode(&params); err != nil {
			errorRender(w, r, http.StatusInternalServerError, err)
			return
		}

		userName := chi.URLParam(r, "user")
		newScore := params.Score
		ctx := r.Context()
		err := s.Client.updateScore(ctx, w, userParams{userName: userName}, int64(newScore))
		if err != nil {
			errorRender(w, r, http.StatusInternalServerError, err)
			return
		}
		render.JSON(w, r, map[string]string{})
	})

	if err := http.ListenAndServe(":"+servicePort, r); err != nil {
		oplog.Err(err)
	}

}
