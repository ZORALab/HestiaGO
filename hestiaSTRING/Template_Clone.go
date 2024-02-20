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
	"text/template"
)

func SN_Template_Clone(engine *Templater, name string) (out *Templater, e hestiaERROR.Error) {
	var err error
	var t *template.Template

	if engine == nil {
		return nil, hestiaERROR.ENTITY_MISSING
	}

	if name == "" {
		return nil, hestiaERROR.BAD_DESCRIPTOR
	}

	if name == engine.Name {
		return nil, hestiaERROR.ENTITY_EXISTS
	}

	t, err = engine.renderer.Clone()
	if err != nil {
		return nil, hestiaERROR.ENTITY_BAD
	}

	return &Templater{
		Name:     name,
		Data:     nil,
		renderer: t,
	}, hestiaERROR.OK
}

func MN_Template_Clone(engine *Templater, name string) (out *Templater, e hestiaERROR.Error) {
	var err error
	var t *template.Template

	if engine == nil {
		return nil, hestiaERROR.ENTITY_MISSING
	}

	if name == "" {
		return nil, hestiaERROR.BAD_DESCRIPTOR
	}

	if name == engine.Name {
		return nil, hestiaERROR.ENTITY_EXISTS
	}

	t, err = engine.renderer.Clone()
	if err != nil {
		return nil, hestiaERROR.ENTITY_BAD
	}

	return &Templater{
		Name:     name,
		Data:     nil,
		renderer: t,
	}, hestiaERROR.OK
}
