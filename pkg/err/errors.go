package errs

import "errors"

var (
	// Client / Request Errors
	// ErrNotFound is returned when a requested entity cannot be located.
	ErrNotFound = errors.New("not found")

	// ErrInvalidArgument is returned when a caller supplies an invalid or out-of-range argument.
	ErrInvalidArgument = errors.New("invalid argument")

	// ErrAlreadyExists is returned when trying to create a resource that already exists.
	ErrAlreadyExists = errors.New("already exists")

	// // ErrConflict is returned when an operation conflicts with current state (e.g. optimistic lock, order already processed).
	// ErrConflict = errors.New("conflict")

	// // ErrPreconditionFailed is returned when system state does not satisfy prerequisites (e.g. empty cart on checkout).
	// ErrPreconditionFailed = errors.New("precondition failed")

	// Auth & Permission
	// ErrUnauthorized is returned when authentication is missing or invalid.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrForbidden is returned when the caller does not have permission.
	ErrForbidden = errors.New("forbidden")

	// Rate Limiting
	// ErrRateLimited is returned when a client exceeds allowed request quotas.
	ErrRateLimited = errors.New("rate limit exceeded")

	// Server / Dependency Errors
	// ErrInternal is returned when an unexpected internal error occurs (e.g. database failure).
	ErrInternal = errors.New("internal error")

	// ErrUnavailable is returned when a downstream service, database, or dependency is temporarily down.
	ErrUnavailable = errors.New("service unavailable")

	// ErrTimeout is returned when an operation or downstream call exceeds its deadline.
	ErrTimeout = errors.New("timeout")

	// ErrNotImplemented is returned when an operation is not supported.
	ErrNotImplemented = errors.New("not implemented")
)
