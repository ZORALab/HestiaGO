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
	"hestia/hestiaTESTING"
	"testing"
)

func test_cases_SN_Parse() []*hestiaTESTING.Scenario {
	return []*hestiaTESTING.Scenario{
		{
			Description: `
test hestiaSTRING/SN_Parse is able to process proper absolute string url.
`,
			Switches: []string{
				cond_STRING_URL_PROPER,
				cond_STRING_URL_ABSOLUTE,
			},
		}, {
			Description: `
test hestiaSTRING/SN_Parse is able to process proper string relative url.
`,
			Switches: []string{
				cond_STRING_URL_PROPER,
				cond_STRING_URL_RELATIVE,
			},
		}, {
			Description: `
test hestiaSTRING/SN_Parse is able to process non-url error string.
`,
			Switches: []string{
				cond_STRING_URL_ERROR,
				cond_STRING_URL_ABSOLUTE,
			},
		}, {
			Description: `
test hestiaSTRING/SN_Parse is able to process with empty string url.
`,
			Switches: []string{
				cond_STRING_URL_EMPTY,
				cond_STRING_URL_ABSOLUTE,
			},
		},
	}
}

func Test_SN_Parse(t *testing.T) {
	scenarios := test_cases_SN_Parse()

	for i, s := range scenarios {
		s.ID = uint64(i)
		s.Name = "hestiaSTRING/SN_Parse API"

		// prepare
		subject := "fox"
		target := "dog"
		input := "A Quick Brown {{ .Subject }} Jumps Over the Lazy {{ .Target -}} ."
		expect := "A Quick Brown " + subject + " Jumps Over the Lazy " + target + "."

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Subject	: %s", subject))
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Target	: %s", target))
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Input	: %s", input))
		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Given Expect	: %s", expect))

		// test
		engine := SN_Template_New("test sample")
		engine.Data = map[string]string{
			"Subject": subject,
			"Target":  target,
		}

		err := SN_Template_ParseString(engine, input)
		if err != hestiaERROR.OK {
			hestiaTESTING.Log(s,
				hestiaTESTING.Format("Got Error while parsing	: '%v'", err))
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_SKIP)
			t.Skip()
			return
		}

		builder := SN_Buffer_New()
		panick := ""
		output := ""

		if hestiaTESTING.HasCondition(s, cond_WRITER_NIL) {
			builder = nil
		}

		if hestiaTESTING.HasCondition(s, cond_ENGINE_NIL) {
			engine = nil
		}

		_panick := hestiaTESTING.Exec(func() any {
			err = SN_Template_Render(builder, engine)
			output = builder.ToString()
			return ""
		})
		panick, _ = _panick.(string)

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Output	: '%s'", output))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Panick	: '%v'", panick))

		hestiaTESTING.Log(s,
			hestiaTESTING.Format("Got Error Out	: '%v'", err))

		// assert
		hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_PASS)

		if !assert_SN_Template_Render(panick) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_SN_Template_Render_output(s, expect, output) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		if !assert_SN_Template_Render_error(s, err) {
			hestiaTESTING.Conclude(s, hestiaTESTING.VERDICT_FAIL)
			t.Fail()
		}

		// report
		t.Logf("%v", hestiaTESTING.ToString(s))
	}
}

func assert_SN_Template_Render_output(s *hestiaTESTING.Scenario, expect string, output string) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_ENGINE_NIL):
		return output == ""
	case hestiaTESTING.HasCondition(s, cond_WRITER_NIL):
		return output == ""
	default:
		return expect == output
	}
}

func assert_SN_Template_Render_error(s *hestiaTESTING.Scenario, err hestiaERROR.Error) bool {
	switch {
	case hestiaTESTING.HasCondition(s, cond_ENGINE_NIL):
		return err != hestiaERROR.OK
	default:
		return err == hestiaERROR.OK
	}
}

func assert_SN_Template_Render(panick string) bool {
	return panick == ""
}
