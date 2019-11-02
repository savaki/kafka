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

// ApiVersionsResponse; ApiKey: 18, Versions: 0-3
type ApiVersionsResponse struct {
	ErrorCode      int16                      // The top-level error code. Versions: 0+
	ApiKeys        []ApiVersionsResponseKey18 // The APIs supported by the broker. Versions: 0+
	ThrottleTimeMs int32                      // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 1+
}

// size of ApiVersionsResponse; Versions: 0-3
func (t ApiVersionsResponse) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16       // ErrorCode
	sz += sizeof.ArrayLength // ApiKeys
	for i := len(t.ApiKeys) - 1; i >= 0; i-- {
		sz += t.ApiKeys[i].size(version)
	}
	if version >= 1 {
		sz += sizeof.Int32 // ThrottleTimeMs
	}
	return sz
}

// encode ApiVersionsResponse; Versions: 0-3
func (t ApiVersionsResponse) Encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode) // ErrorCode
	// ApiKeys
	len1 := len(t.ApiKeys)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.ApiKeys[i].Encode(e, version)
	}
	if version >= 1 {
		e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	}
}

// decode ApiVersionsResponse; Versions: 0-3
func (t *ApiVersionsResponse) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	// ApiKeys
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.ApiKeys = make([]ApiVersionsResponseKey18, n)
		for i := 0; i < n; i++ {
			var item ApiVersionsResponseKey18
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.ApiKeys[i] = item
		}
	}
	if version >= 1 {
		t.ThrottleTimeMs, err = d.Int32()
		if err != nil {
			return err
		}
	}
	return err
}

type ApiVersionsResponseKey18 struct {
	ApiKey     int16 // The API index. Versions: 0+
	MinVersion int16 // The minimum supported version, inclusive. Versions: 0+
	MaxVersion int16 // The maximum supported version, inclusive. Versions: 0+
}

// size of ApiVersionsResponseKey18; Versions: 0-3
func (t ApiVersionsResponseKey18) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16 // ApiKey
	sz += sizeof.Int16 // MinVersion
	sz += sizeof.Int16 // MaxVersion
	return sz
}

// encode ApiVersionsResponseKey18; Versions: 0-3
func (t ApiVersionsResponseKey18) Encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ApiKey)     // ApiKey
	e.PutInt16(t.MinVersion) // MinVersion
	e.PutInt16(t.MaxVersion) // MaxVersion
}

// decode ApiVersionsResponseKey18; Versions: 0-3
func (t *ApiVersionsResponseKey18) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ApiKey, err = d.Int16()
	if err != nil {
		return err
	}
	t.MinVersion, err = d.Int16()
	if err != nil {
		return err
	}
	t.MaxVersion, err = d.Int16()
	if err != nil {
		return err
	}
	return err
}
