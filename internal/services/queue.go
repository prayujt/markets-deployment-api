package services

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"

	"markets-api/internal/environment"
)

func PushToQueue(env *environment.BaseEnvironment, queueName string, payload any, retries *int) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	var task *asynq.Task
	if retries != nil {
		task = asynq.NewTask(queueName, data, asynq.MaxRetry(*retries))
	} else {
		task = asynq.NewTask(queueName, data)
	}
	_, err = env.QueueClient.Enqueue(task)
	return err
}

func SetKey(env *environment.BaseEnvironment, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return env.RedisClient.Set(context.Background(), key, data, 0).Err()
}

func GetKey[T any](env *environment.BaseEnvironment, key string) (*T, error) {
	data, err := env.RedisClient.Get(context.Background(), key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
