package test

import (
	"github.com/swanchain/go-computing-provider/internal/computing"
	"testing"
)

func TestGetToken(t *testing.T) {
	computing.NewSequencer("http://sequencer.storefrontiers.cn/v1").GetToken()
}
