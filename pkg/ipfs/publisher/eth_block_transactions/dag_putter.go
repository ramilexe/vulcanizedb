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

package eth_block_transactions

import (
	"bytes"

	"github.com/ethereum/go-ethereum/core/types"

	ipfs "github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher/util"
)

const (
	EthBlockTransactionCode = 0x93
)

type BlockTransactionsDagPutter struct {
	adder ipfs.Adder
}

func NewBlockTransactionsDagPutter(adder ipfs.Adder) *BlockTransactionsDagPutter {
	return &BlockTransactionsDagPutter{adder: adder}
}

func (bbdp *BlockTransactionsDagPutter) DagPut(body interface{}) ([]string, error) {
	blockBody := body.(*types.Body)
	transactions := blockBody.Transactions
	var cids []string
	for _, transaction := range transactions {
		buffer := new(bytes.Buffer)
		err := transaction.EncodeRLP(buffer)
		if err != nil {
			return nil, err
		}
		transactionCid, err := util.RawToCid(EthBlockTransactionCode, buffer.Bytes())
		if err != nil {
			return nil, err
		}
		transactionNode := &EthTransactionNode{
			Transaction: transaction,
			cid:         transactionCid,
			rawdata:     buffer.Bytes(),
		}
		err = bbdp.adder.Add(transactionNode)
		if err != nil {
			return nil, err
		}
		cids = append(cids, transactionCid.String())
	}
	return cids, nil
}
