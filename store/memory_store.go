package store

// 存储用户token信息
var TokenMemoryStore map[string]bool

func init() {
	TokenMemoryStore = make(map[string]bool)
}
