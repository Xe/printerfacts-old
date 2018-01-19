FROM xena/go:1.9.2
COPY . /go/src/github.com/Xe/printerfacts
RUN apk --no-cache add alpine-sdk protobuf \
 && go get -u github.com/pkg/errors 
WORKDIR /go/src/github.com/Xe/printerfacts
RUN go run ./cmd/mage/main.go build

