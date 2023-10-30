FROM golang:1.19-alpine AS build

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn
WORKDIR /data/cncamp
COPY . .
RUN set -x \
    && go build -o /bin/cncamp module2/main.go \
    && chmod +x /bin/cncamp

FROM alpine:latest
COPY --from=build /bin/cncamp /bin/cncamp
CMD ["/bin/cncamp"]
EXPOSE 80
