package projectfiber

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct{ DB *pgxpool.Pool }

func New(db *pgxpool.Pool) *Handler { return &Handler{DB: db} }

func (h *Handler) Mount(app *fiber.App) {
	v1 := app.Group("/api/v1")
	p := v1.Group("/project")
	p.Get("/", h.List)
}

type dataProject struct {
	Product_id   string `json:"project_id"`
	Product_name string `json:"project_name"`
}

func (h *Handler) List(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
	defer cancel()

	rows, err := h.DB.Query(ctx, `
		SELECT project_id,project_name
		FROM tbl_project_raw
		
	`)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}
	defer rows.Close()

	outData := make([]dataProject, 0, 30)
	for rows.Next() {
		var d dataProject
		if err := rows.Scan(&d.Product_id, &d.Product_name); err != nil {
			return fiber.NewError(fiber.StatusBadGateway, err.Error())
		}
		outData = append(outData, d)
	}

	return c.JSON(outData)
}
