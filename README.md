# Connect

## Go tools required :- task, buf, sqlc, google/yamlfmt

## Todo
- SQL formatter

## 24/10 tasks
- Setup cli
✔ Setup logger (zerolog)
✔ Setup configuration (env)
✔ - Setup postgres (pgx/v5)
✔ - Setup NATS (nats.go)
- Realtime protocol on top of websocket
✔ Events and event schema generated from protobuf definition
NATS and jetstream client(pub-sub, work-queue, kv) generation from protobuf definition
✔ - Setup storage (minio/v7)
- Create db tables and views
✔ - Setup mailer and consumer (smtp)
- Setup Server struct in api/
- Error handling and errordetails package
- Custom error handler to prevent sending backend error messages to the client
- Interceptors
- Caching
✔ Setup connectRPC h2c server
- Startup Context and graceful shutdown, signals, GenServer interface, run server, monitor and restart, listen to os signals and terminate
- Gocilint
- Github actions lint, format & diff, test
- Setup e2e test
- Auth middleware
✔ CORS middleware
- cors only in development: app_mode env variable
✔ Git commit hash/release version integration ref- grafana & let's go further - partially done, automating with taskfile remains
- Setup cspell
✔ Setup docker
✔ Managing tools
- Request validation
- Seeding data
- Document protobuf
- Logger and different pkg from with logging is being done

# 23/10-24/10  Non coding tasks
- Learn about otel, metrics and logging
- Grafana, logs, otel, prometheus
- Checkout links in whatsapp


## TODO NOW AFTER BREAK
- SETUP TOKEN (NOT REQUIRED)