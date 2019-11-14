package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

//Define the metrics we wish to expose
var (
	zcashdBlocks = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "zcash_blocks", Help: "the current number of blocks processed in the server"})
	zcashdDifficulty = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "zcash_difficulty", Help: "the current difficulty"})
	zcashdSizeOnDisk = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "zcash_size_on_disk", Help: "the estimated size of the block and undo files on disk"})
	zcashdVerificationProgress = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "zcash_verification_progress", Help: "estimate of verification progress"})
	zcashdMemPoolSize = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "zcash_mempool_size", Help: "Current tx count"})
	zcashdMemPoolBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "zcash_mempool_bytes", Help: "Sum of all tx sizes"})
	zcashdMemPoolUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "zcash_mempool_usage", Help: "Total memory usage for the mempool"})
	zcashdWalletBalance = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zcash_wallet_balance",
			Help: "Node's wallet balance"},
		[]string{
			"type",
		})
	zcashdChainTips = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zcash_chain_tip",
			Help: "Return information about all known tips in the block tree, including the main chain as well as orphaned branches."},
		[]string{"height", "hash", "branchlen", "status"},
	)
	zcashdPeerInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zcash_peer_info",
			Help: "Returns data about each connected network node."},
		[]string{"id", "addr", "addrlocal", "services", "lastsend", "lastrecv", "bytessent", "bytesrecv", "conntime", "timeoffset", "pingtime", "pingwait", "version", "subver", "inbound", "startingheight", "banscore", "synced_headers", "synced_blocks"},
	)
)

// ZCASH_SOLS = Gauge("zcash_sols", "Estimated network solutions per second")

// ZCASH_ERRORS = Counter("zcash_errors", "Number of errors detected")

// ZCASH_LATEST_BLOCK_SIZE = Gauge("zcash_latest_block_size", "Size of latest block in bytes")
// ZCASH_LATEST_BLOCK_TXS = Gauge("zcash_latest_block_txs", "Number of transactions in latest block")

// ZCASH_TOTAL_BYTES_RECV = Gauge("zcash_total_bytes_recv", "Total bytes received")
// ZCASH_TOTAL_BYTES_SENT = Gauge("zcash_total_bytes_sent", "Total bytes sent")

// ZCASH_LATEST_BLOCK_INPUTS = Gauge("zcash_latest_block_inputs", "Number of inputs in transactions of latest block")
// ZCASH_LATEST_BLOCK_OUTPUTS = Gauge("zcash_latest_block_outputs", "Number of outputs in transactions of latest block")
// ZCASH_LATEST_BLOCK_JOINSPLITS = Gauge("zcash_latest_block_joinsplits", "Number of joinsplits in transactions of latest block")

// ZCASH_NUM_TRANSPARENT_TX = Gauge("zcash_num_transparent_tx", "Number of fully transparent transactions in latest block")
// ZCASH_NUM_SHIELDED_TX = Gauge("zcash_num_shielded_tx", "Number of fully shielded transactions in latest block")
// ZCASH_NUM_MIXED_TX = Gauge("zcash_num_mixed_tx", "Number of mixed transactions in latest block")

func init() {
	//Register metrics with prometheus
	prometheus.MustRegister(zcashdBlocks)
	prometheus.MustRegister(zcashdDifficulty)
	prometheus.MustRegister(zcashdSizeOnDisk)
	prometheus.MustRegister(zcashdVerificationProgress)
	prometheus.MustRegister(zcashdMemPoolSize)
	prometheus.MustRegister(zcashdMemPoolBytes)
	prometheus.MustRegister(zcashdMemPoolUsage)
	prometheus.MustRegister(zcashdWalletBalance)
	prometheus.MustRegister(zcashdChainTips)
	prometheus.MustRegister(zcashdPeerInfo)
}
