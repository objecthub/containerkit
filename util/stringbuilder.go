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

import "fmt"


type StringBuilder interface {
  Prepend(val ...interface{}) StringBuilder
  Append(val ...interface{}) StringBuilder
  PrependStr(text string) StringBuilder
  AppendStr(text string) StringBuilder
  Join(separator string) string
  Clear()
  String() string
}

func NewStringBuilder(text ...interface{}) StringBuilder {
  this := new(stringBuilder)
  this.prepended = make([]string, 0, 4)
  this.appended = make([]string, 0, 8)
  this.Append(text...)
  return this
}

func ToString(val interface{}) string {
  return fmt.Sprint(val)
}

type stringBuilder struct {
  prepended []string
  appended []string
}

func expand(components *[]string, text string) {
  a := *components
  l := len(a)
  if l < cap(a) {
    *components = a[: l + 1]
  } else {
    b := make([]string, l + 1, cap(a) * 2)
    copy(b, a)
    *components = b
  }
  (*components)[l] = text
}

func (this *stringBuilder) Prepend(values ...interface{}) StringBuilder {
  for i := len(values) - 1; i >= 0; i-- {
    text := fmt.Sprint(values[i])
    expand(&this.prepended, text)
  }
  return this
}

func (this *stringBuilder) Append(values ...interface{}) StringBuilder {
  for _, val := range values {
    text := fmt.Sprint(val)
    expand(&this.appended, text)
  }
  return this
}

func (this *stringBuilder) PrependStr(text string) StringBuilder {
  expand(&this.prepended, text)
  return this
}

func (this *stringBuilder) AppendStr(text string) StringBuilder {
  expand(&this.appended, text)
  return this
}

func insert(b *[]byte, bp int, text string) int {
  copy((*b)[bp:], []byte(text))
  return bp + len(text)
}

func (this *stringBuilder) Join(separator string) string {
  n := len(this.prepended)
  m := len(this.appended)
  // handle special cases
  if n + m == 0 {
    return ""
  } else if (n == 0) && (m == 1) {
    return this.appended[0]
  } else if (m == 0) && (n == 1) {
    return this.prepended[0]
  }
  // compute overall length
  l := len(separator) * (m + n - 1)
  for i := 0; i < n; i++ {
    l += len(this.prepended[i])
  }
  for i := 0; i < m; i++ {
    l += len(this.appended[i])
  }
  // copy bytes
  b := make([]byte, l)
  bp := 0
  for i := n - 1; i >= 0; i-- {
    bp = insert(&b, bp, this.prepended[i])
    if i > 0 {
      bp = insert(&b, bp, separator)
    }
  }
  if (bp > 0) && (bp < l) {
    bp = insert(&b, bp, separator)
  }
  for i := 0; i < m; i++ {
    bp = insert(&b, bp, this.appended[i])
    if i + 1 < m {
      bp = insert(&b, bp, separator)
    }
  }
  return string(b)
}

func (this *stringBuilder) Clear() {
  this.appended = make([]string, 0, 8)
  this.prepended = make([]string, 0, 4)
}

func (this *stringBuilder) String() string {
  return this.Join("")
}
