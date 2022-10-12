package main

import (
	"context"
	"encoding/json"
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

type Serving struct {
	Client GameUserOperation
}

func main() {

	s := Serving{
		Client: dbClient{},
	}

	spannerClient, _ := spannerNewClient(spannerString)
	defer spannerClient.Close()

	oplog := httplog.LogEntry(context.Background())
	/* jsonify logging */
	httpLogger := httplog.NewLogger(appName, httplog.Options{JSON: true, LevelFieldName: "severity", Concise: true})

	/* exporter for prometheus */
	m := chiprometheus.NewMiddleware(appName)

	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httplog.RequestLogger(httpLogger))
	r.Use(m)

	r.Handle("/metrics", promhttp.Handler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]string{"Ping": "Pong"})
	})

	r.Post("/api/users/{user:[a-z0-9-.]+}", func(w http.ResponseWriter, r *http.Request) {
		userId, _ := uuid.NewUUID()
		userName := chi.URLParam(r, "user")
		err := s.Client.createUser(w, spannerClient, userParams{userID: userId.String(), userName: userName})
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"ERROR": err.Error()})
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
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"ERROR": err.Error()})
			return
		}

		userName := chi.URLParam(r, "user")
		newScore := params.Score
		err := s.Client.updateScore(w, spannerClient, userParams{userName: userName}, int64(newScore))
		if err != nil {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"ERROR": err.Error()})
			return
		}
		render.JSON(w, r, map[string]string{})
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		oplog.Err(err)
	}

}
