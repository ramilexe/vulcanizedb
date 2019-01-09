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

package publisher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	ipfs "github.com/vulcanize/vulcanizedb/pkg/ipfs/publisher"
	"github.com/vulcanize/vulcanizedb/pkg/ipfs/test_helpers"
	ipfs_publisher "github.com/vulcanize/vulcanizedb/pkg/ipfs/test_helpers/mocks/publisher"
)

var _ = Describe("IPFS publisher", func() {
	It("calls dag put with the passed data", func() {
		mockDagPutter := ipfs_publisher.NewMockDagPutter()
		publisher := ipfs.NewIpfsPublisher(mockDagPutter)
		fakeBytes := []byte{1, 2, 3, 4, 5}

		_, err := publisher.DagPut(fakeBytes)

		Expect(err).NotTo(HaveOccurred())
		Expect(mockDagPutter.Called).To(BeTrue())
		Expect(mockDagPutter.PassedInterface).To(Equal(fakeBytes))
	})

	It("returns error if dag put fails", func() {
		mockDagPutter := ipfs_publisher.NewMockDagPutter()
		mockDagPutter.SetError(test_helpers.FakeError)
		publisher := ipfs.NewIpfsPublisher(mockDagPutter)

		_, err := publisher.DagPut([]byte{1, 2, 3, 4, 5})

		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(test_helpers.FakeError))
	})
})
