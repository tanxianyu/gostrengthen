package week5

import (
	"sync"
	"time"
)

type BucketedCounter struct {
	sync.RWMutex

	SuccessTotal int

	Failtotal int

	Timestamp time.Time
}

func (bucket *BucketedCounter) Record(res bool) {
	bucket.Lock()
	defer bucket.Unlock()

	if res {
		bucket.SuccessTotal++
		return
	}
	bucket.Failtotal++
}

func NewBucket() *BucketedCounter {
	return &BucketedCounter{
		Timestamp: time.Now(),
	}
}
