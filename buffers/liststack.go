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

package buffers

import . "github.com/objecthub/containerkit"
import . "github.com/objecthub/containerkit/impl"


var ListStack StackClass = &listStackClass{}

type listStackClass struct {
}

func (this *listStackClass) Embed(obj Stack) Stack {
  res := new(listStack)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.StackDerived = EmbeddedStack(obj)
  res.list = nil
  res.n = 0
  return res
}

func (this *listStackClass) New(elements... interface{}) Stack {
  res := this.Embed(nil)
  for i := 0; i < len(elements); i++ {
    res.Push(elements[i])
  }
  return res
}

func (this *listStackClass) From(coll Container) Stack {
  res := this.Embed(nil)
  res.PushFrom(coll)
  return res
}

type listStack struct {
  obj Stack
  StackDerived
  list *Cons
  n int
}

func (this *listStack) Size() int {
  return this.n
}

func (this *listStack) Elements() Iterator {
  return this.list.Iterator()
}

func (this *listStack) Push(elem interface{}) {
  this.list = NewCons(elem, this.list)
  this.n++
}

func (this *listStack) Pop() interface{} {
  if this.list == nil {
    panic("listStack: stack empty")
  }
  res := this.list.Head
  this.list = this.list.Tail
  this.n--
  return res
}

func (this *listStack) Peek() interface{} {
  if this.list == nil {
    panic("listStack: stack empty")
  }
  return this.list.Head
}

func (this *listStack) Clear() {
  this.list = nil
  this.n = 0
}

func (this *listStack) Class() StackClass {
  return ListStack
}
