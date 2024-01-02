package handler

import (
	"ejemplo/internal/domain/personas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	service personas.Service
}

func NewPersona(service personas.Service) *Persona {
	return &Persona{
		service: service,
	}
}

func (p *Persona) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		personas, err := p.service.GetAll(ctx)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, personas)
	}
}

func (p *Persona) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		persona, err := p.service.GetById(ctx, int64(id))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, persona)
	}
}
