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


func ImmutableSet(class MutableSetClass) SetClass {
  return &immutableSetClass{class}
}

type immutableSetClass struct {
  class MutableSetClass
}

func (this *immutableSetClass) newImmutableSet() (Set, MutableSet) {
  encapsulated := this.class.New()
  res := new(immutableSet)
  res.Set = encapsulated
  return res, encapsulated
}

func (this *immutableSetClass) Embed(obj Set) Set {
  res, _ := this.newImmutableSet()
  return res
}

func (this *immutableSetClass) New(elements ...interface{}) Set {
  res, encapsulated := this.newImmutableSet()
  encapsulated.Include(elements...)
  return res
}

func (this *immutableSetClass) From(coll Container) Set {
  res, encapsulated := this.newImmutableSet()
  encapsulated.IncludeFrom(coll)
  return res
}

type immutableSet struct {
  Set
}

func (this *immutableSet) ReadOnly() DependentSet {
  return this
}

func (this *immutableSet) Freeze() FiniteContainer {
  return this
}
