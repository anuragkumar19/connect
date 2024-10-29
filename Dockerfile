FROM golang:1.23.1 AS golang-base
WORKDIR /app
RUN go install github.com/go-task/task/v3/cmd/task@latest


FROM golang-base AS build-stage
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target="/root/.cache/go-build" \
    --mount=type=bind,target=. \ 
    CGO_ENABLED=0 GOOS=linux go build -ldflags='-s' -o /connect-server ./cmd/connect-server/

FROM golang-base AS db-migration
COPY database/sql/schema ./database/sql/schema
COPY Taskfile.yml ./Taskfile.yml
RUN task install-goose
CMD [ "task","migrate","--","up" ]


# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

FROM gcr.io/distroless/static-debian12 AS build-release-stage
WORKDIR /
COPY --from=build-stage /connect-server /connect-server
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/connect-server"]