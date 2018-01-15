// Copyright 2018 Andrii Zavorotnii. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package flagutil

import (
	"strings"
	"flag"
)

type StringList []string

var _ flag.Value = &StringList{} // ensure interface

func (sl StringList) String() string {
	return strings.Join(sl, ",")
}

func (sl *StringList) Set(value string) error {
	*sl = append(*sl, value)
	return nil
}