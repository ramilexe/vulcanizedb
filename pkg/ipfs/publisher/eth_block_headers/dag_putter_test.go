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

package eth_block_headers_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher/eth_block_headers"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/test_helpers"
	ipfs "github.com/vulcanize/vulcanizedb/pkg/ipfs/test_helpers/mocks/publisher"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/test_helpers/mocks/wrappers/rlp"
)

var _ = Describe("Creating an IPLD for a block header", func() {
	It("decodes passed bytes into ethereum block header", func() {
		mockDecoder := rlp.NewMockDecoder()
		mockDecoder.SetReturnOut(&types.Header{})
		dagPutter := eth_block_headers.NewBlockHeaderDagPutter(ipfs.NewMockAdder(), mockDecoder)
		fakeBytes := []byte{1, 2, 3, 4, 5}

		_, err := dagPutter.DagPut(fakeBytes)

		Expect(err).NotTo(HaveOccurred())
		mockDecoder.AssertDecodeCalledWith(fakeBytes, &types.Header{})
	})

	It("returns error if decoding fails", func() {
		mockDecoder := rlp.NewMockDecoder()
		mockDecoder.SetReturnOut(&types.Header{})
		mockDecoder.SetError(test_helpers.FakeError)
		dagPutter := eth_block_headers.NewBlockHeaderDagPutter(ipfs.NewMockAdder(), mockDecoder)

		_, err := dagPutter.DagPut([]byte{1, 2, 3, 4, 5})

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(test_helpers.FakeError))
	})

	It("adds ethereum block header to ipfs", func() {
		mockAdder := ipfs.NewMockAdder()
		mockDecoder := rlp.NewMockDecoder()
		mockDecoder.SetReturnOut(&types.Header{})
		dagPutter := eth_block_headers.NewBlockHeaderDagPutter(mockAdder, mockDecoder)
		fakeBytes := []byte{1, 2, 3, 4, 5}

		_, err := dagPutter.DagPut(fakeBytes)

		Expect(err).NotTo(HaveOccurred())
		mockAdder.AssertAddCalled(1, &eth_block_headers.EthBlockHeaderNode{})
	})

	It("returns error if adding to ipfs fails", func() {
		mockAdder := ipfs.NewMockAdder()
		mockAdder.SetError(test_helpers.FakeError)
		mockDecoder := rlp.NewMockDecoder()
		mockDecoder.SetReturnOut(&types.Header{})
		dagPutter := eth_block_headers.NewBlockHeaderDagPutter(mockAdder, mockDecoder)

		_, err := dagPutter.DagPut([]byte{1, 2, 3, 4, 5})

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(test_helpers.FakeError))
	})
})
