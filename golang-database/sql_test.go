package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t * testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES ('kareem', 'Kareem')"
	
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert data to Database")
}


func TestQuerySql(t * testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}


	for rows.Next(){
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}

	defer rows.Close()
	fmt.Println("Success query database")
}


func TestQuerySqlComplex(t * testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}


	for rows.Next(){
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var cratedAt time.Time
		var birthDate sql.NullTime
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &cratedAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("============================")
		fmt.Println("Id :", id,)
		fmt.Println("Name :", name)
		if email.Valid{
			fmt.Println("email :", email)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		if birthDate.Valid{
			fmt.Println("Birth Date :", birthDate)
		}
		fmt.Println("Married :", married)
		fmt.Println("Created At :", cratedAt)
	}

	defer rows.Close()
	fmt.Println("Success query database")
}


func TestSqlInjection(t * testing.T)  {
	db := GetConnection()
	defer db.Close()

	username := "admin' ; #"
	password := "salah"

	ctx := context.Background()

	// this query can injection
	// query := "SELECT username FROM user where username = '" + username + "' and password = '" + password + "' limit 1"

	// solution
	query := "SELECT username FROM user where username = ? and password = ? limit 1"

	fmt.Println(query)
	// rows, err := db.QueryContext(ctx, query)
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}


	if rows.Next(){
		var username string
		rows.Scan(&username)
		err := rows.Scan(&username)
		if err != nil{
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	}else{
		fmt.Println("Gagal Login")
	}

	defer rows.Close()
	fmt.Println("Success query database")
}

func TestExecSqlParameter(t * testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "kareem"
	password := "kareem"

	script := "INSERT INTO user(username, password) VALUES (?, ?)"
	
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new user")
}

func TestAutoIncrement(t * testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "ab@gmail.com"
	comment := "Test comment"

	script := "INSERT INTO comment(email, comment) VALUES (?, ?)"
	
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil{
		panic(err)
	}
	fmt.Println("Success insert new comment with id", insertId)
}

func TestPrepareStatement(t * testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comment(email, comment) VALUES (?, ?)"
	
	statement, err := db.PrepareContext(ctx, script,)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "ab" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini comment ke " + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInserId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}	
		fmt.Println("Comment with id", lastInserId)
	}
}

func TestTransaction(t * testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comment(email, comment) VALUES (?, ?)"


	// do transaction
	for i := 0; i < 10; i++ {
		email := "ab" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini comment ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		lastInserId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}	
		fmt.Println("Comment with id", lastInserId)
	}
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}