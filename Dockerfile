FROM golang:1.14.1
WORKDIR /yam-api
COPY . .
RUN go install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o yam-api .
EXPOSE 80
CMD ["./yam-api"]