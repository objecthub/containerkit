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


type CollectionBase interface {
  Contains(elem interface{}) bool
}

type CollectionDerived interface {
  ContainsAll(elements ...interface{}) bool
  ContainsNone(elements ...interface{}) bool
  ContainsSome(elements ...interface{}) bool
  ContainsAllFrom(elements Container) bool
  ContainsNoneFrom(elements Container) bool
  ContainsSomeFrom(elements Container) bool
  Func() Predicate
}

type Collection interface {
  CollectionBase
  CollectionDerived
}

func EmbeddedCollection(obj Collection) Collection {
  return &collection{obj, obj}
}

type collection struct {
  obj Collection
  CollectionBase
}

func (this *collection) ContainsAll(elements ...interface{}) bool {
  for i := 0; i < len(elements); i++ {
    if !this.obj.Contains(elements[i]) {
      return false
    }
  }
  return true
}

func (this *collection) ContainsNone(elements ...interface{}) bool {
  for i := 0; i < len(elements); i++ {
    if this.obj.Contains(elements[i]) {
      return false
    }
  }
  return true
}

func (this *collection) ContainsSome(elements ...interface{}) bool {
  return !this.obj.ContainsNone(elements...)
}

func (this *collection) ContainsAllFrom(elements Container) bool {
  return elements.ForAll(func (elem interface{}) bool {
    return this.obj.Contains(elem)
  })
}

func (this *collection) ContainsNoneFrom(elements Container) bool {
  return elements.ForAll(func (elem interface{}) bool {
    return !this.obj.Contains(elem)
  })
}

func (this *collection) ContainsSomeFrom(elements Container) bool {
  return elements.Exists(func (elem interface{}) bool {
    return this.obj.Contains(elem)
  })
}

func (this *collection) Func() Predicate {
  return func (x interface{}) bool {
    return this.obj.Contains(x)
  }
}
