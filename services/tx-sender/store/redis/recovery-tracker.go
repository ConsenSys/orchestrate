package redis

import (
	"github.com/consensys/orchestrate/pkg/toolkit/database/redis"
	"github.com/consensys/orchestrate/services/tx-sender/store"
)

type nonceRecoveryTracker struct {
	redisCli *redis.Client
}

func NewNonceRecoveryTracker(redisCli *redis.Client) store.RecoveryTracker {
	return &nonceRecoveryTracker{
		redisCli: redisCli,
	}
}

const recoverTrackerSuf = "recover-tracker"

func (t *nonceRecoveryTracker) Recovering(key string) (count uint64) {
	v, b, err := t.redisCli.LoadUint64(computeKey(key, recoverTrackerSuf))
	if err != nil || !b {
		return 0
	}
	return v
}

func (t *nonceRecoveryTracker) Recover(key string) {
	_ = t.redisCli.Incr(computeKey(key, recoverTrackerSuf))
}

func (t *nonceRecoveryTracker) Recovered(key string) {
	_ = t.redisCli.Delete(computeKey(key, recoverTrackerSuf))
}
