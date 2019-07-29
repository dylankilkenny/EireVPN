package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, token string, planId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/private/plan/%d", planId)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve plan by ID", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		plan := CreatePlan()
		want := 200
		got := makeRequest(t, token, plan.ID)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid token", func(t *testing.T) {
		token := "invalid"
		want := 401
		plan := CreatePlan()
		got := makeRequest(t, token, plan.ID)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Plan not found", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		want := 400
		planID := uint(999)
		got := makeRequest(t, token, planID)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestCreatePlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, token string, plan map[string]interface{}) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		req, _ := http.NewRequest("POST", "/api/private/plan", bytes.NewBuffer(j))
		req.Header.Add("Authorization", "Bearer "+token)
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
		token := GetToken(user)
		want := 200
		got := makeRequest(t, token, plan)
		assertCorrectStatusCode(t, want, got)
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
		token := GetToken(user)
		want := 400
		got := makeRequest(t, token, halfFilledPlan)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid token", func(t *testing.T) {
		plan := map[string]interface{}{
			"name":            "1 Year",
			"type":            "subscription",
			"duration_hours":  0,
			"duration_days":   0,
			"duration_months": 12,
		}
		token := "invalid"
		want := 401
		got := makeRequest(t, token, plan)
		assertCorrectStatusCode(t, want, got)
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
		token := GetToken(user)
		want := 500
		DropPlanTable()
		got := makeRequest(t, token, plan)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestAllPlansRoute(t *testing.T) {

	makeRequest := func(t *testing.T, token string) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/private/plans", nil)
		req.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all plans", func(t *testing.T) {
		_ = CreatePlan()
		user := CreateUser()
		token := GetToken(user)
		want := 200
		got := makeRequest(t, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("No Plans Found", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		want := 400
		got := makeRequest(t, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid token", func(t *testing.T) {
		token := "invalid"
		want := 401
		got := makeRequest(t, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		want := 500
		DropPlanTable()
		got := makeRequest(t, token)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdatePlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, token string, plan map[string]interface{}) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		req, _ := http.NewRequest("PUT", "/api/private/plan", bytes.NewBuffer(j))
		req.Header.Add("Authorization", "Bearer "+token)
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
		token := GetToken(user)
		want := 200
		got := makeRequest(t, token, plan)
		assertCorrectStatusCode(t, want, got)
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
		token := GetToken(user)
		want := 500
		DropPlanTable()
		got := makeRequest(t, token, plan)
		assertCorrectStatusCode(t, want, got)
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
		token := GetToken(user)
		want := 400
		got := makeRequest(t, token, halfFilledPlan)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}

func TestDeletePlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, token string, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/private/plan/%d", id)
		req, _ := http.NewRequest("DELETE", url, nil)
		req.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Delete Plan", func(t *testing.T) {
		plan := CreatePlan()
		user := CreateUser()
		token := GetToken(user)
		want := 200
		got := makeRequest(t, token, plan.ID)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})

	t.Run("Plan not found", func(t *testing.T) {
		user := CreateUser()
		token := GetToken(user)
		want := 400
		planID := uint(999)
		got := makeRequest(t, token, planID)
		assertCorrectStatusCode(t, want, got)
		CreateCleanDB()
	})
}
