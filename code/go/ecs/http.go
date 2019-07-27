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

// Fields related to HTTP activity. Use the `url` field set to store the url of
// the request.
type Http struct {
	// HTTP request method.
	// The field value must be normalized to lowercase for querying. See the
	// documentation section "Implementing ECS".
	RequestMethod string `ecs:"request.method"`

	// The full HTTP request body.
	RequestHeadersContent string `ecs:"request.headers.content"`

	// The full HTTP request body.
	RequestHeader map[string]interface{} `ecs:"request.header"`

	// The full HTTP request body.
	RequestBodyContent string `ecs:"request.body.content"`

	// Referrer for this HTTP request.
	RequestReferrer string `ecs:"request.referrer"`

	// HTTP response status code.
	ResponseStatusCode int64 `ecs:"response.status_code"`

	// The full HTTP response body.
	ResponseHeadersContent string `ecs:"response.headers.content"`

	// The full HTTP response body.
	ResponseBodyContent string `ecs:"response.body.content"`

	// HTTP version.
	Version string `ecs:"version"`

	// Total size in bytes of the request (body and headers).
	RequestBytes int64 `ecs:"request.bytes"`

	// Size in bytes of the request body.
	RequestBodyBytes int64 `ecs:"request.body.bytes"`

	// Total size in bytes of the response (body and headers).
	ResponseBytes int64 `ecs:"response.bytes"`

	// Size in bytes of the response body.
	ResponseBodyBytes int64 `ecs:"response.body.bytes"`
}
