FROM go-opencv:4.5.4

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
WORKDIR /usr/local
ADD . .
RUN go mod download && \
    GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app opencv-test
CMD ["./app"]
