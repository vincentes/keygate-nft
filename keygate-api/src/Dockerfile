FROM golang:latest
MAINTAINER Vicente Bermudez <thevincentes@gmail.com>
# All commands will be executed in the container's /keygate-api directory
WORKDIR /keygate-api
# We'll install air to live reload the server

# Confirm bootstrap.sh is executable
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Finally, we'll set the entrypoint to run air in the /keygate-api directory
ENTRYPOINT ["go", "run", "src/server.go"]