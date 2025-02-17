// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datacoord

import (
	"context"
	"sync"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/milvus-io/milvus/internal/metastore/model"
	"github.com/milvus-io/milvus/internal/mocks"
	"github.com/milvus-io/milvus/internal/proto/indexpb"
	"github.com/milvus-io/milvus/internal/types"
	"github.com/milvus-io/milvus/pkg/util/merr"
)

func TestIndexNodeManager_AddNode(t *testing.T) {
	nm := NewNodeManager(context.Background(), defaultIndexNodeCreatorFunc)
	nodeID, client := nm.PeekClient(&model.SegmentIndex{})
	assert.Equal(t, int64(-1), nodeID)
	assert.Nil(t, client)

	t.Run("success", func(t *testing.T) {
		err := nm.AddNode(1, "indexnode-1")
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		err := nm.AddNode(2, "")
		assert.Error(t, err)
	})
}

func TestIndexNodeManager_PeekClient(t *testing.T) {
	getMockedGetJobStatsClient := func(resp *indexpb.GetJobStatsResponse, err error) types.IndexNodeClient {
		ic := mocks.NewMockIndexNodeClient(t)
		ic.EXPECT().GetJobStats(mock.Anything, mock.Anything, mock.Anything).Return(resp, err)
		return ic
	}

	err := errors.New("error")

	t.Run("multiple unavailable IndexNode", func(t *testing.T) {
		nm := &IndexNodeManager{
			ctx: context.TODO(),
			nodeClients: map[UniqueID]types.IndexNodeClient{
				1: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status: merr.Status(err),
				}, err),
				2: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status: merr.Status(err),
				}, err),
				3: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status: merr.Status(err),
				}, err),
				4: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status: merr.Status(err),
				}, err),
				5: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status: merr.Status(err),
				}, nil),
				6: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status: merr.Status(err),
				}, nil),
				7: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status: merr.Status(err),
				}, nil),
				8: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					TaskSlots: 1,
					Status:    merr.Status(nil),
				}, nil),
				9: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					TaskSlots: 10,
					Status:    merr.Status(nil),
				}, nil),
			},
		}

		nodeID, client := nm.PeekClient(&model.SegmentIndex{})
		assert.NotNil(t, client)
		assert.Contains(t, []UniqueID{8, 9}, nodeID)
	})
}

func TestIndexNodeManager_ClientSupportDisk(t *testing.T) {
	getMockedGetJobStatsClient := func(resp *indexpb.GetJobStatsResponse, err error) types.IndexNodeClient {
		ic := mocks.NewMockIndexNodeClient(t)
		ic.EXPECT().GetJobStats(mock.Anything, mock.Anything, mock.Anything).Return(resp, err)
		return ic
	}

	err := errors.New("error")

	t.Run("support", func(t *testing.T) {
		nm := &IndexNodeManager{
			ctx:  context.Background(),
			lock: sync.RWMutex{},
			nodeClients: map[UniqueID]types.IndexNodeClient{
				1: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status:     merr.Status(nil),
					TaskSlots:  1,
					JobInfos:   nil,
					EnableDisk: true,
				}, nil),
			},
		}

		support := nm.ClientSupportDisk()
		assert.True(t, support)
	})

	t.Run("not support", func(t *testing.T) {
		nm := &IndexNodeManager{
			ctx:  context.Background(),
			lock: sync.RWMutex{},
			nodeClients: map[UniqueID]types.IndexNodeClient{
				1: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status:     merr.Status(nil),
					TaskSlots:  1,
					JobInfos:   nil,
					EnableDisk: false,
				}, nil),
			},
		}

		support := nm.ClientSupportDisk()
		assert.False(t, support)
	})

	t.Run("no indexnode", func(t *testing.T) {
		nm := &IndexNodeManager{
			ctx:         context.Background(),
			lock:        sync.RWMutex{},
			nodeClients: map[UniqueID]types.IndexNodeClient{},
		}

		support := nm.ClientSupportDisk()
		assert.False(t, support)
	})

	t.Run("error", func(t *testing.T) {
		nm := &IndexNodeManager{
			ctx:  context.Background(),
			lock: sync.RWMutex{},
			nodeClients: map[UniqueID]types.IndexNodeClient{
				1: getMockedGetJobStatsClient(nil, err),
			},
		}

		support := nm.ClientSupportDisk()
		assert.False(t, support)
	})

	t.Run("fail reason", func(t *testing.T) {
		nm := &IndexNodeManager{
			ctx:  context.Background(),
			lock: sync.RWMutex{},
			nodeClients: map[UniqueID]types.IndexNodeClient{
				1: getMockedGetJobStatsClient(&indexpb.GetJobStatsResponse{
					Status:     merr.Status(err),
					TaskSlots:  0,
					JobInfos:   nil,
					EnableDisk: false,
				}, nil),
			},
		}

		support := nm.ClientSupportDisk()
		assert.False(t, support)
	})
}

func TestNodeManager_StoppingNode(t *testing.T) {
	nm := NewNodeManager(context.Background(), defaultIndexNodeCreatorFunc)
	err := nm.AddNode(1, "indexnode-1")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(nm.GetAllClients()))

	nm.StoppingNode(1)
	assert.Equal(t, 0, len(nm.GetAllClients()))
	assert.Equal(t, 1, len(nm.stoppingNodes))

	nm.RemoveNode(1)
	assert.Equal(t, 0, len(nm.GetAllClients()))
	assert.Equal(t, 0, len(nm.stoppingNodes))
}
