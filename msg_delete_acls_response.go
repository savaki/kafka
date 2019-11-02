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

// DeleteAclsResponse; ApiKey: 31, Versions: 0-1
type DeleteAclsResponse struct {
	ThrottleTimeMs int32                      // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 0+
	FilterResults  []DeleteAclsFilterResult31 // The results for each filter. Versions: 0+
}

// size of DeleteAclsResponse; Versions: 0-1
func (t DeleteAclsResponse) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32       // ThrottleTimeMs
	sz += sizeof.ArrayLength // FilterResults
	for i := len(t.FilterResults) - 1; i >= 0; i-- {
		sz += t.FilterResults[i].size(version)
	}
	return sz
}

// encode DeleteAclsResponse; Versions: 0-1
func (t DeleteAclsResponse) encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	// FilterResults
	len1 := len(t.FilterResults)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.FilterResults[i].encode(e, version)
	}
}

// decode DeleteAclsResponse; Versions: 0-1
func (t *DeleteAclsResponse) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ThrottleTimeMs, err = d.Int32()
	if err != nil {
		return err
	}
	// FilterResults
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.FilterResults = make([]DeleteAclsFilterResult31, n)
		for i := 0; i < n; i++ {
			var item DeleteAclsFilterResult31
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.FilterResults[i] = item
		}
	}
	return err
}

type DeleteAclsFilterResult31 struct {
	ErrorCode    int16                     // The error code, or 0 if the filter succeeded. Versions: 0+
	ErrorMessage string                    // The error message, or null if the filter succeeded. Versions: 0+
	MatchingAcls []DeleteAclsMatchingAcl31 // The ACLs which matched this filter. Versions: 0+
}

// size of DeleteAclsFilterResult31; Versions: 0-1
func (t DeleteAclsFilterResult31) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16                  // ErrorCode
	sz += sizeof.String(t.ErrorMessage) // ErrorMessage
	sz += sizeof.ArrayLength            // MatchingAcls
	for i := len(t.MatchingAcls) - 1; i >= 0; i-- {
		sz += t.MatchingAcls[i].size(version)
	}
	return sz
}

// encode DeleteAclsFilterResult31; Versions: 0-1
func (t DeleteAclsFilterResult31) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)     // ErrorCode
	e.PutString(t.ErrorMessage) // ErrorMessage
	// MatchingAcls
	len2 := len(t.MatchingAcls)
	e.PutArrayLength(len2)
	for i := 0; i < len2; i++ {
		t.MatchingAcls[i].encode(e, version)
	}
}

// decode DeleteAclsFilterResult31; Versions: 0-1
func (t *DeleteAclsFilterResult31) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.ErrorMessage, err = d.String()
	if err != nil {
		return err
	}
	// MatchingAcls
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.MatchingAcls = make([]DeleteAclsMatchingAcl31, n)
		for i := 0; i < n; i++ {
			var item DeleteAclsMatchingAcl31
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.MatchingAcls[i] = item
		}
	}
	return err
}

type DeleteAclsMatchingAcl31 struct {
	ErrorCode      int16  // The deletion error code, or 0 if the deletion succeeded. Versions: 0+
	ErrorMessage   string // The deletion error message, or null if the deletion succeeded. Versions: 0+
	ResourceType   int8   // The ACL resource type. Versions: 0+
	ResourceName   string // The ACL resource name. Versions: 0+
	PatternType    int8   // The ACL resource pattern type. Versions: 1+
	Principal      string // The ACL principal. Versions: 0+
	Host           string // The ACL host. Versions: 0+
	Operation      int8   // The ACL operation. Versions: 0+
	PermissionType int8   // The ACL permission type. Versions: 0+
}

// size of DeleteAclsMatchingAcl31; Versions: 0-1
func (t DeleteAclsMatchingAcl31) size(version int16) int32 {
	var sz int32
	sz += sizeof.Int16                  // ErrorCode
	sz += sizeof.String(t.ErrorMessage) // ErrorMessage
	sz += sizeof.Int8                   // ResourceType
	sz += sizeof.String(t.ResourceName) // ResourceName
	if version >= 1 {
		sz += sizeof.Int8 // PatternType
	}
	sz += sizeof.String(t.Principal) // Principal
	sz += sizeof.String(t.Host)      // Host
	sz += sizeof.Int8                // Operation
	sz += sizeof.Int8                // PermissionType
	return sz
}

// encode DeleteAclsMatchingAcl31; Versions: 0-1
func (t DeleteAclsMatchingAcl31) encode(e *protocol.Encoder, version int16) {
	e.PutInt16(t.ErrorCode)     // ErrorCode
	e.PutString(t.ErrorMessage) // ErrorMessage
	e.PutInt8(t.ResourceType)   // ResourceType
	e.PutString(t.ResourceName) // ResourceName
	if version >= 1 {
		e.PutInt8(t.PatternType) // PatternType
	}
	e.PutString(t.Principal)    // Principal
	e.PutString(t.Host)         // Host
	e.PutInt8(t.Operation)      // Operation
	e.PutInt8(t.PermissionType) // PermissionType
}

// decode DeleteAclsMatchingAcl31; Versions: 0-1
func (t *DeleteAclsMatchingAcl31) decode(d *protocol.Decoder, version int16) error {
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
	if version >= 1 {
		t.PatternType, err = d.Int8()
		if err != nil {
			return err
		}
	}
	t.Principal, err = d.String()
	if err != nil {
		return err
	}
	t.Host, err = d.String()
	if err != nil {
		return err
	}
	t.Operation, err = d.Int8()
	if err != nil {
		return err
	}
	t.PermissionType, err = d.Int8()
	if err != nil {
		return err
	}
	return err
}
