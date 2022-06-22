package service

import (
	"database/sql"
	"context"

	"github.com/go-playground/validator/v10"

	"github.com/michaelact/firstApi/model/domain"
	"github.com/michaelact/firstApi/repository"
	"github.com/michaelact/firstApi/model/web"
	"github.com/michaelact/firstApi/helper"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTodoService(todoRepository repository.TodoRepository, db *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository, 
		DB:             db, 
		Validate:       validate, 
	}
}

func (self *TodoServiceImpl) Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse {
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := domain.Todo{
		ActivityGroupId: request.ActivityGroupId, 
		Title:           request.Title, 
	}

	todo = self.TodoRepository.Save(ctx, tx, todo)
	return helper.ToTodoResponse(todo)
}

func (self *TodoServiceImpl) Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse {
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)
	
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing todo not found
	todo, err := self.TodoRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// Update existing todo
	todo.ActivityGroupId = request.ActivityGroupId
	todo.Title = request.Title
	todo = self.TodoRepository.Update(ctx, tx, todo)
	return helper.ToTodoResponse(todo)
}

func (self *TodoServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing todo not found
	_, err = self.TodoRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	// Delete existing todo
	self.TodoRepository.Delete(ctx, tx, id)
}

func (self *TodoServiceImpl) FindById(ctx context.Context, id int) web.TodoResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing todo not found
	todo, err := self.TodoRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToTodoResponse(todo)
}

func (self *TodoServiceImpl) FindAll(ctx context.Context) []web.TodoResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	listTodo := self.TodoRepository.FindAll(ctx, tx)
	return helper.ToTodoResponses(listTodo)
}
