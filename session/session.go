// session.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-server-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-cookie")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-cookie")
	if err != nil {
		// Handle error getting session (e.g., log the error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Authentication goes here
	// ...

	// Simulate authentication logic (replace with your actual logic)
	// In this example, we'll just assume successful login
	authenticated := true

	// Set user as authenticated
	session.Values["authenticated"] = authenticated
	session.Save(r, w)

	// Save the updated session data back to the user's browser
	err = session.Save(r, w)
	if err != nil {
		// Handle error saving session (e.g., log the error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Login successful, redirect or display a message
	fmt.Fprintln(w, "Login successful!")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-cookie")
	if err != nil {
		// Handle error getting session (e.g., log the error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)

	// Logout successful, redirect or display a message
	fmt.Fprintln(w, "Logged out!")
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
