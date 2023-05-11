// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	gethvm "github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
	"math/big"
	"sync"
)

// Ensure, that PrecompileEVMMock does implement PrecompileEVM.
// If this is not the case, regenerate this file with moq.
var _ PrecompileEVM = &PrecompileEVMMock{}

// PrecompileEVMMock is a mock implementation of PrecompileEVM.
//
//	func TestSomethingThatUsesPrecompileEVM(t *testing.T) {
//
//		// make and configure a mocked PrecompileEVM
//		mockedPrecompileEVM := &PrecompileEVMMock{
//			CallFunc: func(caller gethvm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
//				panic("mock out the Call method")
//			},
//			CreateFunc: func(caller gethvm.ContractRef, code []byte, gas uint64, value *big.Int) ([]byte, common.Address, uint64, error) {
//				panic("mock out the Create method")
//			},
//			Create2Func: func(caller gethvm.ContractRef, code []byte, gas uint64, endowment *big.Int, salt *uint256.Int) ([]byte, common.Address, uint64, error) {
//				panic("mock out the Create2 method")
//			},
//			GetContextFunc: func() *gethvm.BlockContext {
//				panic("mock out the GetContext method")
//			},
//			GetStateDBFunc: func() state.StateDBI {
//				panic("mock out the GetStateDB method")
//			},
//			StaticCallFunc: func(caller gethvm.ContractRef, addr common.Address, input []byte, gas uint64) ([]byte, uint64, error) {
//				panic("mock out the StaticCall method")
//			},
//		}
//
//		// use mockedPrecompileEVM in code that requires PrecompileEVM
//		// and then make assertions.
//
//	}
type PrecompileEVMMock struct {
	// CallFunc mocks the Call method.
	CallFunc func(caller gethvm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) ([]byte, uint64, error)

	// CreateFunc mocks the Create method.
	CreateFunc func(caller gethvm.ContractRef, code []byte, gas uint64, value *big.Int) ([]byte, common.Address, uint64, error)

	// Create2Func mocks the Create2 method.
	Create2Func func(caller gethvm.ContractRef, code []byte, gas uint64, endowment *big.Int, salt *uint256.Int) ([]byte, common.Address, uint64, error)

	// GetContextFunc mocks the GetContext method.
	GetContextFunc func() *gethvm.BlockContext

	// GetStateDBFunc mocks the GetStateDB method.
	GetStateDBFunc func() state.StateDBI

	// StaticCallFunc mocks the StaticCall method.
	StaticCallFunc func(caller gethvm.ContractRef, addr common.Address, input []byte, gas uint64) ([]byte, uint64, error)

	// calls tracks calls to the methods.
	calls struct {
		// Call holds details about calls to the Call method.
		Call []struct {
			// Caller is the caller argument value.
			Caller gethvm.ContractRef
			// Addr is the addr argument value.
			Addr common.Address
			// Input is the input argument value.
			Input []byte
			// Gas is the gas argument value.
			Gas uint64
			// Value is the value argument value.
			Value *big.Int
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// Caller is the caller argument value.
			Caller gethvm.ContractRef
			// Code is the code argument value.
			Code []byte
			// Gas is the gas argument value.
			Gas uint64
			// Value is the value argument value.
			Value *big.Int
		}
		// Create2 holds details about calls to the Create2 method.
		Create2 []struct {
			// Caller is the caller argument value.
			Caller gethvm.ContractRef
			// Code is the code argument value.
			Code []byte
			// Gas is the gas argument value.
			Gas uint64
			// Endowment is the endowment argument value.
			Endowment *big.Int
			// Salt is the salt argument value.
			Salt *uint256.Int
		}
		// GetContext holds details about calls to the GetContext method.
		GetContext []struct {
		}
		// GetStateDB holds details about calls to the GetStateDB method.
		GetStateDB []struct {
		}
		// StaticCall holds details about calls to the StaticCall method.
		StaticCall []struct {
			// Caller is the caller argument value.
			Caller gethvm.ContractRef
			// Addr is the addr argument value.
			Addr common.Address
			// Input is the input argument value.
			Input []byte
			// Gas is the gas argument value.
			Gas uint64
		}
	}
	lockCall       sync.RWMutex
	lockCreate     sync.RWMutex
	lockCreate2    sync.RWMutex
	lockGetContext sync.RWMutex
	lockGetStateDB sync.RWMutex
	lockStaticCall sync.RWMutex
}

// Call calls CallFunc.
func (mock *PrecompileEVMMock) Call(caller gethvm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) ([]byte, uint64, error) {
	if mock.CallFunc == nil {
		panic("PrecompileEVMMock.CallFunc: method is nil but PrecompileEVM.Call was just called")
	}
	callInfo := struct {
		Caller gethvm.ContractRef
		Addr   common.Address
		Input  []byte
		Gas    uint64
		Value  *big.Int
	}{
		Caller: caller,
		Addr:   addr,
		Input:  input,
		Gas:    gas,
		Value:  value,
	}
	mock.lockCall.Lock()
	mock.calls.Call = append(mock.calls.Call, callInfo)
	mock.lockCall.Unlock()
	return mock.CallFunc(caller, addr, input, gas, value)
}

// CallCalls gets all the calls that were made to Call.
// Check the length with:
//
//	len(mockedPrecompileEVM.CallCalls())
func (mock *PrecompileEVMMock) CallCalls() []struct {
	Caller gethvm.ContractRef
	Addr   common.Address
	Input  []byte
	Gas    uint64
	Value  *big.Int
} {
	var calls []struct {
		Caller gethvm.ContractRef
		Addr   common.Address
		Input  []byte
		Gas    uint64
		Value  *big.Int
	}
	mock.lockCall.RLock()
	calls = mock.calls.Call
	mock.lockCall.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *PrecompileEVMMock) Create(caller gethvm.ContractRef, code []byte, gas uint64, value *big.Int) ([]byte, common.Address, uint64, error) {
	if mock.CreateFunc == nil {
		panic("PrecompileEVMMock.CreateFunc: method is nil but PrecompileEVM.Create was just called")
	}
	callInfo := struct {
		Caller gethvm.ContractRef
		Code   []byte
		Gas    uint64
		Value  *big.Int
	}{
		Caller: caller,
		Code:   code,
		Gas:    gas,
		Value:  value,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(caller, code, gas, value)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedPrecompileEVM.CreateCalls())
func (mock *PrecompileEVMMock) CreateCalls() []struct {
	Caller gethvm.ContractRef
	Code   []byte
	Gas    uint64
	Value  *big.Int
} {
	var calls []struct {
		Caller gethvm.ContractRef
		Code   []byte
		Gas    uint64
		Value  *big.Int
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Create2 calls Create2Func.
func (mock *PrecompileEVMMock) Create2(caller gethvm.ContractRef, code []byte, gas uint64, endowment *big.Int, salt *uint256.Int) ([]byte, common.Address, uint64, error) {
	if mock.Create2Func == nil {
		panic("PrecompileEVMMock.Create2Func: method is nil but PrecompileEVM.Create2 was just called")
	}
	callInfo := struct {
		Caller    gethvm.ContractRef
		Code      []byte
		Gas       uint64
		Endowment *big.Int
		Salt      *uint256.Int
	}{
		Caller:    caller,
		Code:      code,
		Gas:       gas,
		Endowment: endowment,
		Salt:      salt,
	}
	mock.lockCreate2.Lock()
	mock.calls.Create2 = append(mock.calls.Create2, callInfo)
	mock.lockCreate2.Unlock()
	return mock.Create2Func(caller, code, gas, endowment, salt)
}

// Create2Calls gets all the calls that were made to Create2.
// Check the length with:
//
//	len(mockedPrecompileEVM.Create2Calls())
func (mock *PrecompileEVMMock) Create2Calls() []struct {
	Caller    gethvm.ContractRef
	Code      []byte
	Gas       uint64
	Endowment *big.Int
	Salt      *uint256.Int
} {
	var calls []struct {
		Caller    gethvm.ContractRef
		Code      []byte
		Gas       uint64
		Endowment *big.Int
		Salt      *uint256.Int
	}
	mock.lockCreate2.RLock()
	calls = mock.calls.Create2
	mock.lockCreate2.RUnlock()
	return calls
}

// GetContext calls GetContextFunc.
func (mock *PrecompileEVMMock) GetContext() *gethvm.BlockContext {
	if mock.GetContextFunc == nil {
		panic("PrecompileEVMMock.GetContextFunc: method is nil but PrecompileEVM.GetContext was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetContext.Lock()
	mock.calls.GetContext = append(mock.calls.GetContext, callInfo)
	mock.lockGetContext.Unlock()
	return mock.GetContextFunc()
}

// GetContextCalls gets all the calls that were made to GetContext.
// Check the length with:
//
//	len(mockedPrecompileEVM.GetContextCalls())
func (mock *PrecompileEVMMock) GetContextCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetContext.RLock()
	calls = mock.calls.GetContext
	mock.lockGetContext.RUnlock()
	return calls
}

// GetStateDB calls GetStateDBFunc.
func (mock *PrecompileEVMMock) GetStateDB() state.StateDBI {
	if mock.GetStateDBFunc == nil {
		panic("PrecompileEVMMock.GetStateDBFunc: method is nil but PrecompileEVM.GetStateDB was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetStateDB.Lock()
	mock.calls.GetStateDB = append(mock.calls.GetStateDB, callInfo)
	mock.lockGetStateDB.Unlock()
	return mock.GetStateDBFunc()
}

// GetStateDBCalls gets all the calls that were made to GetStateDB.
// Check the length with:
//
//	len(mockedPrecompileEVM.GetStateDBCalls())
func (mock *PrecompileEVMMock) GetStateDBCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetStateDB.RLock()
	calls = mock.calls.GetStateDB
	mock.lockGetStateDB.RUnlock()
	return calls
}

// StaticCall calls StaticCallFunc.
func (mock *PrecompileEVMMock) StaticCall(caller gethvm.ContractRef, addr common.Address, input []byte, gas uint64) ([]byte, uint64, error) {
	if mock.StaticCallFunc == nil {
		panic("PrecompileEVMMock.StaticCallFunc: method is nil but PrecompileEVM.StaticCall was just called")
	}
	callInfo := struct {
		Caller gethvm.ContractRef
		Addr   common.Address
		Input  []byte
		Gas    uint64
	}{
		Caller: caller,
		Addr:   addr,
		Input:  input,
		Gas:    gas,
	}
	mock.lockStaticCall.Lock()
	mock.calls.StaticCall = append(mock.calls.StaticCall, callInfo)
	mock.lockStaticCall.Unlock()
	return mock.StaticCallFunc(caller, addr, input, gas)
}

// StaticCallCalls gets all the calls that were made to StaticCall.
// Check the length with:
//
//	len(mockedPrecompileEVM.StaticCallCalls())
func (mock *PrecompileEVMMock) StaticCallCalls() []struct {
	Caller gethvm.ContractRef
	Addr   common.Address
	Input  []byte
	Gas    uint64
} {
	var calls []struct {
		Caller gethvm.ContractRef
		Addr   common.Address
		Input  []byte
		Gas    uint64
	}
	mock.lockStaticCall.RLock()
	calls = mock.calls.StaticCall
	mock.lockStaticCall.RUnlock()
	return calls
}

// Ensure, that MessageRouterMock does implement MessageRouter.
// If this is not the case, regenerate this file with moq.
var _ MessageRouter = &MessageRouterMock{}

// MessageRouterMock is a mock implementation of MessageRouter.
//
//	func TestSomethingThatUsesMessageRouter(t *testing.T) {
//
//		// make and configure a mocked MessageRouter
//		mockedMessageRouter := &MessageRouterMock{
//			HandlerFunc: func(msg sdk.Msg) func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error) {
//				panic("mock out the Handler method")
//			},
//			HandlerByTypeURLFunc: func(typeURL string) func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error) {
//				panic("mock out the HandlerByTypeURL method")
//			},
//		}
//
//		// use mockedMessageRouter in code that requires MessageRouter
//		// and then make assertions.
//
//	}
type MessageRouterMock struct {
	// HandlerFunc mocks the Handler method.
	HandlerFunc func(msg sdk.Msg) func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error)

	// HandlerByTypeURLFunc mocks the HandlerByTypeURL method.
	HandlerByTypeURLFunc func(typeURL string) func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error)

	// calls tracks calls to the methods.
	calls struct {
		// Handler holds details about calls to the Handler method.
		Handler []struct {
			// Msg is the msg argument value.
			Msg sdk.Msg
		}
		// HandlerByTypeURL holds details about calls to the HandlerByTypeURL method.
		HandlerByTypeURL []struct {
			// TypeURL is the typeURL argument value.
			TypeURL string
		}
	}
	lockHandler          sync.RWMutex
	lockHandlerByTypeURL sync.RWMutex
}

// Handler calls HandlerFunc.
func (mock *MessageRouterMock) Handler(msg sdk.Msg) func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error) {
	if mock.HandlerFunc == nil {
		panic("MessageRouterMock.HandlerFunc: method is nil but MessageRouter.Handler was just called")
	}
	callInfo := struct {
		Msg sdk.Msg
	}{
		Msg: msg,
	}
	mock.lockHandler.Lock()
	mock.calls.Handler = append(mock.calls.Handler, callInfo)
	mock.lockHandler.Unlock()
	return mock.HandlerFunc(msg)
}

// HandlerCalls gets all the calls that were made to Handler.
// Check the length with:
//
//	len(mockedMessageRouter.HandlerCalls())
func (mock *MessageRouterMock) HandlerCalls() []struct {
	Msg sdk.Msg
} {
	var calls []struct {
		Msg sdk.Msg
	}
	mock.lockHandler.RLock()
	calls = mock.calls.Handler
	mock.lockHandler.RUnlock()
	return calls
}

// HandlerByTypeURL calls HandlerByTypeURLFunc.
func (mock *MessageRouterMock) HandlerByTypeURL(typeURL string) func(ctx sdk.Context, req sdk.Msg) (*sdk.Result, error) {
	if mock.HandlerByTypeURLFunc == nil {
		panic("MessageRouterMock.HandlerByTypeURLFunc: method is nil but MessageRouter.HandlerByTypeURL was just called")
	}
	callInfo := struct {
		TypeURL string
	}{
		TypeURL: typeURL,
	}
	mock.lockHandlerByTypeURL.Lock()
	mock.calls.HandlerByTypeURL = append(mock.calls.HandlerByTypeURL, callInfo)
	mock.lockHandlerByTypeURL.Unlock()
	return mock.HandlerByTypeURLFunc(typeURL)
}

// HandlerByTypeURLCalls gets all the calls that were made to HandlerByTypeURL.
// Check the length with:
//
//	len(mockedMessageRouter.HandlerByTypeURLCalls())
func (mock *MessageRouterMock) HandlerByTypeURLCalls() []struct {
	TypeURL string
} {
	var calls []struct {
		TypeURL string
	}
	mock.lockHandlerByTypeURL.RLock()
	calls = mock.calls.HandlerByTypeURL
	mock.lockHandlerByTypeURL.RUnlock()
	return calls
}