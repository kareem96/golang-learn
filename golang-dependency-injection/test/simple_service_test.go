package test

import (
	"golang-dependency-injection/simple"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceError(t *testing.T)  {
	simpleService, err := simple.InitialiazedService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}
func TestSimpleServiceSuccess(t *testing.T)  {
	simpleService, err := simple.InitialiazedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}