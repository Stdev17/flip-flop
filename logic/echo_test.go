package logic

import (
	"testing"
	"fmt"
)

func TestEcho(t *testing.T) {

	result, err := Request()
	if err != nil {
		t.Log(err)
		t.Errorf("Error")
	}

	fmt.Println(result)


	if result != "Hello, World!" {
		t.Log(result)
	}

}
