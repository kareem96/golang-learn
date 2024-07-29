package golangdatabase

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T)  {
	
}


// open connection database in golang
func TestOpenConnection(t *testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name) VALUES ('abdul', 'Karim');")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert data to Database")
}