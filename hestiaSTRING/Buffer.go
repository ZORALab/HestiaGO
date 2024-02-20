// Copyright 2023 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2023 ZORALab Enterprise <tech@zoralab.com>
//
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//                  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package hestiaSTRING

import (
	"hestia/hestiaERROR"
	"strings"
)

type Buffer struct {
	writer *strings.Builder
}

func (b *Buffer) Output(source *[]byte) (err hestiaERROR.Error) {
	_, e := b.writer.Write(*source)
	if e != nil {
		return hestiaERROR.BAD_EXEC
	}

	return hestiaERROR.OK
}

func (b *Buffer) Write(source []byte) (n int, err error) {
	return b.writer.Write(source)
}

func (b *Buffer) String() string {
	return b.writer.String()
}

func (b *Buffer) ToString() string {
	return b.writer.String()
}

func SN_Buffer_New() *Buffer {
	return &Buffer{
		writer: &strings.Builder{},
	}
}

func MN_Buffer_New() *Buffer {
	return &Buffer{
		writer: &strings.Builder{},
	}
}
