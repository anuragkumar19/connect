// Original Source: https://github.com/connectrpc/connect-go/blob/main/code.go
// Copyright 2021-2024 The Connect Authors
// Licensed under the Apache License, Version 2.0

package errors

type ErrType uint32

const (
	// TypeInvalidArgument indicates that client supplied an invalid argument to service.
	TypeInvalidArgument ErrType = iota + 1

	// TypeDeadlineExceeded indicates that deadline expired before the operation
	// could complete.
	TypeDeadlineExceeded

	// TypeNotFound indicates that some requested entity was not found.
	TypeNotFound

	// TypeAlreadyExists indicates that client attempted to create an entity that already exists.
	TypeAlreadyExists

	// TypePermissionDenied indicates that the caller doesn't have permission to
	// execute the specified operation.
	TypePermissionDenied

	// TypeResourceExhausted indicates that some resource has been exhausted. For
	// example, a per-user quota may be exhausted or the entire file system may
	// be full.
	TypeResourceExhausted

	// TypeFailedPrecondition indicates that the system is not in a state
	// required for the operation's execution.
	TypeFailedPrecondition

	// TypeAborted indicates that operation was aborted by the system, usually
	// because of a concurrency issue such as a sequencer check failure or
	// transaction abort.
	TypeAborted

	// TypeUnimplemented indicates that the operation isn't implemented,
	// supported, or enabled in this service.
	TypeUnimplemented

	// TypeInternal indicates that some invariants expected by the underlying
	// system have been broken. This type is reserved for serious errors.
	TypeInternal

	// TypeUnauthenticated indicates that the request does not have valid
	// authentication credentials for the operation.
	TypeUnauthenticated
)
