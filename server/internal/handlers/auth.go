package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/markbates/goth/gothic"
	"github.com/martishin/react-golang-goth-auth/internal/database"
)

func GoogleCallbackHandler(db database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Finalize the authentication process
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			log.Println(err)
			return
		}

		// Save user to the database
		ctx := r.Context()
		userID, err := db.FindOrCreateUser(ctx, &database.User{
			Name:    user.Name,
			Email:   user.Email,
			Picture: user.AvatarURL,
		})
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// Save user ID in the session
		err = gothic.StoreInSession("user_id", userID, r, w)
		if err != nil {
			http.Error(w, "Failed to save session", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// Redirect to the secure area
		redirectSecure := os.Getenv("REDIRECT_SECURE")
		if redirectSecure == "" {
			redirectSecure = "http://localhost:5173/secure"
		}

		http.Redirect(w, r, redirectSecure, http.StatusFound)
	}
}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear session
	err := gothic.Logout(w, r)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	// Redirect to login page
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
