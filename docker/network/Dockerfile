# pull image from registry
FROM golang:1.14.10

RUN go get go.mongodb.org/mongo-driver/mongo

# copy fili project to image directory
COPY main.go /app/main.go

# run the program 
CMD ["go", "run", "/app/main.go"]