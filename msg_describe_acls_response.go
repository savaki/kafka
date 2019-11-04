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

// DescribeAclsResponse; ApiKey: 29, Versions: 0-1
type DescribeAclsResponse struct {
	ThrottleTimeMs int32                    // The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota. Versions: 0+
	ErrorCode      int16                    // The error code, or 0 if there was no error. Versions: 0+
	ErrorMessage   string                   // The error message, or null if there was no error. Versions: 0+
	Resources      []DescribeAclsResource29 // Each Resource that is referenced in an ACL. Versions: 0+
}

// size of DescribeAclsResponse; Versions: 0-1
func (t DescribeAclsResponse) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int32                  // ThrottleTimeMs
	sz += sizeof.Int16                  // ErrorCode
	sz += sizeof.String(t.ErrorMessage) // ErrorMessage
	sz += sizeof.ArrayLength            // Resources
	for i := len(t.Resources) - 1; i >= 0; i-- {
		sz += t.Resources[i].Size(version)
	}
	return sz
}

// encode DescribeAclsResponse; Versions: 0-1
func (t DescribeAclsResponse) Encode(e *protocol.Encoder, version int16) {
	e.PutInt32(t.ThrottleTimeMs) // ThrottleTimeMs
	e.PutInt16(t.ErrorCode)      // ErrorCode
	e.PutString(t.ErrorMessage)  // ErrorMessage
	// Resources
	len3 := len(t.Resources)
	e.PutArrayLength(len3)
	for i := 0; i < len3; i++ {
		t.Resources[i].Encode(e, version)
	}
}

// decode DescribeAclsResponse; Versions: 0-1
func (t *DescribeAclsResponse) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.ThrottleTimeMs, err = d.Int32()
	if err != nil {
		return err
	}
	t.ErrorCode, err = d.Int16()
	if err != nil {
		return err
	}
	t.ErrorMessage, err = d.String()
	if err != nil {
		return err
	}
	// Resources
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Resources = make([]DescribeAclsResource29, n)
		for i := 0; i < n; i++ {
			var item DescribeAclsResource29
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Resources[i] = item
		}
	}
	return err
}

type DescribeAclsResource29 struct {
	Type        int8               // The resource type. Versions: 0+
	Name        string             // The resource name. Versions: 0+
	PatternType int8               // The resource pattern type. Versions: 1+
	Acls        []AclDescription29 // The ACLs. Versions: 0+
}

// size of DescribeAclsResource29; Versions: 0-1
func (t DescribeAclsResource29) Size(version int16) int32 {
	var sz int32
	sz += sizeof.Int8           // Type
	sz += sizeof.String(t.Name) // Name
	if version >= 1 {
		sz += sizeof.Int8 // PatternType
	}
	sz += sizeof.ArrayLength // Acls
	for i := len(t.Acls) - 1; i >= 0; i-- {
		sz += t.Acls[i].Size(version)
	}
	return sz
}

// encode DescribeAclsResource29; Versions: 0-1
func (t DescribeAclsResource29) Encode(e *protocol.Encoder, version int16) {
	e.PutInt8(t.Type)   // Type
	e.PutString(t.Name) // Name
	if version >= 1 {
		e.PutInt8(t.PatternType) // PatternType
	}
	// Acls
	len3 := len(t.Acls)
	e.PutArrayLength(len3)
	for i := 0; i < len3; i++ {
		t.Acls[i].Encode(e, version)
	}
}

// decode DescribeAclsResource29; Versions: 0-1
func (t *DescribeAclsResource29) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Type, err = d.Int8()
	if err != nil {
		return err
	}
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	if version >= 1 {
		t.PatternType, err = d.Int8()
		if err != nil {
			return err
		}
	}
	// Acls
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Acls = make([]AclDescription29, n)
		for i := 0; i < n; i++ {
			var item AclDescription29
			if err := (&item).Decode(d, version); err != nil {
				return err
			}
			t.Acls[i] = item
		}
	}
	return err
}

type AclDescription29 struct {
	Principal      string // The ACL principal. Versions: 0+
	Host           string // The ACL host. Versions: 0+
	Operation      int8   // The ACL operation. Versions: 0+
	PermissionType int8   // The ACL permission type. Versions: 0+
}

// size of AclDescription29; Versions: 0-1
func (t AclDescription29) Size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Principal) // Principal
	sz += sizeof.String(t.Host)      // Host
	sz += sizeof.Int8                // Operation
	sz += sizeof.Int8                // PermissionType
	return sz
}

// encode AclDescription29; Versions: 0-1
func (t AclDescription29) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Principal)    // Principal
	e.PutString(t.Host)         // Host
	e.PutInt8(t.Operation)      // Operation
	e.PutInt8(t.PermissionType) // PermissionType
}

// decode AclDescription29; Versions: 0-1
func (t *AclDescription29) Decode(d *protocol.Decoder, version int16) error {
	var err error
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
