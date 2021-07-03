package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tommynueve/family-feud/data"
)

type Questions struct {
	l *log.Logger
}

func NewQuestions(l *log.Logger) *Questions {
	return &Questions{l}
}

func (q *Questions) getQuestions(rw http.ResponseWriter, r *http.Request) {
	q.l.Println("GET: Questions")

	//getting questions from data
	qs := data.GetQuestions()

	d, err := json.Marshal(qs)
	if err != nil {
		http.Error(rw, "Error getting random question", http.StatusBadRequest)
		return
	}
	rw.Write(d)
}

func (q *Questions) addQuestion(rw http.ResponseWriter, r *http.Request) {
	q.l.Println("POST: Questions")

	newQ := &data.Question{}
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(newQ)

	if err != nil {
		http.Error(rw, "Error decoding json body", http.StatusBadRequest)
		return
	}

	data.AddQuestion(newQ)
}

func (q *Questions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		q.getQuestions(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		q.addQuestion(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}
