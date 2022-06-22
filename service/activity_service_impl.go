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

type ActivityServiceImpl struct {
	ActivityRepository repository.ActivityRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewActivityService(activityRepository repository.ActivityRepository, db *sql.DB, validate *validator.Validate) ActivityService {
	return &ActivityServiceImpl{
		ActivityRepository: activityRepository, 
		DB:                 db, 
		Validate:           validate, 
	}
}

func (self *ActivityServiceImpl) Create(ctx context.Context, request web.ActivityCreateRequest) web.ActivityResponse {
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	activity := domain.Activity{
		Email: request.Email, 
		Title: request.Title, 
	}

	activity = self.ActivityRepository.Save(ctx, tx, activity)
	return helper.ToActivityResponse(activity)
}

func (self *ActivityServiceImpl) Update(ctx context.Context, request web.ActivityUpdateRequest) web.ActivityResponse {
	err := self.Validate.Struct(request)
	helper.PanicIfError(err)
	
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing activity not found
	activity, err := self.ActivityRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// Update existing activity
	activity.Email = request.Email
	activity.Title = request.Title
	activity = self.ActivityRepository.Update(ctx, tx, activity)
	return helper.ToActivityResponse(activity)
}

func (self *ActivityServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing activity not found
	_, err = self.ActivityRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	// Delete existing activity
	self.ActivityRepository.Delete(ctx, tx, id)
}

func (self *ActivityServiceImpl) FindById(ctx context.Context, id int) web.ActivityResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Fail if existing activity not found
	activity, err := self.ActivityRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToActivityResponse(activity)
}

func (self *ActivityServiceImpl) FindAll(ctx context.Context) []web.ActivityResponse {
	tx, err := self.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	listActivity := self.ActivityRepository.FindAll(ctx, tx)
	return helper.ToActivityResponses(listActivity)
}
