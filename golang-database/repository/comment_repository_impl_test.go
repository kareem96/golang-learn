package repository

import (
	belajar_golang_database "golang-database"
	"context"
	"fmt"
	"golang-database/entity"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)



func TestInsert(t *testing.T)  {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Ini tes repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
		
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T)  {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 37)
	if err != nil {
		panic(err)
		
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T)  {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
		
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}