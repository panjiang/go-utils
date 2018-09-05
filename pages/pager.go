package pages

// sizeMax 单页显示最大条数
const sizeMax int = 10

// Pager 分页
type Pager struct {
	Page int
	size int
}

// Size 单页条数
func (p *Pager) Size() int {
	return p.size
}

// Offset 根据页数计算偏移
func (p *Pager) Offset() int {
	return (p.Page - 1) * p.size
}

// NewPager 创建分页对象
func NewPager(page, size int) *Pager {
	if size <= 0 || size > sizeMax {
		size = sizeMax
	}
	return &Pager{Page: page, size: size}
}
