package pointer

func To[T any](v T) *T {
	return &v
}

func Deref[T any](p *T, def T) T {
	if p == nil {
		return def
	}
	return *p
}
