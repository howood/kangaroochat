package entity

import jwt "github.com/golang-jwt/jwt/v4"

// JwtClaims entity
type JwtClaims struct {
	Name       string `json:"name"`
	Admin      bool   `json:"admin"`
	Identifier string `json:"identifier"`
	jwt.StandardClaims
}
