package data_structures

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Calculate the average speed every fixed interval
// Interval is fixed when instance is created
type IntervalSpeedCalculator struct {
	flow    uint64
	speed   uint64 // unit is /second
	t       *time.Ticker
	over    chan struct{}
	closeMu *sync.Mutex
}

func (s *IntervalSpeedCalculator) GetSpeed() uint64 {
	return atomic.LoadUint64(&s.speed)
}

func (s *IntervalSpeedCalculator) AddFlow(n uint64) {
	atomic.AddUint64(&s.flow, n)
}

func (s *IntervalSpeedCalculator) Close() {
	s.closeMu.Lock()
	defer s.closeMu.Unlock()
	select {
	case <-s.over:
		return
	default:
		close(s.over)
		s.t.Stop()
	}
}

func NewIntervalSpeedCalculator(t time.Duration) *IntervalSpeedCalculator {
	delay := time.NewTicker(t)
	s := &IntervalSpeedCalculator{0, 0, delay, make(chan struct{}), new(sync.Mutex)}
	go func(cal *IntervalSpeedCalculator) { // update speed every t
		for {
			select {
			case <-delay.C:
				// set speed/sec
				f := atomic.LoadUint64(&cal.flow)
				atomic.StoreUint64(&cal.speed, f/uint64(t.Seconds()))
				atomic.AddUint64(&cal.flow, -f)
			case <-cal.over:
				return
			}
		}
	}(s)
	return s
}

// Calculate the average speed of the previous period when use GetSpeed()
type AverageSpeedCalculator struct {
	flow               uint64
	unit               uint64
	calculateTimeStamp int64
}

func (asc *AverageSpeedCalculator) GetSpeed() uint64 {
	for { // Spin
		timeStamp := atomic.LoadInt64(&asc.calculateTimeStamp)
		nowTimeStamp := time.Now().UnixNano()
		subTimeStamp := nowTimeStamp - timeStamp
		if atomic.CompareAndSwapInt64(&asc.calculateTimeStamp, timeStamp, nowTimeStamp) {
			return uint64(float64(asc.flow) * (float64(asc.unit) / float64(subTimeStamp)))
		}
		runtime.Gosched()
	}
}

func (asc *AverageSpeedCalculator) AddFlow(n uint64) {
	atomic.AddUint64(&asc.flow, n)
}

func NewAverageSpeedCalculator(unit time.Duration) *AverageSpeedCalculator {
	return &AverageSpeedCalculator{
		0,
		uint64(unit.Nanoseconds()),
		time.Now().UnixNano(),
	}
}
