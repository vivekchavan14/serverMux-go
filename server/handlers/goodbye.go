package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (b *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Goodbye")
}
