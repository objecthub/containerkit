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


var ListBuffer *listBufferClass = &listBufferClass{}

type listBufferClass struct {}

func (this *listBufferClass) Embed(obj Buffer) Buffer {
  res := new(listBuffer)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.BufferDerived = EmbeddedBuffer(obj)
  res.list = NewLinkedList()
  return res
}

func (this *listBufferClass) New(elements... interface{}) Buffer {
  res := this.Embed(nil)
  res.Append(elements...)
  return res
}

func (this *listBufferClass) From(coll Container) Buffer {
  res := this.Embed(nil)
  res.AppendFrom(coll)
  return res
}

type listBuffer struct {
  obj Buffer
  BufferDerived
  list *LinkedList
}

func (this *listBuffer) Size() int {
  return this.list.Length()
}

func (this *listBuffer) At(index int) interface{} {
  return this.list.At(index)
}

func (this *listBuffer) Append(elems ...interface{}) Buffer {
  for _, elem := range elems {
    this.list.InsertTail(elem)
  }
  return this
}

func (this *listBuffer) Class() BufferClass {
  return ListBuffer
}

func (this *listBuffer) Clear() Buffer {
  this.list = NewLinkedList()
  return this
}

func (this *listBuffer) Elements() Iterator {
  return this.list.Iterator()
}
