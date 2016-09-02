package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

type Context struct {
	Namespace
	entities map[int]Entity
}

func NewContext() *Context {
	return &Context{entities: make(map[int]Entity)}
}

func (c *Context) CreateEntity(id int, r bit.Reader) {
}

func (c *Context) GetEntity(id int) Entity {
	return Entity{}
}

func (c *Context) DeleteEntity(id int) {
}

func (c *Context) LeaveEntity(id int) {

}
