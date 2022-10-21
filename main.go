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

	client, err := newClient(ctx, spannerString)
	if err != nil {
		log.Fatal(err)
	}
	defer client.sc.Close()

	s := Serving{
		Client: client,
	}

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

	r.Get("/", s.pingPong)

	r.Get("/api/users/{user:[a-z0-9-.]+}", s.getUsers)

	r.Post("/api/users/{user:[a-z0-9-.]+}", s.createUser)

	r.Put("/api/users/{user_id:[a-z0-9-.]+}", s.updateUser)

	if err := http.ListenAndServe(":"+servicePort, r); err != nil {
		oplog.Err(err)
	}

}

var errorRender = func(w http.ResponseWriter, r *http.Request, httpCode int, err error) {
	render.Status(r, httpCode)
	render.JSON(w, r, map[string]interface{}{"ERROR": err.Error()})
}

func (s Serving) getUsers(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user")
	ctx := r.Context()
	results, err := s.Client.listUsers(ctx, w, userName)
	if err != nil {
		errorRender(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, results)
}

func (s Serving) createUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := uuid.NewUUID()
	userName := chi.URLParam(r, "user")
	ctx := r.Context()
	err := s.Client.createUser(ctx, w, userParams{userID: userId.String(), userName: userName})
	if err != nil {
		errorRender(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, map[string]string{})
}

func (s Serving) updateUser(w http.ResponseWriter, r *http.Request) {
	type bodyParams struct {
		Score int `json:"score"`
	}
	params := bodyParams{}
	jsonDecorder := json.NewDecoder(r.Body)
	if err := jsonDecorder.Decode(&params); err != nil {
		errorRender(w, r, http.StatusInternalServerError, err)
		return
	}

	userID := chi.URLParam(r, "user_id")
	newScore := params.Score
	ctx := r.Context()
	err := s.Client.updateScore(ctx, w, userParams{userID: userID}, int64(newScore))
	if err != nil {
		errorRender(w, r, http.StatusInternalServerError, err)
		return
	}
	render.JSON(w, r, map[string]string{})
}

func (s Serving) pingPong(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"Ping": "Pong"})
}
