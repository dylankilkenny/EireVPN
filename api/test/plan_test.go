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

func TestGetPlanRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, planId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/plans/%d", planId)
		req, _ := http.NewRequest("GET", url, nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve plan by ID", func(t *testing.T) {
		user := CreateAdminUser()
		plan := CreatePlan()
		want := 200
		got := makeRequest(t, user, plan.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Plan not found", func(t *testing.T) {
		user := CreateAdminUser()
		want := 400
		planID := uint(999)
		got := makeRequest(t, user, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestCreatePlanRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, plan map[string]interface{}) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		req, _ := http.NewRequest("POST", "/api/protected/plans/create", bytes.NewBuffer(j))
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Plan Creation", func(t *testing.T) {
		plan := map[string]interface{}{
			"name":           "Test Product in test mode",
			"amount":         500,
			"interval":       "month",
			"interval_count": 1,
			"currency":       "EUR",
			"plan_type":      "PAYG",
		}
		user := CreateAdminUser()
		want := 200
		got := makeRequest(t, user, plan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledPlan := map[string]interface{}{
			"name":           "",
			"amount":         "",
			"interval":       "month",
			"interval_count": 1,
			"currency":       "EUR",
		}
		user := CreateAdminUser()
		want := 400
		got := makeRequest(t, user, halfFilledPlan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Drop table - Internal Server Error", func(t *testing.T) {
		plan := map[string]interface{}{
			"name":           "Test Product in test mode",
			"amount":         500,
			"interval":       "month",
			"interval_count": 1,
			"currency":       "EUR",
			"plan_type":      "PAYG",
		}
		user := CreateAdminUser()
		want := 500
		DropPlanTable()
		got := makeRequest(t, user, plan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestAllPlansRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/plans", nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all plans", func(t *testing.T) {
		_ = CreatePlan()
		user := CreateUser()
		want := 200
		got := makeRequest(t, user)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := CreateUser()
		want := 500
		DropPlanTable()
		got := makeRequest(t, user)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdatePlanRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, plan map[string]interface{}, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		url := fmt.Sprintf("/api/protected/plans/update/%d", id)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Update Plan", func(t *testing.T) {
		plan := map[string]interface{}{
			"name": "Update test plan",
		}
		p := CreatePlan()
		user := CreateAdminUser()
		want := 200
		got := makeRequest(t, user, plan, p.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledPlan := map[string]interface{}{
			"name": "",
		}
		p := CreatePlan()
		user := CreateAdminUser()
		want := 400
		got := makeRequest(t, user, halfFilledPlan, p.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestDeletePlanRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/plans/delete/%d", id)
		req, _ := http.NewRequest("DELETE", url, nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Delete Plan", func(t *testing.T) {
		plan := CreatePlan()
		user := CreateAdminUser()
		want := 200
		got := makeRequest(t, user, plan.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Plan not found", func(t *testing.T) {
		user := CreateAdminUser()
		want := 400
		planID := uint(999)
		got := makeRequest(t, user, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}
