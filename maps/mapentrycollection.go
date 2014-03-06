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
import "github.com/objecthub/containerkit/util"


func KV(key, value interface{}) MapEntry {
  return util.NewPair(key, value)
}

func Invert(x interface{}) interface{} {
  if entry, valid := x.(MapEntry); valid {
    return KV(entry.Value(), entry.Key())
  }
  panic("InvertMapEntry expected a MapEntry value")
}

type MapEntry interface {
  Key() interface{}
  Value() interface{}
  Pair() *util.Pair
  String() string
}

type MapEntryContainerBase interface {
  Container
}

type MapEntryContainerDerived interface {
  Keys() DependentContainer
  Values() DependentContainer
}

type MapEntryContainer interface {
  MapEntryContainerBase
  MapEntryContainerDerived
}

type mapEntryContainerTrait struct {
  obj MapEntryContainer
  MapEntryContainerBase
}

func EmbeddedMapEntryContainer(obj MapEntryContainer) MapEntryContainer {
  return &mapEntryContainerTrait{obj, obj}
}

func (this *mapEntryContainerTrait) Keys() DependentContainer {
  return this.obj.Map(func (entry interface{}) interface{} {
    if e, valid := entry.(MapEntry); valid {
      return e.Key()
    }
    panic("mapEntryContainerTrait.Keys: invalid entry")
  })
}

func (this *mapEntryContainerTrait) Values() DependentContainer {
  return this.obj.Map(func (entry interface{}) interface{} {
    if e, valid := entry.(MapEntry); valid {
      return e.Value()
    }
    panic("mapEntryContainerTrait.Values: invalid entry")
  })
}

func MapEntryPredicate(keyPred Predicate, valuePred Predicate) Predicate {
  return func (x interface{}) bool {
    if entry, valid := x.(MapEntry); valid {
      return keyPred(entry.Key()) && valuePred(entry.Value())
    }
    return false
  }
}

func KeyPredicate(pred Predicate) Predicate {
  return MapEntryPredicate(pred, util.TruePredicate)
}

func ValuePredicate(pred Predicate) Predicate {
  return MapEntryPredicate(util.TruePredicate, pred)
}

func ValueMapping(f Mapping) Mapping {
  return func (x interface{}) interface{} {
    if entry, valid := x.(MapEntry); valid {
      return KV(entry.Key(), f(entry.Value()))
    }
    return nil
  }
}
