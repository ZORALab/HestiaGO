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

func S8_Itoa(input int8) (output string) {
	output, _ = S8_FormatINT8(input, 10, LETTERCASE_LOWER)
	return output
}

func M8_Itoa(input int8) (output string) {
	output, _ = S8_FormatINT8(input, 10, LETTERCASE_LOWER)
	return output
}

func S16_Itoa(input int16) (output string) {
	output, _ = S16_FormatINT16(input, 10, LETTERCASE_LOWER)
	return output
}

func M16_Itoa(input int16) (output string) {
	output, _ = S16_FormatINT16(input, 10, LETTERCASE_LOWER)
	return output
}

func S32_Itoa(input int32) (output string) {
	output, _ = S32_FormatINT32(input, 10, LETTERCASE_LOWER)
	return output
}

func M32_Itoa(input int32) (output string) {
	output, _ = S32_FormatINT32(input, 10, LETTERCASE_LOWER)
	return output
}

func S64_Itoa(input int64) (output string) {
	output, _ = S64_FormatINT64(input, 10, LETTERCASE_LOWER)
	return output
}

func M64_Itoa(input int64) (output string) {
	output, _ = S64_FormatINT64(input, 10, LETTERCASE_LOWER)
	return output
}
