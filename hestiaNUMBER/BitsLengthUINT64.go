// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
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

package hestiaNUMBER

func S64_BitsLengthUINT64(a uint64) uint64 {
	var i, b, c uint64

	if a == 0 {
		return 0
	}

	// loop through bit shifting to find the highest bit
	b = 0
	for i = 1; i <= MAX_UINT64; i <<= 1 {
		a |= (a >> i)

		c = a ^ (a >> 1)
		if b == c {
			// max bit reached
			break
		}

		b = c
	}

	// convert back to decimal counting value
	a = 0
	for b != 0 {
		a++
		b >>= 1
	}

	return a
}

func M64_BitsLengthUINT64(a uint64) uint64 {
	var i, b, c uint64

	if a == 0 {
		return 0
	}

	// loop through bit shifting to find the highest bit
	b = 0
	for i = 1; i <= MAX_UINT64; i <<= 1 {
		a |= (a >> i)

		c = a ^ (a >> 1)
		if b == c {
			// max bit reached
			break
		}

		b = c
	}

	// convert back to decimal counting value
	a = 0
	for b != 0 {
		a++
		b >>= 1
	}

	return a
}
