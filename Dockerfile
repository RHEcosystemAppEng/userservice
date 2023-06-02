FROM registry.access.redhat.com/ubi8/go-toolset:1.18.4 AS builder

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY rest-server.go rest-server.go
COPY .env.openshift.dev .env
COPY signercert.pem signercert.pem

COPY handlers handlers/
COPY middlewares middlewares/
COPY routes routes/
COPY types types/
COPY env env/

# Build
RUN GOOS=linux GOARCH=amd64 go build -a -o keycloak-user-service .

# Build the operator image
FROM registry.access.redhat.com/ubi8-minimal:8.7

WORKDIR /
COPY --from=builder /opt/app-root/src/keycloak-user-service .
COPY --from=builder /opt/app-root/src/signercert.pem .
COPY --from=builder /opt/app-root/src/.env .

USER 65532:65532
ENTRYPOINT ["/keycloak-user-service"]
