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

func M64_FormatFLOAT64(input float64, base uint64, precision uint64,
	lettercase hestiaFMT.Lettercase,
	notation hestiaFMT.Notation) (out string, err hestiaERROR.Error) {
	// Step 1: process all parameters to prevent panic
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	_processLettercase(&lettercase)

	switch {
	case _processNotation(&notation) != hestiaERROR.OK:
		fallthrough
	case notation == NOTATION_ISO6093NR3 && !(base == 2 || base == 10 || base == 16):
		fallthrough
	case notation == NOTATION_ISO6093NR3_AUTO && !(base == 2 || base == 10 || base == 16):
		fallthrough
	case base != 2 && notation == NOTATION_IEEE754:
		return "", hestiaERROR.BAD_DESCRIPTOR
	}

	// Step 2: constructing the data structure for processing
	param := hestiaFMT.ParamsFLOAT64{
		Value:      input,
		Precision:  precision,
		Base:       base,
		Lettercase: lettercase,
		Notation:   notation,
	}

	// Step 3: perform formatting
	return string(hestiaFMT.M64_FormatFLOAT64(&param)), hestiaERROR.OK
}
