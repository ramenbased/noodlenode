package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

//-- Request Stuff

func jsnReq(method string, params interface{}) string {

	var jr map[string]interface{}
	jr = make(map[string]interface{})

	jr["jsonrpc"] = "1.0"
	jr["id"] = ""
	jr["method"] = method
	jr["params"] = params

	jq, _ := json.Marshal(jr)
	fmt.Println("REQUEST:", string(jq))
	return string(jq)
}

func apiReq(jr string, toStruct interface{}) {
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8332/", strings.NewReader(jr))
	req.SetBasicAuth("shnoodle", "78f238pf23z98f9qewfqwfwq89zf2898510970")
	req.Header.Add("content-type", "text/plain;")

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		fmt.Println(e)
	} else {

		defer res.Body.Close()
		api, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal([]byte(api), toStruct)

		//UNSTRUCTURED
		//var resp map[string]interface{}
		//json.NewDecoder(res.Body).Decode(&resp)
		//fmt.Println(resp, "\n", resp["error"])
	}
}

func apiReqRaw(jr string) {
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8332/", strings.NewReader(jr))
	req.SetBasicAuth("shnoodle", "78f238pf23z98f9qewfqwfwq89zf2898510970")
	req.Header.Add("content-type", "text/plain;")

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		fmt.Println(e)
	} else {

		defer res.Body.Close()
		api, _ := ioutil.ReadAll(res.Body)
		//GET RAW JSON
		fmt.Println(string(api))
	}
}

//-- RPCs

func call_GetBestBlockHash() {
	jr := jsnReq("getbestblockhash", []string{})
	apiReq(jr, &GetBestBlockHash)
}

func call_GetBlock(hash string) {
	jr := jsnReq("getblock", []string{hash})
	apiReq(jr, &GetBlock)
}

func call_GetBlockChainInfo() {
	jr := jsnReq("getblockchaininfo", []string{})
	apiReq(jr, &GetBlockChainInfo)
}

func call_GetBlockCount() {
	jr := jsnReq("getblockcount", []string{})
	apiReq(jr, &GetBlockCount)
}

func call_GetBlockHash(height int) {
	jr := jsnReq("getblockhash", []int{height})
	apiReq(jr, &GetBlockHash)
}

func call_GetBlockHeader(hash string) {
	jr := jsnReq("getblockheader", []string{hash})
	apiReq(jr, &GetBlockHeader)
}

func call_GetBlockStats(height int) {
	jr := jsnReq("getblockstats", []int{height})
	apiReq(jr, &GetBlockStats)
}

//To get Result: GetChainTips.Result[1].Hash
func call_GetChainTips() {
	jr := jsnReq("getchaintips", []string{})
	apiReq(jr, &GetChainTips)
}

func call_GetChainTxStats() {
	jr := jsnReq("getchaintxstats", []string{})
	apiReq(jr, &GetChainTxStats)
}

func call_GetDifficulty() {
	jr := jsnReq("getdifficulty", []string{})
	apiReq(jr, &GetDifficulty)
}

//WIP (need txid in mempool to get raw json)
func call_GetMemPoolAncestors(txid string) {
	jr := jsnReq("getmempoolancestors", []string{txid})
	apiReqRaw(jr)
}

//MISSING getmempooldescendants
//MISSING getmempoolentry

func call_GetMemPoolInfo() {
	jr := jsnReq("getmempoolinfo", []string{})
	apiReq(jr, &GetMempoolInfo)
}

//WIP tfw no mempool D:
func call_GetRawMempool() {
	jr := jsnReq("getrawmempool", []bool{true})
	apiReqRaw(jr)
}

//-- Postgres

func pingdb(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection to Database established!")
	}
}

//-- Div

func Er(err error) {
	fmt.Println(err)
}

func main() {
	connStr := "host=localhost user=postgres password=postgres port=5432 dbname=testdb"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		Er(err)
	}
	defer db.Close()
	pingdb(db)

	for i := 0; i <= 10; i++ {
		call_GetBlockStats(i)
		height := GetBlockStats.Result.Height
		hash := GetBlockStats.Result.Blockhash

		tx, err := db.Begin()
		if err != nil {
			Er(err)
		}
		defer tx.Rollback()
		stmt, err := tx.Prepare("INSERT INTO testtable VALUES ($1, $2)")
		if err != nil {
			Er(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(height, hash)
		if err != nil {
			Er(err)
		}
		err = tx.Commit()
		if err != nil {
			Er(err)
		}
	}

	//--- playout
	/*
		var (
			name  string
			based string
		)

		rows, _ := db.Query("select * from homies")
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&name, &based)
			if err != nil {
				Er(err)
			}
			fmt.Println(name, based)
		}
	*/
}
