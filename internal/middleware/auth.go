package middleware

import (
	"context"
	"errors"
	"github.com/rs/zerolog"
	"github.com/yalagtyarzh/online-chat/internal/entity"
	"github.com/yalagtyarzh/online-chat/pkg/token"
	"github.com/yalagtyarzh/online-chat/pkg/utils"
	"net/http"
	"strings"
)

var (
	MissingAuthHeader = errors.New("missing authorization header")
	InvalidBearerType = errors.New("invalid bearer type")
)

func Auth(signer token.Signer, log *zerolog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if values := strings.Split(auth, " "); len(values) != 2 {
			utils.WriteError(w, "Missing authorization header", MissingAuthHeader, http.StatusUnauthorized, log)
			return
		} else {
			bearer, token := values[0], values[1]
			if bearer != "Bearer" {
				utils.WriteError(w, "Invalid bearer type", InvalidBearerType, http.StatusUnauthorized, log)
				return
			}
			userID, err := signer.Verify(token)
			if err != nil {
				utils.WriteError(w, "Verify user error", err, http.StatusUnauthorized, log)
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, entity.ContextKeyUserID, userID)
			r = r.WithContext(ctx)
		}
		next(w, r)
	})
}
