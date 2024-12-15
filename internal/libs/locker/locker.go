package locker

import (
	"context"

	"github.com/Krystian19/quote_management/internal/libs/constants"
	"github.com/Krystian19/quote_management/internal/libs/rds"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type Locker struct {
	lock *redsync.Redsync
	ctx  context.Context
}

func NewLocker(rdds *rds.RDS, ctx context.Context) Locker {
	return Locker{
		lock: redsync.New(goredis.NewPool(rdds.GetClient())),
		ctx:  ctx,
	}
}

func (l *Locker) CreateMutex(name constants.Lock) *redsync.Mutex {
	return l.lock.NewMutex(string(name))
}

func (l *Locker) AcquireLock(mu *redsync.Mutex) error {
	return mu.LockContext(l.ctx)
}

func (l *Locker) ReleaseLock(mu *redsync.Mutex) error {
	_, err := mu.UnlockContext(l.ctx)
	return err
}
