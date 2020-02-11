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

func TestGetUserPlanRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, userId uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/private/userplans/%d", userId)
		req, _ := http.NewRequest("GET", url, nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Retrieve plan by User ID", func(t *testing.T) {
		user := CreateAdminUser()
		plan := CreatePlan()
		_ = CreateUserPlan(plan.ID, user.ID, true)
		want := 200
		got := makeRequest(t, user, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Usser Plan not found", func(t *testing.T) {
		user := CreateAdminUser()

		want := 401
		planID := uint(999)
		got := makeRequest(t, user, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestCreateUserPlanRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, plan map[string]interface{}) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		req, _ := http.NewRequest("POST", "/api/protected/userplans/create", bytes.NewBuffer(j))
		AddTokens(user, req)
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

		want := 200
		got := makeRequest(t, user, plan)
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

		want := 400
		got := makeRequest(t, user, halfFilledPlan)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

}

func TestAllUserPlansRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User) int {
		t.Helper()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/protected/userplans", nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful get all user plans", func(t *testing.T) {
		_ = CreateUserPlan(1, 2, true)
		user := CreateAdminUser()

		want := 200
		got := makeRequest(t, user)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := CreateAdminUser()

		want := 500
		DropUserPlanTable()
		got := makeRequest(t, user)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestUpdateUserPlanRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, plan map[string]interface{}, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		j, _ := json.Marshal(plan)
		url := fmt.Sprintf("/api/protected/userplans/update/%d", id)
		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(j))
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Update User Plan", func(t *testing.T) {
		userplan := map[string]interface{}{
			"active":      "true",
			"start_date":  "2020-02-01 19:18",
			"expiry_date": "2020-09-01 19:18",
		}
		user := CreateAdminUser()
		plan := CreatePlan()
		_ = CreateUserPlan(plan.ID, user.ID, true)

		want := 200
		got := makeRequest(t, user, userplan, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("Invalid form", func(t *testing.T) {
		halfFilledPlan := map[string]interface{}{
			"name": "",
		}
		user := CreateAdminUser()
		plan := CreatePlan()
		_ = CreateUserPlan(plan.ID, user.ID, true)

		want := 400
		got := makeRequest(t, user, halfFilledPlan, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}

func TestDeleteUserPlanRoute(t *testing.T) {
	makeRequest := func(t *testing.T, user *models.User, id uint) int {
		t.Helper()
		w := httptest.NewRecorder()
		url := fmt.Sprintf("/api/protected/userplans/delete/%d", id)
		req, _ := http.NewRequest("DELETE", url, nil)
		AddTokens(user, req)
		r.ServeHTTP(w, req)
		return w.Code
	}

	t.Run("Successful Delete User Plan", func(t *testing.T) {
		user := CreateAdminUser()
		plan := CreatePlan()
		_ = CreateUserPlan(plan.ID, user.ID, true)

		want := 200
		got := makeRequest(t, user, user.ID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})

	t.Run("User Plan not found", func(t *testing.T) {
		user := CreateAdminUser()

		want := 401
		planID := uint(999)
		got := makeRequest(t, user, planID)
		assertCorrectStatus(t, want, got)
		CreateCleanDB()
	})
}
