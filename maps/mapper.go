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

import . "github.com/objecthub/containerkit/util"


type Mapper interface {
  MapperBase
  MapperDerived
}

type MapperBase interface {
  Get(key interface{}) (value interface{}, exists bool)
}

type MapperDerived interface {
  HasKey(key interface{}) bool
  GetValue(key interface{}) interface{}
  GetString(key interface{}) string
  Func() func (interface{}) interface{}
}

func EmbeddedMapper(obj Mapper) Mapper {
  return &mapper{obj, obj}
}

type mapper struct {
  obj Mapper
  MapperBase
}

func (this *mapper) HasKey(key interface{}) bool {
  _, exists := this.obj.Get(key)
  return exists
}

func (this *mapper) GetValue(key interface{}) interface{} {
  if value, exists := this.obj.Get(key); exists {
    return value
  }
  panic("Map.GetValue: mapping does not exist")
}

func (this *mapper) GetString(key interface{}) string {
  return ToString(this.obj.GetValue(key))
}

func (this *mapper) Func() func (interface{}) interface{} {
  return func (key interface{}) interface{} {
    return this.obj.GetValue(key)
  }
}
