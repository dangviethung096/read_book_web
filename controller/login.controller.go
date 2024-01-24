package controller

import (
	"net/http"
	"time"

	"github.com/dangviethung096/core"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"min=8,max=30"`
}

type LoginResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Token   string    `json:"token"`
	Expire  time.Time `json:"expire"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func LoginController(ctx *core.HttpContext, request LoginRequest) (core.HttpResponse, core.HttpError) {
	// Check username and password
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: request.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, core.NewHttpError(http.StatusInternalServerError, http.StatusInternalServerError, "Internal server error", nil)
	}

	response := LoginResponse{
		Status:  "success",
		Message: "Login success",
		Token:   tokenString,
		Expire:  expirationTime,
	}

	// Response token
	return core.NewHttpResponse(http.StatusOK, response), nil

}
