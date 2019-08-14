package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetPlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string, planId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/private/plans/%d", planId)
		req, _ := http.NewRequest("GET", url, nil)
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve plan by ID", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		plan := CreatePlan()
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken, plan.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	// t.Run("Invalid token", func(t *testing.T) {
	// 	token := "invalid"
	// 	want := 401
	// 	plan := CreatePlan()
	// 	got := makeRequest(t, token, plan.ID)
	// 	assertCorrectStatus(t, want, got)
	// 	CreateCleanDB()
	// })

	t.Run("Plan not found", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		planID := uint(999)
		got := makeRequest(t, authToken, refreshToken, csrfToken, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestCreatePlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string, plan map[string]interface{}) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		req, _ := http.NewRequest("POST", "/api/private/plans", bytes.NewBuffer(j))
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Plan Creation", func(t *testing.T) {
		plan := map[string]interface{}{
			"name":            "1 Year",
			"type":            "subscription",
			"duration_hours":  0,
			"duration_days":   0,
			"duration_months": 12,
		}
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken, plan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledPlan := map[string]interface{}{
			"name":            "",
			"type":            "",
			"duration_hours":  0,
			"duration_days":   0,
			"duration_months": 12,
		}
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, refreshToken, csrfToken, halfFilledPlan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		plan := map[string]interface{}{
			"name":            "1 Year",
			"type":            "subscription",
			"duration_hours":  0,
			"duration_days":   0,
			"duration_months": 12,
		}
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 500
		DropPlanTable()
		got := makeRequest(t, authToken, refreshToken, csrfToken, plan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestAllPlansRoute(t *testing.T) {

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

	t.Run("Successful get all plans", func(t *testing.T) {
		_ = CreatePlan()
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("No Plans Found", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, refreshToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 500
		DropPlanTable()
		got := makeRequest(t, authToken, refreshToken, csrfToken)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdatePlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string, plan map[string]interface{}) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		req, _ := http.NewRequest("PUT", "/api/private/plans", bytes.NewBuffer(j))
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Update Plan", func(t *testing.T) {
		plan := map[string]interface{}{
			"name":            "1 Year",
			"type":            "subscription",
			"duration_hours":  0,
			"duration_days":   0,
			"duration_months": 12,
		}
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken, plan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		plan := map[string]interface{}{
			"name":            "1 Year",
			"type":            "subscription",
			"duration_hours":  0,
			"duration_days":   0,
			"duration_months": 12,
		}
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 500
		DropPlanTable()
		got := makeRequest(t, authToken, refreshToken, csrfToken, plan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledPlan := map[string]interface{}{
			"name":            "",
			"type":            "",
			"duration_hours":  0,
			"duration_days":   0,
			"duration_months": 12,
		}
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		got := makeRequest(t, authToken, refreshToken, csrfToken, halfFilledPlan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestDeletePlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, authToken, refreshToken, csrfToken string, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/private/plans/%d", id)
		req, _ := http.NewRequest("DELETE", url, nil)
		authCookie := http.Cookie{Name: "authToken", Value: authToken, Expires: time.Now().Add(time.Minute * 5)}
		refreshCookie := http.Cookie{Name: "refreshToken", Value: refreshToken, Expires: time.Now().Add(time.Minute * 5)}
		req.Header.Set("X-CSRF-Token", csrfToken)
		req.AddCookie(&authCookie)
		req.AddCookie(&refreshCookie)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Delete Plan", func(t *testing.T) {
		plan := CreatePlan()
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 200
		got := makeRequest(t, authToken, refreshToken, csrfToken, plan.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Plan not found", func(t *testing.T) {
		user := CreateUser()
		authToken, refreshToken, csrfToken := GetToken(user)
		want := 400
		planID := uint(999)
		got := makeRequest(t, authToken, refreshToken, csrfToken, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}
