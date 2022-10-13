FROM btwiuse/arch:golang

RUN git clone https://github.com/webteleport/ufo /webteleport/ufo

WORKDIR /webteleport/ufo

RUN go mod tidy

RUN CGO_ENABLED=0 GOBIN=/usr/local/bin go install ./cmd/ufo
