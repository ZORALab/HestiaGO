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

package hestiaURL

import (
	"hestia/hestiaERROR"
	"net/url"
)

func SN_Parse(input string) (output *URL, e hestiaERROR.Error) {
	if input == "" {
		return nil, hestiaERROR.DATA_EMPTY
	}

	u, err := url.Parse(input)
	if err != nil {
		return nil, hestiaERROR.DATA_BAD
	}

	p, _ := u.User.Password()

	output = &URL{
		Scheme:      u.Scheme,
		Opaque:      u.Opaque,
		Host:        u.Host,
		Path:        u.Path,
		RawPath:     u.RawPath,
		Fragment:    u.Fragment,
		RawFragment: u.RawFragment,
		Queries:     map[string][]string(u.Query()),
		RawQueries:  u.RawQuery,

		username: u.User.Username(),
		password: p,
	}

	return output, hestiaERROR.OK
}

func MN_Parse(input string) (output *URL, e hestiaERROR.Error) {
	if input == "" {
		return nil, hestiaERROR.DATA_EMPTY
	}

	u, err := url.Parse(input)
	if err != nil {
		return nil, hestiaERROR.DATA_BAD
	}

	p, _ := u.User.Password()

	output = &URL{
		Scheme:      u.Scheme,
		Opaque:      u.Opaque,
		Host:        u.Host,
		Path:        u.Path,
		RawPath:     u.RawPath,
		Fragment:    u.Fragment,
		RawFragment: u.RawFragment,
		Queries:     map[string][]string(u.Query()),
		RawQueries:  u.RawQuery,

		username: u.User.Username(),
		password: p,
	}

	return output, hestiaERROR.OK
}
