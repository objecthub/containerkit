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


type SetBase interface {
  FiniteContainerBase
  CollectionBase
}

type SetDerived interface {
  FiniteContainerDerived
  CollectionDerived
  ReadOnly() DependentSet
  Union(set Set) DependentSet
  Intersection(set Set) DependentSet
  Difference(set Set) DependentSet
}

type Set interface {
  SetBase
  SetDerived
}

type SetClass interface {
  Embed(obj Set) Set
  New(elements ...interface{}) Set
  From(coll Container) Set
}

func EmbeddedSet(obj Set) Set {
  return &setTrait{obj,
                   obj,
                   EmbeddedFiniteContainer(obj),
                   EmbeddedCollection(obj)}
}

type setTrait struct {
  obj Set
  SetBase
  FiniteContainerDerived
  CollectionDerived
}

func (this *setTrait) ReadOnly() DependentSet {
  return newWrappedSet(this.obj, false)
}

func (this *setTrait) Union(set Set) DependentSet {
  return newUnionSet(this.obj, set)
}

func (this *setTrait) Intersection(set Set) DependentSet {
  return newIntersectionSet(this.obj, set)
}

func (this *setTrait) Difference(set Set) DependentSet {
  return newDifferenceSet(this.obj, set)
}

func (this *setTrait) String() string {
  return "{" + this.FiniteContainerDerived.String() + "}"
}
