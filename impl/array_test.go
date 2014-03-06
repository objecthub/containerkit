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

package impl

import "testing"


func TestEmptyArray(t *testing.T) {
  a := NewArray()
  if a.Length() != 0 {
    t.Errorf("Expected empty array")
  }
}

func TestGetSet(t *testing.T) {
  a := NewArray()
  a.Allocate(0, 2)
  a.Set(0, "zero")
  a.Set(1, "one")
  if a.At(0) != "zero" {
    t.Errorf("Expected zero")
  }
  if a.At(1) != "one" {
    t.Errorf("Expected one")
  }
  if a.Length() != 2 {
    t.Errorf("Length should be 2")
  }
}
