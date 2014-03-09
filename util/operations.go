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

package util

import "math"


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

func Inc(i int) int {
  return i + 1
}

func Dec(i int) int {
  return i - 1
}

func TruePredicate(x interface{}) bool {
  return true
}

func FalsePredicate(x interface{}) bool {
  return false
}

func Negate(pred (func (interface{}) bool)) (func (interface{}) bool) {
  return func (x interface{}) bool {
    return !pred(x)
  }
}

func InvertComparison(comp (func (interface{}, interface{}) int)) (func (interface{}, interface{}) int) {
  return func (x, y interface{}) int {
    return comp(y, x)
  }
}

func Identity(x interface{}) interface{} {
  return x
}

func UniversalEquality(x, y interface{}) bool {
  if x != nil {
    if left, valid := x.(Identifiable); valid {
      return left.Equals(y)
    }
  }
  return x == y
}

func UniversalHash(x interface{}) int {
  switch val := x.(type) {
    case nil:
      return 0
    case bool:
      if val {
        return 1
      }
      return 0
    case byte:
      return int(val)
    case int:
      return val
    case uint:
      return int(val)
    case int32:
      return int(val)
    case int64:
      return int(val ^ (val >> 32))
    case float32:
      return int(math.Float32bits(val))
    case float64:
      return int(math.Float64bits(val))
    case string:
      res := 0
      for _, ch := range val {
        res = res * 31 + int(ch)
      }
      return res
    case Hashable:
      if res, valid := x.(Hashable); valid {
        return res.HashCode()
      }
  }
  panic("UniversalHash: Unknown type")
}


func UniversalComparison(x, y interface{}) int {
  switch this := x.(type) {
    case nil:
      return comparatorCode(y == nil, true)
    case bool:
      if that, valid := y.(bool); valid {
        return comparatorCode(this == that, !this)
      }
    case byte:
      if that, valid := y.(byte); valid {
        return comparatorCode(this == that, this < that)
      }
    case int:
      if that, valid := y.(int); valid {
        return comparatorCode(this == that, this < that)
      }
    case uint:
      if that, valid := y.(uint); valid {
        return comparatorCode(this == that, this < that)
      }
    case int32:
      if that, valid := y.(int32); valid {
        return comparatorCode(this == that, this < that)
      }
    case int64:
      if that, valid := y.(int64); valid {
        return comparatorCode(this == that, this < that)
      }
    case float32:
      if that, valid := y.(float32); valid {
        return comparatorCode(this == that, this < that)
      }
    case float64:
      if that, valid := y.(float64); valid {
        return comparatorCode(this == that, this < that)
      }
    case string:
      if that, valid := y.(string); valid {
        return comparatorCode(this == that, this < that)
      }
    case Comparable:
      return this.Compare(y)
  }
  panic("UniversalComparison: Illegal parameters")
}

func comparatorCode(eq bool, less bool) int {
  if eq {
    return 0
  } else if less {
    return -1
  }
  return 1
}
