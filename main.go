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

	jq, err := json.Marshal(jr)
	Er(err)
	fmt.Println("REQUEST:", string(jq))
	return string(jq)
}

func apiReq(jr string, toStruct interface{}) {
	req, _ := http.NewRequest("POST", "http://127.0.0.1:8332/", strings.NewReader(jr))
	req.SetBasicAuth("shnoodle", "78f238pf23z98f9qewfqwfwq89zf2898510970")
	req.Header.Add("content-type", "text/plain;")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
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

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		defer res.Body.Close()
		api, _ := ioutil.ReadAll(res.Body)
		//GET RAW JSON
		fmt.Println(string(api))
	}
}

//-- RPC - BlockChain

func blockchain_GetBestBlockHash() {
	jr := jsnReq("getbestblockhash", []string{})
	apiReq(jr, &GetBestBlockHash)
}

func blockchain_GetBlock(hash string) {
	jr := jsnReq("getblock", []interface{}{hash, 1})
	apiReq(jr, &GetBlock)
}

func blockchain_GetBlockChainInfo() {
	jr := jsnReq("getblockchaininfo", []string{})
	apiReq(jr, &GetBlockChainInfo)
}

func blockchain_GetBlockCount() {
	jr := jsnReq("getblockcount", []string{})
	apiReq(jr, &GetBlockCount)
}

func blockchain_GetBlockHash(height int) {
	jr := jsnReq("getblockhash", []int{height})
	apiReq(jr, &GetBlockHash)
}

func blockchain_GetBlockHeader(hash string) {
	jr := jsnReq("getblockheader", []string{hash})
	apiReq(jr, &GetBlockHeader)
}

func blockchain_GetBlockStats(height int) {
	jr := jsnReq("getblockstats", []int{height})
	apiReq(jr, &GetBlockStats)
}

func blockchain_GetChainTips() {
	jr := jsnReq("getchaintips", []string{})
	apiReq(jr, &GetChainTips)
}

func blockchain_GetChainTxStats() {
	jr := jsnReq("getchaintxstats", []string{})
	apiReq(jr, &GetChainTxStats)
}

func blockchain_GetDifficulty() {
	jr := jsnReq("getdifficulty", []string{})
	apiReq(jr, &GetDifficulty)
}

func blockchain_GetMempoolAncestors(txid string) {
	jr := jsnReq("getmempoolancestors", []interface{}{txid, false})
	apiReq(jr, &GetMempoolAncestors)
}

func blockchain_GetMempoolDescendants(txid string) {
	jr := jsnReq("getmempooldescendants", []interface{}{txid, false})
	apiReq(jr, &GetMempoolDescendants)
}

func blockchain_GetMempoolEntry(txid string) {
	jr := jsnReq("getmempoolentry", []string{txid})
	apiReq(jr, &GetRawMempoolEntry)
}

func blockchain_GetMempoolInfo() {
	jr := jsnReq("getmempoolinfo", []string{})
	apiReq(jr, &GetMempoolInfo)
}

func blockchain_GetRawMempool() {
	jr := jsnReq("getrawmempool", []bool{false})
	apiReq(jr, &GetRawMempool)
}

func blockchain_GetTxOutsetInfo() {
	jr := jsnReq("gettxoutsetinfo", nil)
	apiReqRaw(jr)
}

func blockchain_GetTxOut(txid string, n int) {
	jr := jsnReq("gettxout", []interface{}{txid, n})
	apiReqRaw(jr)
}

//-- RPC - Raw Transactions

func getrawtxs_GetRawTransaction(txid string, verbose bool, blockhash string) {
	jr := jsnReq("getrawtransaction", []interface{}{txid, verbose, blockhash})
	apiReq(jr, &GetRawTransaction)
}

//-- RPC - Control

func control_GetRPCInfo() {
	jr := jsnReq("getrpcinfo", nil)
	apiReqRaw(jr)
}

//-- Postgres

func pingdb(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to DB! Connections open:", db.Stats().OpenConnections)
	}
}

//-- Div

func Er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	connStr := "host=localhost user=postgres password=postgres port=5432 dbname=noodledb"
	db, err := sql.Open("postgres", connStr)
	Er(err)
	defer db.Close()
	pingdb(db)

	/*
		for i := 1; i <= 3; i++ {
			blockchain_GetBlockStats(i)
			gbs := GetBlockStats.Result

			tx, err := db.Begin()
			Er(err)
			defer tx.Rollback()
			stmt, err := tx.Prepare("INSERT INTO testtable VALUES ($1)")
			Er(err)
			defer stmt.Close()
			_, err = stmt.Exec(
				gbs.Height,
			)
			Er(err)
			err = tx.Commit()
			Er(err)
		}
	*/
	//--- playout
	/*
		var (
			height    int
			blockhash string
			ins       int
			outs      int
		)

		rows, _ := db.Query("select * from testtable")
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&height, &blockhash, &ins, &outs)
			Er(err)
			fmt.Println(height, blockhash, ins, outs)
		}
	*/
}
