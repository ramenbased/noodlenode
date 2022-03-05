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
	//fmt.Println("REQUEST:", string(jq))
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

func GetBestBlockHash() *GetBestBlockHash_ {
	jr := jsnReq("getbestblockhash", []string{})
	var rv GetBestBlockHash_
	apiReq(jr, &rv)
	return &rv
}

func GetBlock(hash string) *GetBlock_ {
	jr := jsnReq("getblock", []interface{}{hash, 1})
	var rv GetBlock_
	apiReq(jr, &rv)
	return &rv
}

func GetBlockChainInfo() *GetBlockChainInfo_ {
	jr := jsnReq("getblockchaininfo", []string{})
	var rv GetBlockChainInfo_
	apiReq(jr, &rv)
	return &rv
}

func GetBlockCount() *GetBlockCount_ {
	jr := jsnReq("getblockcount", []string{})
	var rv GetBlockCount_
	apiReq(jr, &rv)
	return &rv
}

func GetBlockHash(height int) *GetBlockHash_ {
	jr := jsnReq("getblockhash", []int{height})
	var rv GetBlockHash_
	apiReq(jr, &rv)
	return &rv
}

func GetBlockHeader(hash string) *GetBlockHeader_ {
	jr := jsnReq("getblockheader", []string{hash})
	var rv GetBlockHeader_
	apiReq(jr, &rv)
	return &rv
}

func GetBlockStats(height int) *GetBlockStats_ {
	jr := jsnReq("getblockstats", []int{height})
	var rv GetBlockStats_
	apiReq(jr, &rv)
	return &rv
}

func GetChainTips() *GetChainTips_ {
	jr := jsnReq("getchaintips", []string{})
	var rv GetChainTips_
	apiReq(jr, &rv)
	return &rv
}

func GetChainTxStats() *GetChainTxStats_ {
	jr := jsnReq("getchaintxstats", []string{})
	var rv GetChainTxStats_
	apiReq(jr, &rv)
	return &rv
}

func GetDifficulty() *GetDifficulty_ {
	jr := jsnReq("getdifficulty", []string{})
	var rv GetDifficulty_
	apiReq(jr, &rv)
	return &rv
}

func GetMempoolAncestors(txid string) *GetMempoolAncestors_ {
	jr := jsnReq("getmempoolancestors", []interface{}{txid, false})
	var rv GetMempoolAncestors_
	apiReq(jr, &rv)
	return &rv
}

func GetMempoolDescendants(txid string) *GetMempoolDescendants_ {
	jr := jsnReq("getmempooldescendants", []interface{}{txid, false})
	var rv GetMempoolDescendants_
	apiReq(jr, &rv)
	return &rv
}

func GetMempoolEntry(txid string) *GetRawMempoolEntry_ {
	jr := jsnReq("getmempoolentry", []string{txid})
	var rv GetRawMempoolEntry_
	apiReq(jr, &rv)
	return &rv
}

func GetMempoolInfo() *GetMempoolInfo_ {
	jr := jsnReq("getmempoolinfo", []string{})
	var rv GetMempoolInfo_
	apiReq(jr, &rv)
	return &rv
}

func GetRawMempool() *GetRawMempool_ {
	jr := jsnReq("getrawmempool", []bool{false})
	var rv GetRawMempool_
	apiReq(jr, &rv)
	return &rv
}

//WIP
func GetTxOutsetInfo() {
	jr := jsnReq("gettxoutsetinfo", nil)
	apiReqRaw(jr)
}

//WIP
func GetTxOut(txid string, n int) {
	jr := jsnReq("gettxout", []interface{}{txid, n})
	apiReqRaw(jr)
}

//-- RPC - Raw Transactions

func getrawtxs_GetRawTransaction(txid string, verbose bool, blockhash string) *GetRawTransaction_ {
	jr := jsnReq("getrawtransaction", []interface{}{txid, verbose, blockhash})
	var rv GetRawTransaction_
	apiReq(jr, &rv)
	return &rv
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

	best := GetBestBlockHash()
	block := GetBlock(best.Result)
	height := block.Result.Height

	for b := 1; b <= height; b++ {
		stats := GetBlockStats(b)
		fmt.Println(stats.Result.Height, ",", stats.Result.Txs)
	}

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
