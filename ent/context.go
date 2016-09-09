package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

type Context struct {
	*Namespace
	entities map[int]Entity
}

func NewContext() *Context {
	return &Context{Namespace: new(Namespace), entities: make(map[int]Entity)}
}

func (c *Context) CreateEntity(id int, r bit.Reader) error {
	classId := int(c.readClassId(r))
	if len(c.Namespace.classes) == 0 {
		return fmt.Errorf("unable to create entity %d: namespace has no classes", id)
	}
	r.ReadBits(17) // ???
	classV := int(bit.ReadVarInt(r))
	className := c.classIds[classId]
	class := c.Class(className, classV)
	if class == nil {
		return fmt.Errorf("unable to create entity %d: no class found for class name %s, version %d", className, classV)
	}
	Debug.Printf("create entity id: %d classId: %d className: %v class: %v\n", id, classId, className, class)
	e := class.New()
	e.Read(r)
	return nil
}

func (c *Context) UpdateEntity(id int, r bit.Reader) {
	Debug.Printf("update entity id: %d\n", id)
}

func (c *Context) DeleteEntity(id int) {
	Debug.Printf("delete entity id: %d\n", id)
}

func (c *Context) LeaveEntity(id int) {
	Debug.Printf("leave entity id: %d\n", id)
}
