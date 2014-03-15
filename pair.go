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


// Pair represents a tuple encapsulating two values: first and second.
type Pair interface {

  // First() returns the first value of this Pair
  First() interface{}
  
  // Second() returns the second value of this Pair
  Second() interface{}
  
  // Get returns both the first and the second component
  Get() (first interface{}, second interface{})
  
  // Key returns the first component of this Pair. This method is useful if the Pair
  // is used as a key/value pair.
  Key() interface{}
  
  // Value returns the second component of this Pair. This method is useful if the Pair
  // is used as a key/value pair.
  Value() interface{}
  
  // Equals(other) returns true if 'this' and 'other' are equal. Two Pair objects
  // are equal if their first and second values each individually are equals.
  // The notion of equality that is used here is based on function UniversalEquality
  Equals(other interface{}) bool
  
  // Compare(other) returns 0, if 'this' and 'other' are equal; it returns -1 if
  // 'this' < 'other'; it returns 1 if 'this' > 'other'. A Pair 'this' is bigger than
  // a Pair 'other' if the first component of 'this' is bigger than the first component
  // of 'that'. If the first components are equal, then 'this' is bigger than 'other'
  // if the second component of 'this' is bigger than the second component of 'other'.
  Compare(other interface{}) int
  
  // HashCode returns a hash code for this Pair
  HashCode() int
  
  // Swap returns a new Pair with first and second component swapped
  Swap() Pair
  
  // Pair returns the pair itself
  Pair() Pair
  
  // String returns a textual representation of this Pair
  String() string
}

// NewPair returns a new Pair object for the given two components
func NewPair(first, second interface{}) Pair {
  return &pair{first, second}
}

// PairBinop defines a Binop function which encapsulates the two given parameters
// in a new Pair
func PairBinop(first, second interface{}) interface{} {
  return NewPair(first, second)
}

type pair struct {
  first interface{}
  second interface{}
}

func (this *pair) First() interface{} {
  return this.first
}

func (this *pair) Second() interface{} {
  return this.second
}

func (this *pair) Get() (first interface{}, second interface{}) {
  return this.first, this.second
}

func (this *pair) Key() interface{} {
  return this.first
}

func (this *pair) Value() interface{} {
  return this.second
}

func (this *pair) Equals(other interface{}) bool {
  if that, valid := other.(Pair); valid {
    return UniversalEquality(this.first, that.First()) &&
           UniversalEquality(this.second, that.Second())
  }
  return false
}

func (this *pair) Compare(other interface{}) int {
  if that, valid := other.(Pair); valid {
    fstcmp := UniversalComparison(this.first, that.First())
    if (fstcmp == 0) {
      return UniversalComparison(this.second, that.Second())
    } else {
      return fstcmp
    }
  }
  panic("pair.Compare: uncomparable values")
}

func (this *pair) HashCode() int {
  return UniversalHash(this.first) * 31 + UniversalHash(this.second)
}

func (this *pair) Swap() Pair {
  return NewPair(this.second, this.first)
}

func (this *pair) Pair() Pair {
  return this
}

func (this *pair) String() string {
  return util.NewStringBuilder("(", this.first, ", ", this.second, ")").String()
}

