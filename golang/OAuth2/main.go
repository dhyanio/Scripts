package main

import (
	"net/http"
	"golang.org/x/oauth2"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8080/callback",
		ClientID:		os.Getenv("GOOGLE_CLIENT_ID"),
		Client
	}
)