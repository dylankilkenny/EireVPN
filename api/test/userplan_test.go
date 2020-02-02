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

func TestGetUserPlanRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, planId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/userplans/%d", planId)
		req, _ := http.NewRequest("GET", url, nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve plan by ID", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		userplan := CreateUserPlan(1, 2, true)
		want := 200
		got := makeRequest(t, authToken, csrfToken, userplan.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Usser Plan not found", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 401
		planID := uint(999)
		got := makeRequest(t, authToken, csrfToken, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestCreateUserPlanRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, plan map[string]interface{}) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		req, _ := http.NewRequest("POST", "/api/protected/userplans/create", bytes.NewBuffer(j))
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful User Plan Creation", func(t *testing.T) {
		plan := map[string]interface{}{
			"user_id":     1,
			"plan_id":     4,
			"active":      "true",
			"start_date":  "2020-02-01 19:18",
			"expiry_date": "2020-03-01 19:18",
		}
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, plan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledPlan := map[string]interface{}{
			"user_id":     1,
			"plan_id":     4,
			"active":      "",
			"start_date":  "2020-02-01 19:18",
			"expiry_date": "",
		}
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, csrfToken, halfFilledPlan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

}

func TestAllUserPlansRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/protected/userplans", nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all user plans", func(t *testing.T) {
		_ = CreateUserPlan(1, 2, true)
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 500
		DropUserPlanTable()
		got := makeRequest(t, authToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdateUserPlanRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, plan map[string]interface{}, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		url := fmt.Sprintf("/api/protected/userplans/update/%d", id)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Update User Plan", func(t *testing.T) {
		plan := map[string]interface{}{
			"active":      "true",
			"start_date":  "2020-02-01 19:18",
			"expiry_date": "2020-09-01 19:18",
		}
		up := CreateUserPlan(1, 2, true)
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, plan, up.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledPlan := map[string]interface{}{
			"name": "",
		}
		up := CreateUserPlan(1, 2, true)
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, csrfToken, halfFilledPlan, up.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestDeleteUserPlanRoute(t *testing.T) {
	conf := config.GetConfig()
	makeRequest := func(t *testing.T, authToken, csrfToken string, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/userplans/delete/%d", id)
		req, _ := http.NewRequest("DELETE", url, nil)
		authCookie := http.Cookie{Name: conf.App.AuthCookieName, Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Delete User Plan", func(t *testing.T) {
		userPlan := CreateUserPlan(1, 2, true)
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, csrfToken, userPlan.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("User Plan not found", func(t *testing.T) {
		user := CreateAdminUser()
		authToken, csrfToken := GetToken(user)
		want := 401
		planID := uint(999)
		got := makeRequest(t, authToken, csrfToken, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}
