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

package maps

import . "github.com/objecthub/containerkit"
import . "github.com/objecthub/containerkit/sets"


type Map interface {
  MapBase
  MapDerived
}

type MapBase interface {
  FiniteContainerBase
  MapperBase
}

type MapDerived interface {
  FiniteContainerDerived
  MapEntryContainerDerived
  MapperDerived
  KeySet() DependentSet
  ReadOnly() DependentMap
  RestrictTo(domain Set) DependentMap
  MapValues(f Mapping) DependentMap
  Override(base Map) DependentMap
}

// MapClass defines the functionality of Map implementations,
// ie. records that act as Map factories, providing an Embed, New,
// and From method.
type MapClass interface {
  Embed(obj Map) Map
  New(entries... MapEntry) Map
  From(coll Container) Map
}

func EmbeddedMap(obj Map) Map {
  return &mapTrait{obj,
                   obj,
                   EmbeddedFiniteContainer(obj),
                   EmbeddedMapEntryContainer(obj),
                   EmbeddedMapper(obj)}
}

type mapTrait struct {
  obj Map
  MapBase
  FiniteContainerDerived
  MapEntryContainerDerived
  MapperDerived
}

func (this *mapTrait) ReadOnly() DependentMap {
  return wrappedMap(this.obj, false)
}

func (this *mapTrait) RestrictTo(domain Set) DependentMap {
  return newRestrictedMap(this.obj, domain)
}

func (this *mapTrait) MapValues(f Mapping) DependentMap {
  return newValueMappedMap(this.obj, f)
}

func (this *mapTrait) Override(base Map) DependentMap {
  return newExtendedMap(base, this.obj)
}

func (this *mapTrait) KeySet() DependentSet {
  res := new(keySet)
  res.SetDerived = EmbeddedDependentSet(res)
  res.parent = this.obj
  return res
}

func (this *mapTrait) String() string {
  return "{" + this.FiniteContainerDerived.String() + "}"
}

func wrappedMap(encapsulated Map, immutable bool) DependentMap {
  res := new(mapWrapper)
  res.MapDerived = EmbeddedDependentMap(res)
  res.encapsulated = encapsulated
  res.immutable = immutable
  return res
}

type mapWrapper struct {
  MapDerived
  encapsulated Map
  immutable bool
}

func (this *mapWrapper) Size() int {
  return this.encapsulated.Size()
}

func (this *mapWrapper) Get(key interface{}) (value interface{}, exists bool) {
  return this.encapsulated.Get(key)
}

func (this *mapWrapper) Elements() Iterator {
  return this.encapsulated.Elements()
}

func (this *mapWrapper) ReadOnly() DependentMap {
  return this
}

func (this *mapWrapper) Freeze() FiniteContainer {
  if this.immutable {
    return this
  }
  return this.MapDerived.Freeze()
}

type keySet struct {
  SetDerived
  parent Map
}

func (this *keySet) Size() int {
  return this.parent.Size()
}

func (this *keySet) Contains(elem interface{}) bool {
  return this.parent.HasKey(elem)
}

func (this *keySet) Elements() Iterator {
  return NewMappedIterator(func (entry interface{}) interface{} {
    if e, valid := entry.(MapEntry); valid {
      return e.Key()
    }
    panic("mapEntryContainerTrait.Keys: invalid entry")
  }, this.parent.Elements())
}
