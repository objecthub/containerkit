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

package sets

import "testing"


func checkSize(t *testing.T, q MutableSet, size int, name string) {
  qsize := q.Size()
  if qsize != size {
    t.Errorf("Expected size of set %s to be %d; was %d", name, size, qsize)
  }
}

func TestHashSetClass(t *testing.T) {
  s1 := HashSet.New()
  checkSize(t, s1, 0, "s1")
  s2 := HashSet.New(1)
  checkSize(t, s2, 1, "s2")
  s3 := HashSet.New(1, 1)
  checkSize(t, s3, 1, "s3")
  s4 := HashSet.New(1, 2)
  checkSize(t, s4, 2, "s4")
  s5 := HashSet.New(1, 2, 3, 2, 4, 4, 5, 1, 6, 7, 8, 8, 8, 9, 10)
  checkSize(t, s5, 10, "s5")
}
