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


type MutableSetObserver interface {
  Include(subject MutableSet, elem interface{})
  Exclude(subject MutableSet, elem interface{})
}

func ObservableSet(class MutableSetClass, observers Container) MutableSetClass {
  return &observableSetClass{class, observers}
}

type observableSetClass struct {
  class MutableSetClass
  observers Container
}

func (this *observableSetClass) Embed(obj MutableSet) MutableSet {
  res := new(observableSet)
  if obj == nil {
    obj = res
  }
  res.obj = obj
  res.MutableSet = this.class.Embed(obj)
  return res
}

func (this *observableSetClass) New(elements... interface{}) MutableSet {
  res := this.Embed(nil)
  res.Include(elements...)
  return res
}

func (this *observableSetClass) From(coll Container) MutableSet {
  res := this.Embed(nil)
  res.IncludeFrom(coll)
  return res
}

type observableSet struct {
  obj MutableSet
  observers Container
  MutableSet
}

func (this *observableSet) Include(elements... interface{}) {
  this.MutableSet.Include(elements...)
  this.observers.ForEach(func (o interface{}) {
    for _, e := range elements {
      o.(MutableSetObserver).Include(this.obj, e)
    }
  })
}

func (this *observableSet) Exclude(elements... interface{}) {
  this.MutableSet.Exclude(elements...)
  this.observers.ForEach(func (o interface{}) {
    for _, e := range elements {
      o.(MutableSetObserver).Exclude(this.obj, e)
    }
  })
}
