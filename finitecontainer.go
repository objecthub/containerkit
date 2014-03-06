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


type Finite interface {
  Size() int
}

type FiniteContainerBase interface {
  ContainerBase
  Finite
}

type FiniteContainerDerived interface {
  ContainerDerived
}

// Finite containers are containers with a finite number of elements.
// Method Size returns the number of elements encapsulated by the container.
type FiniteContainer interface {
  FiniteContainerBase
  FiniteContainerDerived
}

type FiniteContainerClass interface {
  Embed(obj FiniteContainer) FiniteContainer
  New(elements ...interface{}) FiniteContainer
  From(coll Container) FiniteContainer
}

func EmbeddedFiniteContainer(obj FiniteContainer) FiniteContainer {
  return &finiteContainer{obj, obj, EmbeddedContainer(obj)}
}

type finiteContainer struct {
  obj FiniteContainer
  FiniteContainerBase
  FiniteContainerDerived
}

func (this *finiteContainer) Force() FiniteContainer {
  return this.obj
}

func (this *finiteContainer) IsEmpty() bool {
  return this.obj.Size() == 0
}
