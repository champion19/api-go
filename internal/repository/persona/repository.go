package persona

import (
	"context"
	"database/sql"
	"ejemplo/internal/domain/model"
)

const (
	queryGetAll = "SELECT * FROM personas"
	queryGetById = "SELECT * FROM personas WHERE id = ?"
)

type Repository interface {
	GetAll(ctx context.Context) ([]model.Persona, error)
	GetById(ctx context.Context, id int64) (model.Persona, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]model.Persona, error) {
	rows, err := r.db.QueryContext(ctx, queryGetAll)
	if err != nil {
		return nil, err
	}

	personas := []model.Persona{}

	for rows.Next() {
		persona := model.Persona{}
		err = rows.Scan(&persona.ID, &persona.Nombre, &persona.Apellido, &persona.Edad)
		if err != nil {
			return nil, err
		}

		personas = append(personas, persona)
	}

	return personas, nil
}

func (r *repository) GetById(ctx context.Context, id int64) (model.Persona, error) {
	row := r.db.QueryRowContext(ctx, queryGetById, id)
	persona := model.Persona{}
	err := row.Scan(
		&persona.ID,
		&persona.Nombre,
		&persona.Apellido,
		&persona.Edad,
	)
	
	if err != nil {
		return model.Persona{}, err
	}

	return persona, nil
}