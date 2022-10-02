FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o contest cmd/main.go

ENTRYPOINT [ "./contest" ]

##docker build . -t elecard-test-task 
##docker run --rm  elecard-test-task -m AutoExec