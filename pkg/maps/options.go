package maps

type MapOptions[KEY int | string, VALUE string | float64 | int] map[KEY]VALUE

type Options[T int | string, K string | float64 | int] struct {
	keys    []T
	options MapOptions[T, K]
}

// NewOptions new map
func NewOptions[M MapOptions[K, V], K int | string, V string | float64 | int](m M) *Options[K, V] {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return &Options[K, V]{
		options: MapOptions[K, V](m),
		keys:    r,
	}
}

// Keys get map all key
func (o *Options[T, K]) Keys() []T {
	return o.keys
}

// Option get hash key
func (o *Options[T, K]) Option(key T) K {
	return o.options[key]
}

func (o *Options[T, K]) Options() MapOptions[T, K] {
	return o.options
}
