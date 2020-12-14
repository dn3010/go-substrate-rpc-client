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

package types

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/scale"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
)

// ExtrinsicPayloadV3 is a signing payload for an Extrinsic. For the final encoding, it is variable length based on
// the contents included. Note that `BytesBare` is absolutely critical – we don't want the method (Bytes)
// to have the length prefix included. This means that the data-as-signed is un-decodable,
// but is also doesn't need the extra information, only the pure data (and is not decoded)
// ... The same applies to V1 & V1, if we have a V4, carry move this comment to latest
type ExtrinsicPayloadV1 struct {
	Method BytesBare

	Era                ExtrinsicEra
	Nonce              UCompact
	TransactionPayment TransactionPayment

	SpecVersion        U32
	TransactionVersion U32
	GenesisHash        Hash
	BlockHash          Hash
}

type TransactionPayment struct {
	Tip         UCompact
	FeeExchange OptionFeeExchange
}

type OptionFeeExchange struct {
	HasValue    bool
	FeeExchange FeeExchangeV1
}
type FeeExchangeV1 struct {
	AssetId    UCompact
	MaxPayment UCompact
}

func (fe FeeExchangeV1) Encode(encoder scale.Encoder) error {
	err := encoder.PushByte(0)
	if err != nil {
		return err
	}

	err = encoder.Encode(fe.AssetId)
	if err != nil {
		return err
	}

	err = encoder.Encode(fe.MaxPayment)
	if err != nil {
		return err
	}

	return nil
}

func (fe OptionFeeExchange) Encode(encoder scale.Encoder) error {
	err := encoder.EncodeOption(fe.HasValue, fe.FeeExchange)
	if err != nil {
		return err
	}

	return nil
}

// Sign the extrinsic payload with the given derivation path
func (e ExtrinsicPayloadV1) Sign(signer signature.KeyringPair) (Signature, error) {
	b, err := EncodeToBytes(e)
	if err != nil {
		return Signature{}, err
	}

	hex, err := EncodeToHexString(b)
	if err != nil {
		panic(err)
	}
	println("BUF", hex)

	sig, err := signature.Sign(b, signer.URI)
	return NewSignature(sig), err
}

// func (e ExtrinsicPayloadV1) Encode(encoder scale.Encoder) error {
// 	err := encoder.Encode(e.Method)
// 	if err != nil {
// 		return err
// 	}

// 	err = encoder.Encode(e.Era)
// 	if err != nil {
// 		return err
// 	}

// 	err = encoder.Encode(e.Nonce)
// 	if err != nil {
// 		return err
// 	}

// 	err = encoder.Encode(e.TransactionPayment)
// 	if err != nil {
// 		return err
// 	}

// 	err = encoder.Encode(e.SpecVersion)
// 	if err != nil {
// 		return err
// 	}

// 	err = encoder.Encode(e.TransactionVersion)
// 	if err != nil {
// 		return err
// 	}

// 	err = encoder.Encode(e.GenesisHash)
// 	if err != nil {
// 		return err
// 	}

// 	err = encoder.Encode(e.BlockHash)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Decode does nothing and always returns an error. ExtrinsicPayloadV1 is only used for encoding, not for decoding
func (e *ExtrinsicPayloadV1) Decode(decoder scale.Decoder) error {
	return fmt.Errorf("decoding of ExtrinsicPayloadV1 is not supported")
}
