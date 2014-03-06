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


// Procedures consume an element and produce side-effects using this element
type Procedure func (interface{})

// Predicates are boolean functions that yield true/false based on an element
type Predicate func (interface{}) bool

// Mappings transform an element into a new element (of potentially different type)
type Mapping func (interface{}) interface{}

// A Generator maps a single element into an iterator over elements
type Generator func (interface{}) Iterator

// Equality returns true if the two given elements are considered equals.
type Equality func (interface{}, interface{}) bool

// Hashfunctions map elements to integer values
type Hashfunction func (interface{}) int

// Comparison functions compare the two given elements (a, b) and return
//   -1, if a < b
//    0, if a == b
//   +1, if a > b
type Comparison func (interface{}, interface{}) int

// Binop functions compute a binary operation for the given two elements
type Binop func (interface{}, interface{}) interface{}

// Interface Hashable is implemented by values providing a HashCode method.
// The default generic hash function uses this interface for values that
// don't have a specific predefined hash function (e.g. structs)
type Hashable interface {
  HashCode() int
}

// Interface Comparable is implemented by values providing a Compare method.
// The default generic comparison function uses this interface for values
// that don't have a specific predefined comparison function (e.g. structs)
type Comparable interface {
  Compare(other interface{}) int
}

// Interface Indentifiable is implemented by values providing an Equals
// method which returns true if the given value is equivalent.
type Identifiable interface {
  Equals(other interface{}) bool
}
