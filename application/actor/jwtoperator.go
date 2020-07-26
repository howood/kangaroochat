package actor

import (
	"context"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/howood/kangaroochat/domain/entity"
	"github.com/howood/kangaroochat/domain/repository"
	log "github.com/howood/kangaroochat/infrastructure/logger"
)

// TokenExpired is token's expired
var TokenExpired = os.Getenv("TOKEN_EXPIED")

// TokenSecret define token secrets
var TokenSecret = os.Getenv("TOKEN_SECRET")

const JWTContextKey = "kangaroouser"

// JwtOperator struct
type JwtOperator struct {
	jwtClaims *entity.JwtClaims
	ctx       context.Context
}

// NewJwtOperator creates a new JwtClaimsRepository
func NewJwtOperator(ctx context.Context, username string, admin bool, identifier string) repository.JwtClaimsRepository {
	expired, _ := time.ParseDuration(TokenExpired)
	return &JwtOperator{
		jwtClaims: &entity.JwtClaims{
			Name:       username,
			Admin:      admin,
			Identifier: identifier,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * expired).Unix(),
			},
		},
		ctx: ctx,
	}
}

// CreateToken creates a new token
func (jc *JwtOperator) CreateToken(secret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jc.jwtClaims)
	tokenstring, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Error(jc.ctx, err.Error())
	}
	return tokenstring
}
