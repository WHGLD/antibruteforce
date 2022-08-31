package app

import (
	"context"
	"strconv"
	"time"

	"github.com/WHGLD/antibruteforce/internal/cache"
	"github.com/jonboulle/clockwork"
)

const (
	loginPrefix    = "LOGIN"
	passwordPrefix = "PASSWORD"
	ipPrefix       = "IP"

	RateLimitCacheTTL = time.Minute
)

type RateLimiter struct {
	cache         cache.RedisCache
	clock         clockwork.Clock
	loginLimit    int
	passwordLimit int
	ipLimit       int
}

func NewRateLimiter(
	cache cache.RedisCache,
	clock clockwork.Clock,
	loginLimit int,
	passwordLimit int,
	ipLimit int,
) *RateLimiter {
	return &RateLimiter{
		cache:         cache,
		clock:         clock,
		loginLimit:    loginLimit,
		passwordLimit: passwordLimit,
		ipLimit:       ipLimit,
	}
}

func (rl *RateLimiter) CheckRateLimits(ctx context.Context, login, password, ip string) (bool, error) {
	now := strconv.Itoa(rl.clock.Now().Minute())
	iterate := map[string]string{
		loginPrefix:    login + now,
		passwordPrefix: password + now,
		ipPrefix:       ip + now,
	}

	for prefix, cacheKey := range iterate {
		cacheKey = prefix + ":" + cacheKey
		errIncr := rl.cache.Increment(ctx, cacheKey, RateLimitCacheTTL)
		if errIncr != nil {
			return false, errIncr
		}

		currentLimitCounter, errGet := rl.cache.Get(ctx, cacheKey)
		if errGet != nil {
			return false, errGet
		}

		switch prefix {
		case loginPrefix:
			if currentLimitCounter > rl.loginLimit {
				return false, nil
			}
			break
		case passwordPrefix:
			if currentLimitCounter > rl.passwordLimit {
				return false, nil
			}
			break
		case ipPrefix:
			if currentLimitCounter > rl.ipLimit {
				return false, nil
			}
			break
		}
	}

	return true, nil
}

func (rl *RateLimiter) ResetBuckets(ctx context.Context, login, password, ip string) error {
	now := strconv.Itoa(rl.clock.Now().Minute())

	err := rl.cache.Del(ctx, loginPrefix+":"+login+now)
	if err != nil {
		return err
	}

	err = rl.cache.Del(ctx, passwordPrefix+":"+password+now)
	if err != nil {
		return err
	}

	err = rl.cache.Del(ctx, ipPrefix+":"+ip+now)
	if err != nil {
		return err
	}

	return nil
}
