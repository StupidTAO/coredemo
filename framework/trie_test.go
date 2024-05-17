package framework

import (
	"testing"
)

func Test_filterChildNodes(t *testing.T) {
	root := &node{
		isLast:   false,
		segment:  "",
		handlers: []ControllerHandler{func(c *Context) error { return nil }},
		childs: []*node{
			{
				isLast:   true,
				segment:  "FOO",
				handlers: []ControllerHandler{func(c *Context) error { return nil }},
				childs:   nil,
			},
			{
				isLast:   false,
				segment:  ":id",
				handlers: nil,
				childs:   nil,
			},
		},
	}
	{
		nodes := root.filterChildNodes("FOO")
		if len(nodes) != 2 {
			t.Error("foo error")
		}
	}

	{
		nodes := root.filterChildNodes(":foo")
		if len(nodes) != 2 {
			t.Error(":foo error")
		}
	}
}

func Test_matchNode(t *testing.T) {
	root := &node{
		isLast:   false,
		segment:  "",
		handlers: []ControllerHandler{func(c *Context) error { return nil }},
		childs: []*node{
			{
				isLast:   true,
				segment:  "FOO",
				handlers: nil,
				childs: []*node{
					&node{
						isLast:   true,
						segment:  "BAR",
						handlers: []ControllerHandler{func(c *Context) error { panic("not implemented") }},
						childs:   []*node{},
					},
				},
			},
			{
				isLast:   true,
				segment:  ":id",
				handlers: nil,
				childs:   nil,
			},
		},
	}

	{
		node := root.matchNode("foo/bar")
		if node == nil {
			t.Error("match normal node error")
		}
	}

	{
		node := root.matchNode("test")
		if node == nil {
			t.Error("match test")
		}
	}
}

func TestTree_AddRouter(t *testing.T) {
	tree := NewTree()
	err := tree.AddRouter("/cao/hai/tao", []ControllerHandler{func(c *Context) error { return nil }})
	if err != nil {
		t.Error("/cao/hai/tao add router failed")
	}

	err = tree.AddRouter("/cao/xian/bao", []ControllerHandler{func(c *Context) error { return nil }})
	if err != nil {
		t.Error("/cao/xian/bao add router failed")
	}

	err = tree.AddRouter("/cao/xian/bao", []ControllerHandler{func(c *Context) error { return nil }})
	if err != nil {
		t.Log("/cao/xian/bao add router failed ", err.Error())
	}

	var handler []ControllerHandler
	handler = tree.FindHandler("/cao/hai/tao")
	if handler == nil {
		t.Error("match cao/hai/tao failed")
	}

	handler = tree.FindHandler("/cao/xian/bao")
	if handler == nil {
		t.Error("match cao/xian/bao failed")
	}

	handler = tree.FindHandler("cao/xian/zhong")
	if handler == nil {
		t.Log("match cao/xian/zhong failed")
	}
}
