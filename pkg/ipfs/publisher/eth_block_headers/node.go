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

	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	"gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

type EthBlockHeaderNode struct {
	*types.Header

	cid     *cid.Cid
	rawdata []byte
}

func (ebh *EthBlockHeaderNode) RawData() []byte {
	return ebh.rawdata
}

func (ebh *EthBlockHeaderNode) Cid() cid.Cid {
	return *ebh.cid
}

func (EthBlockHeaderNode) String() string {
	return ""
}

func (EthBlockHeaderNode) Loggable() map[string]interface{} {
	panic("implement me")
}

func (EthBlockHeaderNode) Resolve(path []string) (interface{}, []string, error) {
	panic("implement me")
}

func (EthBlockHeaderNode) Tree(path string, depth int) []string {
	panic("implement me")
}

func (EthBlockHeaderNode) ResolveLink(path []string) (*format.Link, []string, error) {
	panic("implement me")
}

func (EthBlockHeaderNode) Copy() format.Node {
	panic("implement me")
}

func (EthBlockHeaderNode) Links() []*format.Link {
	panic("implement me")
}

func (EthBlockHeaderNode) Stat() (*format.NodeStat, error) {
	panic("implement me")
}

func (EthBlockHeaderNode) Size() (uint64, error) {
	panic("implement me")
}
