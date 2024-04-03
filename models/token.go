package models

import (
	"encoding/base64"
	"jwt-project/consts"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	AccessToken      string // signed access token
	RefreshToken     string // refresh token
	RefreshTokenHash []byte // hashed refresh token for db
	CreatedAt        time.Time
}

func (token *Token) CreateTokenPair(c *gin.Context, userId string) error {
	// creating access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(consts.AccessTokenTTL).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return err
	}

	// creating refresh token from time.Now()
	createdAt := time.Now()

	refreshToken := base64.StdEncoding.EncodeToString([]byte(createdAt.String()))
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	token.AccessToken = accessTokenString
	token.RefreshToken = refreshToken
	token.RefreshTokenHash = hashedToken
	token.CreatedAt = createdAt

	return nil
}
