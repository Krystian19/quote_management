package locker_test

import (
	"context"
	"testing"

	"github.com/Krystian19/quote_management/internal/libs/constants"
	"github.com/Krystian19/quote_management/internal/libs/env"
	"github.com/Krystian19/quote_management/internal/libs/locker"
	"github.com/Krystian19/quote_management/internal/libs/rds"
	"github.com/Krystian19/quote_management/internal/libs/utils"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func getRDSClient() (*rds.RDS, error) {
	return rds.NewRDS(context.Background(), env.GetEnv().REDIS_TEST_URL)
}

func getLockerClient() (*locker.Locker, error) {
	rdds, err := getRDSClient()
	if err != nil {
		return nil, err
	}

	return utils.GetPtr(locker.NewLocker(rdds, context.Background())), nil
}

func TestLockerLockValue(t *testing.T) {
	fake := faker.New()

	lck, err := getLockerClient()
	if err != nil {
		t.Fatal(err)
	}

	key := fake.Hash().SHA256()
	mu := lck.CreateMutex(constants.Lock(key))

	err = lck.AcquireLock(mu)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, lck.AcquireLock(mu))
}

func TestLockerUnLockValue(t *testing.T) {
	fake := faker.New()

	lck, err := getLockerClient()
	if err != nil {
		t.Fatal(err)
	}

	key := fake.Hash().SHA256()
	mu := lck.CreateMutex(constants.Lock(key))

	err = lck.AcquireLock(mu)
	if err != nil {
		t.Fatal(err)
	}

	err = lck.ReleaseLock(mu)
	if err != nil {
		t.Fatal(err)
	}

	require.Nil(t, lck.AcquireLock(mu))
}
