package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

func LoginController(w http.ResponseWriter, r *http.Request) {
	// Check validation
	if r.Method != http.MethodPost {
		NotFound(w)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		InternalServerError(w)
		return
	}

	// get body of request
	var request LoginRequest
	err = json.Unmarshal(bodyBytes, &request)
	if err != nil {
		BadRequest(w)
		return
	}

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
		InternalServerError(w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	response := LoginResponse{
		Status:  "success",
		Message: "Login success",
		Token:   tokenString,
		Expire:  expirationTime,
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		InternalServerError(w)
	}
	// Response token
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(responseBytes))
}
