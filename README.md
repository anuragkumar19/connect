# Connect

## Go tools required :- task

## Todo
- SQL formatter
- Jetstream consumer(current implementations is not good) and publisher protobuf gen
- NATS pub/sub protobuf gen (bus)
- WS protocol from protobuf
- services/.... Init func?? eg. for add listeners, starting cron jobs, starting cleanup func
- NATS realtime kv eg. for shared revoked tokens
- CSpell
- API testing client for connectrpc
- Opentelemetry Tracing
- Metrics

## 24/10 tasks
- Setup cli
- Realtime protocol on top of websocket
✔ Events and event schema generated from protobuf definition
NATS and jetstream client(pub-sub, work-queue, kv) generation from protobuf definition
- Soft delete
- Error handling and errordetails package
- Interceptors
- Caching
- Startup Context and graceful shutdown, signals, GenServer interface, run server, monitor and restart, listen to os signals and terminate
- Gocilint
- Github actions lint, format & diff, test
- Setup e2e test
- Auth middleware
- cors only in development: app_mode env variable
- Setup cspell
- Request validation
- Seeding data
- Document protobuf

