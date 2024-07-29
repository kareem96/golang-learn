package golangvalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func MustValidationUsername(field validator.FieldLevel)  bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value){
			return false
		}
		if len(value) < 5{
			return false
		}
	}
	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidationUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	request := LoginRequest{
		Username: "KAEM",
		Password: "as",
	}

	err := validate.Struct(request)
	if err != nil{
		fmt.Print(err.Error())
	}
}


var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool  {
	length, err := strconv.Atoi(field.Param())
	if err != nil{
		panic(err)
	}
	value := field.Field().String()
	if !regexNumber.MatchString(value){
		return false
	}
	return len(value) == length
}

func TestValidationParamter(t *testing.T)  {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "2134",
		Pin: "123321",
	}


	err := validate.Struct(request)
	if err != nil{
		fmt.Println(err.Error())
	}

}
func TestOrTule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validare:"required"`
	}

	request := Login{
		Username: "123",
		Password: "asdas",
	}

	validate := validator.New()
	err := validate.Struct(request)

	if err != nil{
		fmt.Println(err.Error())
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool  {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("filed not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestValidationIgnoreCase(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email string `validate:"required,email"`
		Phone string `validate:"required,numeric"`
		Name string `validate:"required"`
	}
	request := User{
		Username: "08123",
		Email: "kareem@sample.com",
		Phone: "08123",
		Name: "kareem",
	}
	err := validate.Struct(request)
	if err != nil{
		fmt.Println(err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Email string `validate:"required,email"`
	Phone string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegistrationSucces(level validator.StructLevel)  {
	registerRequest := level.Current().Interface().(RegisterRequest)
	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone{

	}else{
		level.ReportError(registerRequest.Username, "Username", "Username", "Username", "")
	}
}

func TestValidateStructLevel(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegistrationSucces, RegisterRequest{})

	registRequest :=  RegisterRequest{
		Username: "kareem@gmail.coms",
		Email: "kareem@gmail.com",
		Phone: "08123",
		Password: "08123",
	}
	err := validate.Struct(registRequest)
	if err != nil{
		fmt.Println(err.Error())
	}
}