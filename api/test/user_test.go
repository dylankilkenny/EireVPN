package test

import (
	"bytes"
	"eirevpn/api/models"
	"encoding/json"
	"fmt"
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
	makeRequest := func(t *testing.T, user *models.User) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/protected/users", nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all users", func(t *testing.T) {
		user := CreateAdminUser()

		want := 200
		got := makeRequest(t, user)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestGetUserRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, userId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/private/user/get/%d", userId)
		req, _ := http.NewRequest("GET", url, nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve user by ID", func(t *testing.T) {
		user := CreateAdminUser()

		want := 200
		got := makeRequest(t, user, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("User not found", func(t *testing.T) {
		user := CreateAdminUser()

		want := 400
		planID := uint(999)
		got := makeRequest(t, user, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdateUserRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, userupdate map[string]interface{}, userId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(userupdate)
		url := fmt.Sprintf("/api/private/user/update/%d", userId)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		AddTokens(user, req)
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

		want := 200
		got := makeRequest(t, user, plan, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledUser := map[string]interface{}{
			"firstname": "",
		}
		user := CreateAdminUser()

		want := 400
		got := makeRequest(t, user, halfFilledUser, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestChangePasswordRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, passwordupdate map[string]string) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(passwordupdate)
		url := fmt.Sprintf("/api/private/user/changepassword")
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Change Password", func(t *testing.T) {
		passwordupdate := map[string]string{
			"current_password": "password",
			"new_password":     "passwordchanged123",
		}
		user := CreateAdminUser()
		want := 200
		got := makeRequest(t, user, passwordupdate)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		passwordupdate := map[string]string{
			"current_password": "password",
			"new_password":     "",
		}
		user := CreateAdminUser()
		want := 400
		got := makeRequest(t, user, passwordupdate)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}
