// Copyright (c) 2014 The gomqtt Authors. All rights reserved.
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

package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisconnectMessageDecode(t *testing.T) {
	msgBytes := []byte{
		byte(DISCONNECT << 4),
		0,
	}

	msg := NewDisconnectMessage()
	n, err := msg.Decode(msgBytes)

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n, "Error decoding message.")
	require.Equal(t, DISCONNECT, msg.Type(), "Error decoding message.")
}

func TestDisconnectMessageEncode(t *testing.T) {
	msgBytes := []byte{
		byte(DISCONNECT << 4),
		0,
	}

	msg := NewDisconnectMessage()

	dst := make([]byte, 10)
	n, err := msg.Encode(dst)

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n, "Error decoding message.")
	require.Equal(t, msgBytes, dst[:n], "Error decoding message.")
}

// test to ensure encoding and decoding are the same
// decode, encode, and decode again
func TestDisconnectDecodeEncodeEquiv(t *testing.T) {
	msgBytes := []byte{
		byte(DISCONNECT << 4),
		0,
	}

	msg := NewDisconnectMessage()
	n, err := msg.Decode(msgBytes)

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n, "Error decoding message.")

	dst := make([]byte, 100)
	n2, err := msg.Encode(dst)

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n2, "Error decoding message.")
	require.Equal(t, msgBytes, dst[:n2], "Error decoding message.")

	n3, err := msg.Decode(dst)

	require.NoError(t, err, "Error decoding message.")
	require.Equal(t, len(msgBytes), n3, "Error decoding message.")
}

func BenchmarkDisconnectEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msg := NewDisconnectMessage()
		msg.Encode(make([]byte, msg.Len()))
	}
}

func BenchmarkDisconnectDecode(b *testing.B) {
	msgBytes := []byte{
		byte(DISCONNECT << 4),
		0,
	}

	msg := NewDisconnectMessage()

	for i := 0; i < b.N; i++ {
		msg.Decode(msgBytes)
	}
}
