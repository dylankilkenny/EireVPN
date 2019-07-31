package test

import (
	"eirevpn/api/router"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	InitDB()
	r = router.SetupRouter(false)
	code := m.Run()
	os.Exit(code)
}

func TestAuthTokens(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/private/plans", nil)
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Authentification", func(t *testing.T) {
		user := CreateUser()
		_ = CreatePlan()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("No Auth Token", func(t *testing.T) {
		user := CreateUser()
		_ = CreatePlan()
		_, refreshToken, csrfToken := GetToken(user)
		want := 401
		got := makeRequest(t, "", refreshToken, csrfToken)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}
