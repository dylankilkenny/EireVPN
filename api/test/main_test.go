package test

import (
	"eirevpn/api/config"
	"eirevpn/api/logger"
	"eirevpn/api/router"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var logging bool

func TestMain(m *testing.M) {

	flag.BoolVar(&logging, "logging", false, "enable logging")
	flag.Parse()

	config.Init("../config.yaml")
	conf := config.GetConfig()
	config.UseStripeIntegration(false)

	InitDB()
	r = router.Init(conf, logging)
	logger.Init(logging)
	code := m.Run()

	os.Exit(code)
}

func TestAuthTokens(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string) *httptest.ResponseRecorder {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/private/servers", nil)
		if authToken != "" {
			req.AddCookie(&http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)})
		}
		if refreshToken != "" {
			req.AddCookie(&http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)})
		}
		req.Header.Set("X-CSRF-Token", csrfToken)
		r.ServeHTTP(w, req)
		return w
	}

	t.Run("Successful Authentification", func(t *testing.T) {
		user := CreateUser()
		_ = CreateServer()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken)
		assertCorrectStatus(t, want, got.Code)
		CreateCleanDB()
	})

	t.Run("No Auth cookie", func(t *testing.T) {
		user := CreateUser()
		_, refreshToken, csrfToken := GetToken(user)
		wantStatus := 401
		wantCode := "AUTHCOOKMISS"
		resp := makeRequest(t, "", refreshToken, csrfToken)
		apiErr := bindError(resp)
		assertCorrectStatus(t, wantStatus, apiErr.Status)
		assertCorrectCode(t, wantCode, apiErr.Code)
		CreateCleanDB()
	})

	t.Run("No refresh cookie", func(t *testing.T) {
		user := CreateUser()
		authToken, _, csrfToken := GetToken(user)
		wantStatus := 401
		wantCode := "REFCOOKMISS"
		resp := makeRequest(t, authToken, "", csrfToken)
		apiErr := bindError(resp)
		assertCorrectStatus(t, wantStatus, apiErr.Status)
		assertCorrectCode(t, wantCode, apiErr.Code)
		CreateCleanDB()
	})

	t.Run("Token invalid", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		wantStatus := 401
		wantCode := "TOKENINVALID"
		resp := makeRequest(t, authToken+"p", refreshToken+"33333", csrfToken)
		apiErr := bindError(resp)
		assertCorrectStatus(t, wantStatus, apiErr.Status)
		assertCorrectCode(t, wantCode, apiErr.Code)
		CreateCleanDB()
	})

	t.Run("Invalid identifier", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		DeleteIdentifier(user)
		wantStatus := 401
		wantCode := "INVIDENTIFIER"
		resp := makeRequest(t, authToken+"p", refreshToken, csrfToken)
		apiErr := bindError(resp)
		assertCorrectStatus(t, wantStatus, apiErr.Status)
		assertCorrectCode(t, wantCode, apiErr.Code)
		CreateCleanDB()
	})

	t.Run("CSRF Invalid", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, _ := GetToken(user)
		wantStatus := 401
		wantCode := "CSRFTOKEN"
		resp := makeRequest(t, authToken, refreshToken, "")
		apiErr := bindError(resp)
		assertCorrectStatus(t, wantStatus, apiErr.Status)
		assertCorrectCode(t, wantCode, apiErr.Code)
		CreateCleanDB()
	})
}
