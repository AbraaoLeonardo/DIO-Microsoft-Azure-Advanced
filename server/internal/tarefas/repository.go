package tarefas

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Tarefa, error)
	GetByID(ctx context.Context, id string) (*Tarefa, error)
	Create(ctx context.Context, tarefa *Tarefa) (*Tarefa, error)
	Update(ctx context.Context, id string, tarefa *Tarefa) error
	Delete(ctx context.Context, id string) error
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]Tarefa, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, titulo, descricao, status FROM tarefas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tarefas []Tarefa
	for rows.Next() {
		var t Tarefa
		if err := rows.Scan(&t.ID, &t.Titulo, &t.Descricao, &t.Status); err != nil {
			return nil, err
		}
		tarefas = append(tarefas, t)
	}
	return tarefas, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*Tarefa, error) {
	var t Tarefa
	err := r.db.QueryRowContext(ctx,
		"SELECT id, titulo, descricao, status FROM tarefas WHERE id = $1", id).
		Scan(&t.ID, &t.Titulo, &t.Descricao, &t.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *repository) Create(ctx context.Context, tarefa *Tarefa) (*Tarefa, error) {
	var id int64
	err := r.db.QueryRowContext(ctx,
		"INSERT INTO tarefas (titulo, descricao, status) VALUES ($1, $2, $3) RETURNING id",
		tarefa.Titulo, tarefa.Descricao, tarefa.Status).Scan(&id)
	if err != nil {
		return nil, err
	}
	tarefa.ID = id
	return tarefa, nil
}

func (r *repository) Update(ctx context.Context, id string, tarefa *Tarefa) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE tarefas SET titulo=$1, descricao=$2, status=$3 WHERE id=$4",
		tarefa.Titulo, tarefa.Descricao, tarefa.Status, id)
	return err
}

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM tarefas WHERE id = $1", id)
	return err
}
