package core

import (
	"cmdb/pkg/log"
	"cmdb/pkg/trace"
	"go.uber.org/zap"
)

var xlog *log.Xlog

func NewXlog() *log.Xlog {
	if xlog != nil {
		return xlog
	}
	return log.NewXlog(log.WithTrace(trace.GetTp(), zap.InfoLevel))
}

type logRegister struct {
}

func (l logRegister) Set(c *Core) {
	c.Log = NewXlog()
}

func (l logRegister) Unset(c *Core) {
	c.Log = nil
}
