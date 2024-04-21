package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codelikesuraj/gdsc-challenge-day-nineteen-twenty/models"
	"github.com/stretchr/testify/assert"
)

var DB = SetupDB()

func TestRegisterUser(t *testing.T) {
	RefreshDatabase(DB)

	router := SetupRouter(DB)
	rec := httptest.NewRecorder()

	reqBody := models.User{Username: "alskdjfaklsdfjkl", Password: "password"}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf(err.Error())
	}

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(rec, req)

	var registeredUser models.User

	if err := json.Unmarshal(rec.Body.Bytes(), &registeredUser); err != nil {
		t.Fatalf(err.Error())
	}

	t.Run("Returns 201 status", func(t *testing.T) {
		assert.Equal(t, rec.Code, http.StatusCreated, "%s", rec.Body.String())
	})

	t.Run("Contains username", func(t *testing.T) {
		var respBody map[string]models.User
		if err := json.Unmarshal(rec.Body.Bytes(), &respBody); err != nil {
			assert.Fail(t, err.Error())
		}
		assert.Equal(t, respBody["data"].Username, reqBody.Username)
	})
}
