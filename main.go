package main

import "github.com/hacktiv8-ks07-g04/final-project-3/handler"

func main() {
	// ISSUE: No graceful shutdown. StartApp blocks on r.Run() with no SIGTERM/SIGINT
	// handling; the server is killed abruptly on deploy (in-flight requests dropped,
	// DB connection never closed). See handler/app.go:74.
	// ISSUE: No error handling — StartApp has no return value, so startup failures
	// (DB connect panic, port in use) can't be surfaced or retried gracefully.
	handler.StartApp()
}
