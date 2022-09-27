package utils

import (
	"github.com/rs/zerolog"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, body []byte, log *zerolog.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err := w.Write(body)
	if err != nil {
		log.Error().Err(err).Msg("\"WriteResponse\" function error")
	}
}

func WriteError(w http.ResponseWriter, msg string, err error, status int, log *zerolog.Logger) {
	log.Error().Err(err).Msg(msg)
	w.WriteHeader(status)
	_, err = w.Write([]byte(msg))
	if err != nil {
		log.Error().Err(err).Msg("\"WriteError\" function error")
	}
}
