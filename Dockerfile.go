FROM golang:1.12.0-alpine as builder

RUN mkdir /app && apk add --no-cache git

# copy all source files to app folder
ADD cmd/ /app

# make app the current directory all 
# subsequent commands will be executed here

WORKDIR /app

# create the binary executable 

RUN go install && go build -o main  

#create smaller finalized image with binary from lighter image
# using another base image i.e. alpine:latest  

FROM alpine:latest as production

# copy the app directory from our build image to our final image

COPY --from=builder /app . 

# run the binary

CMD ["./main"]
