# zcashd_exporter

The `zcashd_exporter` will poll a zcashd node's rpc endpoint, retreive data, then preent it for prometheus to scrape.

## Getting started

```
go get -v github.com/zcash-hackworks/zcashd_exporter
go install github.com/zcash-hackworks/zcashd_exporter
$GOPATH/bin/zcashd_exporter --help
```

