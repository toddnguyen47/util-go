package getrefresh

import "time"

// It's a bit difficult to make this generic, so you can just base your code on this example.

type SampleInterfaceGetter interface {
	Get() SampleInterface
}

type sampleInterfaceGetterImpl struct {
	item          SampleInterface
	lastRefreshed time.Time
	refreshRate   time.Duration
}

func NewSampleInterfaceGetter(refreshRate time.Duration) SampleInterfaceGetter {
	return &sampleInterfaceGetterImpl{
		item:          nil,
		lastRefreshed: time.Now().UTC(),
		refreshRate:   refreshRate,
	}
}

func (s *sampleInterfaceGetterImpl) Get() SampleInterface {
	now := time.Now().UTC()
	if s.item == nil || now.Sub(s.lastRefreshed) >= s.refreshRate {
		item := newSampleStruct()
		s.setItem(item, now)
	}
	return s.item
}

func (s *sampleInterfaceGetterImpl) setItem(item SampleInterface, now time.Time) {
	s.item = item
	s.lastRefreshed = now
}
