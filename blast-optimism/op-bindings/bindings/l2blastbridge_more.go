// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2BlastBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/mainnet-bridge/L2BlastBridge.sol:L2BlastBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/mainnet-bridge/L2BlastBridge.sol:L2BlastBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/mainnet-bridge/L2BlastBridge.sol:L2BlastBridge\",\"label\":\"spacer_0_2_20\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/mainnet-bridge/L2BlastBridge.sol:L2BlastBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/mainnet-bridge/L2BlastBridge.sol:L2BlastBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"src/mainnet-bridge/L2BlastBridge.sol:L2BlastBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(CrossDomainMessenger)1007\"},{\"astId\":1006,\"contract\":\"src/mainnet-bridge/L2BlastBridge.sol:L2BlastBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_array(t_uint256)46_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)46_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1007\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2BlastBridgeStorageLayout = new(solc.StorageLayout)

var L2BlastBridgeDeployedBin = "0x6080604052600436106100e15760003560e01c80638129fc1c1161007f578063927ede2d11610059578063927ede2d146102c4578063a47a5c35146102e2578063c89701a2146102f5578063e11013dd1461032857600080fd5b80638129fc1c14610249578063870876231461025e5780638f601f661461027e57600080fd5b80633cb747bf116100bb5780633cb747bf1461017a578063540abf73146101b757806354fd4d50146101d75780637f46ddb21461021557600080fd5b80630166a07a1461013457806309fc8843146101545780631635f5fd1461016757600080fd5b3661012f57333b1561010e5760405162461bcd60e51b8152600401610105906118c6565b60405180910390fd5b61012d33333462030d406040518060200160405280600081525061033b565b005b600080fd5b34801561014057600080fd5b5061012d61014f366004611981565b610496565b61012d610162366004611a32565b6106b1565b61012d610175366004611a85565b610718565b34801561018657600080fd5b5060035461019a906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156101c357600080fd5b5061012d6101d2366004611af8565b6109ed565b3480156101e357600080fd5b50610208604051806040016040528060058152602001640312e302e360dc1b81525081565b6040516101ae9190611bc7565b34801561022157600080fd5b5061019a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561025557600080fd5b5061012d610a32565b34801561026a57600080fd5b5061012d610279366004611bda565b610bb1565b34801561028a57600080fd5b506102b6610299366004611c5d565b600260209081526000928352604080842090915290825290205481565b6040519081526020016101ae565b3480156102d057600080fd5b506003546001600160a01b031661019a565b61012d6102f0366004611a85565b610c15565b34801561030157600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061019a565b61012d610336366004611c96565b610ed4565b8234146103b05760405162461bcd60e51b815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c756500006064820152608401610105565b6103bc85858584610f1d565b6003546040516001600160a01b0390911690633dbb202b9085907f000000000000000000000000000000000000000000000000000000000000000090631635f5fd60e01b90610415908b908b9086908a90602401611cf9565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e086901b909216825261045d92918890600401611d36565b6000604051808303818588803b15801561047657600080fd5b505af115801561048a573d6000803e3d6000fd5b50505050505050505050565b6003546001600160a01b031633148015610545575060035460408051636e296e4560e01b815290516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610516573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053a9190611d70565b6001600160a01b0316145b6105615760405162461bcd60e51b815260040161010590611d8d565b61056a87610f70565b156105fc576105798787610fa0565b6105955760405162461bcd60e51b815260040161010590611df4565b6040516340c10f1960e01b81526001600160a01b038581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b1580156105df57600080fd5b505af11580156105f3573d6000803e3d6000fd5b50505050610664565b6001600160a01b038088166000908152600260209081526040808320938a168352929052205461062d908490611e7a565b6001600160a01b038089166000818152600260209081526040808320948c1683529390529190912091909155610664908585611073565b6106a8878787878787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506110d692505050565b50505050505050565b333b156106d05760405162461bcd60e51b8152600401610105906118c6565b6107133333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061033b92505050565b505050565b6003546001600160a01b0316331480156107c7575060035460408051636e296e4560e01b815290516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610798573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107bc9190611d70565b6001600160a01b0316145b6107e35760405162461bcd60e51b815260040161010590611d8d565b8234146108585760405162461bcd60e51b815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e742072657175697265640000000000006064820152608401610105565b306001600160a01b038516036108bc5760405162461bcd60e51b815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201526232b63360e91b6064820152608401610105565b6003546001600160a01b039081169085160361092b5760405162461bcd60e51b815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201526732b9b9b2b733b2b960c11b6064820152608401610105565b61096d85858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061113792505050565b600061098a855a866040518060200160405280600081525061117c565b9050806109e55760405162461bcd60e51b815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e73666572206661696044820152621b195960ea1b6064820152608401610105565b505050505050565b6106a887873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061119692505050565b600054610100900460ff1615808015610a525750600054600160ff909116105b80610a6c5750303b158015610a6c575060005460ff166001145b610acf5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610105565b6000805460ff191660011790558015610af2576000805461ff0019166101001790555b610b026007602160991b0161129b565b60405163099005e760e31b81526002604360981b0190634c802f3890610b3690309060019060009061dead90600401611ea7565b600060405180830381600087803b158015610b5057600080fd5b505af1158015610b64573d6000803e3d6000fd5b505050508015610bae576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b333b15610bd05760405162461bcd60e51b8152600401610105906118c6565b6109e586863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061119692505050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167311110000000000000000000000000000000011101933016001600160a01b031614610cd5576040805162461bcd60e51b81526020600482015260248101919091527f4c32426c6173744272696467653a2066756e6374696f6e2063616e206f6e6c7960448201527f2062652063616c6c65642066726f6d20746865206f74686572206272696467656064820152608401610105565b823414610d4a5760405162461bcd60e51b815260206004820152603960248201527f4c32426c6173744272696467653a20616d6f756e742073656e7420646f65732060448201527f6e6f74206d6174636820616d6f756e74207265717569726564000000000000006064820152608401610105565b306001600160a01b03851603610dad5760405162461bcd60e51b815260206004820152602260248201527f4c32426c6173744272696467653a2063616e6e6f742073656e6420746f207365604482015261363360f11b6064820152608401610105565b6003546001600160a01b0390811690851603610e1b5760405162461bcd60e51b815260206004820152602760248201527f4c32426c6173744272696467653a2063616e6e6f742073656e6420746f206d6560448201526639b9b2b733b2b960c91b6064820152608401610105565b610e5d85858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061113792505050565b6000610e7a855a866040518060200160405280600081525061117c565b9050806109e55760405162461bcd60e51b815260206004820152602260248201527f4c32426c6173744272696467653a20455448207472616e73666572206661696c604482015261195960f21b6064820152608401610105565b610f173385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061033b92505050565b50505050565b826001600160a01b0316846001600160a01b03167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051610f62929190611ef7565b60405180910390a350505050565b6000610f8382631d1d8b6360e01b611328565b80610f9a5750610f9a8263ec4fc8e360e01b611328565b92915050565b6000610fb383631d1d8b6360e01b611328565b1561103557826001600160a01b031663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ff6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101a9190611d70565b6001600160a01b0316826001600160a01b0316149050610f9a565b826001600160a01b031663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ff6573d6000803e3d6000fd5b6040516001600160a01b03831660248201526044810182905261071390849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261134b565b836001600160a01b0316856001600160a01b0316876001600160a01b03167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd86868660405161112793929190611f10565b60405180910390a4505050505050565b826001600160a01b0316846001600160a01b03167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051610f62929190611ef7565b600080600080845160208601878a8af19695505050505050565b6001600160a01b0387166003604360981b011461121b5760405162461bcd60e51b815260206004820152603b60248201527f4c32426c6173744272696467653a206f6e6c7920555344422063616e2062652060448201527f77697468647261776e2066726f6d2074686973206272696467652e00000000006064820152608401610105565b61122c6003604360981b0187610fa0565b61128c5760405162461bcd60e51b815260206004820152602b60248201527f4c32426c6173744272696467653a2077726f6e672072656d6f746520746f6b6560448201526a37103337b9102aa9a2211760a91b6064820152608401610105565b6106a88787878787878761141d565b600054610100900460ff166113065760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610105565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b600061133383611610565b801561134457506113448383611643565b9392505050565b60006113a0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116cc9092919063ffffffff16565b80519091501561071357808060200190518101906113be9190611f40565b6107135760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610105565b61142687610f70565b156114b8576114358787610fa0565b6114515760405162461bcd60e51b815260040161010590611df4565b604051632770a7eb60e21b81526001600160a01b03868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b15801561149b57600080fd5b505af11580156114af573d6000803e3d6000fd5b50505050611525565b6114cd6001600160a01b0388168630866116e3565b6001600160a01b038088166000908152600260209081526040808320938a16835292905220546114fe908490611f62565b6001600160a01b038089166000908152600260209081526040808320938b16835292905220555b61153387878787878661171b565b6003546040516001600160a01b0390911690633dbb202b907f00000000000000000000000000000000000000000000000000000000000000009062b3503d60e11b9061158d908b908d908c908c908c908b90602401611f7a565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199485161790525160e085901b90921682526115d592918790600401611d36565b600060405180830381600087803b1580156115ef57600080fd5b505af1158015611603573d6000803e3d6000fd5b5050505050505050505050565b6000611623826301ffc9a760e01b611643565b8015610f9a575061163c826001600160e01b0319611643565b1592915050565b604080516001600160e01b03198316602480830191909152825180830390910181526044909101909152602080820180516001600160e01b03166301ffc9a760e01b178152825160009392849283928392918391908a617530fa92503d915060005190508280156116b5575060208210155b80156116c15750600081115b979650505050505050565b60606116db848460008561176c565b949350505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610f179085906323b872dd60e01b9060840161109f565b836001600160a01b0316856001600160a01b0316876001600160a01b03167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf86868660405161112793929190611f10565b6060824710156117cd5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610105565b6001600160a01b0385163b6118245760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610105565b600080866001600160a01b031685876040516118409190611fc9565b60006040518083038185875af1925050503d806000811461187d576040519150601f19603f3d011682016040523d82523d6000602084013e611882565b606091505b50915091506116c18282866060831561189c575081611344565b8251156118ac5782518084602001fd5b8160405162461bcd60e51b81526004016101059190611bc7565b60208082526037908201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60408201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606082015260800190565b6001600160a01b0381168114610bae57600080fd5b60008083601f84011261194a57600080fd5b50813567ffffffffffffffff81111561196257600080fd5b60208301915083602082850101111561197a57600080fd5b9250929050565b600080600080600080600060c0888a03121561199c57600080fd5b87356119a781611923565b965060208801356119b781611923565b955060408801356119c781611923565b945060608801356119d781611923565b93506080880135925060a088013567ffffffffffffffff8111156119fa57600080fd5b611a068a828b01611938565b989b979a50959850939692959293505050565b803563ffffffff81168114611a2d57600080fd5b919050565b600080600060408486031215611a4757600080fd5b611a5084611a19565b9250602084013567ffffffffffffffff811115611a6c57600080fd5b611a7886828701611938565b9497909650939450505050565b600080600080600060808688031215611a9d57600080fd5b8535611aa881611923565b94506020860135611ab881611923565b935060408601359250606086013567ffffffffffffffff811115611adb57600080fd5b611ae788828901611938565b969995985093965092949392505050565b600080600080600080600060c0888a031215611b1357600080fd5b8735611b1e81611923565b96506020880135611b2e81611923565b95506040880135611b3e81611923565b945060608801359350611b5360808901611a19565b925060a088013567ffffffffffffffff8111156119fa57600080fd5b60005b83811015611b8a578181015183820152602001611b72565b83811115610f175750506000910152565b60008151808452611bb3816020860160208601611b6f565b601f01601f19169290920160200192915050565b6020815260006113446020830184611b9b565b60008060008060008060a08789031215611bf357600080fd5b8635611bfe81611923565b95506020870135611c0e81611923565b945060408701359350611c2360608801611a19565b9250608087013567ffffffffffffffff811115611c3f57600080fd5b611c4b89828a01611938565b979a9699509497509295939492505050565b60008060408385031215611c7057600080fd5b8235611c7b81611923565b91506020830135611c8b81611923565b809150509250929050565b60008060008060608587031215611cac57600080fd5b8435611cb781611923565b9350611cc560208601611a19565b9250604085013567ffffffffffffffff811115611ce157600080fd5b611ced87828801611938565b95989497509550505050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611d2c90830184611b9b565b9695505050505050565b6001600160a01b0384168152606060208201819052600090611d5a90830185611b9b565b905063ffffffff83166040830152949350505050565b600060208284031215611d8257600080fd5b815161134481611923565b60208082526041908201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60408201527f792062652063616c6c65642066726f6d20746865206f746865722062726964676060820152606560f81b608082015260a00190565b6020808252604a908201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60408201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60608201526937b1b0b6103a37b5b2b760b11b608082015260a00190565b634e487b7160e01b600052601160045260246000fd5b600082821015611e8c57611e8c611e64565b500390565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b038581168252608082019060038610611ec957611ec9611e91565b85602084015260028510611edf57611edf611e91565b84604084015280841660608401525095945050505050565b8281526040602082015260006116db6040830184611b9b565b60018060a01b0384168152826020820152606060408201526000611f376060830184611b9b565b95945050505050565b600060208284031215611f5257600080fd5b8151801515811461134457600080fd5b60008219821115611f7557611f75611e64565b500190565b6001600160a01b03878116825286811660208301528581166040830152841660608201526080810183905260c060a08201819052600090611fbd90830184611b9b565b98975050505050505050565b60008251611fdb818460208701611b6f565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2BlastBridgeStorageLayoutJSON), L2BlastBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2BlastBridge"] = L2BlastBridgeStorageLayout
	deployedBytecodes["L2BlastBridge"] = L2BlastBridgeDeployedBin
}