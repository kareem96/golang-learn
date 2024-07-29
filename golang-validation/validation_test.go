package golangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate = validator.New()
	if validate == nil{
		t.Error("Validate is nil")
	}
}

func TestValidatorVariable(t *testing.T) {
	validate := validator.New()
	user := ""

	err := validate.Var(user,"required")
	if err != nil{
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariable(t *testing.T) {
	validate := validator.New()
	password := "rahasia"
	confirmPassword := "salah"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if  err != nil {
		fmt.Println(err.Error())
	}
}
func TestValidateMulltipleTag(t *testing.T) {
	validate := validator.New()
	user := "12"

	err := validate.Var(user, "eqfield,numeric")
	if  err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest :=  LoginRequest{
		Username: "kareem",
		Password: "kar",
	}
	err := validate.Struct(loginRequest)
	if  err != nil {
		validateErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validateErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}
func TestValidationCrossField(t *testing.T) {
	type RegisterUser struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	loginRequest :=  RegisterUser{
		Username: "kareem@gmail.com",
		Password: "kareem",
		ConfirmPassword: "kareem",
	}

	err := validate.Struct(loginRequest)
	if  err != nil {
		fmt.Println(err.Error())
	}
}
func TestValidationNestedStruct(t *testing.T) {
	type Address struct {
		City string `validate:"required"`
		Country string `validate:"required"`
	}
	type User struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
		Address Address `validate:"required"`
	}

	validate := validator.New()
	loginRequest :=  User{
		Username: "kareem@gmail.com",
		Password: "kareem",
		ConfirmPassword: "kareem",
		Address: Address{
			City: "",
			Country: "",
		},
	}

	err := validate.Struct(loginRequest)
	if  err != nil {
		fmt.Println(err.Error())
	}
}
func TestValidationNestedStructCollection(t *testing.T) {
	type Address struct {
		City string `validate:"required"`
		Country string `validate:"required"`
	}
	type User struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
		Address []Address `validate:"required,dive"`
		Hobbies []string `validate:"dive,required,min=3"`
	}

	validate := validator.New()
	loginRequest :=  User{
		Username: "kareem@.com",
		Password: "kareem",
		ConfirmPassword: "",
		Address: []Address{
			{
				City: "",
				Country: "",
			},
			{
				City: "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Eating",
			"Gaming",
			"Coding",

		},
	}

	err := validate.Struct(loginRequest)
	if  err != nil {
		fmt.Println(err.Error())
	}
}
func TestValidationNestedStructMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Username        string            `validate:"required,email"`
		Password        string            `validate:"required,min=5"`
		ConfirmPassword string            `validate:"required,min=5,eqfield=Password"`
		Address         []Address         `validate:"required,dive"`
		Hobbies         []string          `validate:"dive,required,min=3"`
		School          map[string]School `validate:"dive,keys,required,min=2,endkeys"`
	}

	validate := validator.New()
	loginRequest := User{
		Username:        "kareem@.com",
		Password:        "kareem",
		ConfirmPassword: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Eating",
			"Gaming",
			"Coding",
		},
		School: map[string]School{
			"SD": {
				Name: "S",
			},
			"SMP": {
				Name: "",
			},
			"sas": {
				Name: "",
			},
		},
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationNestedBasicMap(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Username        string            `validate:"required,email"`
		Password        string            `validate:"required,min=5"`
		ConfirmPassword string            `validate:"required,min=5,eqfield=Password"`
		Address         []Address         `validate:"required,dive"`
		Hobbies         []string          `validate:"dive,required,min=3"`
		School          map[string]School `validate:"dive,keys,required,min=2,endkeys"`
		Wallet          map[string]int `validate:"dive,keys,required,min=2,endkeys,required,gt=100"`
	}

	validate := validator.New()
	loginRequest := User{
		Username:        "kareem@.com",
		Password:        "kareem",
		ConfirmPassword: "",
		Address: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Eating",
			"Gaming",
			"Coding",
		},
		School: map[string]School{
			"SD": {
				Name: "S",
			},
			"SMP": {
				Name: "",
			},
			"sas": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BCA": 10,
			"ABC": 0,
			"": 0,
		},
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id 		string `validate:"varchar,min=3"`
		Name 	string `validate:"varchar"`
		Owner 	string `validate:"varchar"`
		Slogan 	string `validate:"varchar"`
	}
	type User struct {
		Username        string            `validate:"required,email"`
		Password        string            `validate:"required,min=5"`
		ConfirmPassword string            `validate:"required,min=5,eqfield=Password"`
		Seller 			[]Seller          `validate:"required,dive"`
	}

	loginRequest := User{
		Username:        "kareem@sample.com",
		Password:        "kareem",
		ConfirmPassword: "kareem",
		Seller: []Seller{
			{
				Id:   	"12",
				Name: 	"asd",
				Owner:	"asd",
				Slogan: "dsv",
			},
			{
				Id:    	"2",
				Name: 	"adfs",
				Owner: 	"adfs",
				Slogan: "adfs",
			},
		},
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}