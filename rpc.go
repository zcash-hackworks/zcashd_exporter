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

// GetPeerInfo Returns data about each connected network node
// https://zcash-rpc.github.io/getpeerinfo.html
type GetPeerInfo []PeerInfo

type PeerInfo struct {
	ID             int     `json:"id"`
	Addr           string  `json:"addr"`
	AddrLocal      string  `json:"addrlocal"`
	Services       string  `json:"services"`
	LastSend       int     `json:"lastsend"`
	LastRecv       int     `json:"lastrecv"`
	BytesSent      int     `json:"bytessent"`
	BytesRecv      int     `json:"bytesrecv"`
	Conntime       int     `json:"conntime"`
	Timeoffset     int     `json:"timeoffset"`
	PingTime       float64 `json:"pingtime"`
	PingWait       int     `json:"pingwait"`
	Version        int     `json:"version"`
	Subver         string  `json:"subver"`
	Inbound        bool    `json:"inbound"`
	Startingheight int     `json:"startingheight"`
	Banscore       int     `json:"banscore"`
	SyncedHeaders  int     `json:"synced_headers"`
	SyncedBlocks   int     `json:"synced_blocks"`
}
