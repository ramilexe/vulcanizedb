// VulcanizeDB
// Copyright © 2018 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package test_helpers

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

var (
	FakeError     = errors.New("failed")
	FakeHash      = common.HexToHash("0x123")
	FakeString    = "Test"
	FakeTrieNodes = [][]byte{{1, 1, 1, 1, 1}, {2, 2, 2, 2, 2}, {3, 3, 3, 3, 3}}
	FakeTrieNode  = []byte{248, 68, 1, 128, 160, 6, 180, 135, 209, 92, 2, 139, 109, 245, 108, 62, 187, 155, 112, 134, 150, 94, 186, 58, 36, 8, 87, 166, 71, 250, 236, 226, 255, 19, 38, 159, 43, 160, 206, 51, 34, 13, 92, 127, 13, 9, 215, 92, 239, 247, 108, 5, 134, 60, 94, 125, 110, 128, 28, 112, 223, 231, 213, 212, 93, 76, 68, 232, 6, 84}
)
