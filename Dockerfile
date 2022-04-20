FROM golang:1.18-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc libc-dev

# Adding nodemon for the hot-reloading development
RUN apk update && apk add nodejs && rm -rf /var/cache/apk/*
RUN apk add npm
RUN npm install -g nodemon

WORKDIR /di-example-go

# Copy only go.mod and go.sum to install dependencies
# The src code will be mounted later
COPY go.mod .
COPY go.sum .

RUN go mod download

CMD ["nodemon", "-e", "go,json,yaml", "--exec", "go run main.go serve", "--signal", "SIGTERM"]
