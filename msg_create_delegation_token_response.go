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

// CreateDelegationTokenResponse; ApiKey: 38, Versions: 0-2
type CreateDelegationTokenResponse struct {
	ErrorCode         int16  // The top-level error, or zero if there was no error. Versions: 0+
	PrincipalType     string // The principal type of the token owner. Versions: 0+
	PrincipalName     string // The name of the token owner. Versions: 0+
	IssueTimestampMs  int64  // When this token was generated. Versions: 0+
	ExpiryTimestampMs int64  // When this token expires. Versions: 0+
	MaxTimestampMs    int64  // The maximum lifetime of this token. Versions: 0+
	TokenId           string // The token UUID. Versions: 0+
	Hmac              []byte // HMAC of the delegation token. Versions: 0+
	ThrottleTimeMs    int32  // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 0+
}

// size of CreateDelegationTokenResponse; Versions: 0-2
func (t CreateDelegationTokenResponse) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16                   // ErrorCode
	sz += sizeof.String(t.PrincipalType) // PrincipalType
	sz += sizeof.String(t.PrincipalName) // PrincipalName
	sz += sizeof.Int64                   // IssueTimestampMs
	sz += sizeof.Int64                   // ExpiryTimestampMs
	sz += sizeof.Int64                   // MaxTimestampMs
	sz += sizeof.String(t.TokenId)       // TokenId
	sz += sizeof.Bytes(t.Hmac)           // Hmac
	sz += sizeof.Int32                   // ThrottleTimeMs
	return sz
}

// encode CreateDelegationTokenResponse; Versions: 0-2
func (t CreateDelegationTokenResponse) Encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)         // ErrorCode
	e.PutString(t.PrincipalType)    // PrincipalType
	e.PutString(t.PrincipalName)    // PrincipalName
	e.PutInt64(t.IssueTimestampMs)  // IssueTimestampMs
	e.PutInt64(t.ExpiryTimestampMs) // ExpiryTimestampMs
	e.PutInt64(t.MaxTimestampMs)    // MaxTimestampMs
	e.PutString(t.TokenId)          // TokenId
	e.PutBytes(t.Hmac)              // Hmac
	e.PutInt32(t.ThrottleTimeMs)    // ThrottleTimeMs
}

// decode CreateDelegationTokenResponse; Versions: 0-2
func (t *CreateDelegationTokenResponse) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.PrincipalType, err = d.String()
	if err != nil {
		return err
	}
	t.PrincipalName, err = d.String()
	if err != nil {
		return err
	}
	t.IssueTimestampMs, err = d.Int64()
	if err != nil {
		return err
	}
	t.ExpiryTimestampMs, err = d.Int64()
	if err != nil {
		return err
	}
	t.MaxTimestampMs, err = d.Int64()
	if err != nil {
		return err
	}
	t.TokenId, err = d.String()
	if err != nil {
		return err
	}
	t.Hmac, err = d.Bytes()
	if err != nil {
		return err
	}
	t.ThrottleTimeMs, err = d.Int32()
	if err != nil {
		return err
	}
	return err
}
