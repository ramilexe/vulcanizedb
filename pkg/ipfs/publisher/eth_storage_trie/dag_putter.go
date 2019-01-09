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

package eth_storage_trie

import (
	ipfs "github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher/util"
)

const (
	EthStorageTrieNodeCode = 0x98
)

type StorageTrieDagPutter struct {
	adder ipfs.Adder
}

func NewStorageTrieDagPutter(adder ipfs.Adder) *StorageTrieDagPutter {
	return &StorageTrieDagPutter{adder: adder}
}

func (stdp StorageTrieDagPutter) DagPut(raw interface{}) ([]string, error) {
	input := raw.([]byte)
	cid, err := util.RawToCid(EthStorageTrieNodeCode, input)
	if err != nil {
		return nil, err
	}
	node := &EthStorageTrieNode{
		cid:     cid,
		rawdata: input,
	}
	err = stdp.adder.Add(node)
	if err != nil {
		return nil, err
	}
	return []string{node.Cid().String()}, nil
}
