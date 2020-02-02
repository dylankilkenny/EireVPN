package test

import (
	"bytes"
	"eirevpn/api/config"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetServerRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, serverId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/servers/%d", serverId)
		req, _ := http.NewRequest("GET", url, nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve server by ID", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		server := CreateServer()
		want := 200
		got := makeRequest(t, authToken, csrfToken, server.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Server not found", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 400
		serverID := uint(999)
		got := makeRequest(t, authToken, csrfToken, serverID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestCreateServerRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken, boundary string, server bytes.Buffer) int {
		t.Helper()
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/api/protected/servers/create", strings.NewReader(server.String()))
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Add("Content-Type", boundary)
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Server Creation", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)
		_ = w.WriteField("country", "United Kingdom")
		_ = w.WriteField("country_code", "UK")
		_ = w.WriteField("ip", "127.0.0.1")
		_ = w.WriteField("type", "Proxy")
		_ = w.WriteField("port", "8080")
		_ = w.WriteField("username", "admin")
		_ = w.WriteField("password", "admin")
		w.Close()
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, w.FormDataContentType(), *buf)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)
		_ = w.WriteField("country", "United Kingdom")
		_ = w.WriteField("country_code", "")
		_ = w.WriteField("ip", "127.0.0.1")
		_ = w.WriteField("type", "")
		_ = w.WriteField("port", "")
		_ = w.WriteField("username", "")
		_ = w.WriteField("password", "")
		w.Close()
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, csrfToken, w.FormDataContentType(), *buf)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Drop table - Internal Server Error", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)
		_ = w.WriteField("country", "United Kingdom")
		_ = w.WriteField("country_code", "UK")
		_ = w.WriteField("ip", "127.0.0.1")
		_ = w.WriteField("type", "Proxy")
		_ = w.WriteField("port", "8080")
		_ = w.WriteField("username", "admin")
		_ = w.WriteField("password", "admin")
		w.Close()
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 500
		DropServerTable()
		got := makeRequest(t, authToken, csrfToken, w.FormDataContentType(), *buf)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestAllServersRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/private/servers", nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all servers", func(t *testing.T) {
		_ = CreateServer()
		user := CreateUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("No Servers Found", func(t *testing.T) {
		user := CreateUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := CreateUser()
		authToken, csrfToken := GetToken(user)
		want := 500
		DropServerTable()
		got := makeRequest(t, authToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdateServerRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, server map[string]interface{}, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(server)
		url := fmt.Sprintf("/api/protected/servers/update/%d", id)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Update Server", func(t *testing.T) {
		server := map[string]interface{}{
			"ip":       "127.0.0.1",
			"port":     11211,
			"username": "user",
			"password": "1212ppp",
		}
		s := CreateServer()
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, server, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledServer := map[string]interface{}{
			"ip":       "",
			"port":     0,
			"username": "",
			"password": "",
		}
		s := CreateServer()
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, csrfToken, halfFilledServer, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestDeleteServerRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/servers/delete/%d", id)
		req, _ := http.NewRequest("DELETE", url, nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Delete Server", func(t *testing.T) {
		server := CreateServer()
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, server.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Server not found", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 400
		serverID := uint(999)
		got := makeRequest(t, authToken, csrfToken, serverID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestConnectServerRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, serverId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/private/servers/connect/%d", serverId)
		req, _ := http.NewRequest("GET", url, nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Connect Server", func(t *testing.T) {
		s := CreateServer()
		plan := CreatePlan()
		user := CreateUser()
		_ = CreateUserPlan(user.ID, plan.ID, true)
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Server Not Found", func(t *testing.T) {
		plan := CreatePlan()
		user := CreateUser()
		_ = CreateUserPlan(user.ID, plan.ID, true)
		authToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, csrfToken, 100)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("User Plan Not Found", func(t *testing.T) {
		s := CreateServer()
		user := CreateUser()
		authToken, csrfToken := GetToken(user)
		want := 401
		got := makeRequest(t, authToken, csrfToken, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("User Plan Expired", func(t *testing.T) {
		plan := CreatePlan()
		user := CreateUser()
		_ = CreateUserPlan(user.ID, plan.ID, false)
		authToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, csrfToken, 100)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

}
