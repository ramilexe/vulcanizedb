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

package eth_block_receipts

import (
	"bytes"

	"github.com/ethereum/go-ethereum/core/types"

	ipfs "github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher/util"

	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
)

type EthBlockReceiptDagPutter struct {
	adder ipfs.Adder
}

func NewEthBlockReceiptDagPutter(adder ipfs.Adder) *EthBlockReceiptDagPutter {
	return &EthBlockReceiptDagPutter{adder: adder}
}

func (dagPutter *EthBlockReceiptDagPutter) DagPut(raw interface{}) ([]string, error) {
	input := raw.(types.Receipts)
	var output []string
	for _, r := range input {
		node, err := getReceiptNode(r)
		if err != nil {
			return nil, err
		}
		err = dagPutter.adder.Add(node)
		if err != nil {
			return nil, err
		}
		output = append(output, node.cid.String())
	}
	return output, nil
}

func getReceiptNode(receipt *types.Receipt) (*EthReceiptNode, error) {
	buffer := new(bytes.Buffer)
	err := receipt.EncodeRLP(buffer)
	if err != nil {
		return nil, err
	}
	receiptCid, err := util.RawToCid(cid.EthTxReceipt, buffer.Bytes())
	if err != nil {
		return nil, err
	}
	node := &EthReceiptNode{
		raw: buffer.Bytes(),
		cid: receiptCid,
	}
	return node, nil
}
