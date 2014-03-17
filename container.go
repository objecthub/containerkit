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
// This file implements the Container trait, i.e. it specifies
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

  // IsEmpty returns true if the container is empty, i.e. has no elements.
  IsEmpty() bool
  
  // Exists returns true if there is at least one element for which the given
  // predicate is true.
  Exists(pred Predicate) bool
  
  // ForAll returns true if the given predicate is true for all elements.
  ForAll(pred Predicate) bool
  
  // ForEach executes the given procedure for all elements.
  ForEach(proc Procedure)
  
  // Filter returns a dependent container containing all elements for which the given
  // predicate is true.
  Filter(pred Predicate) DependentContainer
  
  // Take returns a dependent container encapsulating n elements; it is the first
  // n elements returned from the Elements iterator.
  Take(n int) DependentContainer
  
  // TakeWhile returns a dependent container into which elements are put as long as
  // the predicate returns true.
  TakeWhile(pred Predicate) DependentContainer
  
  // Drop returns a dependent container which contains all elements except for the
  // first n elements returned by the Elements iterator
  Drop(n int) DependentContainer
  
  // DropWhile returns a dependent container into which all elements are put, except
  // for the first elements for which the given predicate returns true.
  DropWhile(pred Predicate) DependentContainer
  
  // Map maps all elements into a new container by applying the given mapping function.
  Map(f Mapping) DependentContainer
  
  // FlatMap applies the given generator to each element and concatenates the
  // containers resulting from the generator invokations.
  FlatMap(g Generator) DependentContainer
  
  // Flatten turns a container of containers into a flat container by concatenating
  // the individual containers.
  Flatten() DependentContainer
  
  // Concat returns a dependent container which contains both the elements from
  // this and the other container.
  Concat(other Container) DependentContainer
  
  // Combine returns a dependent container which combines elements from this and the
  // other container by applying the given binary operation.
  Combine(f Binop, other Container) DependentContainer

  // Zip returns a dependent container which combines elements from this and the
  // other container as Pair objects.
  Zip(other Container) DependentContainer
  
  // FoldLeft aggregates the elements {e1, e2, e3, ..., en} of this container by
  // applying the given binary operation f in the following way:
  // f(... f(f(f(z, e1), e2), e3), ... en)
  FoldLeft(f Binop, z interface{}) interface{}
  
  // FoldRight aggregates the elements {e1, e2, e3, ..., en} of this container by
  // applying the given binary operation f in the following way:
  // f(e1, f(e2, f(e3, ... f(en, z) ...)))
  FoldRight(f Binop, z interface{}) interface{}
  
  Force() FiniteContainer
  Freeze() FiniteContainer
  
  // String returns a textual representation of this container.
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
    switch e := x.(type) {
      case nil:
        return EmptyIterator
      case ContainerBase:
        return e.Elements()
    }
    panic("Container.Flatten: element is not a container")
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
