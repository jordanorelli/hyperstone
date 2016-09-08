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
	classId := int(c.readClassId(r))
	r.ReadBits(17) // ???
	classV := int(bit.ReadVarInt(r))
	className := c.classIds[classId]
	class := c.Class(className, classV)
	e := class.New()
	e.Read(r)
}

func (c *Context) UpdateEntity(id int, r bit.Reader) {

}

func (c *Context) DeleteEntity(id int) {

}

func (c *Context) LeaveEntity(id int) {

}
