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

package transformers

import (
	"log"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher"
)

type EthBlockHeaderTransformer struct {
	blockchain core.BlockChain
	publisher  publisher.Publisher
}

func NewEthBlockHeaderTransformer(bc core.BlockChain, publisher publisher.Publisher) *EthBlockHeaderTransformer {
	return &EthBlockHeaderTransformer{blockchain: bc, publisher: publisher}
}

func (t EthBlockHeaderTransformer) Execute(blockNumbers []int64) error {
	for _, blockNumber := range blockNumbers {
		header, err := t.blockchain.GetHeaderByNumber(blockNumber)
		if err != nil {
			return err
		}
		headerBytes, err := rlp.EncodeToBytes(header)
		if err != nil {
			return err
		}
		output, err := t.publisher.Write(headerBytes)
		if err != nil {
			return NewExecuteError(PutIpldErr, err)
		}
		log.Printf("Created IPLD: %s", output)
	}
	return nil
}
