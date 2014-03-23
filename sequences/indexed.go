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


type Indexed interface {
  IndexedBase
  IndexedDerived
}

type IndexedBase interface {
  Finite
  At(index int) interface{}
}

type IndexedDerived interface {
  First() interface{}
  Last() interface{}
  NextIndexEq(start int, val interface{}) int
  NextIndex(start int, pred func (interface{}) bool) int
  Func() func (int) interface{}
  NativeMap() map[int] interface{}
  Array() []interface{}
}

func EmbeddedIndexed(obj Indexed) Indexed {
  return &indexed{obj, obj}
}

type indexed struct {
  obj Indexed
  IndexedBase
}

func (this *indexed) First() interface{} {
  return this.obj.At(0)
}

func (this *indexed) Last() interface{} {
  return this.obj.At(this.obj.Size() - 1)
}

func (this *indexed) NextIndexEq(start int, val interface{}) int {
  return this.obj.NextIndex(start, func (elem interface{}) bool {
    return elem == val
  })
}

func (this *indexed) NextIndex(start int, pred func (interface{}) bool) int {
  for i := start; i < this.obj.Size(); i++ {
    if pred(this.obj.At(i)) {
      return i
    }
  }
  return -1
}

func (this *indexed) Func() func (int) interface{} {
  return func (i int) interface{} {
    return this.obj.At(i)
  }
}

func (this *indexed) NativeMap() map[int] interface{} {
  res := make(map[int] interface{})
  for i := 0; i < this.obj.Size(); i++ {
    res[i] = this.obj.At(i)
  }
  return res
}

func (this *indexed) Array() []interface{} {
  n := this.obj.Size()
  res := make([]interface{}, n)
  for i := 0; i < n; i++ {
    res[i] = this.obj.At(i)
  }
  return res
}
