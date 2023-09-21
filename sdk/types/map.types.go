package types

type DynamicMap map[string]any

func (f DynamicMap) Get(key string) any {
	if f == nil {
		return ""
	}
	return f[key]
}

func (f DynamicMap) Add(key string, value any) {
	f[key] = value
}

func (f DynamicMap) Del(key string) {
	delete(f, key)
}

func (f DynamicMap) Has(key string) bool {
	_, ok := f[key]
	return ok
}

func (f DynamicMap) Count() int {
	count := 0
	for range f {
		count++
	}
	return count
}
