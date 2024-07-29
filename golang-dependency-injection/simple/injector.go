

//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitialiazedService(isError bool) (*SimpleService, error){
	
	wire.Build(NewSimpleReposiotry, NewSimpleService)
	return nil, nil
}

func InitialiazedDatabaseRepository() *DatabaseRepository{
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitialiazedFooBarService() *FooBarService{
	wire.Build(
		fooSet,
		barSet,
		NewFooBarService,
	)
	return nil
}

// salah
// func InitialiazedHelloService() *HelloService  {
// 	wire.Build(NewHelloService, NewSayHelloImpl)
// 	return nil
// }

var HelloSet = wire.NewSet( //injector set binding
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitialiazedHelloService() *HelloService  {
	wire.Build(HelloSet, NewHelloService)
	return nil
}


func InitializedFooBar() *FooBar {
	wire.Build(
		NewFoo, NewBar,
		// wire.Struct(new(FooBar), "Foo", "Bar"))
		wire.Struct(new(FooBar), "*"))
	return nil
}

var FooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}),
)

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(
		FooBarValueSet,
		// wire.Struct(new(FooBar), "Foo", "Bar"))
		wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration{
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),)
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(
		NewConnection,
		NewFile,
	)
	return nil, nil
}

