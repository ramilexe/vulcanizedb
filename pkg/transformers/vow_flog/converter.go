// Copyright 2018 Vulcanize
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vow_flog

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/shared/constants"
)

type VowFlogConverter struct{}

func (VowFlogConverter) ToModels(ethLogs []types.Log) ([]interface{}, error) {
	var models []interface{}
	for _, ethLog := range ethLogs {
		err := verifyLog(ethLog)
		if err != nil {
			return nil, err
		}

		era := ethLog.Topics[2].Big()

		raw, err := json.Marshal(ethLog)
		if err != nil {
			return nil, err
		}
		model := VowFlogModel{
			Era:              era.String(),
			LogIndex:         ethLog.Index,
			TransactionIndex: ethLog.TxIndex,
			Raw:              raw,
		}
		models = append(models, model)
	}
	return models, nil
}

func verifyLog(log types.Log) error {
	if len(log.Topics) < 3 {
		return errors.New("log missing topics")
	}
	if len(log.Data) < constants.DataItemLength {
		return errors.New("log missing data")
	}
	return nil
}
