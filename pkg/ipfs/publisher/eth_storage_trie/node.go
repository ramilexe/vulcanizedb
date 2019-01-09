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
	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	"gx/ipfs/QmcKKBwfz6FyQdHR2jsXrrF6XeSBXYL86anmWNewpFpoF5/go-ipld-format"
)

type EthStorageTrieNode struct {
	cid     *cid.Cid
	rawdata []byte
}

func (estn *EthStorageTrieNode) RawData() []byte {
	return estn.rawdata
}

func (estn *EthStorageTrieNode) Cid() cid.Cid {
	return *estn.cid
}

func (*EthStorageTrieNode) String() string {
	panic("implement me")
}

func (*EthStorageTrieNode) Loggable() map[string]interface{} {
	panic("implement me")
}

func (*EthStorageTrieNode) Resolve(path []string) (interface{}, []string, error) {
	panic("implement me")
}

func (*EthStorageTrieNode) Tree(path string, depth int) []string {
	panic("implement me")
}

func (*EthStorageTrieNode) ResolveLink(path []string) (*format.Link, []string, error) {
	panic("implement me")
}

func (*EthStorageTrieNode) Copy() format.Node {
	panic("implement me")
}

func (*EthStorageTrieNode) Links() []*format.Link {
	panic("implement me")
}

func (*EthStorageTrieNode) Stat() (*format.NodeStat, error) {
	panic("implement me")
}

func (*EthStorageTrieNode) Size() (uint64, error) {
	panic("implement me")
}
