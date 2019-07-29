package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginRoute(t *testing.T) {

	credentials := map[string]string{"email": "email@email.com", "password": "password"}

	makeRequest := func(t *testing.T, credentials map[string]string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful login", func(t *testing.T) {
		_ = CreateUser()
		want := 200
		got := makeRequest(t, credentials)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Email or Password missing", func(t *testing.T) {
		emptyCredentials := map[string]string{"email": "", "password": ""}
		want := 400
		got := makeRequest(t, emptyCredentials)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Email Not Found", func(t *testing.T) {
		want := 404
		got := makeRequest(t, credentials)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Wrong password", func(t *testing.T) {
		credentials["password"] = "pass"
		_ = CreateUser()
		want := 401
		got := makeRequest(t, credentials)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestSignUpRoute(t *testing.T) {

	signup := map[string]string{"firstname": "dylan", "lastname": "kilkenny", "email": "email@email.com", "password": "password"}

	makeRequest := func(t *testing.T, credentials map[string]string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Sign Up", func(t *testing.T) {
		want := 200
		got := makeRequest(t, signup)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Invalid form", func(t *testing.T) {
		signup = map[string]string{"email": "", "password": ""}
		want := 400
		got := makeRequest(t, signup)
		assertCorrectStatusCode(t, want, got)
	})

	t.Run("Email already exists", func(t *testing.T) {
		_ = CreateUser()
		want := 400
		got := makeRequest(t, signup)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}
