// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const BlastStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/Blast.sol:Blast\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L2/Blast.sol:Blast\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L2/Blast.sol:Blast\",\"label\":\"governorMap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_address)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var BlastStorageLayout = new(solc.StorageLayout)

var BlastDeployedBin = "0x608060405234801561001057600080fd5b50600436106101da5760003560e01c8063a0278c9011610104578063eb59acdc116100a2578063f971966211610071578063f971966214610433578063fafce39e14610446578063fd8c4b9d14610459578063fe9fbb801461047e57600080fd5b8063eb59acdc146103f2578063eb86469814610405578063ec3278e814610418578063f098767a1461042b57600080fd5b8063b71d6dd4116100de578063b71d6dd41461036b578063c8992e611461037e578063dde798a414610391578063e43581b8146103b457600080fd5b8063a0278c9014610315578063a70d11bd1461033c578063aa857d981461036357600080fd5b806354fd4d501161017c57806385ebdc271161014b57806385ebdc271461029b578063860043b6146102dc578063908c8502146102ef578063954fa5ee1461030257600080fd5b806354fd4d5014610263578063662aa11d146102785780637114177a1461028b5780638129fc1c1461029357600080fd5b806337ebe3a8116101b857806337ebe3a8146102225780633ba5713e146102355780634c802f38146102485780634e606c471461025b57600080fd5b80630951888f146101df5780630ca12c4b146102055780632210dfb11461021a575b600080fd5b6101f26101ed366004611853565b610491565b6040519081526020015b60405180910390f35b61021861021336600461188f565b6105a6565b005b6102186105fb565b6102186102303660046118c2565b6106a3565b6102186102433660046118c2565b610771565b6102186102563660046118f9565b6107cd565b610218610952565b61026b6109c6565b6040516101fc919061197b565b6101f261028636600461188f565b610a69565b610218610b55565b610218610c22565b6102c46102a93660046118c2565b6001602052600090815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016101fc565b6101f26102ea36600461188f565b610d2a565b6102186102fd3660046118c2565b610eb8565b6101f261031036600461188f565b610f5a565b6102c47f000000000000000000000000000000000000000000000000000000000000000081565b6102c47f000000000000000000000000000000000000000000000000000000000000000081565b610218611007565b6102186103793660046118c2565b611063565b61021861038c3660046119ae565b6110bf565b6103a461039f3660046118c2565b61123f565b6040516101fc9493929190611a2b565b6103e26103c23660046118c2565b6001600160a01b0390811660009081526001602052604090205416331490565b60405190151581526020016101fc565b6102186104003660046118c2565b6112e1565b6102186104133660046118c2565b611355565b6101f26104263660046118c2565b6113a9565b61021861143e565b6101f2610441366004611853565b61149a565b6101f2610454366004611a4d565b61154e565b61046c6104673660046118c2565b611650565b60405160ff90911681526020016101fc565b6103e261048c3660046118c2565b6116df565b600061049c846116df565b6105005760405162461bcd60e51b815260206004820152602a60248201527f4e6f7420616c6c6f77656420746f20636c61696d20676173206174206d696e20604482015269636c61696d207261746560b01b60648201526084015b60405180910390fd5b604051630951888f60e01b81526001600160a01b0385811660048301528481166024830152604482018490527f00000000000000000000000000000000000000000000000000000000000000001690630951888f906064015b6020604051808303816000875af1158015610578573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059c9190611a8f565b90505b9392505050565b6105af816116df565b6105cb5760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b0390811660009081526001602052604090208054919092166001600160a01b0319909116179055565b610604336116df565b6106205760405162461bcd60e51b81526004016104f790611aa8565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba59061066f903390600090600401611aec565b600060405180830381600087803b15801561068957600080fd5b505af115801561069d573d6000803e3d6000fd5b50505050565b6106ac816116df565b6106c85760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260025b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af1158015610749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076d9190611a8f565b5050565b61077a816116df565b6107965760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260006106fb565b6107d6846116df565b6107f25760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b038481166000908152600160205260409081902080546001600160a01b0319168484161790555163d4810ba560e01b81527f00000000000000000000000000000000000000000000000000000000000000009091169063d4810ba5906108659087908690600401611aec565b600060405180830381600087803b15801561087f57600080fd5b505af1158015610893573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633bdbe9a5858560028111156108d9576108d96119f3565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af1158015610927573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094b9190611a8f565b5050505050565b61095b336116df565b6109775760405162461bcd60e51b81526004016104f790611aa8565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba59061066f903390600190600401611aec565b60606109f17f0000000000000000000000000000000000000000000000000000000000000000611736565b610a1a7f0000000000000000000000000000000000000000000000000000000000000000611736565b610a437f0000000000000000000000000000000000000000000000000000000000000000611736565b604051602001610a5593929190611b09565b604051602081830303815290604052905090565b6000610a74836116df565b610ac05760405162461bcd60e51b815260206004820152601c60248201527f4e6f7420616c6c6f77656420746f20636c61696d206d6178206761730000000060448201526064016104f7565b604051630928b10d60e31b81526001600160a01b03848116600483015283811660248301527f000000000000000000000000000000000000000000000000000000000000000016906349458868906044015b6020604051808303816000875af1158015610b31573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059f9190611a8f565b610b5e336116df565b610b7a5760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a53360005b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af1158015610bfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1f9190611a8f565b50565b600054610100900460ff1615808015610c425750600054600160ff909116105b80610c5c5750303b158015610c5c575060005460ff166001145b610cbf5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104f7565b6000805460ff191660011790558015610ce2576000805461ff0019166101001790555b8015610c1f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6000610d35836116df565b610d815760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420617574686f72697a656420746f20636c61696d207969656c6400000060448201526064016104f7565b60405163e12f3a6160e01b81526001600160a01b0384811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e12f3a6190602401602060405180830381865afa158015610dec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e109190611a8f565b60405163132d974d60e31b81526001600160a01b0386811660048301528581166024830152604482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063996cba68906064016020604051808303816000875af1158015610e8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb09190611a8f565b949350505050565b610ec1816116df565b610edd5760405162461bcd60e51b81526004016104f790611aa8565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba590610f2c908490600190600401611aec565b600060405180830381600087803b158015610f4657600080fd5b505af115801561094b573d6000803e3d6000fd5b6000610f65836116df565b610fb15760405162461bcd60e51b815260206004820152601c60248201527f4e6f7420616c6c6f77656420746f20636c61696d20616c6c206761730000000060448201526064016104f7565b604051635767bba560e01b81526001600160a01b03848116600483015283811660248301527f00000000000000000000000000000000000000000000000000000000000000001690635767bba590604401610b12565b611010336116df565b61102c5760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a5336001610bad565b61106c816116df565b6110885760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a58260016106fb565b6110c8336116df565b6110e45760405162461bcd60e51b81526004016104f790611aa8565b336000818152600160205260409081902080546001600160a01b0319166001600160a01b0385811691909117909155905163d4810ba560e01b81527f00000000000000000000000000000000000000000000000000000000000000009091169163d4810ba59161115991908690600401611aec565b600060405180830381600087803b15801561117357600080fd5b505af1158015611187573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633bdbe9a5338560028111156111cd576111cd6119f3565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260ff1660248201526044016020604051808303816000875af115801561121b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069d9190611a8f565b604051633779e62960e21b81526001600160a01b0382811660048301526000918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063dde798a490602401608060405180830381865afa1580156112ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112d29190611b63565b93509350935093509193509193565b6112ea816116df565b6113065760405162461bcd60e51b81526004016104f790611aa8565b60405163d4810ba560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d4810ba590610f2c908490600090600401611aec565b61135e336116df565b61137a5760405162461bcd60e51b81526004016104f790611aa8565b33600090815260016020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b60405163e12f3a6160e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063e12f3a6190602401602060405180830381865afa158015611414573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114389190611a8f565b92915050565b611447336116df565b6114635760405162461bcd60e51b81526004016104f790611aa8565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016633bdbe9a5336002610bad565b60006114a5846116df565b6114f15760405162461bcd60e51b815260206004820152601d60248201527f4e6f7420617574686f72697a656420746f20636c61696d207969656c6400000060448201526064016104f7565b60405163132d974d60e31b81526001600160a01b0385811660048301528481166024830152604482018490527f0000000000000000000000000000000000000000000000000000000000000000169063996cba6890606401610559565b6000611559856116df565b6115a55760405162461bcd60e51b815260206004820152601860248201527f4e6f7420616c6c6f77656420746f20636c61696d20676173000000000000000060448201526064016104f7565b604051631357a41960e11b81526001600160a01b038681166004830152858116602483015260448201859052606482018490527f000000000000000000000000000000000000000000000000000000000000000016906326af4832906084016020604051808303816000875af1158015611623573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116479190611a8f565b95945050505050565b60405163c44b11f760e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063c44b11f790602401602060405180830381865afa1580156116bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114389190611ba4565b6001600160a01b0380821660009081526001602052604081205490911633148061143857506001600160a01b03808316600090815260016020526040902054161580156114385750506001600160a01b0316331490565b60608160000361175d5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115611787578061177181611bdd565b91506117809050600a83611c0c565b9150611761565b60008167ffffffffffffffff8111156117a2576117a2611c20565b6040519080825280601f01601f1916602001820160405280156117cc576020820181803683370190505b5090505b8415610eb0576117e1600183611c36565b91506117ee600a86611c4d565b6117f9906030611c61565b60f81b81838151811061180e5761180e611c79565b60200101906001600160f81b031916908160001a905350611830600a86611c0c565b94506117d0565b80356001600160a01b038116811461184e57600080fd5b919050565b60008060006060848603121561186857600080fd5b61187184611837565b925061187f60208501611837565b9150604084013590509250925092565b600080604083850312156118a257600080fd5b6118ab83611837565b91506118b960208401611837565b90509250929050565b6000602082840312156118d457600080fd5b61059f82611837565b80356003811061184e57600080fd5b60028110610c1f57600080fd5b6000806000806080858703121561190f57600080fd5b61191885611837565b9350611926602086016118dd565b92506040850135611936816118ec565b915061194460608601611837565b905092959194509250565b60005b8381101561196a578181015183820152602001611952565b8381111561069d5750506000910152565b602081526000825180602084015261199a81604085016020870161194f565b601f01601f19169190910160400192915050565b6000806000606084860312156119c357600080fd5b6119cc846118dd565b925060208401356119dc816118ec565b91506119ea60408501611837565b90509250925092565b634e487b7160e01b600052602160045260246000fd5b60028110611a2757634e487b7160e01b600052602160045260246000fd5b9052565b8481526020810184905260408101839052608081016116476060830184611a09565b60008060008060808587031215611a6357600080fd5b611a6c85611837565b9350611a7a60208601611837565b93969395505050506040820135916060013590565b600060208284031215611aa157600080fd5b5051919050565b60208082526024908201527f6e6f7420617574686f72697a656420746f20636f6e66696775726520636f6e746040820152631c9858dd60e21b606082015260800190565b6001600160a01b03831681526040810161059f6020830184611a09565b60008451611b1b81846020890161194f565b8083019050601760f91b8082528551611b3b816001850160208a0161194f565b60019201918201528351611b5681600284016020880161194f565b0160020195945050505050565b60008060008060808587031215611b7957600080fd5b8451935060208501519250604085015191506060850151611b99816118ec565b939692955090935050565b600060208284031215611bb657600080fd5b815160ff8116811461059f57600080fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611bef57611bef611bc7565b5060010190565b634e487b7160e01b600052601260045260246000fd5b600082611c1b57611c1b611bf6565b500490565b634e487b7160e01b600052604160045260246000fd5b600082821015611c4857611c48611bc7565b500390565b600082611c5c57611c5c611bf6565b500690565b60008219821115611c7457611c74611bc7565b500190565b634e487b7160e01b600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(BlastStorageLayoutJSON), BlastStorageLayout); err != nil {
		panic(err)
	}

	layouts["Blast"] = BlastStorageLayout
	deployedBytecodes["Blast"] = BlastDeployedBin
}