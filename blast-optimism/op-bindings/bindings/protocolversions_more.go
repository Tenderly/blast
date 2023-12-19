// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const ProtocolVersionsStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/ProtocolVersions.sol:ProtocolVersions\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/ProtocolVersions.sol:ProtocolVersions\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/ProtocolVersions.sol:ProtocolVersions\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/ProtocolVersions.sol:ProtocolVersions\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/ProtocolVersions.sol:ProtocolVersions\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var ProtocolVersionsStorageLayout = new(solc.StorageLayout)

var ProtocolVersionsDeployedBin = "0x608060405234801561001057600080fd5b50600436106100d45760003560e01c80638da5cb5b11610081578063f2fde38b1161005b578063f2fde38b146101b8578063f7d12760146101cb578063ffa1ad74146101d357600080fd5b80638da5cb5b14610180578063d798b1ac146101a8578063dc8452cd146101b057600080fd5b80635fd579af116100b25780635fd579af14610152578063715018a6146101655780637a1ac61e1461016d57600080fd5b80630457d6f2146100d9578063206a8300146100ee57806354fd4d5014610109575b600080fd5b6100ec6100e736600461085d565b6101db565b005b6100f66101ef565b6040519081526020015b60405180910390f35b6101456040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161010091906108e1565b6100ec61016036600461085d565b61021d565b6100ec61022e565b6100ec61017b366004610924565b610242565b60335460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610100565b6100f66103ad565b6100f66103e6565b6100ec6101c6366004610957565b610416565b6100f66104ca565b6100f6600081565b6101e36104f9565b6101ec8161057a565b50565b61021a60017f4aaefe95bd84fd3f32700cf3b7566bc944b73138e41958b5785826df2aecace1610972565b81565b6102256104f9565b6101ec81610632565b6102366104f9565b61024060006106ac565b565b600054600390610100900460ff16158015610264575060005460ff8083169116105b6102f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff83161761010017905561032e610723565b61033784610416565b6103408361057a565b61034982610632565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b60006103e16103dd60017fe314dfc40f0025322aacc0ba8ef420b62fb3b702cf01e0cdf3d829117ac2ff1b610972565b5490565b905090565b60006103e16103dd60017f4aaefe95bd84fd3f32700cf3b7566bc944b73138e41958b5785826df2aecace1610972565b61041e6104f9565b73ffffffffffffffffffffffffffffffffffffffff81166104c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102ec565b6101ec816106ac565b61021a60017fe314dfc40f0025322aacc0ba8ef420b62fb3b702cf01e0cdf3d829117ac2ff1b610972565b9055565b60335473ffffffffffffffffffffffffffffffffffffffff163314610240576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102ec565b6105ad6105a860017f4aaefe95bd84fd3f32700cf3b7566bc944b73138e41958b5785826df2aecace1610972565b829055565b6000816040516020016105c291815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060005b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161062691906108e1565b60405180910390a35050565b6106606105a860017fe314dfc40f0025322aacc0ba8ef420b62fb3b702cf01e0cdf3d829117ac2ff1b610972565b60008160405160200161067591815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060016105f5565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166107ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102ec565b610240600054610100900460ff16610854576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016102ec565b610240336106ac565b60006020828403121561086f57600080fd5b5035919050565b6000815180845260005b8181101561089c57602081850181015186830182015201610880565b818111156108ae576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006108f46020830184610876565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461091f57600080fd5b919050565b60008060006060848603121561093957600080fd5b610942846108fb565b95602085013595506040909401359392505050565b60006020828403121561096957600080fd5b6108f4826108fb565b6000828210156109ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50039056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(ProtocolVersionsStorageLayoutJSON), ProtocolVersionsStorageLayout); err != nil {
		panic(err)
	}

	layouts["ProtocolVersions"] = ProtocolVersionsStorageLayout
	deployedBytecodes["ProtocolVersions"] = ProtocolVersionsDeployedBin
}