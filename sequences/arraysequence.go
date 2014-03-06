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


var ArraySequence MutableSequenceClass = &arraySequenceClass{}

var ImmutableArraySequence SequenceClass = ImmutableSequence(ArraySequence)

type arraySequenceClass struct {
}

func (this *arraySequenceClass) Embed(obj MutableSequence) MutableSequence {
  res := new(arraySequence)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.MutableSequenceDerived = EmbeddedMutableSequence(obj)
  res.elements = NewArray()
  return res
}

func (this *arraySequenceClass) New(elements... interface{}) MutableSequence {
  res := this.Embed(nil)
  res.Append(elements...)
  return res
}

func (this *arraySequenceClass) From(coll Container) MutableSequence {
  res := this.Embed(nil)
  res.AppendFrom(coll)
  return res
}

type arraySequence struct {
  obj MutableSequence
  MutableSequenceDerived
  elements Array
}

func (this *arraySequence) Size() int {
  return this.elements.Length()
}

func (this *arraySequence) At(index int) interface{} {
  return this.elements.At(index)
}

func (this *arraySequence) Set(index int, element interface{}) {
  this.elements.Set(index, element)
}

func (this *arraySequence) Allocate(index int, n int, element interface{}) {
  this.elements.Allocate(index, n)
  for i := index; i < index + n; i++ {
    this.elements.Set(i, element)
  }
}

func (this *arraySequence) Delete(index int, n int) {
  this.elements.Delete(index, n)
}

func (this *arraySequence) Class() MutableSequenceClass {
  return ArraySequence
}
