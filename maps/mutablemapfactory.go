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


type MutableMapFactoryBase interface {
  Class() MutableMapClass
}

type MutableMapFactoryContext interface {
  Map
  MutableMapFactoryBase
}

type MutableMapFactoryDerived interface {
  Copy() MutableMap
  Immutable() Map
  ProjectValues(proj Mapping) MutableMap
  FilterKeys(pred Predicate) MutableMap
}

type MutableMapFactory interface {
  MutableMapFactoryContext
  MutableMapFactoryDerived
}

func EmbeddedMutableMapFactory(obj MutableMapFactory) MutableMapFactory {
  return &mapFactory{obj, obj}
}

type mapFactory struct {
  obj MutableMapFactory
  MutableMapFactoryContext
}

func (this *mapFactory) Copy() MutableMap {
  return this.obj.Class().From(this.obj)
}

func (this *mapFactory) Immutable() Map {
  return wrappedMap(this.obj.Copy(), true)
}

func (this *mapFactory) ProjectValues(proj Mapping) MutableMap {
  res := this.obj.Class().New()
  for iter := this.obj.Elements(); iter.HasNext(); {
    if entry, valid := iter.Next().(MapEntry); valid {
      res.Include(entry.Key(), proj(entry.Value()))
    } else {
      panic("mapFactory.ProjectValues: element not a MapEntry")
    }
  }
  return res
}

func (this *mapFactory) FilterKeys(pred Predicate) MutableMap {
  res := this.obj.Class().New()
  for iter := this.obj.Elements(); iter.HasNext(); {
    if entry, valid := iter.Next().(MapEntry); valid {
      if pred(entry.Key()) {
        res.Include(entry.Key(), entry.Value())
      }
    } else {
      panic("mapFactory.ProjectValues: element not a MapEntry")
    }
  }
  return res
}
