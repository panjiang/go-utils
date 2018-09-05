package sqls

import (
	"strings"
)

// QueryWhere 数据库查询条件拼接对象
type QueryWhere struct {
	exprs []string
	args  []interface{}
}

// Add 增加一个条件
func (w *QueryWhere) Add(expr string, arg interface{}) {
	w.exprs = append(w.exprs, expr)
	w.args = append(w.args, arg)
}

// Expr 获取表达式
func (w *QueryWhere) Expr() string {
	return strings.Join(w.exprs, " AND ")
}

// Args 获取参数
func (w *QueryWhere) Args() []interface{} {
	return w.args
}
