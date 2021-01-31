package api

import (
	"fmt"
	"errors"
	"github.com/test/restApi/log"
	"github.com/test/restApi/model"
)

var count int64


func ARankTop(tokenName string)(string, error)  {

	tokenTab := fmt.Sprintf("%s_ar_3", tokenName)

	//check tokenName
	_, err := model.StrFormatTokenId(tokenTab)
	if err != nil {
		return "", err
	}

	arr, err := model.ReadARankValue(model.ERC20, tokenTab, 10)
	if err != nil {
		return "", err
	}

	jsonBytes, err := model.GetRankJson(arr)
	if err != nil {
		return "", err
	}

	count++
	log.Info("ARankTop access count: %d", count)
	return string(jsonBytes), nil
}

func AccountsRankTop(number int) (string, error) {
	//TODO:wait table
	accountTotalTab := "account_ar_test"

	//check number valid
	if number < 0 || number > 100 {
		return "", errors.New("invalid number err: bigger than 100 or less than 0")
	}

	arr, err := model.ReadARankValue(model.ERC20, accountTotalTab, number)
	if err != nil {
		return "", nil
	}

	jsonBytes, err := model.GetRankJson(arr)
	if err != nil {
		return "", nil
	}

	return string(jsonBytes), nil
}
