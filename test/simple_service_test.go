package test

import (
	"belajar-golang-dependency-injection/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService.SimpleRepository)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
