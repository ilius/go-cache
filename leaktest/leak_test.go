package leaktest

import (
	"runtime"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	cache.New(5*time.Minute, 10*time.Minute)
	runtime.GC()
}
