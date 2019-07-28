package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlanRoute(t *testing.T) {

	makeRequest := func(t *testing.T, token string, planId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/plan/%d", planId)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)
		return w.Code
	}

	assertCorrectStatusCode := func(t *testing.T, want, got int) {
		t.Helper()
		ok := assert.Equal(t, want, got)
		if !ok {
			t.Errorf("Status Code is not %v. Got %v", want, got)
		}
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
		req, _ := http.NewRequest("POST", "/api/plan", bytes.NewBuffer(j))
		req.Header.Add("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)
		return w.Code
	}

	assertCorrectStatusCode := func(t *testing.T, want, got int) {
		t.Helper()
		ok := assert.Equal(t, want, got)
		if !ok {
			t.Errorf("Status Code is not %v. Got %v", want, got)
		}
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
}
