package com

import (
	"fmt"
)

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(component Component)
	Print(string)
}

const (
	LeafNode      = iota
	CompositeNode
)

//组合模式的简单工厂，可以看成根据参数决定创建目录，或者文件
//返回即可以描述目录又可以描述文件的interface
func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

//这个是共有部分，叶子和根节点都有这个，细节的不同可以重写相关方法函数
type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(component Component) {
}
func (c *component) Print(string) {
}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

//文件，重写一下Print函数，因为它与目录不一样
func (c *Leaf) Print(pre string) {
	fmt.Printf("%s#%s\n", pre, c.Name())
}

type Composite struct {
	component
	childs []Component
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

//目录才会执行AddChild
func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child)
}

//目录的Print除了打印目录名字，还要遍历它下面所有的文件。
func (c *Composite) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}
