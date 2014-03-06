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

import "testing"


func checkString(t *testing.T, str string, expected string, name string) {
  if str != expected {
    t.Errorf("Expected result \"%s\" of string builder %s to match \"%s%\"", str, name, expected)
  }
}

func TestStringBuilder(t *testing.T) {
  b1 := NewStringBuilder()
  checkString(t, b1.String(), "", "b1")
  b2 := NewStringBuilder("one")
  checkString(t, b2.String(), "one", "b2")
  b3 := NewStringBuilder("one", "two")
  checkString(t, b3.String(), "onetwo", "b3")
  checkString(t, b3.Join(", "), "one, two", "b3")
  b4 := NewStringBuilder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
  checkString(t, b4.Join("-"), "1-2-3-4-5-6-7-8-9-10-11-12-13", "b4")
}
