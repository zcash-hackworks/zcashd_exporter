package main

// GetBlockchainInfo return the zcashd rpc `getblockchaininfo` status
// https://zcash-rpc.github.io/getblockchaininfo.html
type GetBlockchainInfo struct {
	Chain                string  `json:"chain"`
	Blocks               int     `json:"blocks"`
	Difficulty           float64 `json:"difficulty"`
	VerificationProgress float64 `json:"verificationprogress"`
	SizeOnDisk           float64 `json:"size_on_disk"`
}

// GetMemPoolInfo return the zcashd rpc `getmempoolinfo`
// https://zcash-rpc.github.io/getmempoolinfo.html
type GetMemPoolInfo struct {
	Size  float64 `json:"size"`
	Bytes float64 `json:"bytes"`
	Usage float64 `json:"usage"`
}

// ZGetTotalBalance return the node's wallet balances
// https://zcash-rpc.github.io/z_gettotalbalance.html
type ZGetTotalBalance struct {
	Transparent string `json:"transparent"`
	Private     string `json:"private"`
	Total       string `json:"total"`
}

// GetChainTips Return information about all known tips in the block tree, including the main chain as well as orphaned branches.
// https://zcash-rpc.github.io/getchaintips.html
type GetChainTips []ChainTip

type ChainTip struct {
	Height    float64 `json:"height"`
	Hash      string  `json:"hash"`
	Branchlen float64 `json:"branchlen"`
	Status    string  `json:"status"`
}
