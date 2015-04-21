package main

import (
	"fmt"
	"net/http"
)

type Middleware []http.Handler

func (m *Middleware) Add(handler http.Handler) {
	*m = append(*m, handler)
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mw := NewMiddlewareResponseWriter(w)

	for _, handler := range m {
		fmt.Println("Trying handler", handler)
		handler.ServeHTTP(mw, r)
		fmt.Println("Written?", mw.written)
		if mw.written {
			return
		}
	}
	Handle404(mw, r)
}

type MiddlewareResponseWriter struct {
	http.ResponseWriter
	written bool
}

func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWriter {
	return &MiddlewareResponseWriter{
		ResponseWriter: w,
	}
}

func (w *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(bytes)
}

func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}
