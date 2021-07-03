package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Health struct {
	l *log.Logger
}

func NewHealth(l *log.Logger) *Health {
	return &Health{l}
}

func (h *Health) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Alive")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Error!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "I'm alive %s", d)
}
