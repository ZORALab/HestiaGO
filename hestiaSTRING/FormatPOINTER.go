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

func S8_FormatPOINTER[ANY any](input *ANY, base uint8,
	lettercase hestiaFMT.Lettercase) (string, hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	_processLettercase(&lettercase)

	return string(hestiaFMT.S8_FormatPOINTER(input, base, lettercase)), hestiaERROR.OK
}

func S16_FormatPOINTER[ANY any](input *ANY, base uint16,
	lettercase hestiaFMT.Lettercase) (string, hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	_processLettercase(&lettercase)

	return string(hestiaFMT.S16_FormatPOINTER(input, base, lettercase)), hestiaERROR.OK
}

func S32_FormatPOINTER[ANY any](input *ANY, base uint32,
	lettercase hestiaFMT.Lettercase) (string, hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	_processLettercase(&lettercase)

	return string(hestiaFMT.S32_FormatPOINTER(input, base, lettercase)), hestiaERROR.OK
}

func S64_FormatPOINTER[ANY any](input *ANY, base uint64,
	lettercase hestiaFMT.Lettercase) (string, hestiaERROR.Error) {
	if base < 2 || base > 36 {
		return "", hestiaERROR.DATA_INVALID
	}

	_processLettercase(&lettercase)

	return string(hestiaFMT.S64_FormatPOINTER(input, base, lettercase)), hestiaERROR.OK
}
