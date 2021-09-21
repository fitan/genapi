package core

import (
	"cmdb/pkg/log"
)

type Register interface {
	Set(c *Core)
	Unset(c *Core)
}

type Option func(c *Core)

func WithInitRegister(rs ...Register) Option {
	return func(c *Core) {
		for _, r := range rs {
			c.Register(r)
		}
	}
}

type Core struct {
	baseLog *log.Xlog

	objRegister []Register

	Log *log.Xlog

	TraceLog *log.TraceLog

	Api *Api

	Gin *Gin

	Storage

	releaseFn func(x interface{})
}

func (c *Core) Register(register Register) {
	c.objRegister = append(c.objRegister, register)
}

func (c *Core) Init() {
	for _, m := range c.objRegister {
		m.Set(c)
	}
}

func (c *Core) Release() {
	for _, m := range c.objRegister {
		m.Unset(c)
	}

	c.releaseFn(c)
}
