package services

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"

	"markets-api/internal/environment"
)

func PushToQueue(env *environment.ClientEnvironment, queueName string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	task := asynq.NewTask(queueName, data)
	_, err = env.QueueClient.Enqueue(task)
	return err
}

// UNUSED
func SetKey(env *environment.ClientEnvironment, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return env.RedisClient.Set(context.Background(), key, data, 0).Err()
}

func GetKey[T any](env *environment.ClientEnvironment, key string) (*T, error) {
	data, err := env.RedisClient.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
