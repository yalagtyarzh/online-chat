package main

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/yalagtyarzh/online-chat/internal/config"
	"github.com/yalagtyarzh/online-chat/internal/database"
	"github.com/yalagtyarzh/online-chat/pkg/logger"
	"github.com/yalagtyarzh/online-chat/pkg/token"
	"net/http"
	"time"
)

const serviceName = "online-chat"

func main() {
	cfg := config.InitConfig()

	log := logger.InitLogger(serviceName, cfg.PrettyLog, cfg.LogLevel)

	db, err := database.New(cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Connecting to database error")
	}
	defer db.Close()

	signer := token.New(token.SignerOptions{
		Now: func() time.Time {
			return time.Now().UTC()
		},
		Issuer: cfg.JWT.Issuer,
		TTL:    1 * time.Hour,
		Secret: []byte(cfg.JWT.Secret),
	})

	c := chat.New(db, NewRedis(cfg.Redis.Addr, cfg.Redis.Password), log)

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

func NewRedis(addr, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return client
}
