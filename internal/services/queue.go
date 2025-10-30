package services

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

func PushToQueue(client *asynq.Client, queueName string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	task := asynq.NewTask(queueName, data)
	_, err = client.Enqueue(task)
	return err
}

// UNUSED
func SetKey(client *redis.Client, key string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(context.Background(), key, data, 0).Err()
}

func GetKey[T any](client *redis.Client, key string) (*T, error) {
	data, err := client.Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
