package ports

import (
	"context"
)

type IEvents interface {
	Saved(ctx context.Context, message []byte) error
}
