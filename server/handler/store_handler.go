package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yMaatheus/tech-challenge-snet/model"
	"github.com/yMaatheus/tech-challenge-snet/service"
	"go.uber.org/zap"
)

// StoreHandler handles store endpoints
type StoreHandler struct {
	Service service.StoreService
	Logger  *zap.Logger
}

// NewStoreHandler sets up the routes for Store
func NewStoreHandler(e *echo.Echo, svc service.StoreService, logger *zap.Logger) {
	h := &StoreHandler{Service: svc, Logger: logger}
	e.POST("/stores", h.Create)
	e.GET("/stores", h.List)
	e.GET("/stores/:id", h.Get)
	e.PUT("/stores/:id", h.Update)
	e.DELETE("/stores/:id", h.Delete)
}

// CreateStore godoc
// @Summary      Create a new store
// @Tags         stores
// @Accept       json
// @Produce      json
// @Param        store  body     model.Store  true  "Store to create"
// @Success      201    {object} model.Store
// @Failure      400    {object} map[string]string
// @Failure      500    {object} map[string]string
// @Router       /stores [post]
func (h *StoreHandler) Create(c echo.Context) error {
	var store model.Store
	if err := c.Bind(&store); err != nil {
		h.Logger.Warn("Failed to bind store", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := h.Service.Create(c.Request().Context(), &store); err != nil {
		h.Logger.Error("Failed to create store", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create store"})
	}
	h.Logger.Info("Store created", zap.Int64("id", store.ID))
	return c.JSON(http.StatusCreated, store)
}

// ListStores godoc
// @Summary      List all stores
// @Tags         stores
// @Produce      json
// @Success      200  {array}  model.Store
// @Failure      500  {object} map[string]string
// @Router       /stores [get]
func (h *StoreHandler) List(c echo.Context) error {
	stores, err := h.Service.FindAll(c.Request().Context())
	if err != nil {
		h.Logger.Error("Failed to list stores", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch stores"})
	}
	return c.JSON(http.StatusOK, stores)
}

// GetStore godoc
// @Summary      Get store by ID
// @Tags         stores
// @Produce      json
// @Param        id   path      int  true  "Store ID"
// @Success      200  {object}  model.Store
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /stores/{id} [get]
func (h *StoreHandler) Get(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	store, err := h.Service.FindByID(c.Request().Context(), id)
	if err != nil {
		h.Logger.Error("Failed to get store", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch store"})
	}
	if store == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Store not found"})
	}
	return c.JSON(http.StatusOK, store)
}

// UpdateStore godoc
// @Summary      Update a store by ID
// @Tags         stores
// @Accept       json
// @Produce      json
// @Param        id    path     int         true  "Store ID"
// @Param        store body     model.Store true  "Store update"
// @Success      200   {object} map[string]string
// @Failure      400   {object} map[string]string
// @Failure      500   {object} map[string]string
// @Router       /stores/{id} [put]
func (h *StoreHandler) Update(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var store model.Store
	if err := c.Bind(&store); err != nil {
		h.Logger.Warn("Failed to bind store", zap.Error(err))
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	store.ID = id
	if err := h.Service.Update(c.Request().Context(), &store); err != nil {
		h.Logger.Error("Failed to update store", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update store"})
	}
	h.Logger.Info("Store updated", zap.Int64("id", id))
	return c.JSON(http.StatusOK, map[string]string{"message": "Store updated successfully"})
}

// DeleteStore godoc
// @Summary      Delete a store by ID
// @Tags         stores
// @Produce      json
// @Param        id   path      int  true  "Store ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /stores/{id} [delete]
func (h *StoreHandler) Delete(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.Service.Delete(c.Request().Context(), id); err != nil {
		h.Logger.Error("Failed to delete store", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete store"})
	}
	h.Logger.Info("Store deleted", zap.Int64("id", id))
	return c.JSON(http.StatusOK, map[string]string{"message": "Store deleted successfully"})
}
