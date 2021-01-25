// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
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

package author

import (
	"testing"

	"github.com/dn3010/go-cennznet-rpc-client/v2/types"
	"github.com/stretchr/testify/assert"
)

func TestAuthor_SubmitExtrinsic(t *testing.T) {
	ext := types.Extrinsic{Version: 0x84, Signature: types.ExtrinsicSignatureV1{Signer: types.Address{IsAccountID: true, AsAccountID: types.AccountID{0xd4, 0x35, 0x93, 0xc7, 0x15, 0xfd, 0xd3, 0x1c, 0x61, 0x14, 0x1a, 0xbd, 0x4, 0xa9, 0x9f, 0xd6, 0x82, 0x2c, 0x85, 0x58, 0x85, 0x4c, 0xcd, 0xe3, 0x9a, 0x56, 0x84, 0xe7, 0xa5, 0x6d, 0xa2, 0x7d}, IsAccountIndex: false, AsAccountIndex: 0x0}, Signature: types.MultiSignature{IsSr25519: true, AsSr25519: types.Signature{0xc0, 0x42, 0x19, 0x5f, 0x93, 0x25, 0xd, 0x3e, 0xda, 0xa2, 0xe4, 0xa4, 0x2d, 0xcf, 0x4e, 0x41, 0xc1, 0x6c, 0xa7, 0x1c, 0xfc, 0x3a, 0x2b, 0x23, 0x99, 0x8a, 0xd4, 0xec, 0x97, 0x4f, 0x8b, 0x1a, 0xcd, 0xcd, 0xad, 0x97, 0xd1, 0x4b, 0x6d, 0xf5, 0xcb, 0x89, 0x6, 0xff, 0x61, 0xc8, 0x92, 0x17, 0x96, 0x54, 0xa5, 0xec, 0xcc, 0xb, 0x66, 0x85, 0xf6, 0xc1, 0x7f, 0xed, 0x49, 0x21, 0x94, 0x0}}, Era: types.ExtrinsicEra{IsImmortalEra: true, IsMortalEra: false, AsMortalEra: types.MortalEra{First: 0x0, Second: 0x0}}, Nonce: types.NewUCompactFromUInt(0x1), TransactionPayment: types.TransactionPayment{Tip: types.NewUCompactFromUInt(0x0), FeeExchange: types.OptionFeeExchange{HasValue: false}}}, Method: types.Call{CallIndex: types.CallIndex{SectionIndex: 0x6, MethodIndex: 0x0}, Args: types.Args{0xff, 0x8e, 0xaf, 0x4, 0x15, 0x16, 0x87, 0x73, 0x63, 0x26, 0xc9, 0xfe, 0xa1, 0x7e, 0x25, 0xfc, 0x52, 0x87, 0x61, 0x36, 0x93, 0xc9, 0x12, 0x90, 0x9c, 0xb2, 0x26, 0xaa, 0x47, 0x94, 0xf2, 0x6a, 0x48, 0xe5, 0x6c}}} //nolint:lll,dupl
	res, err := author.SubmitExtrinsic(ext)
	assert.NoError(t, err)
	hex, err := types.Hex(res)
	assert.NoError(t, err)
	assert.Equal(t, mockSrv.submitExtrinsicHash, hex)
}
