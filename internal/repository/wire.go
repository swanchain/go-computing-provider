//go:build wireinject
// +build wireinject

package repository

import (
	"github.com/google/wire"
)

func NewJob() Job {
	wire.Build(jobRepository)
	return Job{}
}

func NewTask() Task {
	wire.Build(taskRepository)
	return Task{}
}

func NewProvider() Provider {
	wire.Build(provideRepository)
	return Provider{}
}
