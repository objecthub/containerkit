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
import "github.com/objecthub/containerkit/util"


type DependentSet interface {
  Set
}

func EmbeddedDependentSet(obj DependentSet) SetDerived {
  return &dependentSetTrait{EmbeddedSet(obj)}
}

type dependentSetTrait struct {
  SetDerived
}

func (this *dependentSetTrait) String() string {
  return "<" + this.SetDerived.String() + ">"
}

// Set union

func newUnionSet(fst Set, snd Set) DependentSet {
  res := new(unionSet)
  res.SetDerived = EmbeddedDependentSet(res)
  res.fst = fst
  res.snd = snd
  return res
}

type unionSet struct {
  SetDerived
  fst Set
  snd Set
}

func (this *unionSet) Size() int {
  return CountElements(this.Elements())
}

func (this *unionSet) Contains(elem interface{}) bool {
  return this.fst.Contains(elem) || this.snd.Contains(elem)
}

func (this *unionSet) Elements() Iterator {
  return NewCompositeIterator(
      NewFilterIterator(util.Negate(this.snd.Func()), this.fst.Elements()),
      this.snd.Elements())
}

// Set intersection

func newIntersectionSet(fst Set, snd Set) DependentSet {
  res := new(intersectionSet)
  res.SetDerived = EmbeddedDependentSet(res)
  res.fst = fst
  res.snd = snd
  return res
}

type intersectionSet struct {
  SetDerived
  fst Set
  snd Set
}

func (this *intersectionSet) Size() int {
  return CountElements(this.Elements())
}

func (this *intersectionSet) Contains(elem interface{}) bool {
  return this.fst.Contains(elem) && this.snd.Contains(elem)
}

func (this *intersectionSet) Elements() Iterator {
  return NewFilterIterator(this.snd.Func(), this.fst.Elements())
}

// Set difference

func newDifferenceSet(fst Set, snd Set) DependentSet {
  res := new(differenceSet)
  res.SetDerived = EmbeddedDependentSet(res)
  res.fst = fst
  res.snd = snd
  return res
}

type differenceSet struct {
  SetDerived
  fst Set
  snd Set
}

func (this *differenceSet) Size() int {
  return CountElements(this.Elements())
}

func (this *differenceSet) Contains(elem interface{}) bool {
  return this.fst.Contains(elem) && !this.snd.Contains(elem)
}

func (this *differenceSet) Elements() Iterator {
  return NewFilterIterator(util.Negate(this.snd.Func()), this.fst.Elements())
}

// Set proxy (to hide potential functionality for mutating the set)

func newWrappedSet(encapsulated Set, immutable bool) DependentSet {
  res := new(setWrapper)
  res.SetDerived = EmbeddedDependentSet(res)
  res.encapsulated = encapsulated
  res.immutable = immutable
  return res
}

type setWrapper struct {
  SetDerived
  encapsulated Set
  immutable bool
}

func (this *setWrapper) Size() int {
  return this.encapsulated.Size()
}

func (this *setWrapper) Contains(elem interface{}) bool {
  return this.encapsulated.Contains(elem)
}

func (this *setWrapper) Elements() Iterator {
  return this.encapsulated.Elements()
}

func (this *setWrapper) ReadOnly() DependentSet {
  return this
}

func (this *setWrapper) Freeze() FiniteContainer {
  if this.immutable {
    return this
  }
  return this.SetDerived.Freeze()
}
