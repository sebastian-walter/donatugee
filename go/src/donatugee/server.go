package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"
)

type Server struct {
	donatugee *Donatugee
}

func NewServer(donatugee *Donatugee) *Server {
	s := &Server{
		donatugee: donatugee,
	}

	http.HandleFunc("/challenges", s.challenges)
	return s
}

func (s *Server) start() error {
	addr := ":8081"
	if os.Getenv("ENV") == "production" {
		addr := ":80"
	}
	return http.ListenAndServe(addr, nil)
}

func (s *Server) challenges(resp http.ResponseWriter, r *http.Request) {
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
