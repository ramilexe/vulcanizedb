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

package rlp

import (
	"reflect"

	. "github.com/onsi/gomega"
)

type MockDecoder struct {
	called      bool
	err         error
	passedBytes []byte
	passedOut   interface{}
	returnOut   interface{}
}

func NewMockDecoder() *MockDecoder {
	return &MockDecoder{
		called:      false,
		err:         nil,
		passedBytes: nil,
		passedOut:   nil,
		returnOut:   nil,
	}
}

func (md *MockDecoder) SetError(err error) {
	md.err = err
}

func (md *MockDecoder) SetReturnOut(out interface{}) {
	md.returnOut = out
}

func (md *MockDecoder) Decode(raw []byte, out interface{}) error {
	md.called = true
	md.passedBytes = raw
	md.passedOut = out
	valToAssign := reflect.ValueOf(md.returnOut).Elem()
	reflect.ValueOf(out).Elem().Set(valToAssign)
	return md.err
}

func (md *MockDecoder) AssertDecodeCalledWith(raw []byte, out interface{}) {
	Expect(md.called).To(BeTrue())
	Expect(md.passedBytes).To(Equal(raw))
	Expect(md.passedOut).To(BeAssignableToTypeOf(out))
}
