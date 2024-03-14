FROM btwiuse/arch:golang AS builder-golang

COPY . /webteleport/ufo

WORKDIR /webteleport/ufo

ENV GONOSUMDB="*"

RUN go mod tidy

RUN CGO_ENABLED=0 GOBIN=/usr/local/bin go install -v ./cmd/ufo

FROM btwiuse/arch

COPY --from=builder-golang /usr/local/bin/ufo /usr/bin/ufo

ADD ./tls.crt /
ADD ./tls.key /
ENV CERT=/tls.crt
ENV KEY=/tls.key

ENTRYPOINT ["/usr/bin/ufo"]
