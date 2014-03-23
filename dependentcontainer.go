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


// ============================================================================
// INTERFACE
// ============================================================================

// DependentContainer implementations do not encapsulate mutable state. The elements
// of dependent containers are derived from the state provided by other containers.
// Every time an element gets accessed, the content of a dependent container gets
// recomputed from scratch, but only as needed by the operation.
// 
// DependentContainer extends Container, i.e. it inherits both the base and the
// derived functionality from Container.
type DependentContainer interface {
  DependentContainerBase
  DependentContainerDerived
}

// The base functionality required for all DependentContainer implementations.
// This corresponds to the base functionality of Container.
type DependentContainerBase interface {

  // Inherited from Container
  ContainerBase
}

// The derived functionality implemented by the DependentContainer trait.
// This corresponds to the derived functionality of Container. In addition, two
// access functions to the containers from which the dependent container is derived
// are provided.
type DependentContainerDerived interface {

  // Inherited from Container
  ContainerDerived
  
  // Returns the first Container from which this DependentContainer gets derived.
  first() Container
  
  // Returns the second Container from which this DependentContainer gets derived.
  // If there is no second container, this method returns nil.
  second() Container
}

// Function for embedding the DependentContainer trait into another abstraction
func EmbeddedDependentContainer(obj DependentContainer,
                                first Container,
                                second Container) DependentContainer {
  return &dependentContainer{obj, EmbeddedContainer(obj), first, second}
}


// ============================================================================
// IMPLEMENTATION
// ============================================================================

type dependentContainer struct {
  obj DependentContainer  // The identify of the container abstraction
  Container               // The inherited Container functionality
  fst Container           // Container based on which this DependentContainer is defined
  snd Container           // Container based on which this DependentContainer is defined
}

func (this *dependentContainer) first() Container {
  return this.fst
}

func (this *dependentContainer) second() Container {
  return this.snd
}

func (this *dependentContainer) String() string {
  return "<" + this.Container.String() + ">"
}


// Sliced containers

func newSlicedContainer(base Container,
                        drop int,
                        dropWhile Predicate,
                        takeWhile Predicate,
                        take int) DependentContainer {
  res := new(slicedContainer)
  res.drop = drop
  res.dropWhile = dropWhile
  res.takeWhile = takeWhile
  res.take = take
  res.DependentContainerDerived = EmbeddedDependentContainer(res, base, nil)
  return res
}

type slicedContainer struct {
  drop int
  dropWhile Predicate
  takeWhile Predicate
  take int
  DependentContainerDerived
}

func (this *slicedContainer) Elements() Iterator {
  return NewSlicedIterator(this.first().Elements(),
                           this.drop,
                           this.dropWhile,
                           this.takeWhile,
                           this.take)
}

// Filtered containers

func newFilteredContainer(base Container, pred Predicate) DependentContainer {
  res := new(filteredContainer)
  res.pred = pred
  res.DependentContainerDerived = EmbeddedDependentContainer(res, base, nil)
  return res
}

type filteredContainer struct {
  pred Predicate
  DependentContainerDerived
}

func (this *filteredContainer) Elements() Iterator {
  return NewFilterIterator(this.pred, this.first().Elements())
}

// Mapped containers

func newMappedContainer(base Container, f Mapping) DependentContainer {
  res := new(mappedContainer)
  res.f = f
  res.DependentContainerDerived = EmbeddedDependentContainer(res, base, nil)
  return res
}

type mappedContainer struct {
  f Mapping
  DependentContainerDerived
}

func (this *mappedContainer) Elements() Iterator {
  return &mappedIterator{this.f, this.first().Elements()}
}

// Flat-mapped containers

func newFlatMappedContainer(base Container, g Generator) DependentContainer {
  res := new(flatMappedContainer)
  res.g = g
  res.DependentContainerDerived = EmbeddedDependentContainer(res, base, nil)
  return res
}

type flatMappedContainer struct {
  g Generator
  DependentContainerDerived
}

func (this *flatMappedContainer) Elements() Iterator {
  res := &flatMappedIterator{this.g, this.first().Elements(), nil}
  res.scan()
  return res
}

// Composite containers

func newCompositeContainer(first Container, second Container) DependentContainer {
  res := new(compositeContainer)
  res.DependentContainerDerived = EmbeddedDependentContainer(res, first, second)
  return res
}

type compositeContainer struct {
  DependentContainerDerived
}

func (this *compositeContainer) Elements() Iterator {
  return &compositeIterator{this.first().Elements(), this.second().Elements()}
}

// Combined containers

func newCombinedContainer(first Container,
                           second Container,
                           f Binop) DependentContainer {
  res := new(combinedContainer)
  res.DependentContainerDerived = EmbeddedDependentContainer(res, first, second)
  res.f = f
  return res
}

type combinedContainer struct {
  DependentContainerDerived
  f Binop
}

func (this *combinedContainer) Elements() Iterator {
  return &combinedIterator{this.first().Elements(),
                           this.second().Elements(),
                           this.f}
}
