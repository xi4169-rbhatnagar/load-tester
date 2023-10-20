package persistence

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/xi4169-rbhatnagar/load-tester/modules/server/config"
	"github.com/xi4169-rbhatnagar/load-tester/modules/server/model/types"
)

type RedisUserRequestStore struct {
	types.UserRequestStore
	rdb *redis.Client
}

func NewRedisStore(address, password string, db int) types.UserRequestStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return RedisUserRequestStore{
		rdb: rdb,
	}
}

func (s RedisUserRequestStore) Store(ctx context.Context, req types.UserRequest) error {
	jreq, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("error marshalling input request: %v", err)
	}
	_, err = s.rdb.LPush(ctx, config.RedisUserRequestListName, string(jreq)).Result()
	return err
}
