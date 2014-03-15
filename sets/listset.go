// Copyright 2014 Matthias Zenger. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sets

import . "github.com/objecthub/containerkit"
import . "github.com/objecthub/containerkit/impl"


var ListSet MutableSetClass = ListSetClass(UniversalEquality)

var ImmutableListSet SetClass = ImmutableSet(ListSet)

func ListSetClass(equals Equality) MutableSetClass {
  return &listSetClass{equals}
}

type listSetClass struct {
  equals Equality
}

func (this *listSetClass) Embed(obj MutableSet) MutableSet {
  res := new(listSet)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.eq = this.equals
  res.list = nil
  res.size = 0
  res.MutableSetDerived = EmbeddedMutableSet(obj)
  return res
}

func (this *listSetClass) New(elements ...interface{}) MutableSet {
  res := this.Embed(nil)
  res.Include(elements...)
  return res
}

func (this *listSetClass) From(coll Container) MutableSet {
  res := this.Embed(nil)
  res.IncludeFrom(coll)
  return res
}

type listSet struct {
  obj MutableSet
  eq Equality
  list *Cons
  size int
  MutableSetDerived
}

func (this *listSet) Size() int {
  return this.size
}

func (this *listSet) Contains(elem interface{}) bool {
  return nil != this.list.Find(func (val interface{}) bool {
    return this.eq(elem, val)
  })
}

func (this *listSet) Elements() Iterator {
  return this.list.Iterator()
}

func (this *listSet) Class() MutableSetClass {
  return ListSetClass(this.eq)
}

func (this *listSet) Include(elements ...interface{}) {
  for i := 0; i < len(elements); i++ {
    if nil == this.list.Find(func (val interface{}) bool {
          return this.eq(elements[i], val)
       }) {
      this.list = NewCons(elements[i], this.list)
      this.size++
    }
  }
}

func (this *listSet) Exclude(elements ...interface{}) {
  for i := 0; i < len(elements); i++ {
    if this.size > 0 {
      if this.eq(this.list.Head, elements[i]) {
        this.list = this.list.Tail
        this.size--
      }
      for list := this.list; list.Tail != nil; {
        if this.eq(list.Tail.Head, elements[i]) {
          list.Tail = list.Tail.Tail
          this.size--
        } else {
          list = list.Tail
        }
      }
    }
  }
}

func (this *listSet) Clear() {
  this.list = nil
  this.size = 0
}
