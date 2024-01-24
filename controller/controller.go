package controller

import (
	"net/http"

	"github.com/dangviethung096/core"
)

func InitController() {
	core.RegisterAPI("/api/login", http.MethodPost, LoginController)
}
