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

func test_cases_BitsSizeCPU() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaNUMBER/*N_BitsSizeCPU is able to return value.
`,
			Switches: []string{},
		},
	}
}

func Test_BitsSizeCPU(t *testing.T) {
	scenarios := test_cases_BitsSizeCPU()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaNUMBER/*N_BitsSizeCPU API"

		// prepare

		// test & assert
		output := SN_BitsSizeCPU()
		hestiaTESTING.Log(s, hestiaTESTING.Format("Got Output (SN)	: %d", output))
		if !assert_BitsSizeCPU_output(output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		output = MN_BitsSizeCPU()
		hestiaTESTING.Log(s, hestiaTESTING.Format("Got Output (MN)	: %d", output))
		if !assert_BitsSizeCPU_output(output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_BitsSizeCPU_output(output uint64) bool {
	return output != 0
}
