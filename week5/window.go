package week5

import (
	"sync"
	"time"
)

type Window struct {
	sync.RWMutex
	broken bool
	// 滑动窗口大小
	numBuckets int
	// 桶队列
	buckets []*BucketedCounter
	// 触发熔断的请求总数阈值
	reqThreshold int
	// 出发熔断的失败率阈值
	failedThreshold float64
	// 上次熔断发生时间
	lastBreakTime time.Time
	// 熔断恢复的时间间隔
	brokeTimeGap time.Duration
}

/**
 * @description: 添加一个桶
 * @param  {*}
 * @return {*}
 */
func (r *Window) AddBucket() {
	r.Lock()
	defer r.Unlock()
	r.buckets = append(r.buckets, NewBucket())
	if !(len(r.buckets) < r.numBuckets+1) {
		r.buckets = r.buckets[1:]
	}
}

/**
 * @description: 记录数据
 * @param  {*}
 * @return {*}
 * @param {bool} result
 */
func (r *Window) RecordResult(result bool) {
	if len(r.buckets) == 0 {
		r.AddBucket()
	}
	r.buckets[len(r.buckets)-1].Record(result)
}

/**
 * @description:  启动
 * @param  {*}
 * @return {*}
 */
func (r *Window) Start() {
	go r.AddBucket()
}

/**
 * @description: 根据次数判断是否需要熔断
 * @param  {*}
 * @return {*}
 */
func (r *Window) BreakJudgement() bool {
	r.RLock()
	defer r.RUnlock()
	total := 0
	failed := 0
	for _, v := range r.buckets {
		total += v.SuccessTotal
		failed += v.Failtotal
	}
	if float64(failed)/float64(total) > r.failedThreshold && total > r.reqThreshold {
		return true
	}
	return false
}

/**
 * @description:  监控滑动窗口的总失败次数与是否开启熔断
 * @param  {*}
 * @return {*}
 */
func (r *Window) WhetherOpen() {
	go brokenJudge(r)
}

func brokenJudge(r *Window) {
	for {
		if r.broken {
			if r.OverBrokenTime() {
				r.Lock()
				r.broken = false
				r.Unlock()
			}
			continue
		}
		if r.BreakJudgement() {
			r.Lock()
			r.broken = true
			r.lastBreakTime = time.Now()
			r.Unlock()
		}
	}
}

/**
 * @description: 是否超过熔断间隔期
 * @param  {*}
 * @return {*}
 */
func (r *Window) OverBrokenTime() bool {
	return time.Since(r.lastBreakTime) > r.brokeTimeGap
}
