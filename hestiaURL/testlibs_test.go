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

package hestiaSTRING

// test condition labels
const (
	cond_STRING_URL_PROPER = "using proper URL string"
	cond_STRING_URL_EMPTY  = "using empty URL string"
	cond_STRING_URL_ERROR  = "using non-URL string"

	cond_STRING_URL_FRAGMENT_PROPER = "include proper URL fragment into string"
	cond_STRING_URL_FRAGMENT_EMPTY  = "include empty URL fragment into string"
	cond_STRING_URL_FRAGMENT_ERROR  = "include error-ed URL fragment into string"

	cond_STRING_URL_ABSOLUTE = "using aboslute URL"
	cond_STRING_URL_RELATIVE = "using relative URL"
)

const (
	value_SCHEMA        = "https://"
	value_STRING_HOST   = "www.example.com"
	value_STRING_PORT   = "2325"
	value_QUERY_1_KEY   = "friend"
	value_QUERY_1_VALUE = "ailine"
	value_QUERY_2_KEY   = "friend"
	value_QUERY_2_VALUE = "roland"
	value_QUERY_3_KEY   = "id"
	value_QUERY_3_VALUE = "data123"
)
