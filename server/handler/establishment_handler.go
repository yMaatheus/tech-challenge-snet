package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yMaatheus/tech-challenge-snet/model"
	"github.com/yMaatheus/tech-challenge-snet/service"
	"go.uber.org/zap"
)

// EstablishmentHandler handles HTTP requests for establishments
type EstablishmentHandler struct {
	service service.EstablishmentService
	Logger  *zap.Logger
}

// NewEstablishmentHandler registers establishment routes
func NewEstablishmentHandler(e *echo.Echo, s service.EstablishmentService, logger *zap.Logger) {
	h := &EstablishmentHandler{service: s}
	e.POST("/establishments", h.Create)
	e.GET("/establishments", h.List)
	e.GET("/establishments/:id", h.GetByID)
	e.PUT("/establishments/:id", h.Update)
	e.DELETE("/establishments/:id", h.Delete)
}

// Create godoc
// @Summary      Create a new establishment
// @Description  Creates a new establishment
// @Tags         establishments
// @Accept       json
// @Produce      json
// @Param        establishment  body      model.Establishment  true  "Establishment to create"
// @Success      201            {object}  map[string]interface{}
// @Failure      400            {object}  map[string]interface{}
// @Failure      500            {object}  map[string]interface{}
// @Router       /establishments [post]
func (h *EstablishmentHandler) Create(c echo.Context) error {
	var e model.Establishment
	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid request body"})
	}
	if err := h.service.Create(c.Request().Context(), &e); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Establishment created successfully", "id": e.ID})
}

// List godoc
// @Summary      List all establishments
// @Description  Get all establishments
// @Tags         establishments
// @Produce      json
// @Success      200  {array}   model.Establishment
// @Failure      500  {object}  map[string]interface{}
// @Router       /establishments [get]
func (h *EstablishmentHandler) List(c echo.Context) error {
	establishments, err := h.service.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, establishments)
}

// GetByID godoc
// @Summary      Get establishment by ID
// @Description  Get a specific establishment by its ID
// @Tags         establishments
// @Produce      json
// @Param        id   path      int  true  "Establishment ID"
// @Success      200  {object}  model.EstablishmentWithStores
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /establishments/{id} [get]
func (h *EstablishmentHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
	}
	establishment, err := h.service.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	if establishment == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Establishment not found"})
	}
	return c.JSON(http.StatusOK, establishment)
}

// Update godoc
// @Summary      Update establishment
// @Description  Update an existing establishment by its ID
// @Tags         establishments
// @Accept       json
// @Produce      json
// @Param        id             path      int                 true  "Establishment ID"
// @Param        establishment  body      model.Establishment true  "Establishment data"
// @Success      200            {object}  map[string]interface{}
// @Failure      400            {object}  map[string]interface{}
// @Failure      500            {object}  map[string]interface{}
// @Router       /establishments/{id} [put]
func (h *EstablishmentHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
	}
	var e model.Establishment
	if err := c.Bind(&e); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid request body"})
	}
	e.ID = id
	if err := h.service.Update(c.Request().Context(), &e); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Establishment updated successfully"})
}

// Delete godoc
// @Summary      Delete establishment
// @Description  Delete an establishment by its ID
// @Tags         establishments
// @Produce      json
// @Param        id   path      int  true  "Establishment ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /establishments/{id} [delete]
func (h *EstablishmentHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
	}
	if err := h.service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Establishment deleted successfully"})
}
