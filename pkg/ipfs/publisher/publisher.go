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

package publisher

import "fmt"

type Error struct {
	msg string
	err error
}

func (ie Error) Error() string {
	return fmt.Sprintf("%s: %s", ie.msg, ie.err.Error())
}

type Publisher interface {
	Write(input interface{}) ([]string, error)
}

type BlockDataPublisher struct {
	DagPutter
}

func NewIpfsPublisher(dagPutter DagPutter) *BlockDataPublisher {
	return &BlockDataPublisher{DagPutter: dagPutter}
}

func (ip *BlockDataPublisher) Write(input interface{}) ([]string, error) {
	cids, err := ip.DagPutter.DagPut(input)
	if err != nil {
		return nil, err
	}
	return cids, nil
}
