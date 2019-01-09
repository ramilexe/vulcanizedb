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

import (
	. "github.com/onsi/gomega"

	ipld "gx/ipfs/QmWi2BYBL5gJ3CiAiQchg6rn1A8iBsrWy51EYxvHVjFvLb/go-ipld-format"
)

type MockAdder struct {
	calledCount int
	passedNodes []ipld.Node
	err         error
}

func NewMockAdder() *MockAdder {
	return &MockAdder{
		calledCount: 0,
		passedNodes: nil,
		err:         nil,
	}
}

func (ma *MockAdder) SetError(err error) {
	ma.err = err
}

func (ma *MockAdder) Add(node ipld.Node) error {
	ma.calledCount++
	ma.passedNodes = append(ma.passedNodes, node)
	return ma.err
}

func (ma *MockAdder) AssertAddCalled(times int, nodeType interface{}) {
	Expect(ma.calledCount).To(Equal(times))
	Expect(len(ma.passedNodes)).To(Equal(times))
	for _, passedNode := range ma.passedNodes {
		Expect(passedNode).To(BeAssignableToTypeOf(nodeType))
	}
}
