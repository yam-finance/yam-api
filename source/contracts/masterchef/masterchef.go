// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package masterchef

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MasterchefABI is the input ABI used to generate the binding from.
const MasterchefABI = "[{\"inputs\":[{\"internalType\":\"contractSushiToken\",\"name\":\"_sushi\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sushiPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devaddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingSushi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accSushiPerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushi\",\"outputs\":[{\"internalType\":\"contractSushiToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushiPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MasterchefBin is the compiled bytecode used for deploying new contracts.
var MasterchefBin = "0x6080604052600060085534801561001557600080fd5b50604051611d61380380611d61833981810160405260a081101561003857600080fd5b5080516020820151604083015160608401516080909401519293919290919060006100616100e9565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b039687166001600160a01b03199182161790915560028054959096169416939093179093556004556003556009556100ed565b3390565b611c65806100fc6000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c8063630b5ba1116100de5780638da5cb5b11610097578063b0bcf42a11610071578063b0bcf42a14610419578063d49e77cd14610421578063e2bbb15814610429578063f2fde38b1461044c5761018e565b80638da5cb5b146103a95780638dbb1e3a146103b157806393f1a40b146103d45761018e565b8063630b5ba11461033857806364482f7914610340578063715018a61461036b5780637cd07e47146103735780638aa285501461037b5780638d88a90e146103835761018e565b80631eaaa0451161014b578063454b060811610125578063454b0608146102d957806348cd4cb1146102f657806351eb05a6146102fe5780635312ea8e1461031b5761018e565b80631eaaa0451461025a57806323cf311814610290578063441a3e70146102b65761018e565b8063081e3eda146101935780630a087903146101ad5780631526fe27146101d157806317caf6f11461021e578063195426ec146102265780631aed655314610252575b600080fd5b61019b610472565b60408051918252519081900360200190f35b6101b5610478565b604080516001600160a01b039092168252519081900360200190f35b6101ee600480360360208110156101e757600080fd5b5035610487565b604080516001600160a01b0390951685526020850193909352838301919091526060830152519081900360800190f35b61019b6104c8565b61019b6004803603604081101561023c57600080fd5b50803590602001356001600160a01b03166104ce565b61019b610644565b61028e6004803603606081101561027057600080fd5b508035906001600160a01b036020820135169060400135151561064a565b005b61028e600480360360208110156102a657600080fd5b50356001600160a01b03166107c5565b61028e600480360360408110156102cc57600080fd5b508035906020013561083f565b61028e600480360360208110156102ef57600080fd5b5035610986565b61019b610be2565b61028e6004803603602081101561031457600080fd5b5035610be8565b61028e6004803603602081101561033157600080fd5b5035610e09565b61028e610ea4565b61028e6004803603606081101561035657600080fd5b50803590602081013590604001351515610ec7565b61028e610f98565b6101b561103a565b61019b611049565b61028e6004803603602081101561039957600080fd5b50356001600160a01b031661104e565b6101b56110bb565b61019b600480360360408110156103c757600080fd5b50803590602001356110ca565b610400600480360360408110156103ea57600080fd5b50803590602001356001600160a01b0316611130565b6040805192835260208301919091528051918290030190f35b61019b611154565b6101b561115a565b61028e6004803603604081101561043f57600080fd5b5080359060200135611169565b61028e6004803603602081101561046257600080fd5b50356001600160a01b031661126e565b60065490565b6001546001600160a01b031681565b6006818154811061049457fe5b600091825260209091206004909102018054600182015460028301546003909301546001600160a01b039092169350919084565b60085481565b600080600684815481106104de57fe5b600091825260208083208784526007825260408085206001600160a01b03898116875290845281862060049586029093016003810154815484516370a0823160e01b81523098810198909852935191985093969395939492909116926370a08231926024808301939192829003018186803b15801561055c57600080fd5b505afa158015610570573d6000803e3d6000fd5b505050506040513d602081101561058657600080fd5b505160028501549091504311801561059d57508015155b156106095760006105b28560020154436110ca565b905060006105e56008546105df88600101546105d96004548761136690919063ffffffff16565b90611366565b906113c6565b90506106046105fd846105df8464e8d4a51000611366565b8590611408565b935050505b610637836001015461063164e8d4a510006105df86886000015461136690919063ffffffff16565b90611462565b9450505050505b92915050565b60035481565b6106526114a4565b6000546001600160a01b039081169116146106a2576040805162461bcd60e51b81526020600482018190526024820152600080516020611bb0833981519152604482015290519081900360640190fd5b80156106b0576106b0610ea4565b600060095443116106c3576009546106c5565b435b6008549091506106d59085611408565b600855604080516080810182526001600160a01b0394851681526020810195865290810191825260006060820181815260068054600181018255925291517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f600490920291820180546001600160a01b031916919096161790945593517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40840155517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d418301555090517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4290910155565b6107cd6114a4565b6000546001600160a01b0390811691161461081d576040805162461bcd60e51b81526020600482018190526024820152600080516020611bb0833981519152604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b0392909216919091179055565b60006006838154811061084e57fe5b6000918252602080832086845260078252604080852033865290925292208054600490920290920192508311156108c1576040805162461bcd60e51b81526020600482015260126024820152711dda5d1a191c985dce881b9bdd0819dbdbd960721b604482015290519081900360640190fd5b6108ca84610be8565b60006108f8826001015461063164e8d4a510006105df8760030154876000015461136690919063ffffffff16565b905061090433826114a8565b81546109109085611462565b808355600384015461092d9164e8d4a51000916105df9190611366565b60018301558254610948906001600160a01b03163386611639565b604080518581529051869133917ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b5689181900360200190a35050505050565b6005546001600160a01b03166109da576040805162461bcd60e51b815260206004820152601460248201527336b4b3b930ba329d1037379036b4b3b930ba37b960611b604482015290519081900360640190fd5b6000600682815481106109e957fe5b600091825260208083206004928302018054604080516370a0823160e01b81523095810195909552519195506001600160a01b0316939284926370a0823192602480840193829003018186803b158015610a4257600080fd5b505afa158015610a56573d6000803e3d6000fd5b505050506040513d6020811015610a6c57600080fd5b5051600554909150610a8b906001600160a01b0384811691168361168b565b6005546040805163ce5494bb60e01b81526001600160a01b0385811660048301529151600093929092169163ce5494bb9160248082019260209290919082900301818787803b158015610add57600080fd5b505af1158015610af1573d6000803e3d6000fd5b505050506040513d6020811015610b0757600080fd5b5051604080516370a0823160e01b815230600482015290519192506001600160a01b038316916370a0823191602480820192602092909190829003018186803b158015610b5357600080fd5b505afa158015610b67573d6000803e3d6000fd5b505050506040513d6020811015610b7d57600080fd5b50518214610bc1576040805162461bcd60e51b815260206004820152600c60248201526b1b5a59dc985d194e8818985960a21b604482015290519081900360640190fd5b83546001600160a01b0319166001600160a01b039190911617909255505050565b60095481565b600060068281548110610bf757fe5b9060005260206000209060040201905080600201544311610c185750610e06565b8054604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b158015610c6257600080fd5b505afa158015610c76573d6000803e3d6000fd5b505050506040513d6020811015610c8c57600080fd5b5051905080610ca2575043600290910155610e06565b6000610cb28360020154436110ca565b90506000610cd96008546105df86600101546105d96004548761136690919063ffffffff16565b6001546002549192506001600160a01b03908116916340c10f199116610d0084600a6113c6565b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015610d4657600080fd5b505af1158015610d5a573d6000803e3d6000fd5b5050600154604080516340c10f1960e01b81523060048201526024810186905290516001600160a01b0390921693506340c10f19925060448082019260009290919082900301818387803b158015610db157600080fd5b505af1158015610dc5573d6000803e3d6000fd5b50505050610df3610de8846105df64e8d4a510008561136690919063ffffffff16565b600386015490611408565b6003850155505043600290920191909155505b50565b600060068281548110610e1857fe5b60009182526020808320858452600782526040808520338087529352909320805460049093029093018054909450610e5d926001600160a01b03919091169190611639565b80546040805191825251849133917fbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae05959181900360200190a360008082556001909101555050565b60065460005b81811015610ec357610ebb81610be8565b600101610eaa565b5050565b610ecf6114a4565b6000546001600160a01b03908116911614610f1f576040805162461bcd60e51b81526020600482018190526024820152600080516020611bb0833981519152604482015290519081900360640190fd5b8015610f2d57610f2d610ea4565b610f6a82610f6460068681548110610f4157fe5b90600052602060002090600402016001015460085461146290919063ffffffff16565b90611408565b6008819055508160068481548110610f7e57fe5b906000526020600020906004020160010181905550505050565b610fa06114a4565b6000546001600160a01b03908116911614610ff0576040805162461bcd60e51b81526020600482018190526024820152600080516020611bb0833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6005546001600160a01b031681565b600a81565b6002546001600160a01b03163314611099576040805162461bcd60e51b81526020600482015260096024820152686465763a207775743f60b81b604482015290519081900360640190fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031690565b600060035482116110eb576110e4600a6105d98486611462565b905061063e565b60035483106110fe576110e48284611462565b6110e46111166003548461146290919063ffffffff16565b610f64600a6105d98760035461146290919063ffffffff16565b60076020908152600092835260408084209091529082529020805460019091015482565b60045481565b6002546001600160a01b031681565b60006006838154811061117857fe5b600091825260208083208684526007825260408085203386529092529220600490910290910191506111a984610be8565b8054156111ec5760006111de826001015461063164e8d4a510006105df8760030154876000015461136690919063ffffffff16565b90506111ea33826114a8565b505b8154611203906001600160a01b031633308661179e565b805461120f9084611408565b808255600383015461122c9164e8d4a51000916105df9190611366565b6001820155604080518481529051859133917f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a159181900360200190a350505050565b6112766114a4565b6000546001600160a01b039081169116146112c6576040805162461bcd60e51b81526020600482018190526024820152600080516020611bb0833981519152604482015290519081900360640190fd5b6001600160a01b03811661130b5760405162461bcd60e51b8152600401808060200182810382526026815260200180611b696026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000826113755750600061063e565b8282028284828161138257fe5b04146113bf5760405162461bcd60e51b8152600401808060200182810382526021815260200180611b8f6021913960400191505060405180910390fd5b9392505050565b60006113bf83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506117fe565b6000828201838110156113bf576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60006113bf83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506118a0565b3390565b600154604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b1580156114f357600080fd5b505afa158015611507573d6000803e3d6000fd5b505050506040513d602081101561151d57600080fd5b50519050808211156115b1576001546040805163a9059cbb60e01b81526001600160a01b038681166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561157f57600080fd5b505af1158015611593573d6000803e3d6000fd5b505050506040513d60208110156115a957600080fd5b506116349050565b6001546040805163a9059cbb60e01b81526001600160a01b038681166004830152602482018690529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561160757600080fd5b505af115801561161b573d6000803e3d6000fd5b505050506040513d602081101561163157600080fd5b50505b505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526116349084906118fa565b801580611711575060408051636eb1769f60e11b81523060048201526001600160a01b03848116602483015291519185169163dd62ed3e91604480820192602092909190829003018186803b1580156116e357600080fd5b505afa1580156116f7573d6000803e3d6000fd5b505050506040513d602081101561170d57600080fd5b5051155b61174c5760405162461bcd60e51b8152600401808060200182810382526036815260200180611bfa6036913960400191505060405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b1790526116349084906118fa565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b1790526117f89085906118fa565b50505050565b6000818361188a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561184f578181015183820152602001611837565b50505050905090810190601f16801561187c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161189657fe5b0495945050505050565b600081848411156118f25760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561184f578181015183820152602001611837565b505050900390565b606061194f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166119ab9092919063ffffffff16565b8051909150156116345780806020019051602081101561196e57600080fd5b50516116345760405162461bcd60e51b815260040180806020018281038252602a815260200180611bd0602a913960400191505060405180910390fd5b60606119ba84846000856119c2565b949350505050565b60606119cd85611b2f565b611a1e576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b60208310611a5d5780518252601f199092019160209182019101611a3e565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611abf576040519150601f19603f3d011682016040523d82523d6000602084013e611ac4565b606091505b50915091508115611ad85791506119ba9050565b805115611ae85780518082602001fd5b60405162461bcd60e51b815260206004820181815286516024840152865187939192839260440191908501908083836000831561184f578181015183820152602001611837565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906119ba57505015159291505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565645361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a2646970667358221220dc7149886340d800d71534b66d8dba11ba031f11ea7b498b4fe882e6f441f06c64736f6c634300060c00330000000000000000000000006b3595068778dd592e39a122f4f5a5cf09c90fe2000000000000000000000000f942dba4159cb61f8ad88ca4a83f5204e8f4a6bd0000000000000000000000000000000000000000000000056bc75e2d631000000000000000000000000000000000000000000000000000000000000000a408300000000000000000000000000000000000000000000000000000000000a58ed0"

// DeployMasterchef deploys a new Ethereum contract, binding an instance of Masterchef to it.
func DeployMasterchef(auth *bind.TransactOpts, backend bind.ContractBackend, _sushi common.Address, _devaddr common.Address, _sushiPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int) (common.Address, *types.Transaction, *Masterchef, error) {
	parsed, err := abi.JSON(strings.NewReader(MasterchefABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MasterchefBin), backend, _sushi, _devaddr, _sushiPerBlock, _startBlock, _bonusEndBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Masterchef{MasterchefCaller: MasterchefCaller{contract: contract}, MasterchefTransactor: MasterchefTransactor{contract: contract}, MasterchefFilterer: MasterchefFilterer{contract: contract}}, nil
}

// Masterchef is an auto generated Go binding around an Ethereum contract.
type Masterchef struct {
	MasterchefCaller     // Read-only binding to the contract
	MasterchefTransactor // Write-only binding to the contract
	MasterchefFilterer   // Log filterer for contract events
}

// MasterchefCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterchefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterchefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterchefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterchefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterchefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterchefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterchefSession struct {
	Contract     *Masterchef       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterchefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterchefCallerSession struct {
	Contract *MasterchefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MasterchefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterchefTransactorSession struct {
	Contract     *MasterchefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MasterchefRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterchefRaw struct {
	Contract *Masterchef // Generic contract binding to access the raw methods on
}

// MasterchefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterchefCallerRaw struct {
	Contract *MasterchefCaller // Generic read-only contract binding to access the raw methods on
}

// MasterchefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterchefTransactorRaw struct {
	Contract *MasterchefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterchef creates a new instance of Masterchef, bound to a specific deployed contract.
func NewMasterchef(address common.Address, backend bind.ContractBackend) (*Masterchef, error) {
	contract, err := bindMasterchef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Masterchef{MasterchefCaller: MasterchefCaller{contract: contract}, MasterchefTransactor: MasterchefTransactor{contract: contract}, MasterchefFilterer: MasterchefFilterer{contract: contract}}, nil
}

// NewMasterchefCaller creates a new read-only instance of Masterchef, bound to a specific deployed contract.
func NewMasterchefCaller(address common.Address, caller bind.ContractCaller) (*MasterchefCaller, error) {
	contract, err := bindMasterchef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterchefCaller{contract: contract}, nil
}

// NewMasterchefTransactor creates a new write-only instance of Masterchef, bound to a specific deployed contract.
func NewMasterchefTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterchefTransactor, error) {
	contract, err := bindMasterchef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterchefTransactor{contract: contract}, nil
}

// NewMasterchefFilterer creates a new log filterer instance of Masterchef, bound to a specific deployed contract.
func NewMasterchefFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterchefFilterer, error) {
	contract, err := bindMasterchef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterchefFilterer{contract: contract}, nil
}

// bindMasterchef binds a generic wrapper to an already deployed contract.
func bindMasterchef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasterchefABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Masterchef *MasterchefRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Masterchef.Contract.MasterchefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Masterchef *MasterchefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchef.Contract.MasterchefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Masterchef *MasterchefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Masterchef.Contract.MasterchefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Masterchef *MasterchefCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Masterchef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Masterchef *MasterchefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Masterchef *MasterchefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Masterchef.Contract.contract.Transact(opts, method, params...)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Masterchef *MasterchefCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "BONUS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Masterchef *MasterchefSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Masterchef.Contract.BONUSMULTIPLIER(&_Masterchef.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_Masterchef *MasterchefCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _Masterchef.Contract.BONUSMULTIPLIER(&_Masterchef.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Masterchef *MasterchefCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Masterchef *MasterchefSession) BonusEndBlock() (*big.Int, error) {
	return _Masterchef.Contract.BonusEndBlock(&_Masterchef.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_Masterchef *MasterchefCallerSession) BonusEndBlock() (*big.Int, error) {
	return _Masterchef.Contract.BonusEndBlock(&_Masterchef.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Masterchef *MasterchefCaller) Devaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "devaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Masterchef *MasterchefSession) Devaddr() (common.Address, error) {
	return _Masterchef.Contract.Devaddr(&_Masterchef.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Masterchef *MasterchefCallerSession) Devaddr() (common.Address, error) {
	return _Masterchef.Contract.Devaddr(&_Masterchef.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Masterchef *MasterchefCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Masterchef *MasterchefSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Masterchef.Contract.GetMultiplier(&_Masterchef.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Masterchef *MasterchefCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Masterchef.Contract.GetMultiplier(&_Masterchef.CallOpts, _from, _to)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Masterchef *MasterchefCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Masterchef *MasterchefSession) Migrator() (common.Address, error) {
	return _Masterchef.Contract.Migrator(&_Masterchef.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Masterchef *MasterchefCallerSession) Migrator() (common.Address, error) {
	return _Masterchef.Contract.Migrator(&_Masterchef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchef *MasterchefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchef *MasterchefSession) Owner() (common.Address, error) {
	return _Masterchef.Contract.Owner(&_Masterchef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchef *MasterchefCallerSession) Owner() (common.Address, error) {
	return _Masterchef.Contract.Owner(&_Masterchef.CallOpts)
}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256)
func (_Masterchef *MasterchefCaller) PendingSushi(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "pendingSushi", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256)
func (_Masterchef *MasterchefSession) PendingSushi(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Masterchef.Contract.PendingSushi(&_Masterchef.CallOpts, _pid, _user)
}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256)
func (_Masterchef *MasterchefCallerSession) PendingSushi(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Masterchef.Contract.PendingSushi(&_Masterchef.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accSushiPerShare)
func (_Masterchef *MasterchefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken          common.Address
	AllocPoint       *big.Int
	LastRewardBlock  *big.Int
	AccSushiPerShare *big.Int
}, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken          common.Address
		AllocPoint       *big.Int
		LastRewardBlock  *big.Int
		AccSushiPerShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccSushiPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accSushiPerShare)
func (_Masterchef *MasterchefSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken          common.Address
	AllocPoint       *big.Int
	LastRewardBlock  *big.Int
	AccSushiPerShare *big.Int
}, error) {
	return _Masterchef.Contract.PoolInfo(&_Masterchef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accSushiPerShare)
func (_Masterchef *MasterchefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken          common.Address
	AllocPoint       *big.Int
	LastRewardBlock  *big.Int
	AccSushiPerShare *big.Int
}, error) {
	return _Masterchef.Contract.PoolInfo(&_Masterchef.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Masterchef *MasterchefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Masterchef *MasterchefSession) PoolLength() (*big.Int, error) {
	return _Masterchef.Contract.PoolLength(&_Masterchef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Masterchef *MasterchefCallerSession) PoolLength() (*big.Int, error) {
	return _Masterchef.Contract.PoolLength(&_Masterchef.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Masterchef *MasterchefCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Masterchef *MasterchefSession) StartBlock() (*big.Int, error) {
	return _Masterchef.Contract.StartBlock(&_Masterchef.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Masterchef *MasterchefCallerSession) StartBlock() (*big.Int, error) {
	return _Masterchef.Contract.StartBlock(&_Masterchef.CallOpts)
}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_Masterchef *MasterchefCaller) Sushi(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "sushi")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_Masterchef *MasterchefSession) Sushi() (common.Address, error) {
	return _Masterchef.Contract.Sushi(&_Masterchef.CallOpts)
}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_Masterchef *MasterchefCallerSession) Sushi() (common.Address, error) {
	return _Masterchef.Contract.Sushi(&_Masterchef.CallOpts)
}

// SushiPerBlock is a free data retrieval call binding the contract method 0xb0bcf42a.
//
// Solidity: function sushiPerBlock() view returns(uint256)
func (_Masterchef *MasterchefCaller) SushiPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "sushiPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SushiPerBlock is a free data retrieval call binding the contract method 0xb0bcf42a.
//
// Solidity: function sushiPerBlock() view returns(uint256)
func (_Masterchef *MasterchefSession) SushiPerBlock() (*big.Int, error) {
	return _Masterchef.Contract.SushiPerBlock(&_Masterchef.CallOpts)
}

// SushiPerBlock is a free data retrieval call binding the contract method 0xb0bcf42a.
//
// Solidity: function sushiPerBlock() view returns(uint256)
func (_Masterchef *MasterchefCallerSession) SushiPerBlock() (*big.Int, error) {
	return _Masterchef.Contract.SushiPerBlock(&_Masterchef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchef *MasterchefCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchef *MasterchefSession) TotalAllocPoint() (*big.Int, error) {
	return _Masterchef.Contract.TotalAllocPoint(&_Masterchef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchef *MasterchefCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Masterchef.Contract.TotalAllocPoint(&_Masterchef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Masterchef *MasterchefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Masterchef.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Masterchef *MasterchefSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Masterchef.Contract.UserInfo(&_Masterchef.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Masterchef *MasterchefCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Masterchef.Contract.UserInfo(&_Masterchef.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Masterchef *MasterchefTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Masterchef *MasterchefSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Masterchef.Contract.Add(&_Masterchef.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Masterchef *MasterchefTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Masterchef.Contract.Add(&_Masterchef.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Masterchef *MasterchefTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Masterchef *MasterchefSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.Deposit(&_Masterchef.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Masterchef *MasterchefTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.Deposit(&_Masterchef.TransactOpts, _pid, _amount)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Masterchef *MasterchefTransactor) Dev(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "dev", _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Masterchef *MasterchefSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _Masterchef.Contract.Dev(&_Masterchef.TransactOpts, _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Masterchef *MasterchefTransactorSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _Masterchef.Contract.Dev(&_Masterchef.TransactOpts, _devaddr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Masterchef *MasterchefTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Masterchef *MasterchefSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.EmergencyWithdraw(&_Masterchef.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Masterchef *MasterchefTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.EmergencyWithdraw(&_Masterchef.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Masterchef *MasterchefTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Masterchef *MasterchefSession) MassUpdatePools() (*types.Transaction, error) {
	return _Masterchef.Contract.MassUpdatePools(&_Masterchef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Masterchef *MasterchefTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Masterchef.Contract.MassUpdatePools(&_Masterchef.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Masterchef *MasterchefTransactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Masterchef *MasterchefSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.Migrate(&_Masterchef.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_Masterchef *MasterchefTransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.Migrate(&_Masterchef.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchef *MasterchefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchef *MasterchefSession) RenounceOwnership() (*types.Transaction, error) {
	return _Masterchef.Contract.RenounceOwnership(&_Masterchef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchef *MasterchefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Masterchef.Contract.RenounceOwnership(&_Masterchef.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Masterchef *MasterchefTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Masterchef *MasterchefSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Masterchef.Contract.Set(&_Masterchef.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Masterchef *MasterchefTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Masterchef.Contract.Set(&_Masterchef.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Masterchef *MasterchefTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Masterchef *MasterchefSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Masterchef.Contract.SetMigrator(&_Masterchef.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Masterchef *MasterchefTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Masterchef.Contract.SetMigrator(&_Masterchef.TransactOpts, _migrator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchef *MasterchefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchef *MasterchefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Masterchef.Contract.TransferOwnership(&_Masterchef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchef *MasterchefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Masterchef.Contract.TransferOwnership(&_Masterchef.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Masterchef *MasterchefTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Masterchef *MasterchefSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.UpdatePool(&_Masterchef.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Masterchef *MasterchefTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.UpdatePool(&_Masterchef.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Masterchef *MasterchefTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchef.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Masterchef *MasterchefSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.Withdraw(&_Masterchef.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Masterchef *MasterchefTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchef.Contract.Withdraw(&_Masterchef.TransactOpts, _pid, _amount)
}

// MasterchefDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Masterchef contract.
type MasterchefDepositIterator struct {
	Event *MasterchefDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MasterchefDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterchefDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MasterchefDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MasterchefDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterchefDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterchefDeposit represents a Deposit event raised by the Masterchef contract.
type MasterchefDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterchefDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchef.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterchefDepositIterator{contract: _Masterchef.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MasterchefDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchef.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterchefDeposit)
				if err := _Masterchef.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) ParseDeposit(log types.Log) (*MasterchefDeposit, error) {
	event := new(MasterchefDeposit)
	if err := _Masterchef.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterchefEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Masterchef contract.
type MasterchefEmergencyWithdrawIterator struct {
	Event *MasterchefEmergencyWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MasterchefEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterchefEmergencyWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MasterchefEmergencyWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MasterchefEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterchefEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterchefEmergencyWithdraw represents a EmergencyWithdraw event raised by the Masterchef contract.
type MasterchefEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterchefEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchef.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterchefEmergencyWithdrawIterator{contract: _Masterchef.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *MasterchefEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchef.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterchefEmergencyWithdraw)
				if err := _Masterchef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) ParseEmergencyWithdraw(log types.Log) (*MasterchefEmergencyWithdraw, error) {
	event := new(MasterchefEmergencyWithdraw)
	if err := _Masterchef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterchefOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Masterchef contract.
type MasterchefOwnershipTransferredIterator struct {
	Event *MasterchefOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MasterchefOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterchefOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MasterchefOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MasterchefOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterchefOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterchefOwnershipTransferred represents a OwnershipTransferred event raised by the Masterchef contract.
type MasterchefOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Masterchef *MasterchefFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MasterchefOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Masterchef.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MasterchefOwnershipTransferredIterator{contract: _Masterchef.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Masterchef *MasterchefFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MasterchefOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Masterchef.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterchefOwnershipTransferred)
				if err := _Masterchef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Masterchef *MasterchefFilterer) ParseOwnershipTransferred(log types.Log) (*MasterchefOwnershipTransferred, error) {
	event := new(MasterchefOwnershipTransferred)
	if err := _Masterchef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterchefWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Masterchef contract.
type MasterchefWithdrawIterator struct {
	Event *MasterchefWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MasterchefWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterchefWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MasterchefWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MasterchefWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterchefWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterchefWithdraw represents a Withdraw event raised by the Masterchef contract.
type MasterchefWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterchefWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchef.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterchefWithdrawIterator{contract: _Masterchef.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MasterchefWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchef.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterchefWithdraw)
				if err := _Masterchef.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchef *MasterchefFilterer) ParseWithdraw(log types.Log) (*MasterchefWithdraw, error) {
	event := new(MasterchefWithdraw)
	if err := _Masterchef.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
