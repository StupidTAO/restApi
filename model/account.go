package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/test/restApi/log"
	"io"
	"sync"
)

const (
	ETH = "eth"
	ERC20 = "erc20"
//
//	AOA = "AOA_ar_3"
//	BAT = "BAT_ar_3"
//	BNB = "BNB_ar_3"
//	CRO = "CRO_ar_3"
//	DAI = "DAI_ar_3"
//	ENJ = "ENJ_ar_3"
//	GNT = "GNT_ar_3"
//	HOT = "HOT_ar_3"
//	HT = "HT_ar_3"
//	INB = "INB_ar_3"
//	KCS = "KCS_ar_3"
//	LAMB = "LAMB_ar_3"
//	LEO = "LEO_ar_3"
//	LINK = "LINK_ar_3"
//	MCO = "MCO_ar_3"
//	MKR = "MKR_ar_3"
//	NEXO = "NEXO_ar_3"
//	NPXS = "NPXS_ar_3"
//	PAX = "PAX_ar_3"
//	QNT = "QNT_ar_3"
//	REP = "REP_ar_3"
//	SNT = "SNT_ar_3"
//	SNX = "SNX_ar_3"
//	TUSD = "TUSD_ar_3"
//	USDC = "USDC_ar_3"
//	USDT = "USDT_ar_3"
//	ZB = "ZB_ar_3"
//	ZRX = "ZRX_ar_3"
//	//...
)

var mu sync.Mutex
var count int

//获取AR榜单
func ARankTop3(out io.Writer, tokenName string) {

	comTokenName := fmt.Sprintf("%s_ar_3", tokenName)
	_, err := StrFormatTokenId(comTokenName)
	if err != nil {
		out.Write([]byte(err.Error()))
		return
	}

	arr, err := ReadARankValue(ERC20, comTokenName,10)
	if err != nil {
		out.Write([]byte(err.Error()))
		return
	}

	name := "Accounts rank of %s token"
	titileName := fmt.Sprintf(name, tokenName)
	page, _:= getRankPage(arr, titileName)
	//jsonbytes, _ := GetRankJson(arr)
	out.Write([]byte(page))

	//counter atuo add
	mu.Lock()
	count++
	mu.Unlock()
}

func AccountsRankTop(out io.Writer, number int) {
	//TODO: wait table name
	totalTab := "account_ar_test"
	arr, _ := ReadARankValue(ERC20, totalTab, number)

	pageName := "Accounts rank"
	accountsPage, _ := getRankPage(arr, pageName)
	out.Write([]byte(accountsPage))
}

//从ArangoDB数据库读取AR榜单
func ReadARankValue(database string, table string, number int) ([]MyDocument, error) {
	//read top number

	//dbUrl := "http://10.10.10.10:8529"
	client, err := getDBClient()
	if err != nil {
		return nil, err
	}

	//
	ctx := context.Background()
	//db, err := client.Database(ctx, "erc20")
	db, err := client.Database(ctx, database)
	if err != nil {
		return nil, err
	}

	//query top number data
	ctx1 := driver.WithQueryCount(context.Background())
	query_date := "for tx in %s sort tx.data.date desc collect aggregate max_date=MAX(tx.data.date) return {datetime:max_date}"
	query_datetime := fmt.Sprintf(query_date, table)
	cursor_time, err := db.Query(ctx1, query_datetime, nil)
	if err != nil {
		//handle error
		log.Warning("query db occur error, err: %s", err.Error())
		return nil, err
	}
	//defer cursor_time.Close()
	datetime := ""
	for {
		var result interface{}
		_, err := cursor_time.ReadDocument(ctx, &result)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			//handle other errors
		}
		datetime = result.(map[string]interface{})["datetime"].(string)
	}

	query := "for tx in %s filter tx.data.date==\"%s\" sort tx.data.score desc limit %d return {name:tx.data.address,score:tx.data.score,median:tx.data.median,weight:tx.data.weight,datetime:tx.data.date}"
	query_str := fmt.Sprintf(query, table, datetime, number)
	cursor, err := db.Query(ctx1, query_str, nil)
	if err != nil {
		// handle error
		log.Warning("query db occur error, err: %s", err.Error())
		return nil, err
	}
	defer cursor.Close()

	var docs []MyDocument
	num := 0
	for {
		num++
		var doc MyDocument
		var result interface{}
		_, err := cursor.ReadDocument(ctx, &result)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}
		//doc := result.(MyDocument)
		doc.Name = result.(map[string]interface {})["name"].(string)
		doc.Score = result.(map[string]interface{})["score"].(float64)
		doc.Datetime = result.(map[string]interface{})["datetime"].(string)
		doc.Median = result.(map[string]interface{})["median"].(float64)
		doc.Weight = result.(map[string]interface{})["weight"].(float64)
		doc.Number = int32(num)


		//doc.number = result.(map[string]interface{})["number"].(int32)
		docs = append(docs, doc)
	}

	return docs, nil
}

func getRankPage(form []MyDocument, pageTitle string) (string, error) {
	head := "<h4>Title：%s</h4> <table border=\"1\"><tr><th>Number</th><th>Account</th><th>Score</th><th>Median</th><th>Weight</th><th>Datetime</th></tr>"
	table_head := fmt.Sprintf(head, pageTitle)
	lines := ""
	for i:=0; i<len(form); i++ {
		line := fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%f</td><td>%f</td><td>%f</td><td>%s</td></tr>", form[i].Number, form[i].Name, form[i].Score, form[i].Median, form[i].Weight, form[i].Datetime)
		lines += line
	}
	table_page := "<html>" + table_head + lines + "</table> </html>"

	return table_page, nil
}

func GetRankJson(arr []MyDocument) ([]byte, error) {
	if len(arr) == 0 {
		return []byte{}, errors.New("result is none")
	}

	result, err := json.Marshal(arr)

	return result, err
}