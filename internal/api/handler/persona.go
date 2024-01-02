package handler

import (
	"ejemplo/internal/domain/personas"
	"net/http"

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
