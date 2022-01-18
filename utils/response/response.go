package response

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"error"`
}

type BasicData struct {
	Data string `json:"data"`
}

func JSON(w http.ResponseWriter, status int, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if _, err := w.Write(buf.Bytes()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func String(w http.ResponseWriter, status int, msg string) {
	JSON(w, status, &BasicData{Data: msg})
}

func Error(w http.ResponseWriter, status int, msg string) {
	JSON(w, status, ErrorMessage{
		Message: msg,
	})
}
