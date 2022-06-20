package helper

import (
	"github.com/michaelact/firstApi/model/domain"
	"github.com/michaelact/firstApi/model/web"
)

func ToActivityResponse(activity domain.Activity) web.ActivityResponse {
	return web.ActivityResponse{
		Id:        activity.Id, 
		Email:     activity.Email, 
		Title:     activity.Title, 
		CreatedAt: activity.CreatedAt, 
		UpdatedAt: activity.UpdatedAt, 
		DeletedAt: activity.DeletedAt,
	}
}

func ToActivityResponses(listActivity []domain.Activity) []web.ActivityResponse {
	var activityResponses []web.ActivityResponse
	for _, activity := range listActivity {
		activityResponses = append(activityResponses, ToActivityResponse(activity))
	}

	return activityResponses
}

func ToTodoResponse(todo domain.Todo) web.TodoResponse {
	return web.TodoResponse{
		Id:              todo.Id,
		ActivityGroupId: todo.ActivityGroupId,
		Title:           todo.Title,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.UpdatedAt,
		DeletedAt:       todo.DeletedAt,
	}
}

func ToTodoResponses(listTodo []domain.Todo) []web.TodoResponse {
	var todoResponses []web.TodoResponse
	for _, todo := range listTodo {
		todoResponses = append(todoResponses, ToTodoResponse(todo))
	}

	return todoResponses
}
