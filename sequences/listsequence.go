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

package sequences

import . "github.com/objecthub/containerkit"
import . "github.com/objecthub/containerkit/impl"


var ListSequence MutableSequenceClass = &listSequenceClass{}

var ImmutableListSequence SequenceClass = ImmutableSequence(ListSequence)

type listSequenceClass struct {
}

func (this *listSequenceClass) Embed(obj MutableSequence) MutableSequence {
  res := new(listSequence)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.MutableSequenceDerived = EmbeddedMutableSequence(obj)
  res.list = nil
  res.n = 0
  return res
}

func (this *listSequenceClass) New(elements... interface{}) MutableSequence {
  res := this.Embed(nil)
  res.Append(elements...)
  return res
}

func (this *listSequenceClass) From(coll Container) MutableSequence {
  res := this.Embed(nil)
  res.AppendFrom(coll)
  return res
}

type listSequence struct {
  obj MutableSequence
  MutableSequenceDerived
  list *Cons
  n int
}

func (this *listSequence) Size() int {
  return this.n
}

func (this *listSequence) At(index int) interface{} {
  return this.list.Skip(index).Head
}

func (this *listSequence) Set(index int, element interface{}) {
  this.list.Skip(index).Head = element
}

func (this *listSequence) Allocate(index int, n int, element interface{}) {
  if n > 0 {
    var anchor *Cons = nil
    var tail *Cons = this.list
    if (index > 0) {
      anchor = this.list.Skip(index - 1)
      tail = anchor.Tail
    }
    newlist := NewCons(element, tail)
    for i := 1; i < n; i++ {
      newlist = NewCons(element, newlist)
    }
    if (anchor == nil) {
      this.list = newlist
    } else {
      anchor.Tail = newlist
    }
    this.n += n
  } else if n < 0 {
    panic("listSequence.Allocate: negative number of elements to insert")
  }
}

func (this *listSequence) Delete(index int, n int) {
  if index == 0 {
    for i := 0; i < n; i++ {
      if this.list == nil {
        return
      }
      this.list = this.list.Tail
      this.n--
    }
  } else {
    list := this.list.Skip(index - 1)
    for i := 0; i < n; i++ {
      if list == nil {
        return
      }
      list.Tail = list.Tail.Tail
      this.n--
    }
  }
}

func (this *listSequence) Class() MutableSequenceClass {
  return ListSequence
}
