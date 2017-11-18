package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
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

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/challenges", s.challenges)
	mux.HandleFunc("/api/v1/insert-techfugee", s.insertTechfugee)
	mux.HandleFunc("/api/v1/techfugees", s.techfugees)
	mux.HandleFunc("/api/v1/update-auth", s.updateAuth)
	mux.HandleFunc("/api/v1/add-skills", s.addSkills)

	mux.Handle("/public", http.FileServer(http.Dir("./frontend/public")))
	mux.Handle("/dist", http.FileServer(http.Dir("./frontend/dist")))
	mux.Handle("/", http.FileServer(http.Dir("./frontend")))

	handler := cors.Default().Handler(mux)

	return http.ListenAndServe(":"+addr, handler)
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

func (s *Server) insertChallenge(resp http.ResponseWriter, r *http.Request) {
	idDonator := r.FormValue("id_donator")
	name := r.FormValue("name")
	description := r.FormValue("description")

	challenge, errs := s.donatugee.InsertChallenge(idDonator, name, description)
	if len(errs) > 0 {
		http.Error(resp, fmt.Sprintf("insert: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(challenge)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) updateAuth(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	passed := r.FormValue("passed")

	techfugee, errs := s.donatugee.UpdateAuth(id, passed)
	if len(errs) > 0 {
		http.Error(resp, fmt.Sprintf("update: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", errs), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) techfugees(resp http.ResponseWriter, r *http.Request) {
	techfugees, errs := s.donatugee.Techfugees()
	if errs != nil {
		http.Error(resp, fmt.Sprintf("query: %v", errs), http.StatusInternalServerError)
	}

	js, err := json.Marshal(techfugees)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) insertTechfugee(resp http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	skills := r.FormValue("skills")

	techfugee, errs := s.donatugee.InsertTechfugee(name, email, skills)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("insert: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) techfugee(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	techfugee, errs := s.donatugee.Techfugee(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("get: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) challenge(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	challenge, errs := s.donatugee.Challenge(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("get: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(challenge)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) addSkills(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	skills := r.FormValue("skills")

	techfugee, errs := s.donatugee.Techfugee(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("getfugee: %v", errs), http.StatusInternalServerError)
		return
	}

	techfugee, errs = s.donatugee.UpdateTechfugeeSkills(techfugee, skills)

	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("addskille: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) challenges(resp http.ResponseWriter, r *http.Request) {
	application := Application{
		ApplicationID: 1,
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
		},
		Challenge{
			ChallengeID:  2,
			Applications: []Application{application},
			Name:         "Learn Go in 3 month",
			Image:        "",
			Description:  "go to the Go tour, learn Go and create a small Go app",
		},
	}

	js, err := json.Marshal(challenges)
	if err != nil {
		http.Error(resp, fmt.Sprintf("json: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}
