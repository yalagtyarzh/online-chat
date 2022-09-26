package main

import (
	"github.com/gorilla/mux"
	"github.com/yalagtyarzh/online-chat/internal/config"
	"github.com/yalagtyarzh/online-chat/pkg/logger"
	"net/http"
)

const serviceName = "online-chat"

func main() {
	cfg := config.InitConfig()

	log := logger.InitLogger(serviceName, cfg.PrettyLog, cfg.LogLevel)

	probe := mux.NewRouter()
	go func() {
		probe.HandleFunc("/probe", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		}).Methods(http.MethodGet)

		log.Info().Msgf("Probe started on %s", cfg.ProbeBind)
		if err := http.ListenAndServe(cfg.ProbeBind, probe); err != nil {
			log.Warn().Err(err).Msg("listen probe finished")
		}
	}()

}
