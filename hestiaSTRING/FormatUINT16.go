// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
// Copyright 2009 The Go Authors <https://cs.opensource.google/go/go>
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
	"hestia/hestiaIO/hestiaINTERNAL/hestiaFMT"
)

func S16_FormatUINT16(input uint16, base uint16,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.S16_FormatUINT16(input, base, lettercase)), hestiaERROR.OK
}

func M16_FormatUINT16(input uint16, base uint16,
	lettercase hestiaFMT.Lettercase) (out string, err hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	if input == 0 {
		return "0", hestiaERROR.OK
	}

	_processLettercase(&lettercase)
	return string(hestiaFMT.S16_FormatUINT16(input, base, lettercase)), hestiaERROR.OK
}
