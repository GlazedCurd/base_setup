package main

import (
	"encoding/json"
	"fmt"
	logger "log"
	"net/http"
	"strconv"
)

type LogRecord struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Level   string `json:"level"`
	Handler string `json:"handler"`
}

func log(rec *LogRecord) {
	msg, err := json.Marshal(rec)
	if err != nil {
		logger.Fatalf("log formatting: %s", err)
	}
	logger.Print(string(msg))
}

func main() {
	mux := http.NewServeMux()
	logger.SetFlags(0)
	mux.HandleFunc("/{id}", func(w http.ResponseWriter, req *http.Request) {
		idstr := req.PathValue("id")
		code, err := strconv.Atoi(idstr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid id %s", idstr)
			log(
				&LogRecord{
					Code:    code,
					Message: fmt.Sprintf("Failed to parse %s", err),
					Level:   "error",
					Handler: "default",
				},
			)
			return
		}
		msg := req.URL.Query().Get("msg")
		if msg == "" {
			msg = "default message text"
		}
		log(&LogRecord{
			Code:    code,
			Message: fmt.Sprintf("Msg got \"%s\"", msg),
			Level:   "info",
			Handler: "default",
		})
		fmt.Fprintf(w, "id %s", idstr)
	})
	http.ListenAndServe(":8080", mux)
}
