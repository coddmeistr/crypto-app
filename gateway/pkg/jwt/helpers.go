package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/maxim12233/crypto-app-server/gateway/internal/config"
	"go.uber.org/zap"
)

type helper struct {
	logger *zap.Logger
}

type IHelper interface {
	GenerateJWT(ID uint, roles []uint) (string, error)
}

func NewHelper(logger *zap.Logger) IHelper {
	return &helper{
		logger: logger,
	}
}

func (h *helper) GenerateJWT(ID uint, roles []uint) (string, error) {
	c := config.GetConfig()

	// Generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   ID,
		"roles": roles,
		"exp":   time.Now().Add(time.Hour * time.Duration(c.GetInt("auth.jwtexphours"))).Unix(),
	})

	h.logger.Info("signing jwt token")
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(c.GetString("server.secret_key")))

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}
