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

package containerkit


type Enumeration interface {
  FiniteContainer
}

var Enum enumClass = enumClass{newEnum()}

type enumClass struct {
  empty FiniteContainer
}

func (this *enumClass) Empty() FiniteContainer {
  return this.empty
}

func (this *enumClass) New(elements... interface{}) FiniteContainer {
  res := newEnum()
  if len(elements) > 0 {
    res.elements = make([]interface{}, len(elements))
    copy(res.elements, elements)
  }
  return res
}

func (this *enumClass) From(container Container) FiniteContainer {
  res := newEnum()
  for iter := container.Elements(); iter.HasNext(); {
    res.elements = append(res.elements, iter.Next())
  }
  return res
}

func (this *enumClass) Range(start, end int) FiniteContainer {
  res := newEnum()
  for i := start; i <= end; i++ {
    res.elements = append(res.elements, i)
  }
  return res
}

func newEnum() *enum {
  res := new(enum)
  res.FiniteContainerDerived = EmbeddedContainer(res)
  return res
}

type enum struct {
  FiniteContainerDerived
  elements []interface{}
}

func (this *enum) Size() int {
  return len(this.elements)
}

func (this *enum) Elements() Iterator {
  return &enumIterator{this.elements, 0}
}

type enumIterator struct {
  a []interface{}
  i int
}

func (this *enumIterator) HasNext() bool {
  return this.i < len(this.a)
}

func (this *enumIterator) Next() interface{} {
  if this.HasNext() {
    res := this.a[this.i]
    this.i++
    return res
  }
  panic("enumIterator.Next: no next element")
}
