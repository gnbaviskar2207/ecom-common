package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToMongo(ctx context.Context, url string, logger *slog.Logger) (*mongo.Client, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	const (
		maxAttempts = 5
		baseDelay   = 1 * time.Second
		pingTimeout = 2 * time.Second
	)

	var client *mongo.Client
	var err error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		clientOptions := options.Client().ApplyURI(url).
			SetServerSelectionTimeout(5 * time.Second).
			SetConnectTimeout(5 * time.Second)
		client, err = mongo.Connect(clientOptions)
		if err != nil {
			_ = client.Disconnect(ctx)
			return nil, fmt.Errorf("mongodb error while mongo connect %w", err)
		}

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		pingCtx, pingCancel := context.WithTimeout(ctx, pingTimeout)
		err = client.Ping(pingCtx, nil)
		pingCancel()
		if err == nil {
			return client, nil
		}
		_ = client.Disconnect(ctx)
		if attempt == maxAttempts {
			break
		}
		backoff := baseDelay * time.Duration(attempt-1)
		logger.WarnContext(ctx, "mongodb ping failed, retrying",
			slog.Int("attempt", attempt),
			slog.Int("max_attempts", maxAttempts),
			slog.Duration("backoff", backoff),
			slog.String("error", err.Error()),
		)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(backoff):
			continue
		}
	}

	return nil, fmt.Errorf("could not connect to mongo after %d, err %w ", maxAttempts, err)

}
