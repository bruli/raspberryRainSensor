package rain

type Rain struct {
	value   int
	raining bool
}

func (r Rain) Value() int {
	return r.value
}

func (r Rain) Raining() bool {
	return r.raining
}

func New(value int, raining bool) Rain {
	return Rain{value: value, raining: raining}
}
