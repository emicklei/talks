type StatsdAccess interface {
	Inc(stat string, value int64, rate float32) error
	Gauge(stat string, value int64, rate float32) error
	Timing(stat string, delta int64, rate float32) error
}
