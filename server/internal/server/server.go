package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/martishin/react-golang-goth-auth/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

func init() {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	callbackURL := os.Getenv("GOOGLE_CALLBACK_URL")

	if clientID == "" || clientSecret == "" || callbackURL == "" {
		panic("Google OAuth environment variables are not set in .env")
	}

	// Configure session store
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		sessionKey = "default-session-key" // Use a default key if not set
	}

	store := sessions.NewCookieStore([]byte(sessionKey))

	isProduction := os.Getenv("ENV") == "production"
	domain := os.Getenv("SESSION_COOKIE_DOMAIN")

	store.Options = &sessions.Options{
		HttpOnly: true,
		Secure:   isProduction, // Enable secure cookies in production
		Path:     "/",
		MaxAge:   86400 * 30, // 30 days
		Domain:   domain,
	}

	gothic.Store = store

	// Configure Google provider
	goth.UseProviders(
		google.New(clientID, clientSecret, callbackURL, "email", "profile"),
	)
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
