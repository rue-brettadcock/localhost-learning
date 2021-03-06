FROM golang

WORKDIR /go/src/github.com/rue-brettadcock/localhost-learning
COPY . .


RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["github.com/rue-brettadcock/localhost-learning"]
