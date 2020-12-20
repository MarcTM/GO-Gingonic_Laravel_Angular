FROM golang:1.15
WORKDIR /go/src/go/recipes
COPY . .
RUN bash -c "rm -f go.mod || true rm -f go.sum || true"
RUN bash -c "go mod init go_server && go mod tidy"
RUN bash -c "go get github.com/pilu/fresh"