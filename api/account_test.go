package api

import (
	"fmt"
	"github.com/test/restApi/log"
	"testing"
)

func InitTest() {
	err := log.LogInit()
	if err != nil {
		fmt.Println("panic: log init error")
	}
}

func TestARankTop(t *testing.T) {
	InitTest()

	tokenName := "AOA"
	_, err := ARankTop(tokenName)
	if err != nil {
		t.Log("TestARankTop occur error: ", err)
	}

	tokenName_err := "AOA123"
	_, err = ARankTop(tokenName_err)
	if err != nil {
		t.Log("TestARankTop occur error: ", err)
	}

	str, err := ARankTop("AOA")
	if err != nil {
		t.Error("TestARankTop occur error: ", err)
	}
	if len(str) == 0 {
		t.Error("TestARankTop occur error: ", err)
	}

	t.Log("TestARankTop is pass")
}
