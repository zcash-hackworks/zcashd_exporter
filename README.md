# zcashd_exporter

The `zcashd_exporter` will poll a zcashd node's rpc endpoint, retreive data, then preent it for prometheus to scrape.

## Getting started

```
go get -v github.com/zcash-hackworks/zcashd_exporter
go install github.com/zcash-hackworks/zcashd_exporter
$GOPATH/bin/zcashd_exporter --help
```

## Current metrics

Example:
```
# HELP zcash_blocks the current number of blocks processed in the server
# TYPE zcash_blocks gauge
zcash_blocks 684096
# HELP zcash_chain_tip Return information about all known tips in the block tree, including the main chain as well as orphaned branches.
# TYPE zcash_chain_tip gauge
zcash_chain_tip{branchlen="0.00",hash="00264b6be8b4267e3369a6f22de3146953622a658ed49ccedaeacc12213f66a8",height="684096.00",status="active"} 684096
zcash_chain_tip{branchlen="0.00",hash="0147de0cbeb452f98d2abec0e89f0c834b36d9e94da67f7343110690704b60d7",height="684095.00",status="active"} 684095
zcash_chain_tip{branchlen="1.00",hash="0000627b037a646efa979a3fb4d526a1918ee5ae37cd72a68dd609765a772db7",height="682598.00",status="valid-fork"} 682598
zcash_chain_tip{branchlen="1.00",hash="001326c916c681ef5549915458e6f812946facb59aab2405ead8980104333426",height="679104.00",status="valid-fork"} 679104
zcash_chain_tip{branchlen="1.00",hash="00146600fae51bae06af151e54f0963a5ec7c98acd9efd51c8b8edd76dcd5ddc",height="655232.00",status="valid-fork"} 655232
zcash_chain_tip{branchlen="1.00",hash="00147f707ba0c457d65dff5b29d128ba7c8b56cb7e69b479571e0add5426ac5e",height="643815.00",status="valid-fork"} 643815
zcash_chain_tip{branchlen="1.00",hash="0016ca30a360df0329f40520b4202b38fbd9d278094c6ae4b8040e8211064dd2",height="647517.00",status="valid-fork"} 647517
zcash_chain_tip{branchlen="1.00",hash="00230588c53e1496913631164c7ab328a033d63414c2fea41812f6838fefd3e4",height="660514.00",status="valid-fork"} 660514
zcash_chain_tip{branchlen="1.00",hash="002696194bc1cc6833c141e6e2712312314d192562bcc1ce2460932448d314eb",height="651416.00",status="valid-fork"} 651416
zcash_chain_tip{branchlen="1.00",hash="0166c969a4ac17c346b0f9c185c1ba0b61bf27ae345048890a7fd20cb0caa254",height="672663.00",status="valid-fork"} 672663
# HELP zcash_difficulty the current difficulty
# TYPE zcash_difficulty gauge
zcash_difficulty 5.619307402921726
# HELP zcash_mempool_bytes Sum of all tx sizes
# TYPE zcash_mempool_bytes gauge
zcash_mempool_bytes 0
# HELP zcash_mempool_size Current tx count
# TYPE zcash_mempool_size gauge
zcash_mempool_size 0
# HELP zcash_mempool_usage Total memory usage for the mempool
# TYPE zcash_mempool_usage gauge
zcash_mempool_usage 0
# HELP zcash_peer_info Returns data about each connected network node.
# TYPE zcash_peer_info gauge
zcash_peer_info{addr="138.68.230.245:18233",addrlocal="70.126.192.75:40652",banscore="0",bytesrecv="137667154",bytessent="1797291",conntime="1573741133",id="9",inbound="false",lastrecv="1573744877",lastsend="1573744877",pingtime="0.08",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Mag
icBean:2.0.72/",synced_blocks="684095",synced_headers="684095",timeoffset="0",version="170008"} 9
zcash_peer_info{addr="138.68.230.245:18233",addrlocal="70.126.192.75:40652",banscore="0",bytesrecv="137667486",bytessent="1797477",conntime="1573741133",id="9",inbound="false",lastrecv="1573744976",lastsend="1573744982",pingtime="0.11",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Mag
icBean:2.0.72/",synced_blocks="684095",synced_headers="684095",timeoffset="0",version="170008"} 9
zcash_peer_info{addr="198.199.112.230:18233",addrlocal="70.126.192.75:49030",banscore="0",bytesrecv="169422685",bytessent="1952110",conntime="1573741109",id="1",inbound="false",lastrecv="1573744877",lastsend="1573744877",pingtime="0.09",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Ma
gicBean:2.1.01/",synced_blocks="684095",synced_headers="684095",timeoffset="0",version="170009"} 1
zcash_peer_info{addr="198.199.112.230:18233",addrlocal="70.126.192.75:49030",banscore="0",bytesrecv="169422810",bytessent="1952174",conntime="1573741109",id="1",inbound="false",lastrecv="1573744982",lastsend="1573744951",pingtime="0.09",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Ma
gicBean:2.1.01/",synced_blocks="684096",synced_headers="684096",timeoffset="0",version="170009"} 1
zcash_peer_info{addr="198.58.102.18:18233",addrlocal="70.126.192.75:43386",banscore="0",bytesrecv="137438601",bytessent="1793525",conntime="1573741125",id="7",inbound="false",lastrecv="1573744875",lastsend="1573744877",pingtime="0.05",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Magi
cBean:2.0.7/",synced_blocks="684095",synced_headers="684095",timeoffset="0",version="170008"} 7
zcash_peer_info{addr="198.58.102.18:18233",addrlocal="70.126.192.75:43386",banscore="0",bytesrecv="137443255",bytessent="1794703",conntime="1573741125",id="7",inbound="false",lastrecv="1573744982",lastsend="1573744982",pingtime="0.05",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Magi
cBean:2.0.7/",synced_blocks="684096",synced_headers="684096",timeoffset="0",version="170008"} 7
zcash_peer_info{addr="209.141.47.197:18233",addrlocal="70.126.192.75:47002",banscore="0",bytesrecv="137704215",bytessent="1806296",conntime="1573741111",id="4",inbound="false",lastrecv="1573744877",lastsend="1573744877",pingtime="0.08",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Mag
icBean:2.1.01/",synced_blocks="684095",synced_headers="684095",timeoffset="0",version="170009"} 4
zcash_peer_info{addr="209.141.47.197:18233",addrlocal="70.126.192.75:47002",banscore="0",bytesrecv="137704340",bytessent="1806360",conntime="1573741111",id="4",inbound="false",lastrecv="1573744982",lastsend="1573744953",pingtime="0.08",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Mag
icBean:2.1.01/",synced_blocks="684096",synced_headers="684096",timeoffset="0",version="170009"} 4
zcash_peer_info{addr="34.89.69.21:18233",addrlocal="70.126.192.75:48986",banscore="0",bytesrecv="137794030",bytessent="1834479",conntime="1573741110",id="2",inbound="false",lastrecv="1573744877",lastsend="1573744832",pingtime="0.13",pingwait="0",services="0000000000000005",startingheight="684016",subver="/MagicB
ean:2.0.73/",synced_blocks="684095",synced_headers="684095",timeoffset="0",version="170008"} 2
zcash_peer_info{addr="34.89.69.21:18233",addrlocal="70.126.192.75:48986",banscore="0",bytesrecv="137794155",bytessent="1834543",conntime="1573741110",id="2",inbound="false",lastrecv="1573744982",lastsend="1573744952",pingtime="0.13",pingwait="0",services="0000000000000005",startingheight="684016",subver="/MagicB
ean:2.0.73/",synced_blocks="684096",synced_headers="684096",timeoffset="0",version="170008"} 2
zcash_peer_info{addr="47.98.224.31:18233",addrlocal="70.126.192.75:49818",banscore="0",bytesrecv="137471753",bytessent="1801170",conntime="1573741132",id="8",inbound="false",lastrecv="1573744877",lastsend="1573744877",pingtime="0.79",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Magic
Bean:2.1.0/",synced_blocks="684095",synced_headers="684095",timeoffset="-1",version="170009"} 8
zcash_peer_info{addr="47.98.224.31:18233",addrlocal="70.126.192.75:49818",banscore="0",bytesrecv="137471878",bytessent="1801234",conntime="1573741132",id="8",inbound="false",lastrecv="1573744982",lastsend="1573744977",pingtime="0.29",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Magic
Bean:2.1.0/",synced_blocks="684096",synced_headers="684096",timeoffset="-1",version="170009"} 8
zcash_peer_info{addr="52.192.209.87:18233",addrlocal="70.126.192.75:45892",banscore="0",bytesrecv="137559407",bytessent="1797247",conntime="1573741111",id="3",inbound="false",lastrecv="1573744877",lastsend="1573744877",pingtime="0.23",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Magi
cBean:2.1.01/",synced_blocks="684095",synced_headers="684095",timeoffset="-1",version="170009"} 3
zcash_peer_info{addr="52.192.209.87:18233",addrlocal="70.126.192.75:45892",banscore="0",bytesrecv="137559532",bytessent="1797311",conntime="1573741111",id="3",inbound="false",lastrecv="1573744982",lastsend="1573744954",pingtime="0.21",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Magi
cBean:2.1.01/",synced_blocks="684096",synced_headers="684096",timeoffset="-1",version="170009"} 3
zcash_peer_info{addr="68.183.51.14:18233",addrlocal="70.126.192.75:44424",banscore="0",bytesrecv="138385996",bytessent="1810452",conntime="1573741147",id="10",inbound="false",lastrecv="1573744877",lastsend="1573744878",pingtime="0.05",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Magi
cBean:2.0.73/",synced_blocks="684095",synced_headers="684095",timeoffset="0",version="170008"} 10
zcash_peer_info{addr="68.183.51.14:18233",addrlocal="70.126.192.75:44424",banscore="0",bytesrecv="138386121",bytessent="1810516",conntime="1573741147",id="10",inbound="false",lastrecv="1573744988",lastsend="1573744988",pingtime="0.05",pingwait="0",services="0000000000000005",startingheight="684016",subver="/Magi
cBean:2.0.73/",synced_blocks="684096",synced_headers="684096",timeoffset="0",version="170008"} 10
# HELP zcash_size_on_disk the estimated size of the block and undo files on disk
# TYPE zcash_size_on_disk gauge
zcash_size_on_disk 0
# HELP zcash_verification_progress estimate of verification progress
# TYPE zcash_verification_progress gauge
zcash_verification_progress 0.9999996199585999
# HELP zcash_wallet_balance Node's wallet balance
# TYPE zcash_wallet_balance gauge
zcash_wallet_balance{type="private"} 8.9998
zcash_wallet_balance{type="total"} 13.99979214
zcash_wallet_balance{type="transparent"} 4.99999214
```
