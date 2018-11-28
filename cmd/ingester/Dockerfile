FROM golang:1.11

RUN go install github.com/h12w/msa/cmd/ingester

CMD ["/go/bin/ingester"]
EXPOSE 9001