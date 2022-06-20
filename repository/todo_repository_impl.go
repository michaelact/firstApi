package repository

import (
	"database/sql"
	"context"
	"errors"

	"github.com/michaelact/firstApi/model/domain"
	"github.com/michaelact/firstApi/helper"
)

type TodoRepositoryImpl struct {

}

func (self *TodoRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	// Insert new todo
	SQLPut := "INSERT INTO Todo(activity_group_id, title) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, SQLPut)
	helper.PanicIfError(err)

	// Return created todo
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	todo, err = self.FindById(ctx, tx, int(id))
	return todo
}

func (self *TodoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	// Update existing todo
	SQLPut := "UPDATE Todo SET activity_group_id=?, title=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQLPut, todo.ActivityGroupId, todo.Title, todo.Id)
	helper.PanicIfError(err)

	// Return updated todo
	todo, err = self.FindById(ctx, tx, todo.Id)
	helper.PanicIfError(err)
	return todo	
}

func (self *TodoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	// Delete existing todo
	SQLDel := "DELETE FROM Todo WHERE id=?"
	_, err := tx.ExecContext(ctx, SQLDel, id)
	helper.PanicIfError(err)
}

func (self *TodoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Todo, error) {
	// Extract existing todo
	SQLGet := "SELECT activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at FROM Todo WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQLGet, id)
	helper.PanicIfError(err)

	// Bind all columns value to todo variable
	todo := new(domain.Todo)
	todo.Id = id
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&todo.ActivityGroupId, &todo.Title, &todo.IsActive, &todo.Priority, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("Todo not found.")
	}
}

func (self *TodoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, todo domain.Todo) []domain.Todo {
	// Extract all todolist
	SQLGet := "SELECT id, activity_group_id, is_active, priority, created_at, updated_at, deleted_at FROM Todo"
	rows, err := tx.QueryContext(ctx, SQLGet)
	helper.PanicIfError(err)

	// Iterate all extracted rows
	var listTodo []domain.Todo
	defer rows.Close()
	for rows.Next() {
		todo := new(domain.Todo)
		err := rows.Scan(&todo.Id, &todo.ActivityGroupId, &todo.IsActive, &todo.Priority, &todo.CreatedAt,  &todo.UpdatedAt, &todo.DeletedAt)
		helper.PanicIfError(err)

		listTodo = append(listTodo, todo)
	}

	return listTodo
}
