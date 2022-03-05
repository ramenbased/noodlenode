package main

type GetRawTransaction_ struct {
	Result struct {
		InActiveChain bool   `json:"in_active_chain"`
		Txid          string `json:"txid"`
		Hash          string `json:"hash"`
		Version       int    `json:"version"`
		Size          int    `json:"size"`
		Vsize         int    `json:"vsize"`
		Weight        int    `json:"weight"`
		Locktime      int    `json:"locktime"`
		Vin           []struct {
			Txid      string `json:"txid"`
			Vout      int    `json:"vout"`
			ScriptSig struct {
				Asm string `json:"asm"`
				Hex string `json:"hex"`
			} `json:"scriptSig"`
			Sequence int64 `json:"sequence"`
		} `json:"vin"`
		Vout []struct {
			Value        float64 `json:"value"`
			N            int     `json:"n"`
			ScriptPubKey struct {
				Asm     string `json:"asm"`
				Hex     string `json:"hex"`
				Address string `json:"address"`
				Type    string `json:"type"`
			} `json:"scriptPubKey"`
		} `json:"vout"`
		Hex           string `json:"hex"`
		Blockhash     string `json:"blockhash"`
		Confirmations int    `json:"confirmations"`
		Time          int    `json:"time"`
		Blocktime     int    `json:"blocktime"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}
