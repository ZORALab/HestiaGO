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

import (
	"testing"

	"hestia/hestiaTESTING"
)

func test_cases_M8_BitsTrailingZerosUINT8() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaNUMBER/M8_BitsTrailingZerosUINT8 is able to process 0-bits value.
`,
			Switches: []string{
				cond_BITS_0,
			},
		}, {
			Description: `
test hestiaNUMBER/M8_BitsTrailingZerosUINT8 is able to process 8-bits value.
`,
			Switches: []string{
				cond_BITS_8,
			},
		},
	}
}

func Test_M8_BitsTrailingZerosUINT8(t *testing.T) {
	scenarios := test_cases_M8_BitsTrailingZerosUINT8()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaNUMBER/M8_BitsTrailingZerosUINT8 API"

		// prepare
		subject := uint8(create_sample(s))
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: 0b%b", subject))

		// test
		output := M8_BitsTrailingZerosUINT8(subject)
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Count	: %d", output))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
		if !assert_M8_BitsTrailingZerosUINT8_output(s, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_M8_BitsTrailingZerosUINT8_output(s *hestiaTESTING.Scenario, output uint8) bool {
	if hestiaTESTING.HasCondition(s, cond_BITS_8) {
		return output == (uint8(value_BITS_8_COUNT) - 1)
	}

	if hestiaTESTING.HasCondition(s, cond_BITS_0) {
		return output == 0
	}

	return false
}
