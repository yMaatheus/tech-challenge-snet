package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/yMaatheus/tech-challenge-snet/model"
	"go.uber.org/zap"
)

type mockEstablishmentService struct{}

func (m *mockEstablishmentService) Create(_ context.Context, e *model.Establishment) error {
	e.ID = 1
	return nil
}
func (m *mockEstablishmentService) FindAll(_ context.Context) ([]model.Establishment, error) {
	return []model.Establishment{
		{ID: 1, Name: "Test", Number: "E001"},
	}, nil
}
func (m *mockEstablishmentService) FindByID(_ context.Context, id int64) (*model.Establishment, error) {
	if id == 1 {
		return &model.Establishment{ID: 1, Name: "Test", Number: "E001"}, nil
	}
	return nil, nil
}
func (m *mockEstablishmentService) Update(_ context.Context, e *model.Establishment) error {
	return nil
}
func (m *mockEstablishmentService) Delete(_ context.Context, id int64) error {
	return nil
}

func setupTestEcho() *echo.Echo {
	e := echo.New()
	mockService := &mockEstablishmentService{}
	logger := zap.NewNop()

	NewEstablishmentHandler(e, mockService, logger)
	return e
}

func TestCreateEstablishment(t *testing.T) {
	e := setupTestEcho()

	// Request body
	reqBody, _ := json.Marshal(map[string]interface{}{
		"number": "E001",
		"name":   "Test",
	})
	req := httptest.NewRequest(http.MethodPost, "/establishments", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "Establishment created successfully")
}

func TestListEstablishments(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/establishments", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Test")
}

func TestGetEstablishmentByID(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/establishments/1", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Test")
}

func TestGetEstablishmentByID_NotFound(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/establishments/999", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUpdateEstablishment(t *testing.T) {
	e := setupTestEcho()
	reqBody, _ := json.Marshal(map[string]interface{}{
		"number": "E001",
		"name":   "Test Updated",
	})
	req := httptest.NewRequest(http.MethodPut, "/establishments/1", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteEstablishment(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodDelete, "/establishments/1", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
