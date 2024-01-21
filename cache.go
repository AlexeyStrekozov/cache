package cache

type Cache struct {
	cache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cache: map[string]interface{}{},
	}
}

func (s *Cache) Set(key string, value interface{}) {
	s.cache[key] = value
}

func (s *Cache) Get(key string) interface{} {
	res, exist := s.cache[key]

	if exist {
		return res
	}

	return nil
}

func (s *Cache) Delete(key string) {
	_, exist := s.cache[key]

	if exist {
		delete(s.cache, key)
	}
}
