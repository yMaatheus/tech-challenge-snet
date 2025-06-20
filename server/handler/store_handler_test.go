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

type mockStoreService struct {
	CreateFn   func(ctx context.Context, store *model.Store) error
	FindAllFn  func(ctx context.Context) ([]model.Store, error)
	FindByIDFn func(ctx context.Context, id int64) (*model.Store, error)
	UpdateFn   func(ctx context.Context, store *model.Store) error
	DeleteFn   func(ctx context.Context, id int64) error
}

func (m *mockStoreService) Create(ctx context.Context, store *model.Store) error {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, store)
	}
	store.ID = 1
	return nil
}
func (m *mockStoreService) FindAll(ctx context.Context) ([]model.Store, error) {
	if m.FindAllFn != nil {
		return m.FindAllFn(ctx)
	}
	return []model.Store{
		{ID: 1, Name: "Loja A", Number: "S001", EstablishmentID: 1},
	}, nil
}
func (m *mockStoreService) FindByID(ctx context.Context, id int64) (*model.Store, error) {
	if m.FindByIDFn != nil {
		return m.FindByIDFn(ctx, id)
	}
	if id == 1 {
		return &model.Store{ID: 1, Name: "Loja A", Number: "S001", EstablishmentID: 1}, nil
	}
	return nil, nil
}
func (m *mockStoreService) Update(ctx context.Context, store *model.Store) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(ctx, store)
	}
	return nil
}
func (m *mockStoreService) Delete(ctx context.Context, id int64) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(ctx, id)
	}
	return nil
}

func setupStoreEcho(service service.StoreService) *echo.Echo {
	e := echo.New()
	logger := zap.NewNop()
	NewStoreHandler(e, service, logger)
	return e
}

func TestCreateStore_Success(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	store := model.Store{
		Number: "S001", Name: "StoreTest", CorporateName: "Corp", Address: "Rua", City: "Cidade",
		State: "ST", ZipCode: "12345678", AddressNumber: "10", EstablishmentID: 1,
	}
	body, _ := json.Marshal(store)
	req := httptest.NewRequest(http.MethodPost, "/stores", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "StoreTest")
}

func TestCreateStore_ValidationError(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	body := []byte(`{"name": ""}`)
	req := httptest.NewRequest(http.MethodPost, "/stores", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "validation_error")
}

func TestCreateStore_BadBody(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	req := httptest.NewRequest(http.MethodPost, "/stores", bytes.NewReader([]byte("{invalid_json")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestListStores_Success(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	req := httptest.NewRequest(http.MethodGet, "/stores", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Loja A")
}

func TestListStores_Error(t *testing.T) {
	mockSvc := &mockStoreService{
		FindAllFn: func(ctx context.Context) ([]model.Store, error) { return nil, errors.New("db error") },
	}
	e := setupStoreEcho(mockSvc)
	req := httptest.NewRequest(http.MethodGet, "/stores", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Could not fetch stores")
}

func TestGetStore_Success(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	req := httptest.NewRequest(http.MethodGet, "/stores/1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Loja A")
}

func TestGetStore_NotFound(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	req := httptest.NewRequest(http.MethodGet, "/stores/999", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetStore_InvalidID(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	req := httptest.NewRequest(http.MethodGet, "/stores/abc", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid store ID. Must be a positive integer.")
}

func TestUpdateStore_Success(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	store := model.Store{
		Number: "S002", Name: "Loja Atualizada", CorporateName: "Corp", Address: "Rua", City: "Cidade",
		State: "ST", ZipCode: "12345678", AddressNumber: "20", EstablishmentID: 1,
	}
	body, _ := json.Marshal(store)
	req := httptest.NewRequest(http.MethodPut, "/stores/1", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Store updated successfully")
}

func TestUpdateStore_InvalidID(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	store := model.Store{}
	body, _ := json.Marshal(store)
	req := httptest.NewRequest(http.MethodPut, "/stores/xyz", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid store ID. Must be a positive integer.")
}

func TestUpdateStore_ValidationError(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	body := []byte(`{"name": ""}`)
	req := httptest.NewRequest(http.MethodPut, "/stores/1", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "validation_error")
}

func TestDeleteStore_Success(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	req := httptest.NewRequest(http.MethodDelete, "/stores/1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Store deleted successfully")
}

func TestDeleteStore_InvalidID(t *testing.T) {
	e := setupStoreEcho(&mockStoreService{})
	req := httptest.NewRequest(http.MethodDelete, "/stores/abc", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid store ID. Must be a positive integer.")
}

func TestDeleteStore_Error(t *testing.T) {
	mockSvc := &mockStoreService{
		DeleteFn: func(ctx context.Context, id int64) error { return errors.New("db error") },
	}
	e := setupStoreEcho(mockSvc)
	req := httptest.NewRequest(http.MethodDelete, "/stores/1", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Could not delete store")
}
