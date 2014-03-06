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
import "github.com/objecthub/containerkit/util"


type DependentMap interface {
  Map
}

func EmbeddedDependentMap(obj DependentMap) MapDerived {
  return &dependentMapTrait{EmbeddedMap(obj)}
}

type dependentMapTrait struct {
  MapDerived
}

func (this *dependentMapTrait) String() string {
  return "<" + this.MapDerived.String() + ">"
}

// Map restriction

func newRestrictedMap(mp Map, domain Set) DependentMap {
  res := new(restrictedMap)
  res.MapDerived = EmbeddedDependentMap(res)
  res.fst = mp
  res.domain = domain
  return res
}

type restrictedMap struct {
  MapDerived
  fst Map
  domain Set
}

func (this *restrictedMap) Size() int {
  return CountElements(this.Elements())
}

func (this *restrictedMap) Get(key interface{}) (value interface{}, exists bool) {
  if this.domain.Contains(key) {
    return this.fst.Get(key)
  }
  return nil, false
}

func (this *restrictedMap) Elements() Iterator {
  return NewFilterIterator(KeyPredicate(this.domain.Func()),
                           this.fst.Elements())
}

// Map extension

func newExtendedMap(base Map, overrides Map) DependentMap {
  res := new(extendedMap)
  res.MapDerived = EmbeddedDependentMap(res)
  res.base = base
  res.overrides = overrides
  return res
}

type extendedMap struct {
  MapDerived
  base Map
  overrides Map
}

func (this *extendedMap) Size() int {
  return CountElements(this.Elements())
}

func (this *extendedMap) Get(key interface{}) (value interface{}, exists bool) {
  if val, exists := this.overrides.Get(key); exists {
    return val, true
  }
  return this.base.Get(key)
}

func (this *extendedMap) Elements() Iterator {
  return NewCompositeIterator(
      this.overrides.Elements(),
      NewFilterIterator(KeyPredicate(util.Negate(this.overrides.KeySet().Func())),
                        this.base.Elements()))
}

// Map value mapping

func newValueMappedMap(mp Map, f Mapping) DependentMap {
  res := new(valueMappedMap)
  res.MapDerived = EmbeddedDependentMap(res)
  res.fst = mp
  res.f = f
  return res
}

type valueMappedMap struct {
  MapDerived
  fst Map
  f Mapping
}

func (this *valueMappedMap) Size() int {
  return this.fst.Size()
}

func (this *valueMappedMap) Get(key interface{}) (value interface{}, exists bool) {
  if val, exists := this.fst.Get(key); exists {
    return this.f(val), true
  }
  return nil, false
}

func (this *valueMappedMap) Elements() Iterator {
  return NewMappedIterator(ValueMapping(this.f), this.fst.Elements())
}
