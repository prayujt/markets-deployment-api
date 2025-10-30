package services

import (
	"encoding/json"

	"github.com/hibiken/asynq"
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
