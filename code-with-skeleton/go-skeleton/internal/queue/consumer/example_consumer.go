package consumer

import (
	"context"

	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/entity"
	"github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/helper"
	mongoRepo "github.com/dikadittya/learn-golang/tree/code-with-skeleton/go-skeleton/internal/repository/mongodb"
)

type ExampleQueue struct {
	ctx          context.Context
	logMongoRepo mongoRepo.LogRepository
}

type ExampleConsumer interface {
	Process(payload map[string]interface{}) error
}

func NewExampleConsumer(
	ctx context.Context,
	logMongoRepo mongoRepo.LogRepository,
) ExampleConsumer {
	return &ExampleQueue{ctx, logMongoRepo}
}

func (l *ExampleQueue) Process(payload map[string]interface{}) error {
	var params entity.Log
	params.LoadFromMap(payload)

	helper.Dump(params)

	return nil
}
