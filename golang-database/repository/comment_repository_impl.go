package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository  {
	return &commentRepositoryImpl{DB: db}
}


func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "insert into comment(email, comment) values (?, ?)"
	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	
	if err != nil{
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "select id, email, comment from comment where id = ? limit 1"
	rows, err := repo.DB.QueryContext(ctx, script, id)

	comment := entity.Comment{}

	if err != nil {
		return comment, err

	}
	defer rows.Close()
	if rows.Next(){
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}else{
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
	
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "select id, email, comment from comment"
	rows, err := repo.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}


	defer rows.Close()
	var comments []entity.Comment
	for rows.Next(){
		// ada
		comment := entity.Comment{}

		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)

		comments = append(comments, comment)
	}
	return comments, err
}


