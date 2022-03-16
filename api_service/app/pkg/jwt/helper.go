package jwt

import (
	"github.com/cristalhq/jwt"
	"note_app/api_service/app/pkg/cache"
	"note_app/api_service/app/pkg/logging"
)

var _, Helper = &helper{}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

type RT struct {
	RefreshToken string `json:"refresh_token"`
}

type helper struct {
	Logger  logging.Logger
	RTCache cache.Repository
}
