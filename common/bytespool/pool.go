package bytespool

import "sync"

const (
	numPools = 4
)

var (
	pool     [numPools]sync.Pool
	poolSize [numPools]int32
)

// GetPool returns a sync.Pool that generates bytes array with at least the given size.
// It may return nil if no such pool exists.
//
// v2ray:api:stable
func GetPool(size int32) *sync.Pool {
	for idx, ps := range poolSize {
		if size <= ps {
			return &pool[idx]
		}
	}
	return nil
}
