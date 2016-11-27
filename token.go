// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// VinTrackerABI is the input ABI used to generate the binding from.
const VinTrackerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"vin\",\"type\":\"string\"},{\"name\":\"action\",\"type\":\"string\"}],\"name\":\"addEvent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vin\",\"type\":\"string\"}],\"name\":\"getEventsByVIN\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"history\",\"type\":\"uint256\"}],\"name\":\"getEventById\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"PrintInt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"str\",\"type\":\"string\"}],\"name\":\"PrintString\",\"type\":\"event\"}]"

// VinTrackerBin is the compiled bytecode used for deploying new contracts.
const VinTrackerBin = `0x606060405234610000575b600080555b5b610ac68061001e6000396000f3606060405260e060020a60003504631b26be1d811461003457806356b83f1a146100d6578063b98079bc1461012b575b610000565b34610000576100c4600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284375050604080516020601f89358b018035918201839004830284018301909452808352979998810197919650918201945092508291508401838280828437509496506101a995505050505050565b60408051918252519081900360200190f35b3461000057610129600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437509496506106c495505050505050565b005b346100005761013b60043561094d565b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f16801561019b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60006000600384604051808280519060200190808383829060006004602084601f0104600302600f01f150919091019384525050604051918290036020019091205460ff1615905061023b57600284604051808280519060200190808383829060006004602084601f0104600302600f01f150905001915050908152602001604051809103902060010154905061029b565b6001600385604051808280519060200190808383829060006004602084601f0104600302600f01f150905001915050908152602001604051809103902060006101000a81548160ff021916908360f860020a908102040217905550600090505b60a060405190810160405280828152602001600060008154600101919050819055815260200185815260200133815260200184815260200150600285604051808280519060200190808383829060006004602084601f0104600302600f01f150905001915050908152602001604051809103902060008201518160000155602082015181600101556040820151816002019080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061037557805160ff19168380011785556103a2565b828001600101855582156103a2579182015b828111156103a2578251825591602001919060010190610387565b5b506103c39291505b808211156103bf57600081556001016103ab565b5090565b505060608201518160030160006101000a815481600160a060020a0302191690836c010000000000000000000000009081020402179055506080820151816004019080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061044d57805160ff191683800117855561047a565b8280016001018555821561047a579182015b8281111561047a57825182559160200191906001019061045f565b5b5061049b9291505b808211156103bf57600081556001016103ab565b5090565b5050905050600284604051808280519060200190808383829060006004602084601f0104600302600f01f1509050019150509081526020016040518091039020600160006000548152602001908152602001600020600082015481600001556001820154816001015560028201816002019080546001816001161561010002031660029004828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105595780548555610595565b8280016001018555821561059557600052602060002091601f016020900482015b8281111561059557825482559160010191906001019061057a565b5b506105b69291505b808211156103bf57600081556001016103ab565b5090565b50506003820160009054906101000a9004600160a060020a03168160030160006101000a815481600160a060020a0302191690836c0100000000000000000000000090810204021790555060048201816004019080546001816001161561010002031660029004828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106565780548555610692565b8280016001018555821561069257600052602060002091601f016020900482015b82811115610692578254825591600101919060010190610677565b5b506106b39291505b808211156103bf57600081556001016103ab565b5090565b5050600054935050505b5092915050565b60006000600383604051808280519060200190808383829060006004602084601f0104600302600f01f150919091019384525050604051918290036020019091205460ff161590506108f857600283604051808280519060200190808383829060006004602084601f0104600302600f01f1509190910193845250506040805192839003602090810184208185526004018054600260001961010060018416150201909116049185018290529550600080516020610aa68339815191529386935091829190820190849080156107db5780601f106107b0576101008083540402835291602001916107db565b820191906000526020600020905b8154815290600101906020018083116107be57829003601f168201915b50509250505060405180910390a1600283604051808280519060200190808383829060006004602084601f0104600302600f01f15090500191505090815260200160405180910390206000015490505b60008111156108f35760008181526001602081815260409283902083518281526004909101805460029481161561010002600019011693909304918101829052600080516020610aa68339815191529390918291820190849080156108d15780601f106108a6576101008083540402835291602001916108d1565b820191906000526020600020905b8154815290600101906020018083116108b457829003601f168201915b50509250505060405180910390a160009081526001602052604090205461082b565b610947565b6040805160208082526009908201527f6e6f7420666f756e640000000000000000000000000000000000000000000000818301529051600080516020610aa68339815191529181900360600190a15b5b505050565b60408051602080820183526000808352848152600180835290849020845183815260049091018054600293811615610100026000190116929092049281018390529293600080516020610aa68339815191529391928291820190849080156109f65780601f106109cb576101008083540402835291602001916109f6565b820191906000526020600020905b8154815290600101906020018083116109d957829003601f168201915b50509250505060405180910390a160008281526001602081815260409283902060040180548451600294821615610100026000190190911693909304601f8101839004830284018301909452838352919290830182828015610a995780601f10610a6e57610100808354040283529160200191610a99565b820191906000526020600020905b815481529060010190602001808311610a7c57829003601f168201915b505050505090505b9190505643123f7005ece31cd2478fa2cd0bec5ea2e353c1c3fe9ca390a6de2ab917eac9`

// DeployVinTracker deploys a new Ethereum contract, binding an instance of VinTracker to it.
func DeployVinTracker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VinTracker, error) {
	parsed, err := abi.JSON(strings.NewReader(VinTrackerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VinTrackerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VinTracker{VinTrackerCaller: VinTrackerCaller{contract: contract}, VinTrackerTransactor: VinTrackerTransactor{contract: contract}}, nil
}

// VinTracker is an auto generated Go binding around an Ethereum contract.
type VinTracker struct {
	VinTrackerCaller     // Read-only binding to the contract
	VinTrackerTransactor // Write-only binding to the contract
}

// VinTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VinTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VinTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VinTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VinTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VinTrackerSession struct {
	Contract     *VinTracker       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VinTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VinTrackerCallerSession struct {
	Contract *VinTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// VinTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VinTrackerTransactorSession struct {
	Contract     *VinTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VinTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VinTrackerRaw struct {
	Contract *VinTracker // Generic contract binding to access the raw methods on
}

// VinTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VinTrackerCallerRaw struct {
	Contract *VinTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// VinTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VinTrackerTransactorRaw struct {
	Contract *VinTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVinTracker creates a new instance of VinTracker, bound to a specific deployed contract.
func NewVinTracker(address common.Address, backend bind.ContractBackend) (*VinTracker, error) {
	contract, err := bindVinTracker(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VinTracker{VinTrackerCaller: VinTrackerCaller{contract: contract}, VinTrackerTransactor: VinTrackerTransactor{contract: contract}}, nil
}

// NewVinTrackerCaller creates a new read-only instance of VinTracker, bound to a specific deployed contract.
func NewVinTrackerCaller(address common.Address, caller bind.ContractCaller) (*VinTrackerCaller, error) {
	contract, err := bindVinTracker(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VinTrackerCaller{contract: contract}, nil
}

// NewVinTrackerTransactor creates a new write-only instance of VinTracker, bound to a specific deployed contract.
func NewVinTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*VinTrackerTransactor, error) {
	contract, err := bindVinTracker(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VinTrackerTransactor{contract: contract}, nil
}

// bindVinTracker binds a generic wrapper to an already deployed contract.
func bindVinTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VinTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VinTracker *VinTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VinTracker.Contract.VinTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VinTracker *VinTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VinTracker.Contract.VinTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VinTracker *VinTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VinTracker.Contract.VinTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VinTracker *VinTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VinTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VinTracker *VinTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VinTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VinTracker *VinTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VinTracker.Contract.contract.Transact(opts, method, params...)
}

// AddEvent is a paid mutator transaction binding the contract method 0x1b26be1d.
//
// Solidity: function addEvent(vin string, action string) returns(uint256)
func (_VinTracker *VinTrackerTransactor) AddEvent(opts *bind.TransactOpts, vin string, action string) (*types.Transaction, error) {
	return _VinTracker.contract.Transact(opts, "addEvent", vin, action)
}

// AddEvent is a paid mutator transaction binding the contract method 0x1b26be1d.
//
// Solidity: function addEvent(vin string, action string) returns(uint256)
func (_VinTracker *VinTrackerSession) AddEvent(vin string, action string) (*types.Transaction, error) {
	return _VinTracker.Contract.AddEvent(&_VinTracker.TransactOpts, vin, action)
}

// AddEvent is a paid mutator transaction binding the contract method 0x1b26be1d.
//
// Solidity: function addEvent(vin string, action string) returns(uint256)
func (_VinTracker *VinTrackerTransactorSession) AddEvent(vin string, action string) (*types.Transaction, error) {
	return _VinTracker.Contract.AddEvent(&_VinTracker.TransactOpts, vin, action)
}

// GetEventById is a paid mutator transaction binding the contract method 0xb98079bc.
//
// Solidity: function getEventById(history uint256) returns(string)
func (_VinTracker *VinTrackerTransactor) GetEventById(opts *bind.TransactOpts, history *big.Int) (*types.Transaction, error) {
	return _VinTracker.contract.Transact(opts, "getEventById", history)
}

// GetEventById is a paid mutator transaction binding the contract method 0xb98079bc.
//
// Solidity: function getEventById(history uint256) returns(string)
func (_VinTracker *VinTrackerSession) GetEventById(history *big.Int) (*types.Transaction, error) {
	return _VinTracker.Contract.GetEventById(&_VinTracker.TransactOpts, history)
}

// GetEventById is a paid mutator transaction binding the contract method 0xb98079bc.
//
// Solidity: function getEventById(history uint256) returns(string)
func (_VinTracker *VinTrackerTransactorSession) GetEventById(history *big.Int) (*types.Transaction, error) {
	return _VinTracker.Contract.GetEventById(&_VinTracker.TransactOpts, history)
}

// GetEventsByVIN is a paid mutator transaction binding the contract method 0x56b83f1a.
//
// Solidity: function getEventsByVIN(vin string) returns()
func (_VinTracker *VinTrackerTransactor) GetEventsByVIN(opts *bind.TransactOpts, vin string) (*types.Transaction, error) {
	return _VinTracker.contract.Transact(opts, "getEventsByVIN", vin)
}

// GetEventsByVIN is a paid mutator transaction binding the contract method 0x56b83f1a.
//
// Solidity: function getEventsByVIN(vin string) returns()
func (_VinTracker *VinTrackerSession) GetEventsByVIN(vin string) (*types.Transaction, error) {
	return _VinTracker.Contract.GetEventsByVIN(&_VinTracker.TransactOpts, vin)
}

// GetEventsByVIN is a paid mutator transaction binding the contract method 0x56b83f1a.
//
// Solidity: function getEventsByVIN(vin string) returns()
func (_VinTracker *VinTrackerTransactorSession) GetEventsByVIN(vin string) (*types.Transaction, error) {
	return _VinTracker.Contract.GetEventsByVIN(&_VinTracker.TransactOpts, vin)
}
