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


var ArrayStack StackClass = &arrayStackClass{}

type arrayStackClass struct {}

func (this *arrayStackClass) Embed(obj Stack) Stack {
  res := new(arrayStack)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.StackDerived = EmbeddedStack(obj)
  res.array = NewArray()
  return res
}

func (this *arrayStackClass) New(elements... interface{}) Stack {
  res := this.Embed(nil)
  for i := 0; i < len(elements); i++ {
    res.Push(elements[i])
  }
  return res
}

func (this *arrayStackClass) From(coll Container) Stack {
  res := this.Embed(nil)
  res.PushFrom(coll)
  return res
}

type arrayStack struct {
  obj Stack
  StackDerived
  array Array
}

func (this *arrayStack) Size() int {
  return this.array.Length()
}

func (this *arrayStack) Elements() Iterator {
  return this.array.ReverseIterator()
}

func (this *arrayStack) Push(elem interface{}) {
  this.array.Append(elem)
}

func (this *arrayStack) Pop() interface{} {
  res := this.Peek()
  this.array.Remove(1)
  return res
}

func (this *arrayStack) Peek() interface{} {
  if this.array.Length() == 0 {
    panic("arrayStack: stack empty")
  }
  return this.array.At(this.array.Length() - 1)
}

func (this *arrayStack) Clear() {
  this.array.Clear()
}

func (this *arrayStack) Class() StackClass {
  return ArrayStack
}
