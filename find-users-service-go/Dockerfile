FROM registry.access.redhat.com/ubi8/go-toolset:1.18.4 AS builder

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY environment.go environment.go
COPY rest-server.go rest-server.go
COPY .env .env

COPY handlers/ handlers/
COPY middlewares/ middlewares/
COPY routes/ routes/
COPY types/ types/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o userservice .

# Build the operator image
FROM registry.access.redhat.com/ubi8-minimal:8.7

WORKDIR /
COPY --from=builder /opt/app-root/src/userservice .
USER 65532:65532

ENTRYPOINT ["/userservice"]
