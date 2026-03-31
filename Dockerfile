FROM golang:alpine AS builder

WORKDIR /builder

COPY . .

RUN go mod download

RUN go build -o crm.parroto.com ./cmd/server

RUN apk add --no-cache tzdata

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV TZ=Asia/Ho_Chi_Minh

COPY ./config config

COPY --from=builder /builder/crm.parroto.com /

ENTRYPOINT [ "/crm.parroto.com",  "config/local.yaml"]