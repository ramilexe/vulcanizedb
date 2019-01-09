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

package eth_block_headers

import (
	"github.com/ethereum/go-ethereum/core/types"
	ipld "gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"

	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher/util"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/wrappers/rlp"
)

const (
	EthBlockHeaderCode = 0x90
)

type BlockHeaderDagPutter struct {
	adder   publisher.Adder
	decoder rlp.Decoder
}

func NewBlockHeaderDagPutter(adder publisher.Adder, decoder rlp.Decoder) *BlockHeaderDagPutter {
	return &BlockHeaderDagPutter{adder: adder, decoder: decoder}
}

func (bhdp *BlockHeaderDagPutter) DagPut(raw interface{}) ([]string, error) {
	input := raw.([]byte)
	nd, err := bhdp.getNodeForBlockHeader(input)
	if err != nil {
		return nil, err
	}
	err = bhdp.adder.Add(nd)
	if err != nil {
		return nil, err
	}
	return []string{nd.Cid().String()}, nil
}

func (bhdp *BlockHeaderDagPutter) getNodeForBlockHeader(raw []byte) (ipld.Node, error) {
	var blockHeader types.Header
	err := bhdp.decoder.Decode(raw, &blockHeader)
	if err != nil {
		return nil, err
	}
	blockHeaderCid, err := util.RawToCid(EthBlockHeaderCode, raw)
	if err != nil {
		return nil, err
	}
	return &EthBlockHeaderNode{
		Header:  &blockHeader,
		cid:     blockHeaderCid,
		rawdata: raw,
	}, nil
}
