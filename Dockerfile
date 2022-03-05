FROM csnight/opencv:4.5.4

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV PATH=/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
WORKDIR /usr/local
ADD . .
RUN apt-get update && apt-get install -y --no-install-recommends curl && \
    curl -Lo go.tar.gz https://studygolang.com/dl/golang/go1.17.8.linux-amd64.tar.gz && \
    tar -xzvf go.tar.gz && \
    rm -rf go.tar.gz && \
    go mod download && \
    GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app opencv-test && \
    rm -rf /usr/local/go/ && \
    apt-get autoremove -y curl && apt-get clean && apt-get autoclean
CMD ["./app"]
