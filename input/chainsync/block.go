// Copyright 2023 Blink Labs, LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chainsync

import (
	"fmt"

	"github.com/blinklabs-io/gouroboros/ledger"
	"github.com/fxamacker/cbor/v2"
)

type BlockContext struct {
	BlockNumber uint64           `json:"blockNumber"`
	SlotNumber  uint64           `json:"slotNumber"`
}

type BlockEvent struct {
	BlockBodySize uint64           `json:"blockBodySize"`
	IssuerVkey    string           `json:"issuerVkey"`
	BlockHash     string           `json:"blockHash"`
	BlockCbor     byteSliceJsonHex `json:"blockCbor,omitempty"`
}

func NewBlockContext(block ledger.Block) BlockContext {
	ctx := BlockContext{
		BlockNumber: block.BlockNumber(),
		SlotNumber:  block.SlotNumber(),
	}
	return ctx
}

func NewBlockHeaderContext(block ledger.BlockHeader) BlockContext {
	ctx := BlockContext{
		BlockNumber: block.BlockNumber(),
		SlotNumber:  block.SlotNumber(),
	}
	return ctx
}

func NewBlockEvent(block ledger.Block, includeCbor bool) BlockEvent {
	keyCbor, err := cbor.Marshal(block.IssuerVkey())
	if err != nil {
		panic(err)
	}
	// iss := ledger.NewBlake2b256(block.IssuerVkey())
	evt := BlockEvent{
		BlockBodySize: block.BlockBodySize(),
		BlockHash:     block.Hash(),
		IssuerVkey:    fmt.Sprintf("%x", keyCbor),
		// IssuerVkey:    iss.String(),
	}
	if includeCbor {
		evt.BlockCbor = block.Cbor()
	}
	return evt
}
