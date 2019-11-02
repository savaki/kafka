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

// AddOffsetsToTxnResponse; ApiKey: 25, Versions: 0-1
type AddOffsetsToTxnResponse struct {
	ThrottleTimeMs int32 // Duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 0+
	ErrorCode      int16 // The response error code, or 0 if there was no error. Versions: 0+
}

// size of AddOffsetsToTxnResponse; Versions: 0-1
func (t AddOffsetsToTxnResponse) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32 // ThrottleTimeMs
	sz += sizeof.Int16 // ErrorCode
	return sz
}

// encode AddOffsetsToTxnResponse; Versions: 0-1
func (t AddOffsetsToTxnResponse) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	e.PutInt16(t.ErrorCode)      // ErrorCode
}

// decode AddOffsetsToTxnResponse; Versions: 0-1
func (t *AddOffsetsToTxnResponse) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ThrottleTimeMs, err = d.Int32()
	if err != nil {
		return err
	}
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	return err
}
