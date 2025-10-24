package test

import (
	"belajar-golang-dependency-injection/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
