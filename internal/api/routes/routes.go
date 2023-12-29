package routes

import (
	"database/sql"
	"ejemplo/internal/api/handler"
	"ejemplo/internal/domain/personas"
	"ejemplo/internal/repository/persona"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{
		eng: eng,
		db: db,
	}
}

func (r *router) MapRoutes() {
	r.rg = r.eng.Group("/api/v1")

	r.buildPersonaRoutes()
}

func (r *router) buildPersonaRoutes() {
	repo := persona.NewRepository(r.db)
	service := personas.NewService(repo)
	handler := handler.NewPersona(service)

	pr := r.rg.Group("/personas")

	pr.GET("", handler.GetAll())
}
