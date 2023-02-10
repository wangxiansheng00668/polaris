// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package mock

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:generate moq -out ./logs.mock.go -pkg mock ../../ LogsPlugin

// `NewEmptyLogsPlugin` returns an empty `LogsPluginMock`.
func NewEmptyLogsPlugin() *LogsPluginMock {
	// make and configure a mocked state.LogsPlugin
	return &LogsPluginMock{
		AddLogFunc: func(log *types.Log) {
			panic("mock out the AddLog method")
		},
		BuildLogsAndClearFunc: func(hash1 common.Hash, hash2 common.Hash, v1 uint, v2 uint) []*types.Log {
			panic("mock out the BuildLogsAndClear method")
		},
		FinalizeFunc: func() {
			// no-op
		},
		RegistryKeyFunc: func() string {
			return "emptylogs"
		},
		RevertToSnapshotFunc: func(n int) {
			// no-op
		},
		SnapshotFunc: func() int {
			// no-op
			return 0
		},
	}
}