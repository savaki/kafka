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

// RenewDelegationTokenResponse; ApiKey: 39, Versions: 0-1
type RenewDelegationTokenResponse struct {
	ErrorCode         int16 // The error code, or 0 if there was no error. Versions: 0+
	ExpiryTimestampMs int64 // The timestamp in milliseconds at which this token expires. Versions: 0+
	ThrottleTimeMs    int32 // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 0+
}

// size of RenewDelegationTokenResponse; Versions: 0-1
func (t RenewDelegationTokenResponse) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16 // ErrorCode
	sz += sizeof.Int64 // ExpiryTimestampMs
	sz += sizeof.Int32 // ThrottleTimeMs
	return sz
}

// encode RenewDelegationTokenResponse; Versions: 0-1
func (t RenewDelegationTokenResponse) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)         // ErrorCode
	e.PutInt64(t.ExpiryTimestampMs) // ExpiryTimestampMs
	e.PutInt32(t.ThrottleTimeMs)    // ThrottleTimeMs
}

// decode RenewDelegationTokenResponse; Versions: 0-1
func (t *RenewDelegationTokenResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.ExpiryTimestampMs, err = d.Int64()
	if err != nil {
		return err
	}
	t.ThrottleTimeMs, err = d.Int32()
	if err != nil {
		return err
	}
	return err
}
