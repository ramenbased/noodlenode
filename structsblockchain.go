package main

type GetBestBlockHash_ struct {
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

type GetBlock_ struct {
	Result struct {
		Hash              string   `json:"hash"`
		Confirmations     int      `json:"confirmations"`
		Height            int      `json:"height"`
		Version           int      `json:"version"`
		VersionHex        string   `json:"versionHex"`
		Merkleroot        string   `json:"merkleroot"`
		Time              int      `json:"time"`
		Mediantime        int      `json:"mediantime"`
		Nonce             int      `json:"nonce"`
		Bits              string   `json:"bits"`
		Difficulty        int      `json:"difficulty"`
		Chainwork         string   `json:"chainwork"`
		NTx               int      `json:"nTx"`
		Previousblockhash string   `json:"previousblockhash"`
		Nextblockhash     string   `json:"nextblockhash"`
		Strippedsize      int      `json:"strippedsize"`
		Size              int      `json:"size"`
		Weight            int      `json:"weight"`
		Tx                []string `json:"tx"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

type GetBlockCount_ struct {
	Result int         `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

type GetBlockHash_ struct {
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

type GetBlockChainInfo_ struct {
	Result struct {
		Chain                string  `json:"chain"`
		Blocks               int     `json:"blocks"`
		Headers              int     `json:"headers"`
		Bestblockhash        string  `json:"bestblockhash"`
		Difficulty           float64 `json:"difficulty"`
		Mediantime           int     `json:"mediantime"`
		Verificationprogress float64 `json:"verificationprogress"`
		Initialblockdownload bool    `json:"initialblockdownload"`
		Chainwork            string  `json:"chainwork"`
		SizeOnDisk           int64   `json:"size_on_disk"`
		Pruned               bool    `json:"pruned"`
		Softforks            struct {
			Bip34 struct {
				Type   string `json:"type"`
				Active bool   `json:"active"`
				Height int    `json:"height"`
			} `json:"bip34"`
			Bip66 struct {
				Type   string `json:"type"`
				Active bool   `json:"active"`
				Height int    `json:"height"`
			} `json:"bip66"`
			Bip65 struct {
				Type   string `json:"type"`
				Active bool   `json:"active"`
				Height int    `json:"height"`
			} `json:"bip65"`
			Csv struct {
				Type   string `json:"type"`
				Active bool   `json:"active"`
				Height int    `json:"height"`
			} `json:"csv"`
			Segwit struct {
				Type   string `json:"type"`
				Active bool   `json:"active"`
				Height int    `json:"height"`
			} `json:"segwit"`
			Taproot struct {
				Type string `json:"type"`
				Bip9 struct {
					Status              string `json:"status"`
					StartTime           int    `json:"start_time"`
					Timeout             int    `json:"timeout"`
					Since               int    `json:"since"`
					MinActivationHeight int    `json:"min_activation_height"`
				} `json:"bip9"`
				Active bool `json:"active"`
			} `json:"taproot"`
		} `json:"softforks"`
		Warnings string `json:"warnings"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

type GetBlockHeader_ struct {
	Result struct {
		Hash              string  `json:"hash"`
		Confirmations     int     `json:"confirmations"`
		Height            int     `json:"height"`
		Version           int     `json:"version"`
		VersionHex        string  `json:"versionHex"`
		Merkleroot        string  `json:"merkleroot"`
		Time              int     `json:"time"`
		Mediantime        int     `json:"mediantime"`
		Nonce             int64   `json:"nonce"`
		Bits              string  `json:"bits"`
		Difficulty        float64 `json:"difficulty"`
		Chainwork         string  `json:"chainwork"`
		NTx               int     `json:"nTx"`
		Previousblockhash string  `json:"previousblockhash"`
		Nextblockhash     string  `json:"nextblockhash"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

type GetBlockStats_ struct {
	Result struct {
		Avgfee             int    `json:"avgfee"`
		Avgfeerate         int    `json:"avgfeerate"`
		Avgtxsize          int    `json:"avgtxsize"`
		Blockhash          string `json:"blockhash"`
		FeeratePercentiles []int  `json:"feerate_percentiles"`
		Height             int    `json:"height"`
		Ins                int    `json:"ins"`
		Maxfee             int    `json:"maxfee"`
		Maxfeerate         int    `json:"maxfeerate"`
		Maxtxsize          int    `json:"maxtxsize"`
		Medianfee          int    `json:"medianfee"`
		Mediantime         int    `json:"mediantime"`
		Mediantxsize       int    `json:"mediantxsize"`
		Minfee             int    `json:"minfee"`
		Minfeerate         int    `json:"minfeerate"`
		Mintxsize          int    `json:"mintxsize"`
		Outs               int    `json:"outs"`
		Subsidy            int64  `json:"subsidy"`
		SwtotalSize        int    `json:"swtotal_size"`
		SwtotalWeight      int    `json:"swtotal_weight"`
		Swtxs              int    `json:"swtxs"`
		Time               int    `json:"time"`
		TotalOut           int64  `json:"total_out"`
		TotalSize          int    `json:"total_size"`
		TotalWeight        int    `json:"total_weight"`
		Totalfee           int    `json:"totalfee"`
		Txs                int    `json:"txs"`
		UtxoIncrease       int    `json:"utxo_increase"`
		UtxoSizeInc        int    `json:"utxo_size_inc"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

type GetChainTips_ struct {
	Result []struct {
		Height    int    `json:"height"`
		Hash      string `json:"hash"`
		Branchlen int    `json:"branchlen"`
		Status    string `json:"status"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

type GetChainTxStats_ struct {
	Result struct {
		Time                   int     `json:"time"`
		Txcount                int     `json:"txcount"`
		WindowFinalBlockHash   string  `json:"window_final_block_hash"`
		WindowFinalBlockHeight int     `json:"window_final_block_height"`
		WindowBlockCount       int     `json:"window_block_count"`
		WindowTxCount          int     `json:"window_tx_count"`
		WindowInterval         int     `json:"window_interval"`
		Txrate                 float64 `json:"txrate"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

type GetDifficulty_ struct {
	Result float64     `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

type GetMempoolAncestors_ struct {
	Result []string    `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

type GetMempoolDescendants_ struct {
	Result []string    `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}
type GetRawMempoolEntry_ struct {
	Result struct {
		Fees struct {
			Base       float64 `json:"base"`
			Modified   float64 `json:"modified"`
			Ancestor   float64 `json:"ancestor"`
			Descendant float64 `json:"descendant"`
		} `json:"fees"`
		Vsize             int           `json:"vsize"`
		Weight            int           `json:"weight"`
		Fee               float64       `json:"fee"`
		Modifiedfee       float64       `json:"modifiedfee"`
		Time              int           `json:"time"`
		Height            int           `json:"height"`
		Descendantcount   int           `json:"descendantcount"`
		Descendantsize    int           `json:"descendantsize"`
		Descendantfees    int           `json:"descendantfees"`
		Ancestorcount     int           `json:"ancestorcount"`
		Ancestorsize      int           `json:"ancestorsize"`
		Ancestorfees      int           `json:"ancestorfees"`
		Wtxid             string        `json:"wtxid"`
		Depends           []interface{} `json:"depends"`
		Spentby           []interface{} `json:"spentby"`
		Bip125Replaceable bool          `json:"bip125-replaceable"`
		Unbroadcast       bool          `json:"unbroadcast"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

type GetMempoolInfo_ struct {
	Result struct {
		Loaded           bool    `json:"loaded"`
		Size             int     `json:"size"`
		Bytes            int     `json:"bytes"`
		Usage            int     `json:"usage"`
		TotalFee         float64 `json:"total_fee"`
		Maxmempool       int     `json:"maxmempool"`
		Mempoolminfee    float64 `json:"mempoolminfee"`
		Minrelaytxfee    float64 `json:"minrelaytxfee"`
		Unbroadcastcount int     `json:"unbroadcastcount"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

type GetRawMempool_ struct {
	Result []string    `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}
