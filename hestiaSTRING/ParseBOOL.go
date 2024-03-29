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

func SN_ParseBOOL(input string) (out bool, err hestiaERROR.Error) {
	var errX hestiaFMT.Error

	out, errX = hestiaFMT.SN_ParseBOOL(input)
	switch errX {
	case hestiaFMT.ERROR_INPUT_INVALID:
		err = hestiaERROR.INVALID_ARGUMENT
	default:
		err = hestiaERROR.OK
	}

	return out, err
}

func MN_ParseBOOL(input string) (out bool, err hestiaERROR.Error) {
	var errX hestiaFMT.Error

	out, errX = hestiaFMT.SN_ParseBOOL(input)
	switch errX {
	case hestiaFMT.ERROR_INPUT_INVALID:
		err = hestiaERROR.INVALID_ARGUMENT
	default:
		err = hestiaERROR.OK
	}

	return out, err
}
