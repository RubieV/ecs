// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by scripts/gocodegen.go - DO NOT EDIT.

package ecs

// These fields can represent attacks on applications.
// For example, ModSecurity events, or F5 ASM violations.
type Attack struct {
	// Rule name prefixed by the source system
	ID string `ecs:"id"`

	// Description of the rule data.
	Name string `ecs:"name"`

	// Description of the rule data.
	Description string `ecs:"description"`

	// Payload that the rule is triggered on
	Payload string `ecs:"payload"`

	// The attack category
	Category string `ecs:"category"`

	// The attack category
	Subcategory string `ecs:"subcategory"`

	// Unique path to the system and the location in the system where the
	// payload is inserted to
	Path string `ecs:"path"`

	// Code prefixed by the source system, for example "CRS-942100"
	ApplicableSystems string `ecs:"applicable_systems"`

	// Code prefixed by the source system, for example "CRS-942100"
	Policy string `ecs:"policy"`

	// Code prefixed by the source system, for example "CRS-942100"
	KillchainCategory string `ecs:"killchain.category"`

	// The numeric categroy and subcategory level. The first digit represent
	// the category (1 to 9), the second digit represents a subcategory level
	// (default 0).
	KillchainLevel int64 `ecs:"killchain.level"`

	// The type of certainty that there is an actual positive
	Certainty string `ecs:"certainty"`

	// Stage that the rule is in, from disabled to blocking.
	Mode string `ecs:"mode"`

	// The file or policy location that the rule belongs to
	Configuration string `ecs:"configuration"`
}
