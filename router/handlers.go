package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/test/restApi/api"
	"net/http"
	"github.com/test/restApi/model"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{name: "Write presentation"},
		Todo{name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show: ", todoId)
}

func TokenAccountsRank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tokenName := vars["tokenName"]
	model.ARankTop3(w, tokenName)
}

func ApiTokenAccountsRank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tokenName := vars["tokenName"]

	res, err := api.ARankTop(tokenName)
	if err != nil {
		w.Write([]byte("get api rank top error"))
	}

	w.Write([]byte(res))
}

func ApiAccountsRank(w http.ResponseWriter, r *http.Request) {
	//wait data
	vars := mux.Vars(r)
	number := vars["number"]

	num, _ := strconv.Atoi(number)
	res, err := api.AccountsRankTop(num)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(res))
}

func AccountsRank(w http.ResponseWriter, r *http.Request)  {
	//wait data
	vars := mux.Vars(r)
	number := vars["number"]

	num, _ := strconv.Atoi(number)
	model.AccountsRankTop(w, num)

}