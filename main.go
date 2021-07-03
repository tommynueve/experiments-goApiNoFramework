package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tommynueve/family-feud/handlers"
)

func main() {

	l := log.New(os.Stdout, "quiz-api", log.LstdFlags)
	hh := handlers.NewHealth(l)
	q := handlers.NewQuestions(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/questions", q)

	s := &http.Server{
		Addr:         "localhost:9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt)

	sig := <-sigChan
	l.Println("Termination, shutting down", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeoutContext)
}
