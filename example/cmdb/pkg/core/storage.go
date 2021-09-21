package core

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

type StorageInterface interface {
	Get(key string) (interface{}, error)
	Update(key string) (interface{}, error)
	Delete(key string) (interface{}, error)
}

type Storage struct {
	core *Core
	expiration time.Duration
	db StorageInterface
	cache *redis.Client
}

func (s *Storage) Get(key string) (interface{}, error) {
	data,err := s.cache.Get(s.core.TraceLog.Context(), key).Result()
	if err != nil {
		s.core.TraceLog.Error("redis get key error", zap.Error(err), zap.String("key", key))
		return data, nil
	}
	return s.db.Get(key)
}

func (s *Storage) Update(key string, value string, expiration time.Duration) (interface{}, error) {
	s.cache.GetSet()
	s.cache.Set(s.core.TraceLog.Context(), key, value, expiration)
	data, err := s.cache.(s.core.TraceLog.Context(), key).Result()
	if err != nil {
		s.core.TraceLog.Error("redis update key error", zap.Error(err), zap.String("key", key))
	} else {
		return data, nil
	}
	return s.db.Update(key)
}

func (s *Storage) Delete(key string) (interface{}, error) {
	data, err := s.cache.Get(s.core.TraceLog.Context(), key).Result()
}

