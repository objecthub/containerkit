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

package containerkit

import "github.com/objecthub/containerkit/util"


// A Container is encapsulating a set of elements. All containers
// provide an Elements() method that returns an iterator over all
// elements of the container. The framework derives a number of
// methods like Exists, Forall, ForEach, etc.from the Elements
// method. Those methods are available to all abstractions that
// embedd the Container trait.
// 
// This file implements the Container trait, ie. it specifies
// how the methods in ContainerDerived are getting defined in terms
// of functionality provided by ContainerBase. All functionality in
// ContainerDerived is lazily creating a new Container on top of
// an existing one. Only when the Elements method of the derived
// Container gets invoked, the derived container's elements are
// getting computed on the fly.
type Container interface {
  ContainerBase
  ContainerDerived
}

// The base functionality required for all Container implementations
type ContainerBase interface {
  Elements() Iterator
}

// The derived functionality implemented by the Container trait
type ContainerDerived interface {
  IsEmpty() bool
  Exists(pred Predicate) bool
  ForAll(pred Predicate) bool
  ForEach(proc Procedure)
  Take(n int) DependentContainer
  TakeWhile(pred Predicate) DependentContainer
  Drop(n int) DependentContainer
  DropWhile(pred Predicate) DependentContainer
  Filter(pred Predicate) DependentContainer
  Map(f Mapping) DependentContainer
  FlatMap(g Generator) DependentContainer
  Flatten() DependentContainer
  Concat(other Container) DependentContainer
  Combine(f Binop, other Container) DependentContainer
  Zip(other Container) DependentContainer
  FoldLeft(f Binop, z interface{}) interface{}
  FoldRight(f Binop, z interface{}) interface{}
  Force() FiniteContainer
  Freeze() FiniteContainer
  String() string
}

// Function for embedding the Container trait into another abstraction
func EmbeddedContainer(obj Container) Container {
  return &container{obj, obj}
}

// Implementation of the Container trait.
type container struct {
  obj Container
  ContainerBase
}

func (this *container) IsEmpty() bool {
  return !this.obj.Elements().HasNext()
}

func (this *container) Exists(pred Predicate) bool {
  for iter := this.obj.Elements(); iter.HasNext(); {
    if pred(iter.Next()) {
      return true
    }
  }
  return false
}

func (this *container) ForAll(pred Predicate) bool {
  for iter := this.obj.Elements(); iter.HasNext(); {
    if !pred(iter.Next()) {
      return false
    }
  }
  return true
}

func (this *container) ForEach(proc Procedure) {
  for iter := this.obj.Elements(); iter.HasNext(); {
    proc(iter.Next())
  }
}

func (this *container) Drop(n int) DependentContainer {
  return newSlicedContainer(this.obj, n, FalsePredicate, TruePredicate, 0)
}

func (this *container) DropWhile(pred Predicate) DependentContainer {
  return newSlicedContainer(this.obj, 0, pred, TruePredicate, 0)
}

func (this *container) Take(n int) DependentContainer {
  return newSlicedContainer(this.obj, 0, FalsePredicate, FalsePredicate, n)
}

func (this *container) TakeWhile(pred Predicate) DependentContainer {
  return newSlicedContainer(this.obj, 0, FalsePredicate, pred, 0)
}

func (this *container) Filter(pred Predicate) DependentContainer {
  return newFilteredContainer(this.obj, pred)
}

func (this *container) Map(f Mapping) DependentContainer {
  return newMappedContainer(this.obj, f)
}

func (this *container) FlatMap(g Generator) DependentContainer {
  return newFlatMappedContainer(this.obj, g)
}

func (this *container) Flatten() DependentContainer {
  return this.obj.FlatMap(func (x interface{}) Iterator {
    return x.(Iterator)
  })
}

func (this *container) Concat(other Container) DependentContainer {
  return newCompositeContainer(this.obj, other)
}

func (this *container) Combine(f Binop, other Container) DependentContainer {
  return newCombinedContainer(this.obj, other, f)
}

func (this *container) Zip(other Container) DependentContainer {
  return this.Combine(PairBinop, other)
}

func (this *container) FoldLeft(f Binop, z interface{}) interface{} {
  res := z
  for iter := this.obj.Elements(); iter.HasNext(); {
    res = f(res, iter.Next())
  }
  return res
}

func (this *container) FoldRight(f Binop, z interface{}) interface{} {
  return foldRight(this.obj.Elements(), f, z)
}

func (this *container) Force() FiniteContainer {
  return Enum.From(this.obj)
}

func (this *container) Freeze() FiniteContainer {
  return Enum.From(this.obj)
}

func (this *container) String() string {
  builder := util.NewStringBuilder()
  for iter := this.obj.Elements(); iter.HasNext(); {
    builder.Append(iter.Next())
  }
  return builder.Join(", ")
}

func foldRight(iter Iterator, f Binop, z interface{}) interface{} {
  if iter.HasNext() {
    return f(iter.Next(), foldRight(iter, f, z))
  }
  return z
}
