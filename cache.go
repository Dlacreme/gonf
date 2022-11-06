package gonf

// Cached loaded key/value
var cache map[string]string = make(map[string]string)

func pushCache(key string, value string) {
	cache[key] = value
}
