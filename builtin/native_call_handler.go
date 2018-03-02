package builtin

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethparams "github.com/ethereum/go-ethereum/params"
	"github.com/vechain/thor/builtin/abi"
	"github.com/vechain/thor/builtin/gen"
	"github.com/vechain/thor/builtin/native"
	"github.com/vechain/thor/state"
	"github.com/vechain/thor/thor"
	"github.com/vechain/thor/vm"
)

func mustLoadNativeABI(name string) *abi.ABI {
	data := gen.MustAsset("compiled/" + name + "Native.abi")
	abi, err := abi.New(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	return abi
}

var nativeCalls = map[thor.Address]struct {
	nativeABI *abi.ABI
	calls     map[string]*native.Callable
}{
	Params.Address: {
		mustLoadNativeABI(Params.name),
		map[string]*native.Callable{
			"nativeGetExecutor": {
				Gas: ethparams.SloadGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					return []interface{}{Executor.Address}, nil
				}},
			"nativeGet": {
				Gas: ethparams.SloadGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var key common.Hash
					env.Args(&key)
					v := Params.Get(env.State, thor.Hash(key))
					return []interface{}{v}, nil
				}},
			"nativeSet": {
				Gas:            ethparams.SstoreSetGas,
				RequiredCaller: &Params.Address,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var args struct {
						Key   common.Hash
						Value *big.Int
					}
					env.Args(&args)
					Params.Set(env.State, thor.Hash(args.Key), args.Value)
					return nil, nil
				},
			}},
	},
	Authority.Address: {
		mustLoadNativeABI(Authority.name),
		map[string]*native.Callable{
			"nativeGetExecutor": {
				Gas: ethparams.SloadGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					return []interface{}{Executor.Address}, nil
				}},
			"nativeAdd": {
				Gas:            ethparams.SstoreSetGas * 3,
				RequiredCaller: &Authority.Address,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var args struct {
						Addr     common.Address
						Identity common.Hash
					}
					env.Args(&args)
					ok := Authority.Add(env.State, thor.Address(args.Addr), thor.Hash(args.Identity))
					return []interface{}{ok}, nil
				}},
			"nativeRemove": {
				Gas:            ethparams.SstoreClearGas * 3,
				RequiredCaller: &Authority.Address,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var addr common.Address
					env.Args(&addr)
					ok := Authority.Remove(env.State, thor.Address(addr))
					return []interface{}{ok}, nil
				}},
			"nativeStatus": {
				Gas: ethparams.SloadGas * 3,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var addr common.Address
					env.Args(&addr)
					ok, identity, status := Authority.Status(env.State, thor.Address(addr))
					return []interface{}{ok, identity, status}, nil
				}},
			"nativeCount": {
				Gas: ethparams.SloadGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					count := Authority.Count(env.State)
					return []interface{}{count}, nil
				}},
		},
	},
	Energy.Address: {
		mustLoadNativeABI(Energy.name),
		map[string]*native.Callable{
			"nativeGetExecutor": {
				Gas: ethparams.SloadGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					return []interface{}{Executor.Address}, nil
				}},
			"nativeGetTotalSupply": {
				Gas: 3000,
				Proc: func(env *native.Env) ([]interface{}, error) {
					supply := Energy.GetTotalSupply(env.State, env.VMContext.Time)
					return []interface{}{supply}, nil
				}},
			"nativeGetTotalBurned": {
				Gas: ethparams.SloadGas * 2,
				Proc: func(env *native.Env) ([]interface{}, error) {
					burned := Energy.GetTotalBurned(env.State)
					return []interface{}{burned}, nil
				}},
			"nativeGetBalance": {
				Gas: 2000,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var addr common.Address
					env.Args(&addr)
					bal := Energy.GetBalance(env.State, env.VMContext.Time, thor.Address(addr))
					return []interface{}{bal}, nil
				}},
			"nativeAddBalance": {
				Gas:            ethparams.SstoreSetGas,
				RequiredCaller: &Energy.Address,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var args struct {
						Addr   common.Address
						Amount *big.Int
					}
					env.Args(&args)
					Energy.AddBalance(env.State, env.VMContext.Time, thor.Address(args.Addr), args.Amount)
					return nil, nil
				},
			},
			"nativeSubBalance": {
				Gas:            ethparams.SstoreResetGas,
				RequiredCaller: &Energy.Address,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var args struct {
						Addr   common.Address
						Amount *big.Int
					}
					env.Args(&args)
					ok := Energy.SubBalance(env.State, env.VMContext.Time, thor.Address(args.Addr), args.Amount)
					return []interface{}{ok}, nil
				},
			},
			"nativeAdjustGrowthRate": {
				Gas:            ethparams.SstoreSetGas,
				RequiredCaller: &Energy.Address,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var rate *big.Int
					env.Args(&rate)
					Energy.AdjustGrowthRate(env.State, env.VMContext.Time, rate)
					return nil, nil
				},
			},
			"nativeApproveConsumption": {
				Gas:            ethparams.SstoreSetGas,
				RequiredCaller: &Energy.Address,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var args struct {
						ContractAddr common.Address
						Caller       common.Address
						Credit       *big.Int
						RecoveryRate *big.Int
						Expiration   uint64
					}
					env.Args(&args)
					Energy.ApproveConsumption(env.State, env.VMContext.Time,
						thor.Address(args.ContractAddr), thor.Address(args.Caller), args.Credit, args.RecoveryRate, args.Expiration)
					return nil, nil
				},
			},
			"nativeGetConsumptionAllowance": {
				Gas: ethparams.SloadGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var args struct {
						ContractAddr common.Address
						Caller       common.Address
					}
					env.Args(&args)
					remained := Energy.GetConsumptionAllowance(env.State, env.VMContext.Time,
						thor.Address(args.ContractAddr), thor.Address(args.Caller))
					return []interface{}{remained}, nil
				},
			},
			"nativeSetSupplier": {
				Gas: ethparams.SstoreSetGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var args struct {
						ContractAddr common.Address
						Supplier     common.Address
						Agreed       bool
					}
					env.Args(&args)
					Energy.SetSupplier(env.State, thor.Address(args.ContractAddr), thor.Address(args.Supplier), args.Agreed)
					return nil, nil
				},
			},
			"nativeGetSupplier": {
				Gas: ethparams.SloadGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var contractAddr common.Address
					env.Args(&contractAddr)
					supplier, ok := Energy.GetSupplier(env.State, thor.Address(contractAddr))
					return []interface{}{supplier, ok}, nil
				},
			},
			"nativeSetContractMaster": {
				Gas:            ethparams.SstoreResetGas,
				RequiredCaller: &Energy.Address,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var args struct {
						ContractAddr common.Address
						Master       common.Address
					}
					env.Args(&args)
					Energy.SetContractMaster(env.State, thor.Address(args.ContractAddr), thor.Address(args.Master))
					return nil, nil
				},
			},
			"nativeGetContractMaster": {
				Gas: ethparams.SloadGas,
				Proc: func(env *native.Env) ([]interface{}, error) {
					var contractAddr common.Address
					env.Args(&contractAddr)
					master := Energy.GetContractMaster(env.State, thor.Address(contractAddr))
					return []interface{}{master}, nil
				},
			},
		},
	},
}

func init() {
	for _, contract := range nativeCalls {
		for name, call := range contract.calls {
			call.MethodCodec = contract.nativeABI.MustForMethod(name)
		}
	}
}

func HandleNativeCall(state *state.State, vmCtx *vm.Context, to thor.Address, input []byte) func(useGas func(gas uint64) bool, caller thor.Address) ([]byte, error) {
	contract, found := nativeCalls[to]
	if !found {
		return nil
	}

	name, err := contract.nativeABI.MethodName(input)
	if err != nil {
		return nil
	}

	if call, ok := contract.calls[name]; ok {
		return func(useGas func(gas uint64) bool, caller thor.Address) ([]byte, error) {
			return call.Call(state, vmCtx, caller, useGas, input)
		}
	}
	return nil
}
