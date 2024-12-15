package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	REDIS_TEST_URL                   string
	RABBIT_MQ_URL                    string
	CLERK_DB_URL                     string
	CLERK_DB_NAME                    string
	CLERK_DB_NAME_TEST               string
	CLERK_EVENTS_REDIS_URL           string
	CLERK_API_PORT                   int
	CLERK_WRITER_WORKER_COUNT        int
	CLERK_WRITER_LOCKER_REDIS_URL    string
	TAINTED_TXS_CACHE_REDIS_URL      string
	SYNCHRONIZER_SYNC_PERCENTAGE     uint64
	TAINT_SHORT_IS_TAINT_MODE_ACTIVE bool
	TAINT_LONG_WORKER_COUNT          int
	BTC_NODE_MAINNET_URL             string
	BTC_NODE_MAINNET_USER            string
	BTC_NODE_MAINNET_PASSWORD        string
	BTC_NODE_MAINNET_ZMQ_URL         string
	BTC_NODE_TESTNET_URL             string
	BTC_NODE_TESTNET_USER            string
	BTC_NODE_TESTNET_PASSWORD        string
	BTC_NODE_TESTNET_ZMQ_URL         string
	BTC_NODE_REGTEST_URL             string
	BTC_NODE_REGTEST_USER            string
	BTC_NODE_REGTEST_PASSWORD        string
	BTC_NODE_REGTEST_ZMQ_URL         string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func getBoolEnv(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

func getUint64Env(key string, fallback uint64) uint64 {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.ParseUint(value, 10, 64); err == nil {
			return intValue
		}
	}
	return fallback
}

func GetEnv() Env {
	// ignore file if we are currently in the CI container
	if !getBoolEnv("IS_CI", false) {
		godotenv.Load("../../../.env")
	}

	return Env{
		REDIS_TEST_URL:                   getEnv("REDIS_TEST_URL", "aurora.btc.test.redis:6379"),
		RABBIT_MQ_URL:                    getEnv("RABBIT_MQ_URL", "amqp://user:pass@aurora.btc.rabbitmq:5672"),
		CLERK_DB_URL:                     getEnv("CLERK_DB_URL", "mongodb://root:pass@aurora.btc.clerk.mongodb:27017/"),
		CLERK_DB_NAME:                    getEnv("CLERK_DB_NAME", "clerk_db"),
		CLERK_DB_NAME_TEST:               getEnv("CLERK_DB_NAME_TEST", "clerk_db_test"),
		CLERK_EVENTS_REDIS_URL:           getEnv("CLERK_EVENTS_REDIS_URL", "aurora.btc.clerk.events.redis:6379"),
		CLERK_API_PORT:                   getIntEnv("CLERK_API_PORT", 9040),
		CLERK_WRITER_WORKER_COUNT:        getIntEnv("CLERK_WRITER_WORKER_COUNT", 4),
		CLERK_WRITER_LOCKER_REDIS_URL:    getEnv("CLERK_WRITER_LOCKER_REDIS_URL", "aurora.btc.writer.locker.redis:6379"),
		TAINTED_TXS_CACHE_REDIS_URL:      getEnv("TAINTED_TXS_CACHE_REDIS_URL", "aurora.btc.tainted.txs.cache.redis:6379"),
		SYNCHRONIZER_SYNC_PERCENTAGE:     getUint64Env("SYNCHRONIZER_SYNC_PERCENTAGE", 100),
		TAINT_SHORT_IS_TAINT_MODE_ACTIVE: getBoolEnv("TAINT_SHORT_IS_TAINT_MODE_ACTIVE", true),
		TAINT_LONG_WORKER_COUNT:          getIntEnv("TAINT_LONG_WORKER_COUNT", 4),
		BTC_NODE_MAINNET_URL:             getEnv("BTC_NODE_MAINNET_URL", "aurora.btc.clerk.node:8332"),
		BTC_NODE_MAINNET_USER:            getEnv("BTC_NODE_MAINNET_USER", "user"),
		BTC_NODE_MAINNET_PASSWORD:        getEnv("BTC_NODE_MAINNET_PASSWORD", "pass"),
		BTC_NODE_MAINNET_ZMQ_URL:         getEnv("BTC_NODE_MAINNET_ZMQ_URL", "tcp://aurora.btc.clerk.node:29000"),
		BTC_NODE_TESTNET_URL:             getEnv("BTC_NODE_TESTNET_URL", "aurora.btc.clerk.node:18332"),
		BTC_NODE_TESTNET_USER:            getEnv("BTC_NODE_TESTNET_USER", "user"),
		BTC_NODE_TESTNET_PASSWORD:        getEnv("BTC_NODE_TESTNET_PASSWORD", "pass"),
		BTC_NODE_TESTNET_ZMQ_URL:         getEnv("BTC_NODE_TESTNET_ZMQ_URL", "tcp://aurora.btc.clerk.node:29000"),
		BTC_NODE_REGTEST_URL:             getEnv("BTC_NODE_REGTEST_URL", "aurora.btc.clerk.node.ci:18443"),
		BTC_NODE_REGTEST_USER:            getEnv("BTC_NODE_REGTEST_USER", "user"),
		BTC_NODE_REGTEST_PASSWORD:        getEnv("BTC_NODE_REGTEST_PASSWORD", "pass"),
		BTC_NODE_REGTEST_ZMQ_URL:         getEnv("BTC_NODE_REGTEST_ZMQ_URL", "tcp://aurora.btc.clerk.node:29000"),
	}
}
