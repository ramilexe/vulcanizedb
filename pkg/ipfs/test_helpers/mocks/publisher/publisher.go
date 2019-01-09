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
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/ipfs/test_helpers"
)

type MockPublisher struct {
	err              error
	passedBlockDatas []interface{}
	returnStrings    [][]string
}

func NewMockPublisher() *MockPublisher {
	return &MockPublisher{
		err:              nil,
		passedBlockDatas: []interface{}{},
		returnStrings:    nil,
	}
}

func (publisher *MockPublisher) SetReturnStrings(returnBytes [][]string) {
	publisher.returnStrings = returnBytes
}

func (publisher *MockPublisher) SetError(err error) {
	publisher.err = err
}

func (publisher *MockPublisher) Write(input interface{}) ([]string, error) {
	publisher.passedBlockDatas = append(publisher.passedBlockDatas, input)
	if publisher.err != nil {
		return nil, publisher.err
	}
	var stringsToReturn []string
	if len(publisher.returnStrings) > 0 {
		stringsToReturn = publisher.returnStrings[0]
		if len(publisher.returnStrings) > 1 {
			publisher.returnStrings = publisher.returnStrings[1:]
		} else {
			publisher.returnStrings = [][]string{{test_helpers.FakeString}}
		}
	} else {
		stringsToReturn = []string{test_helpers.FakeString}
	}
	return stringsToReturn, nil
}

func (publisher *MockPublisher) AssertWriteCalledWithBytes(inputs [][]byte) {
	for i := 0; i < len(inputs); i++ {
		Expect(publisher.passedBlockDatas).To(ContainElement(inputs[i]))
	}
	for i := 0; i < len(publisher.passedBlockDatas); i++ {
		Expect(inputs).To(ContainElement(publisher.passedBlockDatas[i]))
	}
}

func (publisher *MockPublisher) AssertWriteCalledWithInterfaces(interfaces []interface{}) {
	for i := 0; i < len(interfaces); i++ {
		Expect(publisher.passedBlockDatas).To(ContainElement(interfaces[i]))
	}
	for i := 0; i < len(publisher.passedBlockDatas); i++ {
		Expect(interfaces).To(ContainElement(publisher.passedBlockDatas[i]))
	}
}

func (publisher *MockPublisher) AssertWriteCalledWithBodies(bodies []*types.Body) {
	var expected []*types.Body
	for i := 0; i < len(publisher.passedBlockDatas); i++ {
		expected = append(expected, publisher.passedBlockDatas[i].(*types.Body))
	}
	Expect(expected).To(Equal(bodies))
}
