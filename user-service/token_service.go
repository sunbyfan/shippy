package main

import (
	pb "shippy/user-service/proto/user"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("sunfanfsasdweqeadad")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

// Decode a token string into a token object
func (src *TokenService) Decode(token string) (*CustomClaims, error) {
	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	//Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "go.micro.src.user",
		},
	}

	//Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Sign token and return
	return token.SignedString(key)
}
