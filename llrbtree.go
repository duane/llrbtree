package llrbtree

import (
  "errors"
)

type Tree struct {
  Root *Node
}

type Key interface {
  Equals(other interface{}) bool
  Less(other interface{}) bool
  String() string
}

type Color int

var KeyNotFoundError = errors.New("Key not found")

const (
  RED Color = iota
  BLACK
)

type Node struct {
  Key         Key
  Value       interface{}
  Left, Right *Node
  Color       Color
}

func (node *Node) insert(key Key, value interface{}) *Node {
  if node == nil {
    return NewNode(key, value)
  }

  if key.Equals(node.Key) {
    node.Value = value
  } else if key.Less(node.Key) {
    node.Left = node.Left.insert(key, value)
  } else {
    node.Right = node.Right.insert(key, value)
  }

  return node
}

func NewNode(key Key, value interface{}) *Node {
  return &Node{Key: key, Value: value, Color: RED}
}

func NewTree() *Tree {
  return &Tree{}
}

func (t *Tree) Search(key Key) (value interface{}) {
  x := t.Root
  for x != nil {
    if x.Key.Equals(key) {
      return x.Value
    } else if key.Less(x.Key) {
      x = x.Left
    } else {
      x = x.Right
    }
  }
  return
}

func (t *Tree) Insert(key Key, value interface{}) {
  t.Root = t.Root.insert(key, value)
}

func (t *Tree) Put(key Key, value []byte) (err error) {
  t.Insert(key, value)
  return
}

func (t *Tree) Get(key Key) (value []byte, err error) {
  found := t.Search(key)
  if found == nil {
    err = KeyNotFoundError
    return
  }
  value = found.([]byte)
  return
}
