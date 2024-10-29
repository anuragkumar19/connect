# Connect

## Go tools required :- task, buf, sqlc, google/yamlfmt

## Todo
- SQL formatter

## 24/10 tasks
- Setup cli
- Setup logger (zerolog)
- Setup configuration (env and cli flags)
✔ - Setup postgres (pgx/v5)
✔ - Setup NATS (nats.go)
- Realtime protocol on top of websocket
- Events and event schema generated from protobuf definition
✔ - Setup storage (minio/v7)
- Create db tables and views
✔ - Setup mailer and consumer (smtp)
- Setup Server struct in api/
- Error handling and errordetails package
- Custom error handler to prevent sending backend error messages to the client
- Interceptors
- Caching
- Setup connectRPC h2c server
- Startup Context and graceful shutdown, signals, GenServer interface, run server, monitor and restart, listen to os signals and terminate
- Gocilint
- Github actions lint, format & diff, test
- Setup e2e test
- Auth middleware
- CORS middleware
- Git commit hash/release version integration ref- grafana & let's go further - partially done, automating with taskfile remains
- Setup cspell
- Setup docker
- Live reload for generators and go build
- Managing tools
- Request validation
- Seeding data
- Document protobuf

# 23/10-24/10  Non coding tasks
- Figure out db schema for user and authentication
- Auth and MFA flows
- Learn about otel, metrics and logging
- Grafana, logs, otel, prometheus
- Checkout links in whatsapp


## TODO NOW AFTER BREAK
- DOCKER GO BUILD CACHE AND MODULE DOWNLOAD CACHE VOLUME MOUNT
- DOCKER COMPOSE WITH DB/NATS
- ✔ SETUP CONFIG LOADING AND MANAGEMENT
- SETUP TOKEN (NOT REQUIRED)