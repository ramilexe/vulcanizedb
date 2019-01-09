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

type MockDagPutter struct {
	Called          bool
	PassedInterface interface{}
	Err             error
}

func NewMockDagPutter() *MockDagPutter {
	return &MockDagPutter{
		Called:          false,
		PassedInterface: nil,
		Err:             nil,
	}
}

func (mdp *MockDagPutter) SetError(err error) {
	mdp.Err = err
}

func (mdp *MockDagPutter) DagPut(raw interface{}) ([]string, error) {
	mdp.Called = true
	mdp.PassedInterface = raw
	return nil, mdp.Err
}
