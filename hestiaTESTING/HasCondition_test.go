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

package hestiaTESTING

import (
	"testing"
)

func _testHasConditionScenarios() []*Scenario {
	return []*Scenario{
		{
			Description: `
Test HasCondition() is able to work properly with proper Scenario settings.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test HasCondition() is able to work properly with empty Name setting.
`,
			Switches: []string{
				cond_EMPTY_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test HasCondition() is able to work properly with empty Switches setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_EMPTY_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test HasCondition() is able to work properly with nil Switches setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_NIL_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test HasCondition() is able to work properly with empty log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_EMPTY_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test HasCondition() is able to work properly with nil log setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_NIL_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test HasCondition() is able to work properly with empty description setting.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_EMPTY_DESCRIPTION,
				cond_NIL_LOG,
				cond_PROPER_SWITCHES,
				cond_PROPER_VERDICT,
			},
		}, {
			Description: `
Test HasCondition() is able to panic when nil Scenario is supplied.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_SUPPLY_NIL_SCENARIO,
				cond_PROPER_VERDICT,
				expect_PANIC,
			},
		}, {
			Description: `
Test HasCondition() is able to work properly when verdict is set to unknown.
`,
			Switches: []string{
				cond_PROPER_NAME,
				cond_PROPER_DESCRIPTION,
				cond_PROPER_LOG,
				cond_PROPER_SWITCHES,
				cond_SUPPLY_NIL_SCENARIO,
				cond_UNKNOWN_VERDICT,
				expect_PANIC,
			},
		},
	}
}

func TestHasConditionAPI(t *testing.T) {
	scenarios := _testHasConditionScenarios()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = suite_HAS_CONDITION_API

		// prepare
		ts := &Scenario{}
		testlib_ConfigureName(s, ts)
		testlib_ConfigureDescription(s, ts)
		testlib_ConfigureLog(s, ts)
		testlib_ConfigureSwitches(s, ts)
		testlib_ConfigureVerdict(s, ts)

		// test
		output := false
		_panick := Exec(func() any {
			if !HasCondition(s, cond_SUPPLY_NIL_SCENARIO) {
				output = HasCondition(ts, value_SWITCH_1)
			} else {
				output = HasCondition(nil, value_SWITCH_1)
			}
			return ""
		})
		panick := _panick.(string)

		// log output
		Log(s, Format("Test Scenario's Switches	= %#v", ts.Switches))
		Log(s, Format("Got Output		= %v", output))
		Log(s, Format("Got Panic		= %q", panick))

		// assert
		if !testlib_AssertPanic(s, panick) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if !assert_HasCondition(s, output) {
			Conclude(s, VERDICT_FAIL)
			t.Fail()
		}

		if Conclusion(s) != VERDICT_FAIL {
			Conclude(s, VERDICT_PASS)
		}

		// report
		t.Logf("%s", ToString(s))
	}
}

func assert_HasCondition(s *Scenario, output bool) bool {
	switch {
	case HasCondition(s, cond_EMPTY_SWITCHES),
		HasCondition(s, cond_NIL_SWITCHES),
		HasCondition(s, cond_SUPPLY_NIL_SCENARIO):
		return output == false
	}

	return output == true
}
