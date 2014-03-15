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
// FUNCTION TYPES
// ============================================================================

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

// Binop functions compute a binary operation for the given two parameters
type Binop func (interface{}, interface{}) interface{}

// Incrementor functions compute a successor for a given integer
type Incrementor func (int) int


// ============================================================================
// DEFAULT FUNCTION IMPLEMENTATIONS
// ============================================================================

// Identity defines an identity mapping function
func Identity(x interface{}) interface{} {
  return x
}

// InvertComparison returns a new inverted Comparison function for the given
// Comparison function
func InvertComparison(comp Comparison) Comparison {
  return func (x, y interface{}) int {
    return comp(y, x)
  }
}

// Inc defines an incrementor function which increments the parameter by one
func Inc(i int) int {
  return i + 1
}

// Dec defines an incrementor function which decrements the parameter by one
func Dec(i int) int {
  return i - 1
}

// TruePredicate defines a predicate function which always returns true
func TruePredicate(x interface{}) bool {
  return true
}

// FalsePredicate defines a predicate function which always returns false
func FalsePredicate(x interface{}) bool {
  return false
}

// Negate defines a function which given a predicate, returns the negated version
// of this predicate
func Negate(pred Predicate) Predicate {
  return func (x interface{}) bool {
    return !pred(x)
  }
}
