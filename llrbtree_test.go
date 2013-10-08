package llrbtree

import (
  "testing"
)

type StringKey string

func (k *StringKey) Equals(other interface{}) bool {
  return string(*k) == string(*(other.(*StringKey)))
}

func (k *StringKey) Less(other interface{}) bool {
  return string(*k) < string(*(other.(*StringKey)))
}

func (k *StringKey) String() string {
  return string(*k)
}

func TestPut(t *testing.T) {
  tree := NewTree()
  key_str := StringKey("key")
  err := tree.Put(&key_str, []byte("value"))
  if err != nil {
    t.Fatal(err)
  }
}

func TestPutGet(t *testing.T) {
  tree := NewTree()
  key_str := StringKey("one")
  err := tree.Put(&key_str, []byte{1})
  if err != nil {
    t.Fatal(err)
  }

  value, err := tree.Get(&key_str)
  if err != nil {
    t.Fatal(err)
  }

  if len(value) != 1 || value[0] != 1 {
    t.Fail()
  }
}
