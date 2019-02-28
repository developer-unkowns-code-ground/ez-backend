package category

import (
	"net/http"

	"github.com/labstack/echo"
)

// HTTPHandler struct
type HTTPHandler struct {
}

// Category Struct
// ex
type Category struct {
	ID int64 `json:"id"`
}

func NewHttpCategoryHandler(e *echo.Echo) {
	handler := &HTTPHandler{}

	categories := []Category{
		{ID: 1},
	}

	e.GET("/api/categories", func(c echo.Context) error {
		return c.JSON(http.StatusOK, categories)
	})
	e.GET("/api/category/:id", handler.GetCategory)
	e.DELETE("/api/categori/:id", handler.DeleteCategory)
}

func (h *HTTPHandler) GetCategory(c echo.Context) error {

	category := &Category{ID: 1}
	return c.JSON(http.StatusOK, category)
}

func (h *HTTPHandler) DeleteCategory(c echo.Context) error {
	return c.JSON(204, nil)
}
