package main

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}

	return w.gw.Write(b)
}

type GZipHandler struct {
	next http.Handler
}

func (h *GZipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")

	if strings.Contains(encodings, "gzip") {
		h.serveGzipped(w, r)
	} else if strings.Contains(encodings, "deflate") {
		panic("Deflate not implemented")
	} else {
		h.ServeHTTP(w, r)
	}
}

func (h *GZipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	w.Header().Set("Content-Encoding", "gzip")
	h.next.ServeHTTP(GzipResponseWriter{
		gw:             gzw,
		ResponseWriter: w,
	}, r)
}

func (h *GZipHandler) servePlain(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r)
}
