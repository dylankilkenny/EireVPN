package test

import (
	"bytes"
	"eirevpn/api/models"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetServerRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, serverId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/servers/%d", serverId)
		req, _ := http.NewRequest("GET", url, nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve server by ID", func(t *testing.T) {
		user := CreateAdminUser()
		server := CreateServer()
		want := 200
		got := makeRequest(t, user, server.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Server not found", func(t *testing.T) {
		user := CreateAdminUser()
		want := 400
		serverID := uint(999)
		got := makeRequest(t, user, serverID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestCreateServerRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, boundary string, server bytes.Buffer) int {
		t.Helper()
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/api/protected/servers/create", strings.NewReader(server.String()))
		req.Header.Add("Content-Type", boundary)
		AddTokens(user, req)
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
		want := 200
		got := makeRequest(t, user, w.FormDataContentType(), *buf)
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
		want := 400
		got := makeRequest(t, user, w.FormDataContentType(), *buf)
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
		want := 500
		DropServerTable()
		got := makeRequest(t, user, w.FormDataContentType(), *buf)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestAllServersRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/private/servers", nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all servers", func(t *testing.T) {
		_ = CreateServer()
		user := CreateUser()
		want := 200
		got := makeRequest(t, user)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("No Servers Found", func(t *testing.T) {
		user := CreateUser()
		want := 200
		got := makeRequest(t, user)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := CreateUser()
		want := 500
		DropServerTable()
		got := makeRequest(t, user)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdateServerRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, server map[string]interface{}, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(server)
		url := fmt.Sprintf("/api/protected/servers/update/%d", id)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		AddTokens(user, req)
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
		want := 200
		got := makeRequest(t, user, server, s.ID)
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
		want := 400
		got := makeRequest(t, user, halfFilledServer, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestDeleteServerRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/servers/delete/%d", id)
		req, _ := http.NewRequest("DELETE", url, nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Delete Server", func(t *testing.T) {
		server := CreateServer()
		user := CreateAdminUser()
		want := 200
		got := makeRequest(t, user, server.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Server not found", func(t *testing.T) {
		user := CreateAdminUser()
		want := 400
		serverID := uint(999)
		got := makeRequest(t, user, serverID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestConnectServerRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, serverId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/private/servers/connect/%d", serverId)
		req, _ := http.NewRequest("GET", url, nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Connect Server", func(t *testing.T) {
		s := CreateServer()
		plan := CreatePlan()
		user := CreateUser()
		_ = CreateUserPlan(user.ID, plan.ID, true)
		want := 200
		got := makeRequest(t, user, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Server Not Found", func(t *testing.T) {
		plan := CreatePlan()
		user := CreateUser()
		_ = CreateUserPlan(user.ID, plan.ID, true)
		want := 400
		got := makeRequest(t, user, 100)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("User Plan Not Found", func(t *testing.T) {
		s := CreateServer()
		user := CreateUser()
		want := 401
		got := makeRequest(t, user, s.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("User Plan Expired", func(t *testing.T) {
		plan := CreatePlan()
		user := CreateUser()
		_ = CreateUserPlan(user.ID, plan.ID, false)
		want := 400
		got := makeRequest(t, user, 100)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

}
