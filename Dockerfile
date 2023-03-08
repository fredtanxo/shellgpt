FROM golang AS builder
WORKDIR /shellgpt
COPY . .
RUN make

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /shellgpt/bin/shellgpt ./shellgpt
ENTRYPOINT ["./shellgpt"]
