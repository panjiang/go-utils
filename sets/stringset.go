package sets

// StringSet 字符串无序集合
type StringSet struct {
	m map[string]struct{}
}

// Add 添加元素
func (ss StringSet) Add(s string) {
	ss.m[s] = struct{}{}
}

// Contains 包含
func (ss StringSet) Contains(s string) bool {
	_, ok := ss.m[s]
	return ok
}

// Items 获取所有元素
func (ss StringSet) Items() []string {
	sl := make([]string, 0, len(ss.m))
	for k := range ss.m {
		sl = append(sl, k)
	}
	return sl
}

// Size 长度
func (ss StringSet) Size() int {
	return len(ss.m)
}

// NewStringSet 创建
func NewStringSet() *StringSet {
	return &StringSet{
		m: make(map[string]struct{}, 0),
	}
}
