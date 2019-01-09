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

package eth_block_receipts_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher/eth_block_receipts"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/test_helpers"
	ipfs "github.com/vulcanize/vulcanizedb/pkg/ipfs/test_helpers/mocks/publisher"
)

var _ = Describe("Eth block receipts dag putter", func() {
	It("adds a node for each receipt", func() {
		mockAdder := ipfs.NewMockAdder()
		dagPutter := eth_block_receipts.NewEthBlockReceiptDagPutter(mockAdder)
		fakeReceipts := types.Receipts{
			&types.Receipt{},
			&types.Receipt{},
		}

		_, err := dagPutter.DagPut(fakeReceipts)

		Expect(err).NotTo(HaveOccurred())
		mockAdder.AssertAddCalled(2, &eth_block_receipts.EthReceiptNode{})
	})

	It("returns error if adding node fails", func() {
		mockAdder := ipfs.NewMockAdder()
		mockAdder.SetError(test_helpers.FakeError)
		dagPutter := eth_block_receipts.NewEthBlockReceiptDagPutter(mockAdder)
		fakeReceipts := types.Receipts{
			&types.Receipt{},
			&types.Receipt{},
		}

		_, err := dagPutter.DagPut(fakeReceipts)

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(test_helpers.FakeError))
	})
})
