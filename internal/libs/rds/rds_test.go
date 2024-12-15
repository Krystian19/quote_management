package rds_test

import (
	"context"
	"testing"
	"time"

	"github.com/Krystian19/quote_management/internal/libs/env"
	"github.com/Krystian19/quote_management/internal/libs/rds"
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

func getRDSClient() (*rds.RDS, error) {
	return rds.NewRDS(context.Background(), env.GetEnv().REDIS_TEST_URL)
}

func TestRDSSetValue(t *testing.T) {
	fake := faker.New()

	cache, err := getRDSClient()
	if err != nil {
		t.Fatal(err)
	}

	key := fake.Hash().SHA256()
	val := fake.Hash().SHA256()

	if err := cache.SetValueInCache(key, val, time.Minute); err != nil {
		t.Fatal(err)
	}

	foundValue, err := cache.GetValueFromCache(key)
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, foundValue)
	require.Equal(t, val, *foundValue)
}

func TestRDSSetValueWithExpiry(t *testing.T) {
	fake := faker.New()
	timeToWait := time.Millisecond * 200

	cache, err := getRDSClient()
	if err != nil {
		t.Fatal(err)
	}

	key := fake.Hash().SHA256()
	val := fake.Hash().SHA256()

	if err := cache.SetValueInCache(key, val, timeToWait); err != nil {
		t.Fatal(err)
	}

	foundValue, _ := cache.GetValueFromCache(key)
	require.NotNil(t, foundValue)

	time.Sleep(timeToWait * 2)

	foundValue, _ = cache.GetValueFromCache(key)
	require.Nil(t, foundValue)
}

func TestRDSIsValueInCache(t *testing.T) {
	fake := faker.New()

	cache, err := getRDSClient()
	if err != nil {
		t.Fatal(err)
	}

	key := fake.Hash().SHA256()
	val := fake.Hash().SHA256()

	ok, err := cache.IsValueInCache(key)
	if err != nil {
		t.Fatal(err)
	}

	require.False(t, ok)

	if err := cache.SetValueInCache(key, val, time.Minute); err != nil {
		t.Fatal(err)
	}

	ok, err = cache.IsValueInCache(key)
	if err != nil {
		t.Fatal(err)
	}

	require.True(t, ok)
}

func TestRDSRemoveValue(t *testing.T) {
	fake := faker.New()

	cache, err := getRDSClient()
	if err != nil {
		t.Fatal(err)
	}

	key := fake.Hash().SHA256()
	val := fake.Hash().SHA256()

	if err := cache.SetValueInCache(key, val, time.Minute); err != nil {
		t.Fatal(err)
	}

	ok, _ := cache.IsValueInCache(key)
	require.True(t, ok)

	if err := cache.RemoveFromCache(key); err != nil {
		t.Fatal(err)
	}

	ok, _ = cache.IsValueInCache(key)
	require.False(t, ok)
}

func TestRDSGetClient(t *testing.T) {
	cache, err := getRDSClient()
	if err != nil {
		t.Fatal(err)
	}

	require.NotNil(t, cache.GetClient())
}
