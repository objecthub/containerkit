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

package maps

import "testing"


func TestNativeMapClass(t *testing.T) {
  m1 := NativeMap.New()
  checkSize(t, m1, 0, "m1")
  m2 := NativeMap.New(KV("one", 1))
  checkSize(t, m2, 1, "m2")
  m3 := NativeMap.New(KV("one", 1), KV("two", 2))
  checkSize(t, m3, 2, "m3")
  m4 := NativeMap.New(KV("one", 1), KV("two", 2), KV("one", 10))
  checkSize(t, m4, 2, "m4")
  m5 := NativeMap.New(KV("one", 1), KV("two", 2), KV("three", 3), KV("four", 4), KV("five", 5),
                      KV("six", 6), KV("seven", 7), KV("eight", 8), KV("nine", 9), KV("ten", 10))
  checkSize(t, m5, 10, "m5")
}
