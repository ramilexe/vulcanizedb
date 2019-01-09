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
	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	"gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

type EthReceiptNode struct {
	raw []byte
	cid *cid.Cid
}

func (node *EthReceiptNode) RawData() []byte {
	return node.raw
}

func (node *EthReceiptNode) Cid() cid.Cid {
	return *node.cid
}

func (*EthReceiptNode) String() string {
	panic("implement me")
}

func (*EthReceiptNode) Loggable() map[string]interface{} {
	panic("implement me")
}

func (*EthReceiptNode) Resolve(path []string) (interface{}, []string, error) {
	panic("implement me")
}

func (*EthReceiptNode) Tree(path string, depth int) []string {
	panic("implement me")
}

func (*EthReceiptNode) ResolveLink(path []string) (*format.Link, []string, error) {
	panic("implement me")
}

func (*EthReceiptNode) Copy() format.Node {
	panic("implement me")
}

func (*EthReceiptNode) Links() []*format.Link {
	panic("implement me")
}

func (*EthReceiptNode) Stat() (*format.NodeStat, error) {
	panic("implement me")
}

func (*EthReceiptNode) Size() (uint64, error) {
	panic("implement me")
}
