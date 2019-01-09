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

package eth_state_trie

import (
	ipfs "github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher/util"
)

const (
	EthStateTrieNodeCode = 0x96
)

type StateTrieDagPutter struct {
	adder ipfs.Adder
}

func NewStateTrieDagPutter(adder ipfs.Adder) *StateTrieDagPutter {
	return &StateTrieDagPutter{adder: adder}
}

func (stdp StateTrieDagPutter) DagPut(raw interface{}) ([]string, error) {
	input := raw.([]byte)
	stateTrieNode, err := stdp.getStateTrieNode(input)
	if err != nil {
		return nil, err
	}
	err = stdp.adder.Add(stateTrieNode)
	if err != nil {
		return nil, err
	}
	return []string{stateTrieNode.Cid().String()}, nil
}

func (stdp StateTrieDagPutter) getStateTrieNode(raw []byte) (*EthStateTrieNode, error) {
	stateTrieNodeCid, err := util.RawToCid(EthStateTrieNodeCode, raw)
	if err != nil {
		return nil, err
	}
	stateTrieNode := &EthStateTrieNode{
		cid:     stateTrieNodeCid,
		rawdata: raw,
	}
	return stateTrieNode, nil
}
