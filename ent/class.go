package ent

import (
	"fmt"
)

type class struct {
	name    string
	version int
	fields  []field
}

func (c class) String() string { return c.typeName() }

func (c class) typeName() string {
	return fmt.Sprintf("%s.%d", c.name, c.version)
}

func (c *class) nü() value {
	return &entity{class: c, slots: make([]value, len(c.fields))}
}

type classHistory struct {
	versions map[int]*class
	oldest   *class
	newest   *class
}

func (h *classHistory) add(c *class) {
	if h.oldest == nil || c.version < h.oldest.version {
		h.oldest = c
	}
	if h.newest == nil || c.version > h.newest.version {
		h.newest = c
	}
	if h.versions == nil {
		h.versions = make(map[int]*class)
	}
	h.versions[c.version] = c
}

func (h *classHistory) version(v int) *class {
	if h.versions == nil {
		return nil
	}
	return h.versions[v]
}

func classType(spec *typeSpec, env *Env) tÿpe {
	if spec.serializer != "" {
		c := env.classVersion(spec.serializer, spec.serializerV)
		if c != nil {
			return c
		}
		return typeError("unable to find class named %s with version %d", spec.serializer, spec.serializerV)
	}
	return nil
}
