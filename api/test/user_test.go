package test

import (
	"bytes"
	"eirevpn/api/config"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoginRoute(t *testing.T) {

	credentials := map[string]string{"email": "email@email.com", "password": "password"}

	makeRequest := func(t *testing.T, credentials map[string]string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/user/login", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful login", func(t *testing.T) {
		_ = CreateUser()
		want := 200
		got := makeRequest(t, credentials)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Email or Password missing", func(t *testing.T) {
		emptyCredentials := map[string]string{"email": "", "password": ""}
		want := 400
		got := makeRequest(t, emptyCredentials)
		assertCorrectStatus(t, want, got)
	})

	t.Run("Email Not Found", func(t *testing.T) {
		want := 400
		got := makeRequest(t, credentials)
		assertCorrectStatus(t, want, got)
	})

	t.Run("Wrong password", func(t *testing.T) {
		credentials["password"] = "pass"
		_ = CreateUser()
		want := 401
		got := makeRequest(t, credentials)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestSignUpRoute(t *testing.T) {

	signup := map[string]string{"firstname": "dylan", "lastname": "kilkenny", "email": "email@email.com", "password": "password"}

	makeRequest := func(t *testing.T, credentials map[string]string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(credentials)
		req, _ := http.NewRequest("POST", "/api/user/signup", bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Sign Up", func(t *testing.T) {
		want := 200
		got := makeRequest(t, signup)
		assertCorrectStatus(t, want, got)
	})

	t.Run("Invalid form", func(t *testing.T) {
		signup = map[string]string{"email": "", "password": ""}
		want := 400
		got := makeRequest(t, signup)
		assertCorrectStatus(t, want, got)
	})

	t.Run("Email already exists", func(t *testing.T) {
		_ = CreateUser()
		want := 400
		got := makeRequest(t, signup)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestAllUsersRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/protected/users", nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all users", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestGetUserRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, userId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/user/%d", userId)
		req, _ := http.NewRequest("GET", url, nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve user by ID", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("User not found", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 400
		planID := uint(999)
		got := makeRequest(t, authToken, csrfToken, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdateUserRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, user map[string]interface{}, userId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(user)
		url := fmt.Sprintf("/api/protected/user/update/%d", userId)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Update User", func(t *testing.T) {
		plan := map[string]interface{}{
			"firstname": "Simon",
			"lastname":  "Wilson",
			"email":     "sw@email.com",
		}
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, plan, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledUser := map[string]interface{}{
			"firstname": "",
		}
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, csrfToken, halfFilledUser, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}
