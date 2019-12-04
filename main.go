package main

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ybbus/jsonrpc"
	"github.com/zcash-hackworks/zcashd_exporter/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address on which to expose metrics and web interface.",
	).Default(":9100").String()
	rpcHost = kingpin.Flag(
		"rpc.host",
		"Host address for RPC endpoint.",
	).Default("127.0.0.1").String()
	rpcPort = kingpin.Flag(
		"rpc.port",
		"Post for RPC endpoint",
	).Default("18232").String()
	rpcUser = kingpin.Flag(
		"rpc.user",
		"User for RPC endpoint auth.",
	).Default("zcashrpc").String()
	rpcPassword = kingpin.Flag(
		"rpc.password",
		"Password for RPC endpoint auth.",
	).Default("notsecure").String()
)

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	log.Infoln("Starting zcashd_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>Zcashd Exporter</title></head>
		<body>
		<h1>Zcashd Exporter</h1>
		<p><a href="/metrics">Metrics</a></p>
		</body>
		</html>`))
	})
	go getBlockchainInfo()
	go getMemPoolInfo()
	go getWalletInfo()
	go getPeerInfo()
	log.Infoln("Listening on", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}

}

func getBlockchainInfo() {
	basicAuth := base64.StdEncoding.EncodeToString([]byte(*rpcUser + ":" + *rpcPassword))
	rpcClient := jsonrpc.NewClientWithOpts("http://"+*rpcHost+":"+*rpcPort,
		&jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"Authorization": "Basic " + basicAuth,
			}})
	var blockinfo *GetBlockchainInfo

	for {
		if err := rpcClient.CallFor(&blockinfo, "getblockchaininfo"); err != nil {
			log.Warnln("Error calling getblockchaininfo", err)
		} else {
			zcashdBlocks.Set(float64(blockinfo.Blocks))
			zcashdDifficulty.Set(blockinfo.Difficulty)
			zcashdVerificationProgress.Set(blockinfo.VerificationProgress)
		}
		time.Sleep(time.Duration(30) * time.Second)
	}

}

func getMemPoolInfo() {
	basicAuth := base64.StdEncoding.EncodeToString([]byte(*rpcUser + ":" + *rpcPassword))
	rpcClient := jsonrpc.NewClientWithOpts("http://"+*rpcHost+":"+*rpcPort,
		&jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"Authorization": "Basic " + basicAuth,
			}})
	var mempoolinfo *GetMemPoolInfo

	for {
		if err := rpcClient.CallFor(&mempoolinfo, "getmempoolinfo"); err != nil {
			log.Warnln("Error calling getmempoolinfo", err)
		} else {
			zcashdMemPoolSize.Set(float64(mempoolinfo.Size))
			zcashdMemPoolBytes.Set(mempoolinfo.Bytes)
			zcashdMemPoolUsage.Set(mempoolinfo.Usage)
		}
		time.Sleep(time.Duration(30) * time.Second)
	}

}

func getWalletInfo() {
	basicAuth := base64.StdEncoding.EncodeToString([]byte(*rpcUser + ":" + *rpcPassword))
	rpcClient := jsonrpc.NewClientWithOpts("http://"+*rpcHost+":"+*rpcPort,
		&jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"Authorization": "Basic " + basicAuth,
			}})
	var walletinfo *ZGetTotalBalance

	for {
		if err := rpcClient.CallFor(&walletinfo, "z_gettotalbalance"); err != nil {
			log.Warnln("Error calling z_gettotalbalance", err)
		} else {
			if t, err := strconv.ParseFloat(walletinfo.Transparent, 64); err == nil {
				zcashdWalletBalance.WithLabelValues("transparent").Set(t)
			}
			if p, err := strconv.ParseFloat(walletinfo.Private, 64); err == nil {
				zcashdWalletBalance.WithLabelValues("private").Set(p)
			}
			if total, err := strconv.ParseFloat(walletinfo.Total, 64); err == nil {
				zcashdWalletBalance.WithLabelValues("total").Set(total)
			}
		}
		time.Sleep(time.Duration(30) * time.Second)
	}

}

func getPeerInfo() {
	basicAuth := base64.StdEncoding.EncodeToString([]byte(*rpcUser + ":" + *rpcPassword))
	rpcClient := jsonrpc.NewClientWithOpts("http://"+*rpcHost+":"+*rpcPort,
		&jsonrpc.RPCClientOpts{
			CustomHeaders: map[string]string{
				"Authorization": "Basic " + basicAuth,
			}})
	var peerinfo *GetPeerInfo

	for {
		if err := rpcClient.CallFor(&peerinfo, "getpeerinfo"); err != nil {
			log.Warnln("Error calling getchaintips", err)
		} else {
			for _, pi := range *peerinfo {
				log.Infoln("Got peerinfo: ", pi.Addr)
				zcashdPeerVerion.WithLabelValues(
					pi.Addr,
					pi.AddrLocal,
					strconv.FormatBool(pi.Inbound),
					strconv.Itoa(pi.Banscore),
					pi.Subver,
				).Set(float64(pi.Version))
				zcashdPeerConnTime.WithLabelValues(
					pi.Addr,
					pi.AddrLocal,
					strconv.FormatBool(pi.Inbound),
					strconv.Itoa(pi.Banscore),
					pi.Subver,
				).Set(float64(pi.Conntime))
				zcashdPeerBytesSent.WithLabelValues(
					pi.Addr,
					pi.AddrLocal,
					strconv.FormatBool(pi.Inbound),
					strconv.Itoa(pi.Banscore),
					pi.Subver,
				).Set(float64(pi.BytesSent))
				zcashdPeerBytesRecv.WithLabelValues(
					pi.Addr,
					pi.AddrLocal,
					strconv.FormatBool(pi.Inbound),
					strconv.Itoa(pi.Banscore),
					pi.Subver,
				).Set(float64(pi.BytesRecv))
				// zcashdPeerInfo.WithLabelValues(
				// 	strconv.Itoa(pi.ID),
				// 	pi.Addr,
				// 	pi.AddrLocal,
				// 	pi.Services,
				// 	strconv.Itoa(pi.LastSend),
				// 	strconv.Itoa(pi.LastRecv),
				// 	strconv.Itoa(pi.BytesSent),
				// 	strconv.Itoa(pi.BytesRecv),
				// 	strconv.Itoa(pi.Conntime),
				// 	strconv.Itoa(pi.Timeoffset),
				// 	strconv.FormatFloat(pi.PingTime, 'f', 2, 64),
				// 	strconv.Itoa(pi.PingWait),
				// 	strconv.Itoa(pi.Version),
				// 	pi.Subver,
				// 	strconv.FormatBool(pi.Inbound),
				// 	strconv.Itoa(pi.Startingheight),
				// 	strconv.Itoa(pi.Banscore),
				// 	strconv.Itoa(pi.SyncedHeaders),
				// 	strconv.Itoa(pi.SyncedBlocks),
				// ).Set(float64(pi.ID))
			}
		}
		time.Sleep(time.Duration(30) * time.Second)
	}

}
