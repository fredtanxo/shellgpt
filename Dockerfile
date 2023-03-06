FROM golang
WORKDIR /shellgpt
COPY . .
RUN make build

FROM alpine
COPY --from=0 /shellgpt/bin/shellgpt ./shellgpt
ENTRYPOINT ["./shellgpt"]
