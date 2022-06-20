package repository

import (
	"database/sql"
	"context"
	"errors"

	"github.com/michaelact/firstApi/model/domain"
	"github.com/michaelact/firstApi/helper"
)

type ActivityRepositoryImpl struct {

}

func (self *ActivityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, activity domain.Activity) domain.Activity {
	// Insert new activity
	SQLPut := "INSERT INTO Activity(title, email) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, SQLPut, activity.Title, activity.Email)
	helper.PanicIfError(err)

	// Return created activity
	id, err := result.LastInsertId()
	activity, err = self.FindById(ctx, tx, int(id))
	return activity
}

func (self *ActivityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, activity domain.Activity) domain.Activity {
	// Update existing activity
	SQLPut := "UPDATE Activity SET title=?, email=?, updated_at=NOW() WHERE id=?"
	_, err := tx.ExecContext(ctx, SQLPut, activity.Title, activity.Email, activity.UpdatedAt, activity.Id)
	helper.PanicIfError(err)

	// Return updated activity
	activity, err = self.FindById(ctx, tx, activity.Id)
	helper.PanicIfError(err)
	return activity
}

func (self *ActivityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	// Delete existing activity
	SQLDel := "DELETE FROM Activity WHERE id=?"	
	result, err := tx.ExecContext(ctx, SQLDel, id)
	helper.PanicIfError(err)
}

func (self *ActivityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Activity, error) {
	// Extract existing activity
	SQLGet := "SELECT title, email, created_at, updated_at, deleted_at FROM Activity WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQLGet, id)
	helper.PanicIfError(err)

	// Bind all columns value to activity variable
	activity := new(domain.Activity)
	activity.Id = id
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&activity.Title, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)
		helper.PanicIfError(err)
		return activity, nil
	} else {
		return activity, errors.New("Activity not found")
	}
}

func (self *ActivityRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, activity domain.Activity) []domain.Activity {
	// Extract all activity
	SQLGet := "SELECT id, title, email, created_at, updated_at, deleted_at FROM Activity"
	rows, err := tx.QueryContext(ctx. SQLGet)
	helper.PanicIfError(err)

	// Iterate all extracted rows
	var listActivity []domain.Activity
	defer rows.Close()
	for rows.Next() {
		activity := new(domain.Activity)
		err := rows.Scan(&activity.Id, &activity.Email, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)
		helper.PanicIfError(err)

		listActivity = append(listActivity, activity)
	}

	return listActivity
}
