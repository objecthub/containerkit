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


var ArrayBuffer *arrayBufferClass = &arrayBufferClass{}

type arrayBufferClass struct {}

func (this *arrayBufferClass) Embed(obj Buffer) Buffer {
  res := new(arrayBuffer)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.BufferDerived = EmbeddedBuffer(obj)
  res.appended = NewArray()
  return res
}

func (this *arrayBufferClass) New(elements... interface{}) Buffer {
  res := this.Embed(nil)
  res.Append(elements...)
  return res
}

func (this *arrayBufferClass) From(coll Container) Buffer {
  res := this.Embed(nil)
  res.AppendFrom(coll)
  return res
}

type arrayBuffer struct {
  obj Buffer
  BufferDerived
  appended Array
}

func (this *arrayBuffer) Size() int {
  return this.appended.Length()
}

func (this *arrayBuffer) At(index int) interface{} {
  return this.appended.At(index)
}

func (this *arrayBuffer) Append(elems ...interface{}) Buffer {
  l := this.appended.Length()
  this.appended.Allocate(l, len(elems))
  for i := len(elems) - 1; i >= 0; i-- {
    this.appended.Set(l + i, elems[i])
  }
  return this
}

func (this *arrayBuffer) Class() BufferClass {
  return ArrayBuffer
}

func (this *arrayBuffer) Clear() Buffer {
  this.appended = NewArray()
  return this
}

func (this *arrayBuffer) Array() []interface{} {
  res := make([]interface{}, this.appended.Length())
  copy(res, this.appended)
  return res
}
