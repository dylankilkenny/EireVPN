package test

import (
	"bytes"
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

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string, serverId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/servers/%d", serverId)
		req, _ := http.NewRequest("GET", url, nil)
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve server by ID", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		server := CreateServer()
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken, server.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Server not found", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		serverID := uint(999)
		got := makeRequest(t, authToken, refreshToken, csrfToken, serverID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestCreateServerRoute(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken, boundary string, server bytes.Buffer) int {
		t.Helper()
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/api/protected/servers/create", strings.NewReader(server.String()))
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Add("Content-Type", boundary)
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
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
		w.Close()
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken, w.FormDataContentType(), *buf)
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
		w.Close()
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, refreshToken, csrfToken, w.FormDataContentType(), *buf)
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
		w.Close()
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 500
		DropServerTable()
		got := makeRequest(t, authToken, refreshToken, csrfToken, w.FormDataContentType(), *buf)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestAllServersRoute(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/private/servers", nil)
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all servers", func(t *testing.T) {
		_ = CreateServer()
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("No Servers Found", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 500
		DropServerTable()
		got := makeRequest(t, authToken, refreshToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdateServerRoute(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string, server map[string]interface{}, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(server)
		url := fmt.Sprintf("/api/protected/servers/update/%d", id)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Update Server", func(t *testing.T) {
		server := map[string]interface{}{
			"ip": "127.0.0.1",
		}
		s := CreateServer()
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken, server, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledServer := map[string]interface{}{
			"ip": "",
		}
		s := CreateServer()
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, refreshToken, csrfToken, halfFilledServer, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestDeleteServerRoute(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/servers/delete/%d", id)
		req, _ := http.NewRequest("DELETE", url, nil)
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Delete Server", func(t *testing.T) {
		server := CreateServer()
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken, server.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Server not found", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		serverID := uint(999)
		got := makeRequest(t, authToken, refreshToken, csrfToken, serverID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}
