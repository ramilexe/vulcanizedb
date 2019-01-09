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
	"github.com/ethereum/go-ethereum/core/types"

	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	"gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

type EthTransactionNode struct {
	*types.Transaction

	cid     *cid.Cid
	rawdata []byte
}

func (etn *EthTransactionNode) RawData() []byte {
	return etn.rawdata
}

func (etn *EthTransactionNode) Cid() cid.Cid {
	return *etn.cid
}

func (EthTransactionNode) String() string {
	return ""
}

func (EthTransactionNode) Loggable() map[string]interface{} {
	panic("implement me")
}

func (EthTransactionNode) Resolve(path []string) (interface{}, []string, error) {
	panic("implement me")
}

func (EthTransactionNode) Tree(path string, depth int) []string {
	panic("implement me")
}

func (EthTransactionNode) ResolveLink(path []string) (*format.Link, []string, error) {
	panic("implement me")
}

func (EthTransactionNode) Copy() format.Node {
	panic("implement me")
}

func (EthTransactionNode) Links() []*format.Link {
	panic("implement me")
}

func (EthTransactionNode) Stat() (*format.NodeStat, error) {
	panic("implement me")
}

func (EthTransactionNode) Size() (uint64, error) {
	panic("implement me")
}
