package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/yMaatheus/tech-challenge-snet/model"
	"github.com/yMaatheus/tech-challenge-snet/service"
	"go.uber.org/zap"
)

type mockEstablishmentService struct {
	createErr     error
	findAllErr    error
	findByIDErr   error
	updateErr     error
	deleteErr     error
	returnNilOnID bool
}

func (m *mockEstablishmentService) Create(_ context.Context, e *model.Establishment) error {
	if m.createErr != nil {
		return m.createErr
	}
	e.ID = 1
	return nil
}
func (m *mockEstablishmentService) FindAll(_ context.Context) ([]model.EstablishmentWithStoresTotal, error) {
	if m.findAllErr != nil {
		return nil, m.findAllErr
	}
	return []model.EstablishmentWithStoresTotal{
		{
			ID:            1,
			Number:        "E001",
			Name:          "Test",
			CorporateName: "Corp Test",
			Address:       "Rua Teste",
			AddressNumber: "10",
			City:          "Cidade Teste",
			State:         "ST",
			ZipCode:       "12345678",
			StoresTotal:   2,
		},
	}, nil
}
func (m *mockEstablishmentService) FindByID(_ context.Context, id int64) (*model.EstablishmentWithStores, error) {
	if m.findByIDErr != nil {
		return nil, m.findByIDErr
	}
	if m.returnNilOnID || id == 999 {
		return nil, nil
	}
	return &model.EstablishmentWithStores{
		ID:            1,
		Number:        "E001",
		Name:          "Test",
		CorporateName: "Corp Test",
		Address:       "Rua 1",
		AddressNumber: "10",
		City:          "Cidade",
		State:         "ST",
		ZipCode:       "12345-000",
		Stores:        []model.Store{},
	}, nil
}
func (m *mockEstablishmentService) Update(_ context.Context, e *model.Establishment) error {
	return m.updateErr
}
func (m *mockEstablishmentService) Delete(_ context.Context, id int64) error {
	return m.deleteErr
}

func setupTestEchoWithService(svc service.EstablishmentService) *echo.Echo {
	e := echo.New()
	logger := zap.NewNop()
	NewEstablishmentHandler(e, svc, logger)
	return e
}

func setupTestEcho() *echo.Echo {
	return setupTestEchoWithService(&mockEstablishmentService{})
}

func TestCreateEstablishment(t *testing.T) {
	e := setupTestEcho()

	reqBody, _ := json.Marshal(map[string]interface{}{
		"number":         "E001",
		"name":           "Test",
		"corporate_name": "Corp Test",
		"address":        "Rua Teste",
		"address_number": "10",
		"city":           "Cidade Teste",
		"state":          "ST",
		"zip_code":       "12345678",
	})
	req := httptest.NewRequest(http.MethodPost, "/establishments", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "Establishment created successfully")
}

func TestCreateEstablishment_InvalidJSON(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodPost, "/establishments", bytes.NewReader([]byte(`{invalid`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid request body")
}

func TestCreateEstablishment_ValidationError(t *testing.T) {
	e := setupTestEcho()
	reqBody, _ := json.Marshal(map[string]interface{}{"number": ""})
	req := httptest.NewRequest(http.MethodPost, "/establishments", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "validation_error")
}

func TestCreateEstablishment_ServiceError(t *testing.T) {
	mockSvc := &mockEstablishmentService{createErr: errors.New("fail create")}
	e := setupTestEchoWithService(mockSvc)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"number":         "E001",
		"name":           "Test",
		"corporate_name": "Corp Test",
		"address":        "Rua Teste",
		"address_number": "10",
		"city":           "Cidade Teste",
		"state":          "ST",
		"zip_code":       "12345678",
	})
	req := httptest.NewRequest(http.MethodPost, "/establishments", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "fail create")
}

func TestListEstablishments(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/establishments", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Test")
	assert.Contains(t, rec.Body.String(), "storesTotal")
}

func TestListEstablishments_ServiceError(t *testing.T) {
	mockSvc := &mockEstablishmentService{findAllErr: errors.New("fail list")}
	e := setupTestEchoWithService(mockSvc)
	req := httptest.NewRequest(http.MethodGet, "/establishments", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "fail list")
}

func TestGetEstablishmentByID(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/establishments/1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"id":1`)
	assert.Contains(t, rec.Body.String(), `"stores":`)
}

func TestGetEstablishmentByID_BadID(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/establishments/bad", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid establishment ID. Must be a positive integer.")
}

func TestGetEstablishmentByID_NotFound(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodGet, "/establishments/999", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, rec.Body.String(), "Establishment not found")
}

func TestGetEstablishmentByID_ServiceError(t *testing.T) {
	mockSvc := &mockEstablishmentService{findByIDErr: errors.New("find error")}
	e := setupTestEchoWithService(mockSvc)
	req := httptest.NewRequest(http.MethodGet, "/establishments/1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "find error")
}

func TestUpdateEstablishment(t *testing.T) {
	e := setupTestEcho()
	reqBody, _ := json.Marshal(map[string]interface{}{
		"number":         "E001",
		"name":           "Test Updated",
		"corporate_name": "Corp Test",
		"address":        "Rua Teste",
		"address_number": "10",
		"city":           "Cidade Teste",
		"state":          "ST",
		"zip_code":       "12345678",
	})
	req := httptest.NewRequest(http.MethodPut, "/establishments/1", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "updated successfully")
}

func TestUpdateEstablishment_BadID(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodPut, "/establishments/bad", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid establishment ID. Must be a positive integer.")
}

func TestUpdateEstablishment_InvalidBody(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodPut, "/establishments/1", bytes.NewReader([]byte("{invalid")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid request body")
}

func TestUpdateEstablishment_ValidationError(t *testing.T) {
	e := setupTestEcho()
	reqBody := []byte(`{"number":""}`)
	req := httptest.NewRequest(http.MethodPut, "/establishments/1", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "validation_error")
}

func TestUpdateEstablishment_ServiceError(t *testing.T) {
	mockSvc := &mockEstablishmentService{updateErr: errors.New("fail update")}
	e := setupTestEchoWithService(mockSvc)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"number":         "E001",
		"name":           "Test Updated",
		"corporate_name": "Corp Test",
		"address":        "Rua Teste",
		"address_number": "10",
		"city":           "Cidade Teste",
		"state":          "ST",
		"zip_code":       "12345678",
	})
	req := httptest.NewRequest(http.MethodPut, "/establishments/1", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "fail update")
}

func TestDeleteEstablishment(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodDelete, "/establishments/1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "deleted successfully")
}

func TestDeleteEstablishment_BadID(t *testing.T) {
	e := setupTestEcho()
	req := httptest.NewRequest(http.MethodDelete, "/establishments/bad", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid establishment ID. Must be a positive integer.")
}

func TestDeleteEstablishment_ServiceError(t *testing.T) {
	mockSvc := &mockEstablishmentService{deleteErr: errors.New("fail delete")}
	e := setupTestEchoWithService(mockSvc)
	req := httptest.NewRequest(http.MethodDelete, "/establishments/1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "fail delete")
}
