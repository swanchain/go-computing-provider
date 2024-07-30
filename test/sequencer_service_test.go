package test

import (
	"github.com/swanchain/go-computing-provider/internal/computing"
	"testing"
)

func TestQueryTask(t *testing.T) {
	taskIds := []int64{35913, 35915, 35921}
	computing.NewSequencer().QueryTask(taskIds...)
}
