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

// EndTxnRequest; ApiKey: 26, Versions: 0-1
type EndTxnRequest struct {
	TransactionalId string // The ID of the transaction to end. Versions: 0+
	ProducerId      int64  // The producer ID. Versions: 0+
	ProducerEpoch   int16  // The current epoch associated with the producer. Versions: 0+
	Committed       bool   // True if the transaction was committed, false if it was aborted. Versions: 0+
}

// size of EndTxnRequest; Versions: 0-1
func (t EndTxnRequest) size(version int16) int32 {
	var sz int32
	sz += sizeof.String(t.TransactionalId) // TransactionalId
	sz += sizeof.Int64                     // ProducerId
	sz += sizeof.Int16                     // ProducerEpoch
	sz += sizeof.Bool                      // Committed
	return sz
}

// encode EndTxnRequest; Versions: 0-1
func (t EndTxnRequest) Encode(e *protocol.Encoder, version int16) {
	e.PutString(t.TransactionalId) // TransactionalId
	e.PutInt64(t.ProducerId)       // ProducerId
	e.PutInt16(t.ProducerEpoch)    // ProducerEpoch
	e.PutBool(t.Committed)         // Committed
}

// decode EndTxnRequest; Versions: 0-1
func (t *EndTxnRequest) Decode(d *protocol.Decoder, version int16) error {
	var err error
	t.TransactionalId, err = d.String()
	if err != nil {
		return err
	}
	t.ProducerId, err = d.Int64()
	if err != nil {
		return err
	}
	t.ProducerEpoch, err = d.Int16()
	if err != nil {
		return err
	}
	t.Committed, err = d.Bool()
	if err != nil {
		return err
	}
	return err
}
