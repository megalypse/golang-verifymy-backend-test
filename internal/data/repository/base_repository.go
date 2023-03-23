package repository

import "context"

type baseRepository interface {
	NewConnection(context.Context) Closable
}
