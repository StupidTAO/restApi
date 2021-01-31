package model

import (
	"encoding/json"
	"testing"
)

func TestReadARankValue(t *testing.T) {
	tokenTab := "AOA_ar_3"
	number := 10

	arr, err := ReadARankValue("erc20", tokenTab, number)
	if err != nil {
		t.Error("testReadARankValue occur err: ", err)
	}
	if len(arr) != number {
		t.Error("testReadARankValue occur err: result number not equal ", number)
	}

	t.Log("testReadARankValue is pass")
}

func TestGetRankJson(t *testing.T) {
	var docs []MyDocument
	tokenTab := "AOA_ar_3"
	number := 10

	arr, _ := ReadARankValue("erc20", tokenTab, number)
	jsonBytes, err := GetRankJson(arr)
	if err != nil {
		t.Error("TestGetRankJson occur err: ", err)
	}

	if len(jsonBytes) == 0 {
		t.Error("TestGetRankJson occur err: ", err)
	}

	err = json.Unmarshal(jsonBytes, &docs)
	if err != nil {
		t.Error("TestGetRankJson occur err: ", err)
	}

	t.Log("TestGetRankJson is pass")
}

func TestGetRankJsonNil(t *testing.T) {
	var docs []MyDocument
	_, err := GetRankJson(docs)
	if err != nil {
		t.Log("TestGetRankJsonNil array is empty")
	}
}