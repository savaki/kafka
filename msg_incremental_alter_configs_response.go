// Code generated by kafka-protocol-gen. DO NOT EDIT.
//
// Copyright 2019 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kafka

import (
	"github.com/savaki/kafka/protocol"
	"github.com/savaki/kafka/sizeof"
)

// IncrementalAlterConfigsResponse; ApiKey: 44, Versions: 0-1
type IncrementalAlterConfigsResponse struct {
	ThrottleTimeMs int32                            // Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 0+
	Responses      []AlterConfigsResourceResponse44 // The responses for each resource. Versions: 0+
}

// size of IncrementalAlterConfigsResponse; Versions: 0-1
func (t IncrementalAlterConfigsResponse) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32       // ThrottleTimeMs
	sz += sizeof.ArrayLength // Responses
	for i := len(t.Responses) - 1; i >= 0; i-- {
		sz += t.Responses[i].Size(version)
	}
	return sz
}

// encode IncrementalAlterConfigsResponse; Versions: 0-1
func (t IncrementalAlterConfigsResponse) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	// Responses
	len1 := len(t.Responses)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Responses[i].Encode(e, version)
	}
}

// decode IncrementalAlterConfigsResponse; Versions: 0-1
func (t *IncrementalAlterConfigsResponse) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ThrottleTimeMs, err = d.Int32()
	if err != nil {
		return err
	}
	// Responses
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Responses = make([]AlterConfigsResourceResponse44, n)
		for i := 0; i < n; i++ {
			var item AlterConfigsResourceResponse44
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Responses[i] = item
		}
	}
	return err
}

type AlterConfigsResourceResponse44 struct {
	ErrorCode    int16  // The resource error code. Versions: 0+
	ErrorMessage string // The resource error message, or null if there was no error. Versions: 0+
	ResourceType int8   // The resource type. Versions: 0+
	ResourceName string // The resource name. Versions: 0+
}

// size of AlterConfigsResourceResponse44; Versions: 0-1
func (t AlterConfigsResourceResponse44) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16                  // ErrorCode
	sz += sizeof.String(t.ErrorMessage) // ErrorMessage
	sz += sizeof.Int8                   // ResourceType
	sz += sizeof.String(t.ResourceName) // ResourceName
	return sz
}

// encode AlterConfigsResourceResponse44; Versions: 0-1
func (t AlterConfigsResourceResponse44) Encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)     // ErrorCode
	e.PutString(t.ErrorMessage) // ErrorMessage
	e.PutInt8(t.ResourceType)   // ResourceType
	e.PutString(t.ResourceName) // ResourceName
}

// decode AlterConfigsResourceResponse44; Versions: 0-1
func (t *AlterConfigsResourceResponse44) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.ErrorMessage, err = d.String()
	if err != nil {
		return err
	}
	t.ResourceType, err = d.Int8()
	if err != nil {
		return err
	}
	t.ResourceName, err = d.String()
	if err != nil {
		return err
	}
	return err
}
