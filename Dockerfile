FROM golang:1.13

ADD . /go/src/github.com/zcash-hackworks/zcashd_exporter

RUN go get github.com/zcash-hackworks/zcashd_exporter
RUN go install github.com/zcash-hackworks/zcashd_exporter

ENTRYPOINT ["zcashd_exporter"]
CMD ["--help"]