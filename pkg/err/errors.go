package errs

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func ToGRPCStatusError(ctx context.Context, err error, logger *slog.Logger, msg string) error {
	if err == nil {
		return nil
	}
	if _, ok := status.FromError(err); ok {
		return err
	}

	switch {
	case errors.Is(err, ErrNotFound):
		logger.ErrorContext(ctx, "not found", slog.String("msg", msg), slog.Any("error", err))
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, ErrAlreadyExists):
		logger.ErrorContext(ctx, "already exists", slog.String("msg", msg), slog.Any("error", err))
		return status.Error(codes.AlreadyExists, err.Error())
		// case errors.Is(err, ErrAborted):
		// logger.ErrorContext(ctx, "internal server error", slog.String("msg", msg), slog.Any("error", err))
	// 	return status.Error(codes.Aborted, err.Error())
	case errors.Is(err, ErrInvalidArgument):
		logger.ErrorContext(ctx, "invalid arguement", slog.String("msg", msg), slog.Any("error", err))
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, ErrUnauthorized):
		logger.ErrorContext(ctx, "unauthorized", slog.String("msg", msg), slog.Any("error", err))
		return status.Error(codes.PermissionDenied, err.Error())
	// case errors.Is(err, ErrForbidden):
	// logger.ErrorContext(ctx, "internal server error", slog.String("msg", msg), slog.Any("error", err))
	// 	return status.Error(codes., err.Error())
	case errors.Is(err, ErrRateLimited):
		logger.ErrorContext(ctx, "rate limited", slog.String("msg", msg), slog.Any("error", err))
		return status.Error(codes.ResourceExhausted, err.Error())

	case errors.Is(err, ErrUnavailable):
		logger.ErrorContext(ctx, "unavailable", slog.String("msg", msg), slog.Any("error", err))
		return status.Error(codes.Unavailable, err.Error())

	case errors.Is(err, ErrTimeout):
		logger.ErrorContext(ctx, "timeout", slog.String("msg", msg), slog.Any("error", err))
		return status.Error(codes.DeadlineExceeded, err.Error())

	default:
		logger.ErrorContext(ctx, "internal server error", slog.String("msg", msg), slog.Any("error", err))
		return status.Error(codes.Internal, "internal server error")
	}
}
