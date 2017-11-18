package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	donatugee *Donatugee
}

func NewServer(donatugee *Donatugee) *Server {
	s := &Server{
		donatugee: donatugee,
	}

	return s
}

func (s *Server) start() error {
	addr := "8081"
	if os.Getenv("ENV") == "production" {
		addr = os.Getenv("PORT")
	}

	r := mux.NewRouter()

	// Note: In a larger application, we'd likely extract our route-building logic into our handlers
	// package, given the coupling between them.

	// It's important that this is before your catch-all route ("/")
	api := r.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/challenges", GetChallenges).Methods("GET")
	// Optional: Use a custom 404 handler for our API paths.
	// api.NotFoundHandler = JSONNotFound

	// Serve static assets directly.
	r.PathPrefix("/dist/build.js").HandlerFunc(IndexHandler("./frontend/dist/build.js"))
	r.PathPrefix("/public").Handler(http.FileServer(http.Dir("./frontend/public/")))
	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(IndexHandler("./frontend/index.html"))

	srv := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, r),
		Addr:    "0.0.0.0:" + addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}


	return srv.ListenAndServe()
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

func GetChallenges(resp http.ResponseWriter, r *http.Request) {
	application := Application{
		ApplicationID: 1,
		Created:       time.Now(),
		Modified:      time.Now(),
	}

	// techfugee := Techfugee{
	// 	TechfugeeID:  1,
	// 	Applications: []Application{application},
	// 	Name:         "Michael Foo",
	// 	Email:        "michaelfoo@gmail.com",
	// 	Created:      time.Now(),
	// 	Modified:     time.Now(),
	// }

	challenges := []Challenge{
		Challenge{
			ChallengeID:  1,
			Applications: []Application{},
			Name:         "Learn PHP in 3 month",
			Image:        "",
			Description:  "go to laracast, learn PHP and pitch us what you learned",
			Created:      time.Now(),
			Modified:     time.Now(),
		},
		Challenge{
			ChallengeID:  2,
			Applications: []Application{application},
			Name:         "Learn Go in 3 month",
			Image:        "",
			Description:  "go to the Go tour, learn Go and create a small Go app",
			Created:      time.Now(),
			Modified:     time.Now(),
		},
	}

	js, err := json.Marshal(challenges)
	if err != nil {
		http.Error(resp, fmt.Sprintf("json: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}
