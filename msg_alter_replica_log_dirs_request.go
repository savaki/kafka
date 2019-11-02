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

// AlterReplicaLogDirsRequest; ApiKey: 34, Versions: 0-1
type AlterReplicaLogDirsRequest struct {
	Dirs []AlterReplicaLogDir34 // The alterations to make for each directory. Versions: 0+
}

// size of AlterReplicaLogDirsRequest; Versions: 0-1
func (t AlterReplicaLogDirsRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.ArrayLength // Dirs
	for i := len(t.Dirs) - 1; i >= 0; i-- {
		sz += t.Dirs[i].size(version)
	}
	return sz
}

// encode AlterReplicaLogDirsRequest; Versions: 0-1
func (t AlterReplicaLogDirsRequest) encode(e *protocol.Encoder, version int16) {
	// Dirs
	len0 := len(t.Dirs)
	e.PutArrayLength(len0)
	for i := 0; i < len0; i++ {
		t.Dirs[i].encode(e, version)
	}
}

// decode AlterReplicaLogDirsRequest; Versions: 0-1
func (t *AlterReplicaLogDirsRequest) decode(d *protocol.Decoder, version int16) error {
	var err error
	// Dirs
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Dirs = make([]AlterReplicaLogDir34, n)
		for i := 0; i < n; i++ {
			var item AlterReplicaLogDir34
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Dirs[i] = item
		}
	}
	return err
}

type AlterReplicaLogDir34 struct {
	Path   string                      // The absolute directory path. Versions: 0+
	Topics []AlterReplicaLogDirTopic34 // The topics to add to the directory. Versions: 0+
}

// size of AlterReplicaLogDir34; Versions: 0-1
func (t AlterReplicaLogDir34) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Path) // Path
	sz += sizeof.ArrayLength    // Topics
	for i := len(t.Topics) - 1; i >= 0; i-- {
		sz += t.Topics[i].size(version)
	}
	return sz
}

// encode AlterReplicaLogDir34; Versions: 0-1
func (t AlterReplicaLogDir34) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Path) // Path
	// Topics
	len1 := len(t.Topics)
	e.PutArrayLength(len1)
	for i := 0; i < len1; i++ {
		t.Topics[i].encode(e, version)
	}
}

// decode AlterReplicaLogDir34; Versions: 0-1
func (t *AlterReplicaLogDir34) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Path, err = d.String()
	if err != nil {
		return err
	}
	// Topics
	if n, err := d.ArrayLength(); err != nil {
		return err
	} else if n >= 0 {
		t.Topics = make([]AlterReplicaLogDirTopic34, n)
		for i := 0; i < n; i++ {
			var item AlterReplicaLogDirTopic34
			if err := (&item).decode(d, version); err != nil {
				return err
			}
			t.Topics[i] = item
		}
	}
	return err
}

type AlterReplicaLogDirTopic34 struct {
	Name       string  // The topic name. Versions: 0+
	Partitions []int32 // The partition indexes. Versions: 0+
}

// size of AlterReplicaLogDirTopic34; Versions: 0-1
func (t AlterReplicaLogDirTopic34) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.Name)           // Name
	sz += sizeof.Int32Array(t.Partitions) // Partitions
	return sz
}

// encode AlterReplicaLogDirTopic34; Versions: 0-1
func (t AlterReplicaLogDirTopic34) encode(e *protocol.Encoder, version int16) {
	e.PutString(t.Name)           // Name
	e.PutInt32Array(t.Partitions) // Partitions
}

// decode AlterReplicaLogDirTopic34; Versions: 0-1
func (t *AlterReplicaLogDirTopic34) decode(d *protocol.Decoder, version int16) error {
	var err error
	t.Name, err = d.String()
	if err != nil {
		return err
	}
	t.Partitions, err = d.Int32Array()
	if err != nil {
		return err
	}
	return err
}
