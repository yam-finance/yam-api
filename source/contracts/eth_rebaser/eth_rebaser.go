// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_rebaser

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

// EthRebaserABI is the input ABI used to generate the binding from.
const EthRebaserABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"yamAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reserveToken_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reservesContract_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"public_goods_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"public_goods_perc_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"MintAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDeviationThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDeviationThreshold\",\"type\":\"uint256\"}],\"name\":\"NewDeviationThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGov\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGov\",\"type\":\"address\"}],\"name\":\"NewGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSlippageFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSlippageFactor\",\"type\":\"uint256\"}],\"name\":\"NewMaxSlippageFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingGov\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingGov\",\"type\":\"address\"}],\"name\":\"NewPendingGov\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRebaseMintPerc\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRebaseMintPerc\",\"type\":\"uint256\"}],\"name\":\"NewRebaseMintPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldReserveContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newReserveContract\",\"type\":\"address\"}],\"name\":\"NewReserveContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservesAdded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yamsSold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yamsFromReserves\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yamsToReserves\",\"type\":\"uint256\"}],\"name\":\"TreasuryIncreased\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_MINT_PERC_PARAM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_SLIPPAGE_PARAM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptGov\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingGov_\",\"type\":\"address\"}],\"name\":\"_setPendingGov\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate_rebasing\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"uniSyncPairs_\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"balGulpPairs_\",\"type\":\"address[]\"}],\"name\":\"addSyncPairs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balGulpPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockTimestampLast\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deviationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eth_usdc_pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalGulpPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentTWAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUniSyncPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inRebaseWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"init_twap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isToken0\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRebaseTimestampSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxSlippageFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minRebaseTimeIntervalSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingGov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceCumulativeLastETHUSDC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceCumulativeLastYAMETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"public_goods\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"public_goods_perc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rebase\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rebaseDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rebaseLag\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rebaseMintPerc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rebaseWindowLengthSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rebaseWindowOffsetSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rebasingActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeBalPair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeUniPair\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reservesContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deviationThreshold_\",\"type\":\"uint256\"}],\"name\":\"setDeviationThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSlippageFactor_\",\"type\":\"uint256\"}],\"name\":\"setMaxSlippageFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rebaseLag_\",\"type\":\"uint256\"}],\"name\":\"setRebaseLag\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rebaseMintPerc_\",\"type\":\"uint256\"}],\"name\":\"setRebaseMintPerc\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minRebaseTimeIntervalSec_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebaseWindowOffsetSec_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebaseWindowLengthSec_\",\"type\":\"uint256\"}],\"name\":\"setRebaseTimingParameters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservesContract_\",\"type\":\"address\"}],\"name\":\"setReserveContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetRate_\",\"type\":\"uint256\"}],\"name\":\"setTargetRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setTransactionEnabled\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeOfTWAPInit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"trade_pair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uniSyncPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"yamAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EthRebaserBin is the compiled bytecode used for deploying new contracts.
var EthRebaserBin = "0x60806040523480156200001157600080fd5b5060405162005317380380620053178339810160408190526200003491620002dc565b61a8c06007556170806009556000806200005888886001600160e01b03620001a316565b91509150876001600160a01b0316826001600160a01b031614156200008a57601b805460ff1916600117905562000095565b601b805460ff191690555b620000ab8683836001600160e01b036200023e16565b601180546001600160a01b0319166001600160a01b039290921691909117905573a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48620000ed87828a6200023e565b601280546001600160a01b03199081166001600160a01b039384161790915560108054821698831698909817909755600f80548816998216999099179098555050600e8054851697871697909717909655601880548416929095169190911790935560199290925550665c4b2e47a5b000601a55670de0b6b3a7640000600455601460035567016345785d8a000060055566b1a2bc2ec50000600655610e10600a55600180543392169190911790555062000591565b600080826001600160a01b0316846001600160a01b03161415620001e45760405162461bcd60e51b8152600401620001db906200050a565b60405180910390fd5b826001600160a01b0316846001600160a01b0316106200020657828462000209565b83835b90925090506001600160a01b038216620002375760405162461bcd60e51b8152600401620001db906200051c565b9250929050565b600080806200025785856001600160e01b03620001a316565b91509150858282604051602001620002719291906200049a565b604051602081830303815290604052805190602001206040516020016200029a929190620004c4565b60408051601f1981840301815291905280516020909101209695505050505050565b8051620002c9816200056c565b92915050565b8051620002c98162000586565b60008060008060008060c08789031215620002f657600080fd5b6000620003048989620002bc565b96505060206200031789828a01620002bc565b95505060406200032a89828a01620002bc565b94505060606200033d89828a01620002bc565b93505060806200035089828a01620002bc565b92505060a06200036389828a01620002cf565b9150509295509295509295565b620003856200037f826200053c565b62000558565b82525050565b620003856200039a8262000549565b62000549565b6000620003af60208362000537565b7fe18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303815260200192915050565b6000620003ea6025836200052e565b7f556e697377617056324c6962726172793a204944454e544943414c5f41444452815264455353455360d81b602082015260400192915050565b60006200043360018362000537565b7fff00000000000000000000000000000000000000000000000000000000000000815260010192915050565b60006200046e601e836200052e565b7f556e697377617056324c6962726172793a205a45524f5f414444524553530000815260200192915050565b6000620004a8828562000370565b601482019150620004ba828462000370565b5060140192915050565b6000620004d18262000424565b9150620004df828562000370565b601482019150620004f182846200038b565b6020820191506200050282620003a0565b949350505050565b60208082528101620002c981620003db565b60208082528101620002c9816200045f565b90815260200190565b919050565b6000620002c9826200054c565b90565b6001600160a01b031690565b6000620002c9826000620002c98260601b90565b62000577816200053c565b81146200058357600080fd5b50565b620005778162000549565b614d7680620005a16000396000f3fe608060405234801561001057600080fd5b50600436106103785760003560e01c80636e9dde99116101d3578063cd87782611610104578063dcf93f32116100a2578063ec342ad01161007c578063ec342ad014610649578063f12b83b914610651578063f4325d6714610659578063fb0ce7d21461066157610378565b8063dcf93f3214610631578063dfebe32814610639578063e4751f7c1461064157610378565b8063cf1b927e116100de578063cf1b927e14610611578063d72cdafc14610619578063d94ad83714610621578063dcbff61e1461062957610378565b8063cd877826146105e3578063cdabdaac146105eb578063cea9d26f146105fe57610378565b80639671ecff11610171578063b532be181161014b578063b532be181461059e578063b60e1e3e146105b3578063c5700a02146105c6578063cc8fd393146105db57610378565b80639671ecff146105615780639ace38c214610574578063af14052c1461059657610378565b8063832a3035116101ad578063832a3035146105415780638d76a5bb14610549578063900cf0cf146105515780639466120f1461055957610378565b80636e9dde99146105135780637052b9021461052657806373f03dff1461052e57610378565b80632d3a4af9116102ad5780634f2b96291161024b57806357466c8b1161022557806357466c8b146104e857806363f6d4c8146104f05780636406ca5f146104f85780636bf9ace71461050057610378565b80634f2b9629146104af578063527a52c8146104c257806353a15edc146104d557610378565b806346c3bd1f1161028757806346c3bd1f146104845780634bda2e20146104975780634dc95de11461049f5780634e66f8ae146104a757610378565b80632d3a4af91461046c5780633a68eaf6146104745780633a93069b1461047c57610378565b806314eb3f241161031a5780631cab801c116102f45780631cab801c146104365780631e0cd44e1461043e57806320ce838914610451578063252408101461046457610378565b806314eb3f24146103fd57806316250fd4146104105780631b58ac4a1461042357610378565b806310d1e85c1161035657806310d1e85c146103b8578063111d0498146103cd578063126e19be146103e257806312d43a51146103f557610378565b8063021018991461037d5780630978af4f1461039b5780630dd2c45e146103a3575b600080fd5b610385610669565b6040516103929190614ad7565b60405180910390f35b61038561066f565b6103ab610675565b6040516103929190614853565b6103cb6103c6366004613e6c565b610691565b005b6103d5610d87565b60405161039291906148d0565b6103cb6103f0366004613e16565b610d98565b6103ab610eda565b6103ab61040b366004613fcd565b610ef6565b6103cb61041e366004614039565b610f2a565b6103cb610431366004613fcd565b610f83565b6103856110b3565b6103ab61044c366004613fcd565b6110b9565b6103cb61045f366004613fcd565b6110c6565b6103ab6110fc565b6103ab611118565b6103ab611134565b610385611150565b6103cb610492366004613fcd565b611156565b6103cb611333565b6103d5611416565b6103cb61141f565b6103cb6104bd366004613fcd565b6114c8565b6103cb6104d0366004613dab565b6115f3565b6103cb6104e3366004613fcd565b61169b565b6103cb611708565b610385611823565b610385611829565b6103cb61050e366004613fcd565b61182f565b6103cb610521366004614009565b6118a1565b61038561194f565b6103cb61053c366004613dab565b611955565b6103856119f1565b610385611b9e565b610385611baa565b610385611bb0565b6103cb61056f366004613ee5565b611bb6565b610587610582366004613fcd565b611ce8565b604051610392939291906148de565b6103cb611dd8565b6105a661220c565b60405161039291906148bf565b6103cb6105c1366004613fcd565b61227e565b6105ce6122f0565b6040516103929190614b8b565b6103856122fc565b610385612302565b6103cb6105f9366004613fcd565b612308565b6103d561060c366004613dc9565b61233e565b610385612377565b6103ab61237d565b610385612399565b61038561239f565b6103d56123a5565b6103ab6123ae565b6103856123ca565b6103856123d6565b6105a66123e2565b6103ab612452565b61038561246e565b60075481565b60165481565b60115473ffffffffffffffffffffffffffffffffffffffff1681565b60115473ffffffffffffffffffffffffffffffffffffffff1633146106eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a49565b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8416301461073a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614ab9565b6107426139cc565b818060200190516107569190810190613f6c565b600e54602082015191925073ffffffffffffffffffffffffffffffffffffffff1690156108f15760105460115460208401516040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808616946323b872dd946107e0949183169392169160040161487c565b602060405180830381600087803b1580156107fa57600080fd5b505af115801561080e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108329190810190613f4e565b508151602083015110156108ec576011546020830151835173ffffffffffffffffffffffffffffffffffffffff808516936340c10f199391169161087b9163ffffffff61247416565b6040518363ffffffff1660e01b81526004016108989291906148a4565b602060405180830381600087803b1580156108b257600080fd5b505af11580156108c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108ea9190810190613f4e565b505b6109a0565b60115482516040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116936340c10f199361094c9392909116916004016148a4565b602060405180830381600087803b15801561096657600080fd5b505af115801561097a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061099e9190810190613f4e565b505b604082015115610a5c5760105460408381015190517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116936340c10f1993610a089392909116916004016148a4565b602060405180830381600087803b158015610a2257600080fd5b505af1158015610a36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610a5a9190810190613f4e565b505b601b5460ff1615610bfe5760185473ffffffffffffffffffffffffffffffffffffffff1615801590610a9057506000601954115b15610b85576000610ac4670de0b6b3a7640000610ab8601954886124bf90919063ffffffff16565b9063ffffffff61251316565b600f54601054919250610b009173ffffffffffffffffffffffffffffffffffffffff9182169116610afb888563ffffffff61247416565b612555565b600f54601854610b2a9173ffffffffffffffffffffffffffffffffffffffff908116911683612555565b7fb335015c214ae37ed112cc5eb042235c0ea40a7617987e1bd847839143872350610b5b868363ffffffff61247416565b845160208601516040808801519051610b779493929190614b56565b60405180910390a150610bf9565b600f54601054610baf9173ffffffffffffffffffffffffffffffffffffffff908116911686612555565b7fb335015c214ae37ed112cc5eb042235c0ea40a7617987e1bd84783914387235084836000015184602001518560400151604051610bf09493929190614b56565b60405180910390a15b610d7f565b60185473ffffffffffffffffffffffffffffffffffffffff1615801590610c2757506000601954115b15610d0b576000610c4f670de0b6b3a7640000610ab8601954896124bf90919063ffffffff16565b600f54601054919250610c869173ffffffffffffffffffffffffffffffffffffffff9182169116610afb898563ffffffff61247416565b600f54601854610cb09173ffffffffffffffffffffffffffffffffffffffff908116911683612555565b7fb335015c214ae37ed112cc5eb042235c0ea40a7617987e1bd847839143872350610ce1878363ffffffff61247416565b845160208601516040808801519051610cfd9493929190614b56565b60405180910390a150610d7f565b600f54601054610d359173ffffffffffffffffffffffffffffffffffffffff908116911687612555565b7fb335015c214ae37ed112cc5eb042235c0ea40a7617987e1bd84783914387235085836000015184602001518560400151604051610d769493929190614b56565b60405180910390a15b505050505050565b6000610d91612612565b5060015b90565b60015473ffffffffffffffffffffffffffffffffffffffff163314610dbc57600080fd5b600060405180606001604052806001151581526020018573ffffffffffffffffffffffffffffffffffffffff16815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093909452505083546001818101808755958352602092839020855160029093020180548487015173ffffffffffffffffffffffffffffffffffffffff16610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff9415157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909216919091179390931692909217825560408501518051929450610ed193918501929101906139ed565b50505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60138181548110610f0357fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60015473ffffffffffffffffffffffffffffffffffffffff163314610f4e57600080fd5b60008311610f5b57600080fd5b828210610f6757600080fd5b8281830110610f7557600080fd5b600792909255600955600a55565b60015473ffffffffffffffffffffffffffffffffffffffff163314610fa757600080fd5b6014548110610fb5576110b0565b805b6014547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181101561107c5760148160010181548110610ff357fe5b6000918252602090912001546014805473ffffffffffffffffffffffffffffffffffffffff909216918390811061102657fe5b600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055600101610fb7565b5060148054906110ae907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8301613a6b565b505b50565b60055481565b60148181548110610f0357fe5b60015473ffffffffffffffffffffffffffffffffffffffff1633146110ea57600080fd5b600081116110f757600080fd5b600355565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60125473ffffffffffffffffffffffffffffffffffffffff1681565b60105473ffffffffffffffffffffffffffffffffffffffff1681565b60085481565b60015473ffffffffffffffffffffffffffffffffffffffff16331461117a57600080fd5b60005481106111b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a89565b6000547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0181101561130257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811061121157fe5b90600052602060002090600202016000828154811061122c57fe5b600091825260209091208254600292830290910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff90921615159190911780825583547fffffffffffffffffffffff0000000000000000000000000000000000000000ff9091166101009182900473ffffffffffffffffffffffffffffffffffffffff1682021782556001808501805493946112fe948387019492938116159092027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190911604613a8f565b5050505b60008054906110ae907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8301613b04565b60025473ffffffffffffffffffffffffffffffffffffffff163314611384576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e2906149e9565b60018054600280547fffffffffffffffffffffffff000000000000000000000000000000000000000080841673ffffffffffffffffffffffffffffffffffffffff83811691909117958690559116909155604051918116927f1f14cfc03e486d23acee577b07bc0b3b23f4888c91fcdba5e0fef5a2549d55239261140b9285921690614861565b60405180910390a150565b600c5460ff1681565b6000600d541161145b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614969565b61a8c0600d540142101561149b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e2906149a9565b600c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60015473ffffffffffffffffffffffffffffffffffffffff1633146114ec57600080fd5b60135481106114fa576110b0565b805b6013547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018110156115c1576013816001018154811061153857fe5b6000918252602090912001546013805473ffffffffffffffffffffffffffffffffffffffff909216918390811061156b57fe5b600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556001016114fc565b5060138054906110ae907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8301613a6b565b60015473ffffffffffffffffffffffffffffffffffffffff16331461161757600080fd5b6010805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008316179092556040519116907fce840d2205f08f33375689943da5da9fdfde146fcbb5553b17910a60c8284a209061168f9083908590614861565b60405180910390a15050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146116bf57600080fd5b6000600654116116ce57600080fd5b60068054908290556040517f2a5cda4d16fba415b52d90b59ee30d4cb16494da9fd1ee51c4d5bac4a1f75bbe9061168f9083908590614b20565b600d5415611742576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a79565b601154601b5460009182916117709173ffffffffffffffffffffffffffffffffffffffff169060ff166126ff565b601254919350915060009061179b9073ffffffffffffffffffffffffffffffffffffffff16826126ff565b50905060008263ffffffff16116117de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614999565b6015805463ffffffff9093167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000000090931683179055601692909255601791909155600d55565b60035481565b61a8c081565b60015473ffffffffffffffffffffffffffffffffffffffff16331461185357600080fd5b6703782dace9d90000811061186757600080fd5b60058054908290556040517f59b3ffce759ec92c629beee27554d8fbc2ca1a05020fa0cf500c890c172094be9061168f9083908590614b20565b60015473ffffffffffffffffffffffffffffffffffffffff1633146118c557600080fd5b6000548210611900576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614989565b806000838154811061190e57fe5b6000918252602090912060029091020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790555050565b60095481565b60015473ffffffffffffffffffffffffffffffffffffffff16331461197957600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008316179092556040519116907f6163d5b9efd962645dd649e6e48a61bcb0f9df00997a2398b80d135a9ab0c61e9061168f9083908590614861565b601154601b5460009182918291611a239173ffffffffffffffffffffffffffffffffffffffff9091169060ff166126ff565b6012549193509150600090611a4e9073ffffffffffffffffffffffffffffffffffffffff16826126ff565b5060155460165491925063ffffffff908116840391600091831690860381611a7257fe5b047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905060008263ffffffff16601754850381611aa857fe5b047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690506000807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff77ffffffffffffffffffffffffffffffffffffffffffffffff16841115611b2457670de0b6b3a7640000607085901c029150611b35565b670de0b6b3a7640000840260701c91505b77ffffffffffffffffffffffffffffffffffffffffffffffff831115611b6a5750607082901c670de0b6b3a764000002611b7a565b50670de0b6b3a7640000820260701c5b611b91620f4240610ab8848463ffffffff6124bf16565b9850505050505050505090565b6701a35734e8f4380081565b600b5481565b600a5481565b60015473ffffffffffffffffffffffffffffffffffffffff163314611bda57600080fd5b60005b8251811015611c5e576013838281518110611bf457fe5b60209081029190910181015182546001808201855560009485529290932090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915501611bdd565b5060005b8151811015611ce3576014828281518110611c7957fe5b60209081029190910181015182546001808201855560009485529290932090920180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9093169290921790915501611c62565b505050565b60008181548110611cf557fe5b6000918252602091829020600291820201805460018083018054604080516101009483161585027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190921696909604601f810188900488028201880190965285815260ff841697509190920473ffffffffffffffffffffffffffffffffffffffff169492939092830182828015611dce5780601f10611da357610100808354040283529160200191611dce565b820191906000526020600020905b815481529060010190602001808311611db157829003601f168201915b5050505050905083565b333214611e11576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a59565b611e19612612565b42611e3160075460085461294690919063ffffffff16565b10611e3b57600080fd5b611e72600954611e66611e596007544261298590919063ffffffff16565b429063ffffffff61247416565b9063ffffffff61294616565b600855600b54611e8990600163ffffffff61294616565b600b556000611e966129c7565b9050600080611ea483612b12565b60035491935091508290611ebf90829063ffffffff61251316565b600e5490915073ffffffffffffffffffffffffffffffffffffffff16821561203f578073ffffffffffffffffffffffffffffffffffffffff166311d3e6c46040518163ffffffff1660e01b815260040160206040518083038186803b158015611f2757600080fd5b505afa158015611f3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611f5f9190810190613feb565b612008670de0b6b3a7640000610ab8611f7e828763ffffffff61294616565b8573ffffffffffffffffffffffffffffffffffffffff1663b6fa85766040518163ffffffff1660e01b815260040160206040518083038186803b158015611fc457600080fd5b505afa158015611fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611ffc9190810190613feb565b9063ffffffff6124bf16565b1061203f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e2906149c9565b60008173ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561208757600080fd5b505afa15801561209b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506120bf9190810190613feb565b9050600084156121205760006120ec670de0b6b3a7640000610ab8600554886124bf90919063ffffffff16565b90506120fe858263ffffffff61247416565b945061211c670de0b6b3a7640000610ab8858463ffffffff6124bf16565b9150505b600b546040517f7af548c100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851691637af548c191612178919088908a90600401614b2e565b602060405180830381600087803b15801561219257600080fd5b505af11580156121a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506121ca9190810190613feb565b507f8e37eb2ee3a6f3c8b13b8973588daad75a4ce752de14c00006bd8247f4e212e8816040516121fa9190614ad7565b60405180910390a1610ed18187612b95565b606080601380548060200260200160405190810160405280929190818152602001828054801561227257602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311612247575b50939550505050505090565b60015473ffffffffffffffffffffffffffffffffffffffff1633146122a257600080fd5b6701a35734e8f4380081106122b657600080fd5b601a8054908290556040517fe21b25c4eda0340cd924f3247795d0acde6c304b68ae77657bb2d4e840198bf89061168f9083908590614b20565b60155463ffffffff1681565b60045481565b600d5481565b60015473ffffffffffffffffffffffffffffffffffffffff16331461232c57600080fd5b6000811161233957600080fd5b600455565b60015460009073ffffffffffffffffffffffffffffffffffffffff16331461236557600080fd5b612370848484612555565b9392505050565b601a5481565b600e5473ffffffffffffffffffffffffffffffffffffffff1681565b60065481565b60175481565b601b5460ff1681565b60185473ffffffffffffffffffffffffffffffffffffffff1681565b6703782dace9d9000081565b670de0b6b3a764000081565b60608060148054806020026020016040519081016040528092919081815260200182805480156122725760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116122475750939550505050505090565b600f5473ffffffffffffffffffffffffffffffffffffffff1681565b60195481565b60006124b683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612ebc565b90505b92915050565b6000826124ce575060006124b9565b828202828482816124db57fe5b04146124b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e2906149d9565b60006124b683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250612f02565b604051611ce39084907fa9059cbb000000000000000000000000000000000000000000000000000000009061259090869086906024016148a4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152612f53565b600c5460ff1661264e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e2906149f9565b60095460075461266590429063ffffffff61298516565b101561269d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a19565b600a546009546126b29163ffffffff61294616565b6007546126c690429063ffffffff61298516565b106126fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a29565b565b60008061270a613009565b905060008060008673ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561275757600080fd5b505afa15801561276b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061278f9190810190613f8a565b925092509250851561286e578673ffffffffffffffffffffffffffffffffffffffff16635909c0d56040518163ffffffff1660e01b815260040160206040518083038186803b1580156127e157600080fd5b505afa1580156127f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506128199190810190613feb565b94508363ffffffff168163ffffffff16146128695780840363ffffffff81166128428486613013565b517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16029590950194505b61293c565b8673ffffffffffffffffffffffffffffffffffffffff16635a3d54936040518163ffffffff1660e01b815260040160206040518083038186803b1580156128b457600080fd5b505afa1580156128c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506128ec9190810190613feb565b94508363ffffffff168163ffffffff161461293c5780840363ffffffff81166129158585613013565b517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16029590950194505b5050509250929050565b6000828201838110156124b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614979565b60006124b683836040518060400160405280601881526020017f536166654d6174683a206d6f64756c6f206279207a65726f00000000000000008152506130d6565b601154601b54600091829182916129f99173ffffffffffffffffffffffffffffffffffffffff9091169060ff166126ff565b6012549193509150600090612a249073ffffffffffffffffffffffffffffffffffffffff16826126ff565b5060155460165491925063ffffffff908116840391600091831690860381612a4857fe5b047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905060008263ffffffff16601754850381612a7e57fe5b60168890556017869055601580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8916179055047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905060008077ffffffffffffffffffffffffffffffffffffffffffffffff841115611b2457670de0b6b3a7640000607085901c029150611b35565b600080612b1e83613124565b15612b2e57506000905080612b90565b600454831115612b6657600454612b5b90610ab8670de0b6b3a7640000611ffc878463ffffffff61247416565b600191509150612b90565b600454612b8990610ab8670de0b6b3a7640000611ffc838863ffffffff61247416565b6000915091505b915091565b60005b601354811015612c415760138181548110612baf57fe5b6000918252602082200154604080517ffff6cae9000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9092169263fff6cae99260048084019382900301818387803b158015612c1d57600080fd5b505af1158015612c31573d6000803e3d6000fd5b505060019092019150612b989050565b5060005b601454811015612cfd5760148181548110612c5c57fe5b600091825260209091200154600e546040517f8c28cbe800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff92831692638c28cbe892612cbf92911690600401614853565b600060405180830381600087803b158015612cd957600080fd5b505af1158015612ced573d6000803e3d6000fd5b505060019092019150612c459050565b508115612d0e57612d0e82826131a0565b60005b600054811015611ce3576000808281548110612d2957fe5b60009182526020909120600290910201805490915060ff1615612eb35780546001808301805460408051602060026101009685161587027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190941693909304601f8101849004840282018401909252818152600095612e1c95900473ffffffffffffffffffffffffffffffffffffffff16939092909190830182828015612e125780601f10612de757610100808354040283529160200191612e12565b820191906000526020600020905b815481529060010190602001808311612df557829003601f168201915b5050505050613700565b905080612eb157815460405161010090910473ffffffffffffffffffffffffffffffffffffffff16907f8091ecaaa54ebb82e02d36c2c336528e0fcb9b3430fc1291ac88295032b9c26390612e779086906001870190614ae5565b60405180910390a26040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a99565b505b50600101612d11565b60008184841115612efa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e29190614958565b505050900390565b60008183612f3d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e29190614958565b506000838581612f4957fe5b0495945050505050565b6060612fb5826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166137239092919063ffffffff16565b805190915015611ce35780806020019051612fd39190810190613f4e565b611ce3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a39565b63ffffffff421690565b61301b613b30565b6000826dffffffffffffffffffffffffffff1611613065576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614aa9565b6040805160208101909152806dffffffffffffffffffffffffffff84167bffffffffffffffffffffffffffff0000000000000000000000000000607087901b16816130ac57fe5b047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815250905092915050565b60008183613111576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e29190614958565b5082848161311b57fe5b06949350505050565b60008061314a670de0b6b3a7640000610ab86006546004546124bf90919063ffffffff16565b9050600454831015801561317157508061316f6004548561247490919063ffffffff16565b105b8061237057506004548310801561237057506004548190613198908563ffffffff61247416565b109392505050565b601154600e54604080517f0902f1ac000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416939092169160009182918591630902f1ac916004808301926060929190829003018186803b15801561321857600080fd5b505afa15801561322c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506132509190810190613f8a565b506010546040517f70a082310000000000000000000000000000000000000000000000000000000081526dffffffffffffffffffffffffffff938416955091909216925060009173ffffffffffffffffffffffffffffffffffffffff808716926370a08231926132c4921690600401614853565b60206040518083038186803b1580156132dc57600080fd5b505afa1580156132f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506133149190810190613feb565b9050600061332384848961373a565b905061332d6139cc565b506040805160608101825282815260208101849052600091810191909152601b5460ff161561354457613366898463ffffffff61294616565b82111561342657600061337c848b0187876137e7565b8a85018352602080840186905260405191925073ffffffffffffffffffffffffffffffffffffffff8a169163022c0d9f91600091859130916133c091899101614ac9565b6040516020818303038152906040526040518563ffffffff1660e01b81526004016133ee9493929190614914565b600060405180830381600087803b15801561340857600080fd5b505af115801561341c573d6000803e3d6000fd5b505050505061353f565b8282111561348957600061343b8387876137e7565b905061344f8a85850363ffffffff61247416565b8260400181815250508773ffffffffffffffffffffffffffffffffffffffff1663022c0d9f60008330866040516020016133c09190614ac9565b60006134968387876137e7565b905082826020018181525050898260400181815250508773ffffffffffffffffffffffffffffffffffffffff1663022c0d9f60008330866040516020016134dd9190614ac9565b6040516020818303038152906040526040518563ffffffff1660e01b815260040161350b9493929190614914565b600060405180830381600087803b15801561352557600080fd5b505af1158015613539573d6000803e3d6000fd5b50505050505b6136f5565b613554898463ffffffff61294616565b8211156135dc57600061356a848b0186886137e7565b8a85018352602080840186905260405191925073ffffffffffffffffffffffffffffffffffffffff8a169163022c0d9f91849160009130916135ae91899101614ac9565b6040516020818303038152906040526040518563ffffffff1660e01b815260040161350b9493929190614b05565b8282111561363f5760006135f18386886137e7565b90506136058a85850363ffffffff61247416565b8260400181815250508773ffffffffffffffffffffffffffffffffffffffff1663022c0d9f82600030866040516020016135ae9190614ac9565b600061364c8386886137e7565b905082826020018181525050898260400181815250508773ffffffffffffffffffffffffffffffffffffffff1663022c0d9f82600030866040516020016136939190614ac9565b6040516020818303038152906040526040518563ffffffff1660e01b81526004016136c19493929190614b05565b600060405180830381600087803b1580156136db57600080fd5b505af11580156136ef573d6000803e3d6000fd5b50505050505b505050505050505050565b6000806040516020840160008286518360008a6187965a03f19695505050505050565b606061373284846000856138c1565b949350505050565b601b5460009060ff161561379b5767016345785d8a0000821061377f57613778670de0b6b3a7640000610ab8601a54876124bf90919063ffffffff16565b9050612370565b6137786729a2241af62c0000610ab8868563ffffffff6124bf16565b67016345785d8a000082106137cb57613778670de0b6b3a7640000610ab8601a54866124bf90919063ffffffff16565b6137786729a2241af62c0000610ab8858563ffffffff6124bf16565b6000808411613822576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a69565b6000831180156138325750600082115b613868576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e2906149b9565b600061387c856103e563ffffffff6124bf16565b90506000613890828563ffffffff6124bf16565b905060006138aa83611e66886103e863ffffffff6124bf16565b90508082816138b557fe5b04979650505050505050565b60606138cc856139c6565b613902576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e290614a09565b600060608673ffffffffffffffffffffffffffffffffffffffff16858760405161392c9190614847565b60006040518083038185875af1925050503d8060008114613969576040519150601f19603f3d011682016040523d82523d6000602084013e61396e565b606091505b509150915081156139825791506137329050565b8051156139925780518082602001fd5b836040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e29190614958565b3b151590565b60405180606001604052806000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10613a2e57805160ff1916838001178555613a5b565b82800160010185558215613a5b579182015b82811115613a5b578251825591602001919060010190613a40565b50613a67929150613b42565b5090565b815481835581811115611ce357600083815260209020611ce3918101908301613b42565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10613ac85780548555613a5b565b82800160010185558215613a5b57600052602060002091601f016020900482015b82811115613a5b578254825591600101919060010190613ae9565b815481835581811115611ce357600202816002028360005260206000209182019101611ce39190613b5c565b60408051602081019091526000815290565b610d9591905b80821115613a675760008155600101613b48565b610d9591905b80821115613a675780547fffffffffffffffffffffff0000000000000000000000000000000000000000001681556000613b9f6001830182613ba8565b50600201613b62565b50805460018160011615610100020316600290046000825580601f10613bce57506110b0565b601f0160209004906000526020600020908101906110b09190613b42565b80356124b981614cfb565b600082601f830112613c0857600080fd5b8135613c1b613c1682614bc0565b614b99565b91508181835260208401935060208101905083856020840282011115613c4057600080fd5b60005b83811015613c6c5781613c568882613bec565b8452506020928301929190910190600101613c43565b5050505092915050565b80356124b981614d0f565b80516124b981614d0f565b60008083601f840112613c9e57600080fd5b50813567ffffffffffffffff811115613cb657600080fd5b602083019150836001820283011115613cce57600080fd5b9250929050565b600082601f830112613ce657600080fd5b8135613cf4613c1682614be1565b91508082526020830160208301858383011115613d1057600080fd5b613d1b838284614c9b565b50505092915050565b600060608284031215613d3657600080fd5b613d406060614b99565b90506000613d4e8484613d95565b8252506020613d5f84848301613d95565b6020830152506040613d7384828501613d95565b60408301525092915050565b80516124b981614d18565b80356124b981614d21565b80516124b981614d21565b80516124b981614d2a565b600060208284031215613dbd57600080fd5b60006137328484613bec565b600080600060608486031215613dde57600080fd5b6000613dea8686613bec565b9350506020613dfb86828701613bec565b9250506040613e0c86828701613d8a565b9150509250925092565b600080600060408486031215613e2b57600080fd5b6000613e378686613bec565b935050602084013567ffffffffffffffff811115613e5457600080fd5b613e6086828701613c8c565b92509250509250925092565b60008060008060808587031215613e8257600080fd5b6000613e8e8787613bec565b9450506020613e9f87828801613d8a565b9350506040613eb087828801613d8a565b925050606085013567ffffffffffffffff811115613ecd57600080fd5b613ed987828801613cd5565b91505092959194509250565b60008060408385031215613ef857600080fd5b823567ffffffffffffffff811115613f0f57600080fd5b613f1b85828601613bf7565b925050602083013567ffffffffffffffff811115613f3857600080fd5b613f4485828601613bf7565b9150509250929050565b600060208284031215613f6057600080fd5b60006137328484613c81565b600060608284031215613f7e57600080fd5b60006137328484613d24565b600080600060608486031215613f9f57600080fd5b6000613fab8686613d7f565b9350506020613fbc86828701613d7f565b9250506040613e0c86828701613da0565b600060208284031215613fdf57600080fd5b60006137328484613d8a565b600060208284031215613ffd57600080fd5b60006137328484613d95565b6000806040838503121561401c57600080fd5b60006140288585613d8a565b9250506020613f4485828601613c76565b60008060006060848603121561404e57600080fd5b600061405a8686613d8a565b9350506020613dfb86828701613d8a565b6000614077838361407f565b505060200190565b61408881614c4b565b82525050565b600061409982614c39565b6140a38185614c3d565b93506140ae83614c27565b8060005b838110156140dc5781516140c6888261406b565b97506140d183614c27565b9250506001016140b2565b509495945050505050565b61408881614c56565b60006140fb82614c39565b6141058185614c3d565b9350614115818560208601614ca7565b61411e81614cd3565b9093019392505050565b600061413382614c39565b61413d8185614c46565b935061414d818560208601614ca7565b9290920192915050565b60008154600181166000811461417457600181146141b8576141f7565b607f60028304166141858187614c3d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00841681529550506020850192506141f7565b600282046141c68187614c3d565b95506141d185614c2d565b60005b828110156141f0578154888201526001909101906020016141d4565b8701945050505b505092915050565b61408881614c90565b6000614215602783614c3d565b7f74776170207761736e7420696e74697469617465642c2063616c6c20696e697481527f5f74776170282900000000000000000000000000000000000000000000000000602082015260400192915050565b6000614274601b83614c3d565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000815260200192915050565b60006142ad602883614c3d565b7f696e646578206d75737420626520696e2072616e6765206f662073746f72656481527f207478206c697374000000000000000000000000000000000000000000000000602082015260400192915050565b600061430c600983614c3d565b7f6e6f207472616465730000000000000000000000000000000000000000000000815260200192915050565b6000614345600a83614c3d565b7f21656e645f64656c617900000000000000000000000000000000000000000000815260200192915050565b600061437e602883614c3d565b7f556e697377617056324c6962726172793a20494e53554646494349454e545f4c81527f4951554944495459000000000000000000000000000000000000000000000000602082015260400192915050565b60006143dd602283614c3d565b7f6e6577207363616c696e6720666163746f722077696c6c20626520746f6f206281527f6967000000000000000000000000000000000000000000000000000000000000602082015260400192915050565b600061443c602183614c3d565b7f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f81527f7700000000000000000000000000000000000000000000000000000000000000602082015260400192915050565b600061449b600883614c3d565b7f2170656e64696e67000000000000000000000000000000000000000000000000815260200192915050565b60006144d4601383614c3d565b7f7265626173696e67206e6f742061637469766500000000000000000000000000815260200192915050565b600061450d601d83614c3d565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000815260200192915050565b6000614546600983614c3d565b7f746f6f206561726c790000000000000000000000000000000000000000000000815260200192915050565b600061457f600883614c3d565b7f746f6f206c617465000000000000000000000000000000000000000000000000815260200192915050565b60006145b8602a83614c3d565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e81527f6f74207375636365656400000000000000000000000000000000000000000000602082015260400192915050565b6000614617600e83614c3d565b7f626164206d73672e73656e646572000000000000000000000000000000000000815260200192915050565b6000614650600483614c3d565b7f21454f4100000000000000000000000000000000000000000000000000000000815260200192915050565b6000614689602b83614c3d565b7f556e697377617056324c6962726172793a20494e53554646494349454e545f4981527f4e5055545f414d4f554e54000000000000000000000000000000000000000000602082015260400192915050565b60006146e8601183614c3d565b7f616c726561647920616374697661746564000000000000000000000000000000815260200192915050565b6000614721601383614c3d565b7f696e646578206f7574206f6620626f756e647300000000000000000000000000815260200192915050565b600061475a601283614c3d565b7f5472616e73616374696f6e204661696c65640000000000000000000000000000815260200192915050565b6000614793601783614c3d565b7f4669786564506f696e743a204449565f42595f5a45524f000000000000000000815260200192915050565b60006147cc600a83614c3d565b7f626164206f726967696e00000000000000000000000000000000000000000000815260200192915050565b805160608301906148098482614835565b50602082015161481c6020850182614835565b50604082015161482f6040850182614835565b50505050565b61408881610d95565b61408881614c87565b60006123708284614128565b602081016124b9828461407f565b6040810161486f828561407f565b612370602083018461407f565b6060810161488a828661407f565b614897602083018561407f565b6137326040830184614835565b604081016148b2828561407f565b6123706020830184614835565b602080825281016124b6818461408e565b602081016124b982846140e7565b606081016148ec82866140e7565b6148f9602083018561407f565b818103604083015261490b81846140f0565b95945050505050565b6080810161492282876141ff565b61492f6020830186614835565b61493c604083018561407f565b818103606083015261494e81846140f0565b9695505050505050565b602080825281016124b681846140f0565b602080825281016124b981614208565b602080825281016124b981614267565b602080825281016124b9816142a0565b602080825281016124b9816142ff565b602080825281016124b981614338565b602080825281016124b981614371565b602080825281016124b9816143d0565b602080825281016124b98161442f565b602080825281016124b98161448e565b602080825281016124b9816144c7565b602080825281016124b981614500565b602080825281016124b981614539565b602080825281016124b981614572565b602080825281016124b9816145ab565b602080825281016124b98161460a565b602080825281016124b981614643565b602080825281016124b98161467c565b602080825281016124b9816146db565b602080825281016124b981614714565b602080825281016124b98161474d565b602080825281016124b981614786565b602080825281016124b9816147bf565b606081016124b982846147f8565b602081016124b98284614835565b60408101614af38285614835565b81810360208301526137328184614157565b60808101614b138287614835565b61492f60208301866141ff565b604081016148b28285614835565b60608101614b3c8286614835565b614b496020830185614835565b61373260408301846140e7565b60808101614b648287614835565b614b716020830186614835565b614b7e6040830185614835565b61490b6060830184614835565b602081016124b9828461483e565b60405181810167ffffffffffffffff81118282101715614bb857600080fd5b604052919050565b600067ffffffffffffffff821115614bd757600080fd5b5060209081020190565b600067ffffffffffffffff821115614bf857600080fd5b506020601f919091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160190565b60200190565b60009081526020902090565b5190565b90815260200190565b919050565b60006124b982614c6e565b151590565b6dffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b63ffffffff1690565b60006124b982610d95565b82818337506000910152565b60005b83811015614cc2578181015183820152602001614caa565b8381111561482f5750506000910152565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690565b614d0481614c4b565b81146110b057600080fd5b614d0481614c56565b614d0481614c5b565b614d0481610d95565b614d0481614c8756fea365627a7a7231582026cca9128a3b8e76bfd6381a9c727776a240529eee69271242e5a71ade9f57516c6578706572696d656e74616cf564736f6c634300050f00400000000000000000000000000aacfbec6a24756c20d41914f2caba817c0d8521000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000c0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac00000000000000000000000097990b693835da58a281636296d2bf02787dea17000000000000000000000000de21f729137c5af1b01d73af1dc21effa2b8a0d6000000000000000000000000000000000000000000000000002386f26fc10000"

// DeployEthRebaser deploys a new Ethereum contract, binding an instance of EthRebaser to it.
func DeployEthRebaser(auth *bind.TransactOpts, backend bind.ContractBackend, yamAddress_ common.Address, reserveToken_ common.Address, factory common.Address, reservesContract_ common.Address, public_goods_ common.Address, public_goods_perc_ *big.Int) (common.Address, *types.Transaction, *EthRebaser, error) {
	parsed, err := abi.JSON(strings.NewReader(EthRebaserABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthRebaserBin), backend, yamAddress_, reserveToken_, factory, reservesContract_, public_goods_, public_goods_perc_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthRebaser{EthRebaserCaller: EthRebaserCaller{contract: contract}, EthRebaserTransactor: EthRebaserTransactor{contract: contract}, EthRebaserFilterer: EthRebaserFilterer{contract: contract}}, nil
}

// EthRebaser is an auto generated Go binding around an Ethereum contract.
type EthRebaser struct {
	EthRebaserCaller     // Read-only binding to the contract
	EthRebaserTransactor // Write-only binding to the contract
	EthRebaserFilterer   // Log filterer for contract events
}

// EthRebaserCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthRebaserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthRebaserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthRebaserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthRebaserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthRebaserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthRebaserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthRebaserSession struct {
	Contract     *EthRebaser       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthRebaserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthRebaserCallerSession struct {
	Contract *EthRebaserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EthRebaserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthRebaserTransactorSession struct {
	Contract     *EthRebaserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EthRebaserRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthRebaserRaw struct {
	Contract *EthRebaser // Generic contract binding to access the raw methods on
}

// EthRebaserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthRebaserCallerRaw struct {
	Contract *EthRebaserCaller // Generic read-only contract binding to access the raw methods on
}

// EthRebaserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthRebaserTransactorRaw struct {
	Contract *EthRebaserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthRebaser creates a new instance of EthRebaser, bound to a specific deployed contract.
func NewEthRebaser(address common.Address, backend bind.ContractBackend) (*EthRebaser, error) {
	contract, err := bindEthRebaser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthRebaser{EthRebaserCaller: EthRebaserCaller{contract: contract}, EthRebaserTransactor: EthRebaserTransactor{contract: contract}, EthRebaserFilterer: EthRebaserFilterer{contract: contract}}, nil
}

// NewEthRebaserCaller creates a new read-only instance of EthRebaser, bound to a specific deployed contract.
func NewEthRebaserCaller(address common.Address, caller bind.ContractCaller) (*EthRebaserCaller, error) {
	contract, err := bindEthRebaser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthRebaserCaller{contract: contract}, nil
}

// NewEthRebaserTransactor creates a new write-only instance of EthRebaser, bound to a specific deployed contract.
func NewEthRebaserTransactor(address common.Address, transactor bind.ContractTransactor) (*EthRebaserTransactor, error) {
	contract, err := bindEthRebaser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthRebaserTransactor{contract: contract}, nil
}

// NewEthRebaserFilterer creates a new log filterer instance of EthRebaser, bound to a specific deployed contract.
func NewEthRebaserFilterer(address common.Address, filterer bind.ContractFilterer) (*EthRebaserFilterer, error) {
	contract, err := bindEthRebaser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthRebaserFilterer{contract: contract}, nil
}

// bindEthRebaser binds a generic wrapper to an already deployed contract.
func bindEthRebaser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthRebaserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthRebaser *EthRebaserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthRebaser.Contract.EthRebaserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthRebaser *EthRebaserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthRebaser.Contract.EthRebaserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthRebaser *EthRebaserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthRebaser.Contract.EthRebaserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthRebaser *EthRebaserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthRebaser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthRebaser *EthRebaserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthRebaser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthRebaser *EthRebaserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthRebaser.Contract.contract.Transact(opts, method, params...)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) BASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_EthRebaser *EthRebaserSession) BASE() (*big.Int, error) {
	return _EthRebaser.Contract.BASE(&_EthRebaser.CallOpts)
}

// BASE is a free data retrieval call binding the contract method 0xec342ad0.
//
// Solidity: function BASE() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) BASE() (*big.Int, error) {
	return _EthRebaser.Contract.BASE(&_EthRebaser.CallOpts)
}

// MAXMINTPERCPARAM is a free data retrieval call binding the contract method 0xe4751f7c.
//
// Solidity: function MAX_MINT_PERC_PARAM() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) MAXMINTPERCPARAM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "MAX_MINT_PERC_PARAM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMINTPERCPARAM is a free data retrieval call binding the contract method 0xe4751f7c.
//
// Solidity: function MAX_MINT_PERC_PARAM() view returns(uint256)
func (_EthRebaser *EthRebaserSession) MAXMINTPERCPARAM() (*big.Int, error) {
	return _EthRebaser.Contract.MAXMINTPERCPARAM(&_EthRebaser.CallOpts)
}

// MAXMINTPERCPARAM is a free data retrieval call binding the contract method 0xe4751f7c.
//
// Solidity: function MAX_MINT_PERC_PARAM() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) MAXMINTPERCPARAM() (*big.Int, error) {
	return _EthRebaser.Contract.MAXMINTPERCPARAM(&_EthRebaser.CallOpts)
}

// MAXSLIPPAGEPARAM is a free data retrieval call binding the contract method 0x8d76a5bb.
//
// Solidity: function MAX_SLIPPAGE_PARAM() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) MAXSLIPPAGEPARAM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "MAX_SLIPPAGE_PARAM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSLIPPAGEPARAM is a free data retrieval call binding the contract method 0x8d76a5bb.
//
// Solidity: function MAX_SLIPPAGE_PARAM() view returns(uint256)
func (_EthRebaser *EthRebaserSession) MAXSLIPPAGEPARAM() (*big.Int, error) {
	return _EthRebaser.Contract.MAXSLIPPAGEPARAM(&_EthRebaser.CallOpts)
}

// MAXSLIPPAGEPARAM is a free data retrieval call binding the contract method 0x8d76a5bb.
//
// Solidity: function MAX_SLIPPAGE_PARAM() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) MAXSLIPPAGEPARAM() (*big.Int, error) {
	return _EthRebaser.Contract.MAXSLIPPAGEPARAM(&_EthRebaser.CallOpts)
}

// BalGulpPairs is a free data retrieval call binding the contract method 0x1e0cd44e.
//
// Solidity: function balGulpPairs(uint256 ) view returns(address)
func (_EthRebaser *EthRebaserCaller) BalGulpPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "balGulpPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalGulpPairs is a free data retrieval call binding the contract method 0x1e0cd44e.
//
// Solidity: function balGulpPairs(uint256 ) view returns(address)
func (_EthRebaser *EthRebaserSession) BalGulpPairs(arg0 *big.Int) (common.Address, error) {
	return _EthRebaser.Contract.BalGulpPairs(&_EthRebaser.CallOpts, arg0)
}

// BalGulpPairs is a free data retrieval call binding the contract method 0x1e0cd44e.
//
// Solidity: function balGulpPairs(uint256 ) view returns(address)
func (_EthRebaser *EthRebaserCallerSession) BalGulpPairs(arg0 *big.Int) (common.Address, error) {
	return _EthRebaser.Contract.BalGulpPairs(&_EthRebaser.CallOpts, arg0)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_EthRebaser *EthRebaserCaller) BlockTimestampLast(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "blockTimestampLast")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_EthRebaser *EthRebaserSession) BlockTimestampLast() (uint32, error) {
	return _EthRebaser.Contract.BlockTimestampLast(&_EthRebaser.CallOpts)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_EthRebaser *EthRebaserCallerSession) BlockTimestampLast() (uint32, error) {
	return _EthRebaser.Contract.BlockTimestampLast(&_EthRebaser.CallOpts)
}

// DeviationThreshold is a free data retrieval call binding the contract method 0xd94ad837.
//
// Solidity: function deviationThreshold() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) DeviationThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "deviationThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeviationThreshold is a free data retrieval call binding the contract method 0xd94ad837.
//
// Solidity: function deviationThreshold() view returns(uint256)
func (_EthRebaser *EthRebaserSession) DeviationThreshold() (*big.Int, error) {
	return _EthRebaser.Contract.DeviationThreshold(&_EthRebaser.CallOpts)
}

// DeviationThreshold is a free data retrieval call binding the contract method 0xd94ad837.
//
// Solidity: function deviationThreshold() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) DeviationThreshold() (*big.Int, error) {
	return _EthRebaser.Contract.DeviationThreshold(&_EthRebaser.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_EthRebaser *EthRebaserSession) Epoch() (*big.Int, error) {
	return _EthRebaser.Contract.Epoch(&_EthRebaser.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) Epoch() (*big.Int, error) {
	return _EthRebaser.Contract.Epoch(&_EthRebaser.CallOpts)
}

// EthUsdcPair is a free data retrieval call binding the contract method 0x2d3a4af9.
//
// Solidity: function eth_usdc_pair() view returns(address)
func (_EthRebaser *EthRebaserCaller) EthUsdcPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "eth_usdc_pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthUsdcPair is a free data retrieval call binding the contract method 0x2d3a4af9.
//
// Solidity: function eth_usdc_pair() view returns(address)
func (_EthRebaser *EthRebaserSession) EthUsdcPair() (common.Address, error) {
	return _EthRebaser.Contract.EthUsdcPair(&_EthRebaser.CallOpts)
}

// EthUsdcPair is a free data retrieval call binding the contract method 0x2d3a4af9.
//
// Solidity: function eth_usdc_pair() view returns(address)
func (_EthRebaser *EthRebaserCallerSession) EthUsdcPair() (common.Address, error) {
	return _EthRebaser.Contract.EthUsdcPair(&_EthRebaser.CallOpts)
}

// GetBalGulpPairs is a free data retrieval call binding the contract method 0xf12b83b9.
//
// Solidity: function getBalGulpPairs() view returns(address[])
func (_EthRebaser *EthRebaserCaller) GetBalGulpPairs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "getBalGulpPairs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBalGulpPairs is a free data retrieval call binding the contract method 0xf12b83b9.
//
// Solidity: function getBalGulpPairs() view returns(address[])
func (_EthRebaser *EthRebaserSession) GetBalGulpPairs() ([]common.Address, error) {
	return _EthRebaser.Contract.GetBalGulpPairs(&_EthRebaser.CallOpts)
}

// GetBalGulpPairs is a free data retrieval call binding the contract method 0xf12b83b9.
//
// Solidity: function getBalGulpPairs() view returns(address[])
func (_EthRebaser *EthRebaserCallerSession) GetBalGulpPairs() ([]common.Address, error) {
	return _EthRebaser.Contract.GetBalGulpPairs(&_EthRebaser.CallOpts)
}

// GetCurrentTWAP is a free data retrieval call binding the contract method 0x832a3035.
//
// Solidity: function getCurrentTWAP() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) GetCurrentTWAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "getCurrentTWAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTWAP is a free data retrieval call binding the contract method 0x832a3035.
//
// Solidity: function getCurrentTWAP() view returns(uint256)
func (_EthRebaser *EthRebaserSession) GetCurrentTWAP() (*big.Int, error) {
	return _EthRebaser.Contract.GetCurrentTWAP(&_EthRebaser.CallOpts)
}

// GetCurrentTWAP is a free data retrieval call binding the contract method 0x832a3035.
//
// Solidity: function getCurrentTWAP() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) GetCurrentTWAP() (*big.Int, error) {
	return _EthRebaser.Contract.GetCurrentTWAP(&_EthRebaser.CallOpts)
}

// GetUniSyncPairs is a free data retrieval call binding the contract method 0xb532be18.
//
// Solidity: function getUniSyncPairs() view returns(address[])
func (_EthRebaser *EthRebaserCaller) GetUniSyncPairs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "getUniSyncPairs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUniSyncPairs is a free data retrieval call binding the contract method 0xb532be18.
//
// Solidity: function getUniSyncPairs() view returns(address[])
func (_EthRebaser *EthRebaserSession) GetUniSyncPairs() ([]common.Address, error) {
	return _EthRebaser.Contract.GetUniSyncPairs(&_EthRebaser.CallOpts)
}

// GetUniSyncPairs is a free data retrieval call binding the contract method 0xb532be18.
//
// Solidity: function getUniSyncPairs() view returns(address[])
func (_EthRebaser *EthRebaserCallerSession) GetUniSyncPairs() ([]common.Address, error) {
	return _EthRebaser.Contract.GetUniSyncPairs(&_EthRebaser.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_EthRebaser *EthRebaserCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_EthRebaser *EthRebaserSession) Gov() (common.Address, error) {
	return _EthRebaser.Contract.Gov(&_EthRebaser.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_EthRebaser *EthRebaserCallerSession) Gov() (common.Address, error) {
	return _EthRebaser.Contract.Gov(&_EthRebaser.CallOpts)
}

// InRebaseWindow is a free data retrieval call binding the contract method 0x111d0498.
//
// Solidity: function inRebaseWindow() view returns(bool)
func (_EthRebaser *EthRebaserCaller) InRebaseWindow(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "inRebaseWindow")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRebaseWindow is a free data retrieval call binding the contract method 0x111d0498.
//
// Solidity: function inRebaseWindow() view returns(bool)
func (_EthRebaser *EthRebaserSession) InRebaseWindow() (bool, error) {
	return _EthRebaser.Contract.InRebaseWindow(&_EthRebaser.CallOpts)
}

// InRebaseWindow is a free data retrieval call binding the contract method 0x111d0498.
//
// Solidity: function inRebaseWindow() view returns(bool)
func (_EthRebaser *EthRebaserCallerSession) InRebaseWindow() (bool, error) {
	return _EthRebaser.Contract.InRebaseWindow(&_EthRebaser.CallOpts)
}

// IsToken0 is a free data retrieval call binding the contract method 0xdcf93f32.
//
// Solidity: function isToken0() view returns(bool)
func (_EthRebaser *EthRebaserCaller) IsToken0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "isToken0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsToken0 is a free data retrieval call binding the contract method 0xdcf93f32.
//
// Solidity: function isToken0() view returns(bool)
func (_EthRebaser *EthRebaserSession) IsToken0() (bool, error) {
	return _EthRebaser.Contract.IsToken0(&_EthRebaser.CallOpts)
}

// IsToken0 is a free data retrieval call binding the contract method 0xdcf93f32.
//
// Solidity: function isToken0() view returns(bool)
func (_EthRebaser *EthRebaserCallerSession) IsToken0() (bool, error) {
	return _EthRebaser.Contract.IsToken0(&_EthRebaser.CallOpts)
}

// LastRebaseTimestampSec is a free data retrieval call binding the contract method 0x3a93069b.
//
// Solidity: function lastRebaseTimestampSec() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) LastRebaseTimestampSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "lastRebaseTimestampSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRebaseTimestampSec is a free data retrieval call binding the contract method 0x3a93069b.
//
// Solidity: function lastRebaseTimestampSec() view returns(uint256)
func (_EthRebaser *EthRebaserSession) LastRebaseTimestampSec() (*big.Int, error) {
	return _EthRebaser.Contract.LastRebaseTimestampSec(&_EthRebaser.CallOpts)
}

// LastRebaseTimestampSec is a free data retrieval call binding the contract method 0x3a93069b.
//
// Solidity: function lastRebaseTimestampSec() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) LastRebaseTimestampSec() (*big.Int, error) {
	return _EthRebaser.Contract.LastRebaseTimestampSec(&_EthRebaser.CallOpts)
}

// MaxSlippageFactor is a free data retrieval call binding the contract method 0xcf1b927e.
//
// Solidity: function maxSlippageFactor() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) MaxSlippageFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "maxSlippageFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSlippageFactor is a free data retrieval call binding the contract method 0xcf1b927e.
//
// Solidity: function maxSlippageFactor() view returns(uint256)
func (_EthRebaser *EthRebaserSession) MaxSlippageFactor() (*big.Int, error) {
	return _EthRebaser.Contract.MaxSlippageFactor(&_EthRebaser.CallOpts)
}

// MaxSlippageFactor is a free data retrieval call binding the contract method 0xcf1b927e.
//
// Solidity: function maxSlippageFactor() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) MaxSlippageFactor() (*big.Int, error) {
	return _EthRebaser.Contract.MaxSlippageFactor(&_EthRebaser.CallOpts)
}

// MinRebaseTimeIntervalSec is a free data retrieval call binding the contract method 0x02101899.
//
// Solidity: function minRebaseTimeIntervalSec() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) MinRebaseTimeIntervalSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "minRebaseTimeIntervalSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinRebaseTimeIntervalSec is a free data retrieval call binding the contract method 0x02101899.
//
// Solidity: function minRebaseTimeIntervalSec() view returns(uint256)
func (_EthRebaser *EthRebaserSession) MinRebaseTimeIntervalSec() (*big.Int, error) {
	return _EthRebaser.Contract.MinRebaseTimeIntervalSec(&_EthRebaser.CallOpts)
}

// MinRebaseTimeIntervalSec is a free data retrieval call binding the contract method 0x02101899.
//
// Solidity: function minRebaseTimeIntervalSec() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) MinRebaseTimeIntervalSec() (*big.Int, error) {
	return _EthRebaser.Contract.MinRebaseTimeIntervalSec(&_EthRebaser.CallOpts)
}

// PendingGov is a free data retrieval call binding the contract method 0x25240810.
//
// Solidity: function pendingGov() view returns(address)
func (_EthRebaser *EthRebaserCaller) PendingGov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "pendingGov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGov is a free data retrieval call binding the contract method 0x25240810.
//
// Solidity: function pendingGov() view returns(address)
func (_EthRebaser *EthRebaserSession) PendingGov() (common.Address, error) {
	return _EthRebaser.Contract.PendingGov(&_EthRebaser.CallOpts)
}

// PendingGov is a free data retrieval call binding the contract method 0x25240810.
//
// Solidity: function pendingGov() view returns(address)
func (_EthRebaser *EthRebaserCallerSession) PendingGov() (common.Address, error) {
	return _EthRebaser.Contract.PendingGov(&_EthRebaser.CallOpts)
}

// PriceCumulativeLastETHUSDC is a free data retrieval call binding the contract method 0xdcbff61e.
//
// Solidity: function priceCumulativeLastETHUSDC() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) PriceCumulativeLastETHUSDC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "priceCumulativeLastETHUSDC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceCumulativeLastETHUSDC is a free data retrieval call binding the contract method 0xdcbff61e.
//
// Solidity: function priceCumulativeLastETHUSDC() view returns(uint256)
func (_EthRebaser *EthRebaserSession) PriceCumulativeLastETHUSDC() (*big.Int, error) {
	return _EthRebaser.Contract.PriceCumulativeLastETHUSDC(&_EthRebaser.CallOpts)
}

// PriceCumulativeLastETHUSDC is a free data retrieval call binding the contract method 0xdcbff61e.
//
// Solidity: function priceCumulativeLastETHUSDC() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) PriceCumulativeLastETHUSDC() (*big.Int, error) {
	return _EthRebaser.Contract.PriceCumulativeLastETHUSDC(&_EthRebaser.CallOpts)
}

// PriceCumulativeLastYAMETH is a free data retrieval call binding the contract method 0x0978af4f.
//
// Solidity: function priceCumulativeLastYAMETH() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) PriceCumulativeLastYAMETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "priceCumulativeLastYAMETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceCumulativeLastYAMETH is a free data retrieval call binding the contract method 0x0978af4f.
//
// Solidity: function priceCumulativeLastYAMETH() view returns(uint256)
func (_EthRebaser *EthRebaserSession) PriceCumulativeLastYAMETH() (*big.Int, error) {
	return _EthRebaser.Contract.PriceCumulativeLastYAMETH(&_EthRebaser.CallOpts)
}

// PriceCumulativeLastYAMETH is a free data retrieval call binding the contract method 0x0978af4f.
//
// Solidity: function priceCumulativeLastYAMETH() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) PriceCumulativeLastYAMETH() (*big.Int, error) {
	return _EthRebaser.Contract.PriceCumulativeLastYAMETH(&_EthRebaser.CallOpts)
}

// PublicGoods is a free data retrieval call binding the contract method 0xdfebe328.
//
// Solidity: function public_goods() view returns(address)
func (_EthRebaser *EthRebaserCaller) PublicGoods(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "public_goods")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PublicGoods is a free data retrieval call binding the contract method 0xdfebe328.
//
// Solidity: function public_goods() view returns(address)
func (_EthRebaser *EthRebaserSession) PublicGoods() (common.Address, error) {
	return _EthRebaser.Contract.PublicGoods(&_EthRebaser.CallOpts)
}

// PublicGoods is a free data retrieval call binding the contract method 0xdfebe328.
//
// Solidity: function public_goods() view returns(address)
func (_EthRebaser *EthRebaserCallerSession) PublicGoods() (common.Address, error) {
	return _EthRebaser.Contract.PublicGoods(&_EthRebaser.CallOpts)
}

// PublicGoodsPerc is a free data retrieval call binding the contract method 0xfb0ce7d2.
//
// Solidity: function public_goods_perc() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) PublicGoodsPerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "public_goods_perc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicGoodsPerc is a free data retrieval call binding the contract method 0xfb0ce7d2.
//
// Solidity: function public_goods_perc() view returns(uint256)
func (_EthRebaser *EthRebaserSession) PublicGoodsPerc() (*big.Int, error) {
	return _EthRebaser.Contract.PublicGoodsPerc(&_EthRebaser.CallOpts)
}

// PublicGoodsPerc is a free data retrieval call binding the contract method 0xfb0ce7d2.
//
// Solidity: function public_goods_perc() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) PublicGoodsPerc() (*big.Int, error) {
	return _EthRebaser.Contract.PublicGoodsPerc(&_EthRebaser.CallOpts)
}

// RebaseDelay is a free data retrieval call binding the contract method 0x6406ca5f.
//
// Solidity: function rebaseDelay() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) RebaseDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "rebaseDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseDelay is a free data retrieval call binding the contract method 0x6406ca5f.
//
// Solidity: function rebaseDelay() view returns(uint256)
func (_EthRebaser *EthRebaserSession) RebaseDelay() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseDelay(&_EthRebaser.CallOpts)
}

// RebaseDelay is a free data retrieval call binding the contract method 0x6406ca5f.
//
// Solidity: function rebaseDelay() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) RebaseDelay() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseDelay(&_EthRebaser.CallOpts)
}

// RebaseLag is a free data retrieval call binding the contract method 0x63f6d4c8.
//
// Solidity: function rebaseLag() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) RebaseLag(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "rebaseLag")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseLag is a free data retrieval call binding the contract method 0x63f6d4c8.
//
// Solidity: function rebaseLag() view returns(uint256)
func (_EthRebaser *EthRebaserSession) RebaseLag() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseLag(&_EthRebaser.CallOpts)
}

// RebaseLag is a free data retrieval call binding the contract method 0x63f6d4c8.
//
// Solidity: function rebaseLag() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) RebaseLag() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseLag(&_EthRebaser.CallOpts)
}

// RebaseMintPerc is a free data retrieval call binding the contract method 0x1cab801c.
//
// Solidity: function rebaseMintPerc() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) RebaseMintPerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "rebaseMintPerc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseMintPerc is a free data retrieval call binding the contract method 0x1cab801c.
//
// Solidity: function rebaseMintPerc() view returns(uint256)
func (_EthRebaser *EthRebaserSession) RebaseMintPerc() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseMintPerc(&_EthRebaser.CallOpts)
}

// RebaseMintPerc is a free data retrieval call binding the contract method 0x1cab801c.
//
// Solidity: function rebaseMintPerc() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) RebaseMintPerc() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseMintPerc(&_EthRebaser.CallOpts)
}

// RebaseWindowLengthSec is a free data retrieval call binding the contract method 0x9466120f.
//
// Solidity: function rebaseWindowLengthSec() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) RebaseWindowLengthSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "rebaseWindowLengthSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseWindowLengthSec is a free data retrieval call binding the contract method 0x9466120f.
//
// Solidity: function rebaseWindowLengthSec() view returns(uint256)
func (_EthRebaser *EthRebaserSession) RebaseWindowLengthSec() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseWindowLengthSec(&_EthRebaser.CallOpts)
}

// RebaseWindowLengthSec is a free data retrieval call binding the contract method 0x9466120f.
//
// Solidity: function rebaseWindowLengthSec() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) RebaseWindowLengthSec() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseWindowLengthSec(&_EthRebaser.CallOpts)
}

// RebaseWindowOffsetSec is a free data retrieval call binding the contract method 0x7052b902.
//
// Solidity: function rebaseWindowOffsetSec() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) RebaseWindowOffsetSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "rebaseWindowOffsetSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RebaseWindowOffsetSec is a free data retrieval call binding the contract method 0x7052b902.
//
// Solidity: function rebaseWindowOffsetSec() view returns(uint256)
func (_EthRebaser *EthRebaserSession) RebaseWindowOffsetSec() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseWindowOffsetSec(&_EthRebaser.CallOpts)
}

// RebaseWindowOffsetSec is a free data retrieval call binding the contract method 0x7052b902.
//
// Solidity: function rebaseWindowOffsetSec() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) RebaseWindowOffsetSec() (*big.Int, error) {
	return _EthRebaser.Contract.RebaseWindowOffsetSec(&_EthRebaser.CallOpts)
}

// RebasingActive is a free data retrieval call binding the contract method 0x4dc95de1.
//
// Solidity: function rebasingActive() view returns(bool)
func (_EthRebaser *EthRebaserCaller) RebasingActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "rebasingActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RebasingActive is a free data retrieval call binding the contract method 0x4dc95de1.
//
// Solidity: function rebasingActive() view returns(bool)
func (_EthRebaser *EthRebaserSession) RebasingActive() (bool, error) {
	return _EthRebaser.Contract.RebasingActive(&_EthRebaser.CallOpts)
}

// RebasingActive is a free data retrieval call binding the contract method 0x4dc95de1.
//
// Solidity: function rebasingActive() view returns(bool)
func (_EthRebaser *EthRebaserCallerSession) RebasingActive() (bool, error) {
	return _EthRebaser.Contract.RebasingActive(&_EthRebaser.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_EthRebaser *EthRebaserCaller) ReserveToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "reserveToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_EthRebaser *EthRebaserSession) ReserveToken() (common.Address, error) {
	return _EthRebaser.Contract.ReserveToken(&_EthRebaser.CallOpts)
}

// ReserveToken is a free data retrieval call binding the contract method 0xf4325d67.
//
// Solidity: function reserveToken() view returns(address)
func (_EthRebaser *EthRebaserCallerSession) ReserveToken() (common.Address, error) {
	return _EthRebaser.Contract.ReserveToken(&_EthRebaser.CallOpts)
}

// ReservesContract is a free data retrieval call binding the contract method 0x3a68eaf6.
//
// Solidity: function reservesContract() view returns(address)
func (_EthRebaser *EthRebaserCaller) ReservesContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "reservesContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReservesContract is a free data retrieval call binding the contract method 0x3a68eaf6.
//
// Solidity: function reservesContract() view returns(address)
func (_EthRebaser *EthRebaserSession) ReservesContract() (common.Address, error) {
	return _EthRebaser.Contract.ReservesContract(&_EthRebaser.CallOpts)
}

// ReservesContract is a free data retrieval call binding the contract method 0x3a68eaf6.
//
// Solidity: function reservesContract() view returns(address)
func (_EthRebaser *EthRebaserCallerSession) ReservesContract() (common.Address, error) {
	return _EthRebaser.Contract.ReservesContract(&_EthRebaser.CallOpts)
}

// TargetRate is a free data retrieval call binding the contract method 0xcc8fd393.
//
// Solidity: function targetRate() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) TargetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "targetRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetRate is a free data retrieval call binding the contract method 0xcc8fd393.
//
// Solidity: function targetRate() view returns(uint256)
func (_EthRebaser *EthRebaserSession) TargetRate() (*big.Int, error) {
	return _EthRebaser.Contract.TargetRate(&_EthRebaser.CallOpts)
}

// TargetRate is a free data retrieval call binding the contract method 0xcc8fd393.
//
// Solidity: function targetRate() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) TargetRate() (*big.Int, error) {
	return _EthRebaser.Contract.TargetRate(&_EthRebaser.CallOpts)
}

// TimeOfTWAPInit is a free data retrieval call binding the contract method 0xcd877826.
//
// Solidity: function timeOfTWAPInit() view returns(uint256)
func (_EthRebaser *EthRebaserCaller) TimeOfTWAPInit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "timeOfTWAPInit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeOfTWAPInit is a free data retrieval call binding the contract method 0xcd877826.
//
// Solidity: function timeOfTWAPInit() view returns(uint256)
func (_EthRebaser *EthRebaserSession) TimeOfTWAPInit() (*big.Int, error) {
	return _EthRebaser.Contract.TimeOfTWAPInit(&_EthRebaser.CallOpts)
}

// TimeOfTWAPInit is a free data retrieval call binding the contract method 0xcd877826.
//
// Solidity: function timeOfTWAPInit() view returns(uint256)
func (_EthRebaser *EthRebaserCallerSession) TimeOfTWAPInit() (*big.Int, error) {
	return _EthRebaser.Contract.TimeOfTWAPInit(&_EthRebaser.CallOpts)
}

// TradePair is a free data retrieval call binding the contract method 0x0dd2c45e.
//
// Solidity: function trade_pair() view returns(address)
func (_EthRebaser *EthRebaserCaller) TradePair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "trade_pair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradePair is a free data retrieval call binding the contract method 0x0dd2c45e.
//
// Solidity: function trade_pair() view returns(address)
func (_EthRebaser *EthRebaserSession) TradePair() (common.Address, error) {
	return _EthRebaser.Contract.TradePair(&_EthRebaser.CallOpts)
}

// TradePair is a free data retrieval call binding the contract method 0x0dd2c45e.
//
// Solidity: function trade_pair() view returns(address)
func (_EthRebaser *EthRebaserCallerSession) TradePair() (common.Address, error) {
	return _EthRebaser.Contract.TradePair(&_EthRebaser.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bool enabled, address destination, bytes data)
func (_EthRebaser *EthRebaserCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Enabled     bool
	Destination common.Address
	Data        []byte
}, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Enabled     bool
		Destination common.Address
		Data        []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Enabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Destination = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bool enabled, address destination, bytes data)
func (_EthRebaser *EthRebaserSession) Transactions(arg0 *big.Int) (struct {
	Enabled     bool
	Destination common.Address
	Data        []byte
}, error) {
	return _EthRebaser.Contract.Transactions(&_EthRebaser.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(bool enabled, address destination, bytes data)
func (_EthRebaser *EthRebaserCallerSession) Transactions(arg0 *big.Int) (struct {
	Enabled     bool
	Destination common.Address
	Data        []byte
}, error) {
	return _EthRebaser.Contract.Transactions(&_EthRebaser.CallOpts, arg0)
}

// UniSyncPairs is a free data retrieval call binding the contract method 0x14eb3f24.
//
// Solidity: function uniSyncPairs(uint256 ) view returns(address)
func (_EthRebaser *EthRebaserCaller) UniSyncPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "uniSyncPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniSyncPairs is a free data retrieval call binding the contract method 0x14eb3f24.
//
// Solidity: function uniSyncPairs(uint256 ) view returns(address)
func (_EthRebaser *EthRebaserSession) UniSyncPairs(arg0 *big.Int) (common.Address, error) {
	return _EthRebaser.Contract.UniSyncPairs(&_EthRebaser.CallOpts, arg0)
}

// UniSyncPairs is a free data retrieval call binding the contract method 0x14eb3f24.
//
// Solidity: function uniSyncPairs(uint256 ) view returns(address)
func (_EthRebaser *EthRebaserCallerSession) UniSyncPairs(arg0 *big.Int) (common.Address, error) {
	return _EthRebaser.Contract.UniSyncPairs(&_EthRebaser.CallOpts, arg0)
}

// YamAddress is a free data retrieval call binding the contract method 0xd72cdafc.
//
// Solidity: function yamAddress() view returns(address)
func (_EthRebaser *EthRebaserCaller) YamAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthRebaser.contract.Call(opts, &out, "yamAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YamAddress is a free data retrieval call binding the contract method 0xd72cdafc.
//
// Solidity: function yamAddress() view returns(address)
func (_EthRebaser *EthRebaserSession) YamAddress() (common.Address, error) {
	return _EthRebaser.Contract.YamAddress(&_EthRebaser.CallOpts)
}

// YamAddress is a free data retrieval call binding the contract method 0xd72cdafc.
//
// Solidity: function yamAddress() view returns(address)
func (_EthRebaser *EthRebaserCallerSession) YamAddress() (common.Address, error) {
	return _EthRebaser.Contract.YamAddress(&_EthRebaser.CallOpts)
}

// AcceptGov is a paid mutator transaction binding the contract method 0x4bda2e20.
//
// Solidity: function _acceptGov() returns()
func (_EthRebaser *EthRebaserTransactor) AcceptGov(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "_acceptGov")
}

// AcceptGov is a paid mutator transaction binding the contract method 0x4bda2e20.
//
// Solidity: function _acceptGov() returns()
func (_EthRebaser *EthRebaserSession) AcceptGov() (*types.Transaction, error) {
	return _EthRebaser.Contract.AcceptGov(&_EthRebaser.TransactOpts)
}

// AcceptGov is a paid mutator transaction binding the contract method 0x4bda2e20.
//
// Solidity: function _acceptGov() returns()
func (_EthRebaser *EthRebaserTransactorSession) AcceptGov() (*types.Transaction, error) {
	return _EthRebaser.Contract.AcceptGov(&_EthRebaser.TransactOpts)
}

// SetPendingGov is a paid mutator transaction binding the contract method 0x73f03dff.
//
// Solidity: function _setPendingGov(address pendingGov_) returns()
func (_EthRebaser *EthRebaserTransactor) SetPendingGov(opts *bind.TransactOpts, pendingGov_ common.Address) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "_setPendingGov", pendingGov_)
}

// SetPendingGov is a paid mutator transaction binding the contract method 0x73f03dff.
//
// Solidity: function _setPendingGov(address pendingGov_) returns()
func (_EthRebaser *EthRebaserSession) SetPendingGov(pendingGov_ common.Address) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetPendingGov(&_EthRebaser.TransactOpts, pendingGov_)
}

// SetPendingGov is a paid mutator transaction binding the contract method 0x73f03dff.
//
// Solidity: function _setPendingGov(address pendingGov_) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetPendingGov(pendingGov_ common.Address) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetPendingGov(&_EthRebaser.TransactOpts, pendingGov_)
}

// ActivateRebasing is a paid mutator transaction binding the contract method 0x4e66f8ae.
//
// Solidity: function activate_rebasing() returns()
func (_EthRebaser *EthRebaserTransactor) ActivateRebasing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "activate_rebasing")
}

// ActivateRebasing is a paid mutator transaction binding the contract method 0x4e66f8ae.
//
// Solidity: function activate_rebasing() returns()
func (_EthRebaser *EthRebaserSession) ActivateRebasing() (*types.Transaction, error) {
	return _EthRebaser.Contract.ActivateRebasing(&_EthRebaser.TransactOpts)
}

// ActivateRebasing is a paid mutator transaction binding the contract method 0x4e66f8ae.
//
// Solidity: function activate_rebasing() returns()
func (_EthRebaser *EthRebaserTransactorSession) ActivateRebasing() (*types.Transaction, error) {
	return _EthRebaser.Contract.ActivateRebasing(&_EthRebaser.TransactOpts)
}

// AddSyncPairs is a paid mutator transaction binding the contract method 0x9671ecff.
//
// Solidity: function addSyncPairs(address[] uniSyncPairs_, address[] balGulpPairs_) returns()
func (_EthRebaser *EthRebaserTransactor) AddSyncPairs(opts *bind.TransactOpts, uniSyncPairs_ []common.Address, balGulpPairs_ []common.Address) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "addSyncPairs", uniSyncPairs_, balGulpPairs_)
}

// AddSyncPairs is a paid mutator transaction binding the contract method 0x9671ecff.
//
// Solidity: function addSyncPairs(address[] uniSyncPairs_, address[] balGulpPairs_) returns()
func (_EthRebaser *EthRebaserSession) AddSyncPairs(uniSyncPairs_ []common.Address, balGulpPairs_ []common.Address) (*types.Transaction, error) {
	return _EthRebaser.Contract.AddSyncPairs(&_EthRebaser.TransactOpts, uniSyncPairs_, balGulpPairs_)
}

// AddSyncPairs is a paid mutator transaction binding the contract method 0x9671ecff.
//
// Solidity: function addSyncPairs(address[] uniSyncPairs_, address[] balGulpPairs_) returns()
func (_EthRebaser *EthRebaserTransactorSession) AddSyncPairs(uniSyncPairs_ []common.Address, balGulpPairs_ []common.Address) (*types.Transaction, error) {
	return _EthRebaser.Contract.AddSyncPairs(&_EthRebaser.TransactOpts, uniSyncPairs_, balGulpPairs_)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x126e19be.
//
// Solidity: function addTransaction(address destination, bytes data) returns()
func (_EthRebaser *EthRebaserTransactor) AddTransaction(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "addTransaction", destination, data)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x126e19be.
//
// Solidity: function addTransaction(address destination, bytes data) returns()
func (_EthRebaser *EthRebaserSession) AddTransaction(destination common.Address, data []byte) (*types.Transaction, error) {
	return _EthRebaser.Contract.AddTransaction(&_EthRebaser.TransactOpts, destination, data)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x126e19be.
//
// Solidity: function addTransaction(address destination, bytes data) returns()
func (_EthRebaser *EthRebaserTransactorSession) AddTransaction(destination common.Address, data []byte) (*types.Transaction, error) {
	return _EthRebaser.Contract.AddTransaction(&_EthRebaser.TransactOpts, destination, data)
}

// InitTwap is a paid mutator transaction binding the contract method 0x57466c8b.
//
// Solidity: function init_twap() returns()
func (_EthRebaser *EthRebaserTransactor) InitTwap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "init_twap")
}

// InitTwap is a paid mutator transaction binding the contract method 0x57466c8b.
//
// Solidity: function init_twap() returns()
func (_EthRebaser *EthRebaserSession) InitTwap() (*types.Transaction, error) {
	return _EthRebaser.Contract.InitTwap(&_EthRebaser.TransactOpts)
}

// InitTwap is a paid mutator transaction binding the contract method 0x57466c8b.
//
// Solidity: function init_twap() returns()
func (_EthRebaser *EthRebaserTransactorSession) InitTwap() (*types.Transaction, error) {
	return _EthRebaser.Contract.InitTwap(&_EthRebaser.TransactOpts)
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns()
func (_EthRebaser *EthRebaserTransactor) Rebase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "rebase")
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns()
func (_EthRebaser *EthRebaserSession) Rebase() (*types.Transaction, error) {
	return _EthRebaser.Contract.Rebase(&_EthRebaser.TransactOpts)
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns()
func (_EthRebaser *EthRebaserTransactorSession) Rebase() (*types.Transaction, error) {
	return _EthRebaser.Contract.Rebase(&_EthRebaser.TransactOpts)
}

// RemoveBalPair is a paid mutator transaction binding the contract method 0x1b58ac4a.
//
// Solidity: function removeBalPair(uint256 index) returns()
func (_EthRebaser *EthRebaserTransactor) RemoveBalPair(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "removeBalPair", index)
}

// RemoveBalPair is a paid mutator transaction binding the contract method 0x1b58ac4a.
//
// Solidity: function removeBalPair(uint256 index) returns()
func (_EthRebaser *EthRebaserSession) RemoveBalPair(index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.RemoveBalPair(&_EthRebaser.TransactOpts, index)
}

// RemoveBalPair is a paid mutator transaction binding the contract method 0x1b58ac4a.
//
// Solidity: function removeBalPair(uint256 index) returns()
func (_EthRebaser *EthRebaserTransactorSession) RemoveBalPair(index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.RemoveBalPair(&_EthRebaser.TransactOpts, index)
}

// RemoveTransaction is a paid mutator transaction binding the contract method 0x46c3bd1f.
//
// Solidity: function removeTransaction(uint256 index) returns()
func (_EthRebaser *EthRebaserTransactor) RemoveTransaction(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "removeTransaction", index)
}

// RemoveTransaction is a paid mutator transaction binding the contract method 0x46c3bd1f.
//
// Solidity: function removeTransaction(uint256 index) returns()
func (_EthRebaser *EthRebaserSession) RemoveTransaction(index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.RemoveTransaction(&_EthRebaser.TransactOpts, index)
}

// RemoveTransaction is a paid mutator transaction binding the contract method 0x46c3bd1f.
//
// Solidity: function removeTransaction(uint256 index) returns()
func (_EthRebaser *EthRebaserTransactorSession) RemoveTransaction(index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.RemoveTransaction(&_EthRebaser.TransactOpts, index)
}

// RemoveUniPair is a paid mutator transaction binding the contract method 0x4f2b9629.
//
// Solidity: function removeUniPair(uint256 index) returns()
func (_EthRebaser *EthRebaserTransactor) RemoveUniPair(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "removeUniPair", index)
}

// RemoveUniPair is a paid mutator transaction binding the contract method 0x4f2b9629.
//
// Solidity: function removeUniPair(uint256 index) returns()
func (_EthRebaser *EthRebaserSession) RemoveUniPair(index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.RemoveUniPair(&_EthRebaser.TransactOpts, index)
}

// RemoveUniPair is a paid mutator transaction binding the contract method 0x4f2b9629.
//
// Solidity: function removeUniPair(uint256 index) returns()
func (_EthRebaser *EthRebaserTransactorSession) RemoveUniPair(index *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.RemoveUniPair(&_EthRebaser.TransactOpts, index)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns(bool)
func (_EthRebaser *EthRebaserTransactor) RescueTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "rescueTokens", token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns(bool)
func (_EthRebaser *EthRebaserSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.RescueTokens(&_EthRebaser.TransactOpts, token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns(bool)
func (_EthRebaser *EthRebaserTransactorSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.RescueTokens(&_EthRebaser.TransactOpts, token, to, amount)
}

// SetDeviationThreshold is a paid mutator transaction binding the contract method 0x53a15edc.
//
// Solidity: function setDeviationThreshold(uint256 deviationThreshold_) returns()
func (_EthRebaser *EthRebaserTransactor) SetDeviationThreshold(opts *bind.TransactOpts, deviationThreshold_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "setDeviationThreshold", deviationThreshold_)
}

// SetDeviationThreshold is a paid mutator transaction binding the contract method 0x53a15edc.
//
// Solidity: function setDeviationThreshold(uint256 deviationThreshold_) returns()
func (_EthRebaser *EthRebaserSession) SetDeviationThreshold(deviationThreshold_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetDeviationThreshold(&_EthRebaser.TransactOpts, deviationThreshold_)
}

// SetDeviationThreshold is a paid mutator transaction binding the contract method 0x53a15edc.
//
// Solidity: function setDeviationThreshold(uint256 deviationThreshold_) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetDeviationThreshold(deviationThreshold_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetDeviationThreshold(&_EthRebaser.TransactOpts, deviationThreshold_)
}

// SetMaxSlippageFactor is a paid mutator transaction binding the contract method 0xb60e1e3e.
//
// Solidity: function setMaxSlippageFactor(uint256 maxSlippageFactor_) returns()
func (_EthRebaser *EthRebaserTransactor) SetMaxSlippageFactor(opts *bind.TransactOpts, maxSlippageFactor_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "setMaxSlippageFactor", maxSlippageFactor_)
}

// SetMaxSlippageFactor is a paid mutator transaction binding the contract method 0xb60e1e3e.
//
// Solidity: function setMaxSlippageFactor(uint256 maxSlippageFactor_) returns()
func (_EthRebaser *EthRebaserSession) SetMaxSlippageFactor(maxSlippageFactor_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetMaxSlippageFactor(&_EthRebaser.TransactOpts, maxSlippageFactor_)
}

// SetMaxSlippageFactor is a paid mutator transaction binding the contract method 0xb60e1e3e.
//
// Solidity: function setMaxSlippageFactor(uint256 maxSlippageFactor_) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetMaxSlippageFactor(maxSlippageFactor_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetMaxSlippageFactor(&_EthRebaser.TransactOpts, maxSlippageFactor_)
}

// SetRebaseLag is a paid mutator transaction binding the contract method 0x20ce8389.
//
// Solidity: function setRebaseLag(uint256 rebaseLag_) returns()
func (_EthRebaser *EthRebaserTransactor) SetRebaseLag(opts *bind.TransactOpts, rebaseLag_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "setRebaseLag", rebaseLag_)
}

// SetRebaseLag is a paid mutator transaction binding the contract method 0x20ce8389.
//
// Solidity: function setRebaseLag(uint256 rebaseLag_) returns()
func (_EthRebaser *EthRebaserSession) SetRebaseLag(rebaseLag_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetRebaseLag(&_EthRebaser.TransactOpts, rebaseLag_)
}

// SetRebaseLag is a paid mutator transaction binding the contract method 0x20ce8389.
//
// Solidity: function setRebaseLag(uint256 rebaseLag_) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetRebaseLag(rebaseLag_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetRebaseLag(&_EthRebaser.TransactOpts, rebaseLag_)
}

// SetRebaseMintPerc is a paid mutator transaction binding the contract method 0x6bf9ace7.
//
// Solidity: function setRebaseMintPerc(uint256 rebaseMintPerc_) returns()
func (_EthRebaser *EthRebaserTransactor) SetRebaseMintPerc(opts *bind.TransactOpts, rebaseMintPerc_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "setRebaseMintPerc", rebaseMintPerc_)
}

// SetRebaseMintPerc is a paid mutator transaction binding the contract method 0x6bf9ace7.
//
// Solidity: function setRebaseMintPerc(uint256 rebaseMintPerc_) returns()
func (_EthRebaser *EthRebaserSession) SetRebaseMintPerc(rebaseMintPerc_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetRebaseMintPerc(&_EthRebaser.TransactOpts, rebaseMintPerc_)
}

// SetRebaseMintPerc is a paid mutator transaction binding the contract method 0x6bf9ace7.
//
// Solidity: function setRebaseMintPerc(uint256 rebaseMintPerc_) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetRebaseMintPerc(rebaseMintPerc_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetRebaseMintPerc(&_EthRebaser.TransactOpts, rebaseMintPerc_)
}

// SetRebaseTimingParameters is a paid mutator transaction binding the contract method 0x16250fd4.
//
// Solidity: function setRebaseTimingParameters(uint256 minRebaseTimeIntervalSec_, uint256 rebaseWindowOffsetSec_, uint256 rebaseWindowLengthSec_) returns()
func (_EthRebaser *EthRebaserTransactor) SetRebaseTimingParameters(opts *bind.TransactOpts, minRebaseTimeIntervalSec_ *big.Int, rebaseWindowOffsetSec_ *big.Int, rebaseWindowLengthSec_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "setRebaseTimingParameters", minRebaseTimeIntervalSec_, rebaseWindowOffsetSec_, rebaseWindowLengthSec_)
}

// SetRebaseTimingParameters is a paid mutator transaction binding the contract method 0x16250fd4.
//
// Solidity: function setRebaseTimingParameters(uint256 minRebaseTimeIntervalSec_, uint256 rebaseWindowOffsetSec_, uint256 rebaseWindowLengthSec_) returns()
func (_EthRebaser *EthRebaserSession) SetRebaseTimingParameters(minRebaseTimeIntervalSec_ *big.Int, rebaseWindowOffsetSec_ *big.Int, rebaseWindowLengthSec_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetRebaseTimingParameters(&_EthRebaser.TransactOpts, minRebaseTimeIntervalSec_, rebaseWindowOffsetSec_, rebaseWindowLengthSec_)
}

// SetRebaseTimingParameters is a paid mutator transaction binding the contract method 0x16250fd4.
//
// Solidity: function setRebaseTimingParameters(uint256 minRebaseTimeIntervalSec_, uint256 rebaseWindowOffsetSec_, uint256 rebaseWindowLengthSec_) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetRebaseTimingParameters(minRebaseTimeIntervalSec_ *big.Int, rebaseWindowOffsetSec_ *big.Int, rebaseWindowLengthSec_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetRebaseTimingParameters(&_EthRebaser.TransactOpts, minRebaseTimeIntervalSec_, rebaseWindowOffsetSec_, rebaseWindowLengthSec_)
}

// SetReserveContract is a paid mutator transaction binding the contract method 0x527a52c8.
//
// Solidity: function setReserveContract(address reservesContract_) returns()
func (_EthRebaser *EthRebaserTransactor) SetReserveContract(opts *bind.TransactOpts, reservesContract_ common.Address) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "setReserveContract", reservesContract_)
}

// SetReserveContract is a paid mutator transaction binding the contract method 0x527a52c8.
//
// Solidity: function setReserveContract(address reservesContract_) returns()
func (_EthRebaser *EthRebaserSession) SetReserveContract(reservesContract_ common.Address) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetReserveContract(&_EthRebaser.TransactOpts, reservesContract_)
}

// SetReserveContract is a paid mutator transaction binding the contract method 0x527a52c8.
//
// Solidity: function setReserveContract(address reservesContract_) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetReserveContract(reservesContract_ common.Address) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetReserveContract(&_EthRebaser.TransactOpts, reservesContract_)
}

// SetTargetRate is a paid mutator transaction binding the contract method 0xcdabdaac.
//
// Solidity: function setTargetRate(uint256 targetRate_) returns()
func (_EthRebaser *EthRebaserTransactor) SetTargetRate(opts *bind.TransactOpts, targetRate_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "setTargetRate", targetRate_)
}

// SetTargetRate is a paid mutator transaction binding the contract method 0xcdabdaac.
//
// Solidity: function setTargetRate(uint256 targetRate_) returns()
func (_EthRebaser *EthRebaserSession) SetTargetRate(targetRate_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetTargetRate(&_EthRebaser.TransactOpts, targetRate_)
}

// SetTargetRate is a paid mutator transaction binding the contract method 0xcdabdaac.
//
// Solidity: function setTargetRate(uint256 targetRate_) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetTargetRate(targetRate_ *big.Int) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetTargetRate(&_EthRebaser.TransactOpts, targetRate_)
}

// SetTransactionEnabled is a paid mutator transaction binding the contract method 0x6e9dde99.
//
// Solidity: function setTransactionEnabled(uint256 index, bool enabled) returns()
func (_EthRebaser *EthRebaserTransactor) SetTransactionEnabled(opts *bind.TransactOpts, index *big.Int, enabled bool) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "setTransactionEnabled", index, enabled)
}

// SetTransactionEnabled is a paid mutator transaction binding the contract method 0x6e9dde99.
//
// Solidity: function setTransactionEnabled(uint256 index, bool enabled) returns()
func (_EthRebaser *EthRebaserSession) SetTransactionEnabled(index *big.Int, enabled bool) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetTransactionEnabled(&_EthRebaser.TransactOpts, index, enabled)
}

// SetTransactionEnabled is a paid mutator transaction binding the contract method 0x6e9dde99.
//
// Solidity: function setTransactionEnabled(uint256 index, bool enabled) returns()
func (_EthRebaser *EthRebaserTransactorSession) SetTransactionEnabled(index *big.Int, enabled bool) (*types.Transaction, error) {
	return _EthRebaser.Contract.SetTransactionEnabled(&_EthRebaser.TransactOpts, index, enabled)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_EthRebaser *EthRebaserTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _EthRebaser.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_EthRebaser *EthRebaserSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _EthRebaser.Contract.UniswapV2Call(&_EthRebaser.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_EthRebaser *EthRebaserTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _EthRebaser.Contract.UniswapV2Call(&_EthRebaser.TransactOpts, sender, amount0, amount1, data)
}

// EthRebaserMintAmountIterator is returned from FilterMintAmount and is used to iterate over the raw logs and unpacked data for MintAmount events raised by the EthRebaser contract.
type EthRebaserMintAmountIterator struct {
	Event *EthRebaserMintAmount // Event containing the contract specifics and raw log

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
func (it *EthRebaserMintAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserMintAmount)
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
		it.Event = new(EthRebaserMintAmount)
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
func (it *EthRebaserMintAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserMintAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserMintAmount represents a MintAmount event raised by the EthRebaser contract.
type EthRebaserMintAmount struct {
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintAmount is a free log retrieval operation binding the contract event 0x8e37eb2ee3a6f3c8b13b8973588daad75a4ce752de14c00006bd8247f4e212e8.
//
// Solidity: event MintAmount(uint256 mintAmount)
func (_EthRebaser *EthRebaserFilterer) FilterMintAmount(opts *bind.FilterOpts) (*EthRebaserMintAmountIterator, error) {

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "MintAmount")
	if err != nil {
		return nil, err
	}
	return &EthRebaserMintAmountIterator{contract: _EthRebaser.contract, event: "MintAmount", logs: logs, sub: sub}, nil
}

// WatchMintAmount is a free log subscription operation binding the contract event 0x8e37eb2ee3a6f3c8b13b8973588daad75a4ce752de14c00006bd8247f4e212e8.
//
// Solidity: event MintAmount(uint256 mintAmount)
func (_EthRebaser *EthRebaserFilterer) WatchMintAmount(opts *bind.WatchOpts, sink chan<- *EthRebaserMintAmount) (event.Subscription, error) {

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "MintAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserMintAmount)
				if err := _EthRebaser.contract.UnpackLog(event, "MintAmount", log); err != nil {
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

// ParseMintAmount is a log parse operation binding the contract event 0x8e37eb2ee3a6f3c8b13b8973588daad75a4ce752de14c00006bd8247f4e212e8.
//
// Solidity: event MintAmount(uint256 mintAmount)
func (_EthRebaser *EthRebaserFilterer) ParseMintAmount(log types.Log) (*EthRebaserMintAmount, error) {
	event := new(EthRebaserMintAmount)
	if err := _EthRebaser.contract.UnpackLog(event, "MintAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthRebaserNewDeviationThresholdIterator is returned from FilterNewDeviationThreshold and is used to iterate over the raw logs and unpacked data for NewDeviationThreshold events raised by the EthRebaser contract.
type EthRebaserNewDeviationThresholdIterator struct {
	Event *EthRebaserNewDeviationThreshold // Event containing the contract specifics and raw log

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
func (it *EthRebaserNewDeviationThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserNewDeviationThreshold)
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
		it.Event = new(EthRebaserNewDeviationThreshold)
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
func (it *EthRebaserNewDeviationThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserNewDeviationThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserNewDeviationThreshold represents a NewDeviationThreshold event raised by the EthRebaser contract.
type EthRebaserNewDeviationThreshold struct {
	OldDeviationThreshold *big.Int
	NewDeviationThreshold *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewDeviationThreshold is a free log retrieval operation binding the contract event 0x2a5cda4d16fba415b52d90b59ee30d4cb16494da9fd1ee51c4d5bac4a1f75bbe.
//
// Solidity: event NewDeviationThreshold(uint256 oldDeviationThreshold, uint256 newDeviationThreshold)
func (_EthRebaser *EthRebaserFilterer) FilterNewDeviationThreshold(opts *bind.FilterOpts) (*EthRebaserNewDeviationThresholdIterator, error) {

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "NewDeviationThreshold")
	if err != nil {
		return nil, err
	}
	return &EthRebaserNewDeviationThresholdIterator{contract: _EthRebaser.contract, event: "NewDeviationThreshold", logs: logs, sub: sub}, nil
}

// WatchNewDeviationThreshold is a free log subscription operation binding the contract event 0x2a5cda4d16fba415b52d90b59ee30d4cb16494da9fd1ee51c4d5bac4a1f75bbe.
//
// Solidity: event NewDeviationThreshold(uint256 oldDeviationThreshold, uint256 newDeviationThreshold)
func (_EthRebaser *EthRebaserFilterer) WatchNewDeviationThreshold(opts *bind.WatchOpts, sink chan<- *EthRebaserNewDeviationThreshold) (event.Subscription, error) {

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "NewDeviationThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserNewDeviationThreshold)
				if err := _EthRebaser.contract.UnpackLog(event, "NewDeviationThreshold", log); err != nil {
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

// ParseNewDeviationThreshold is a log parse operation binding the contract event 0x2a5cda4d16fba415b52d90b59ee30d4cb16494da9fd1ee51c4d5bac4a1f75bbe.
//
// Solidity: event NewDeviationThreshold(uint256 oldDeviationThreshold, uint256 newDeviationThreshold)
func (_EthRebaser *EthRebaserFilterer) ParseNewDeviationThreshold(log types.Log) (*EthRebaserNewDeviationThreshold, error) {
	event := new(EthRebaserNewDeviationThreshold)
	if err := _EthRebaser.contract.UnpackLog(event, "NewDeviationThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthRebaserNewGovIterator is returned from FilterNewGov and is used to iterate over the raw logs and unpacked data for NewGov events raised by the EthRebaser contract.
type EthRebaserNewGovIterator struct {
	Event *EthRebaserNewGov // Event containing the contract specifics and raw log

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
func (it *EthRebaserNewGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserNewGov)
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
		it.Event = new(EthRebaserNewGov)
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
func (it *EthRebaserNewGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserNewGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserNewGov represents a NewGov event raised by the EthRebaser contract.
type EthRebaserNewGov struct {
	OldGov common.Address
	NewGov common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewGov is a free log retrieval operation binding the contract event 0x1f14cfc03e486d23acee577b07bc0b3b23f4888c91fcdba5e0fef5a2549d5523.
//
// Solidity: event NewGov(address oldGov, address newGov)
func (_EthRebaser *EthRebaserFilterer) FilterNewGov(opts *bind.FilterOpts) (*EthRebaserNewGovIterator, error) {

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "NewGov")
	if err != nil {
		return nil, err
	}
	return &EthRebaserNewGovIterator{contract: _EthRebaser.contract, event: "NewGov", logs: logs, sub: sub}, nil
}

// WatchNewGov is a free log subscription operation binding the contract event 0x1f14cfc03e486d23acee577b07bc0b3b23f4888c91fcdba5e0fef5a2549d5523.
//
// Solidity: event NewGov(address oldGov, address newGov)
func (_EthRebaser *EthRebaserFilterer) WatchNewGov(opts *bind.WatchOpts, sink chan<- *EthRebaserNewGov) (event.Subscription, error) {

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "NewGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserNewGov)
				if err := _EthRebaser.contract.UnpackLog(event, "NewGov", log); err != nil {
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

// ParseNewGov is a log parse operation binding the contract event 0x1f14cfc03e486d23acee577b07bc0b3b23f4888c91fcdba5e0fef5a2549d5523.
//
// Solidity: event NewGov(address oldGov, address newGov)
func (_EthRebaser *EthRebaserFilterer) ParseNewGov(log types.Log) (*EthRebaserNewGov, error) {
	event := new(EthRebaserNewGov)
	if err := _EthRebaser.contract.UnpackLog(event, "NewGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthRebaserNewMaxSlippageFactorIterator is returned from FilterNewMaxSlippageFactor and is used to iterate over the raw logs and unpacked data for NewMaxSlippageFactor events raised by the EthRebaser contract.
type EthRebaserNewMaxSlippageFactorIterator struct {
	Event *EthRebaserNewMaxSlippageFactor // Event containing the contract specifics and raw log

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
func (it *EthRebaserNewMaxSlippageFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserNewMaxSlippageFactor)
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
		it.Event = new(EthRebaserNewMaxSlippageFactor)
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
func (it *EthRebaserNewMaxSlippageFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserNewMaxSlippageFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserNewMaxSlippageFactor represents a NewMaxSlippageFactor event raised by the EthRebaser contract.
type EthRebaserNewMaxSlippageFactor struct {
	OldSlippageFactor *big.Int
	NewSlippageFactor *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewMaxSlippageFactor is a free log retrieval operation binding the contract event 0xe21b25c4eda0340cd924f3247795d0acde6c304b68ae77657bb2d4e840198bf8.
//
// Solidity: event NewMaxSlippageFactor(uint256 oldSlippageFactor, uint256 newSlippageFactor)
func (_EthRebaser *EthRebaserFilterer) FilterNewMaxSlippageFactor(opts *bind.FilterOpts) (*EthRebaserNewMaxSlippageFactorIterator, error) {

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "NewMaxSlippageFactor")
	if err != nil {
		return nil, err
	}
	return &EthRebaserNewMaxSlippageFactorIterator{contract: _EthRebaser.contract, event: "NewMaxSlippageFactor", logs: logs, sub: sub}, nil
}

// WatchNewMaxSlippageFactor is a free log subscription operation binding the contract event 0xe21b25c4eda0340cd924f3247795d0acde6c304b68ae77657bb2d4e840198bf8.
//
// Solidity: event NewMaxSlippageFactor(uint256 oldSlippageFactor, uint256 newSlippageFactor)
func (_EthRebaser *EthRebaserFilterer) WatchNewMaxSlippageFactor(opts *bind.WatchOpts, sink chan<- *EthRebaserNewMaxSlippageFactor) (event.Subscription, error) {

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "NewMaxSlippageFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserNewMaxSlippageFactor)
				if err := _EthRebaser.contract.UnpackLog(event, "NewMaxSlippageFactor", log); err != nil {
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

// ParseNewMaxSlippageFactor is a log parse operation binding the contract event 0xe21b25c4eda0340cd924f3247795d0acde6c304b68ae77657bb2d4e840198bf8.
//
// Solidity: event NewMaxSlippageFactor(uint256 oldSlippageFactor, uint256 newSlippageFactor)
func (_EthRebaser *EthRebaserFilterer) ParseNewMaxSlippageFactor(log types.Log) (*EthRebaserNewMaxSlippageFactor, error) {
	event := new(EthRebaserNewMaxSlippageFactor)
	if err := _EthRebaser.contract.UnpackLog(event, "NewMaxSlippageFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthRebaserNewPendingGovIterator is returned from FilterNewPendingGov and is used to iterate over the raw logs and unpacked data for NewPendingGov events raised by the EthRebaser contract.
type EthRebaserNewPendingGovIterator struct {
	Event *EthRebaserNewPendingGov // Event containing the contract specifics and raw log

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
func (it *EthRebaserNewPendingGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserNewPendingGov)
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
		it.Event = new(EthRebaserNewPendingGov)
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
func (it *EthRebaserNewPendingGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserNewPendingGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserNewPendingGov represents a NewPendingGov event raised by the EthRebaser contract.
type EthRebaserNewPendingGov struct {
	OldPendingGov common.Address
	NewPendingGov common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewPendingGov is a free log retrieval operation binding the contract event 0x6163d5b9efd962645dd649e6e48a61bcb0f9df00997a2398b80d135a9ab0c61e.
//
// Solidity: event NewPendingGov(address oldPendingGov, address newPendingGov)
func (_EthRebaser *EthRebaserFilterer) FilterNewPendingGov(opts *bind.FilterOpts) (*EthRebaserNewPendingGovIterator, error) {

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "NewPendingGov")
	if err != nil {
		return nil, err
	}
	return &EthRebaserNewPendingGovIterator{contract: _EthRebaser.contract, event: "NewPendingGov", logs: logs, sub: sub}, nil
}

// WatchNewPendingGov is a free log subscription operation binding the contract event 0x6163d5b9efd962645dd649e6e48a61bcb0f9df00997a2398b80d135a9ab0c61e.
//
// Solidity: event NewPendingGov(address oldPendingGov, address newPendingGov)
func (_EthRebaser *EthRebaserFilterer) WatchNewPendingGov(opts *bind.WatchOpts, sink chan<- *EthRebaserNewPendingGov) (event.Subscription, error) {

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "NewPendingGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserNewPendingGov)
				if err := _EthRebaser.contract.UnpackLog(event, "NewPendingGov", log); err != nil {
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

// ParseNewPendingGov is a log parse operation binding the contract event 0x6163d5b9efd962645dd649e6e48a61bcb0f9df00997a2398b80d135a9ab0c61e.
//
// Solidity: event NewPendingGov(address oldPendingGov, address newPendingGov)
func (_EthRebaser *EthRebaserFilterer) ParseNewPendingGov(log types.Log) (*EthRebaserNewPendingGov, error) {
	event := new(EthRebaserNewPendingGov)
	if err := _EthRebaser.contract.UnpackLog(event, "NewPendingGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthRebaserNewRebaseMintPercentIterator is returned from FilterNewRebaseMintPercent and is used to iterate over the raw logs and unpacked data for NewRebaseMintPercent events raised by the EthRebaser contract.
type EthRebaserNewRebaseMintPercentIterator struct {
	Event *EthRebaserNewRebaseMintPercent // Event containing the contract specifics and raw log

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
func (it *EthRebaserNewRebaseMintPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserNewRebaseMintPercent)
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
		it.Event = new(EthRebaserNewRebaseMintPercent)
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
func (it *EthRebaserNewRebaseMintPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserNewRebaseMintPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserNewRebaseMintPercent represents a NewRebaseMintPercent event raised by the EthRebaser contract.
type EthRebaserNewRebaseMintPercent struct {
	OldRebaseMintPerc *big.Int
	NewRebaseMintPerc *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRebaseMintPercent is a free log retrieval operation binding the contract event 0x59b3ffce759ec92c629beee27554d8fbc2ca1a05020fa0cf500c890c172094be.
//
// Solidity: event NewRebaseMintPercent(uint256 oldRebaseMintPerc, uint256 newRebaseMintPerc)
func (_EthRebaser *EthRebaserFilterer) FilterNewRebaseMintPercent(opts *bind.FilterOpts) (*EthRebaserNewRebaseMintPercentIterator, error) {

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "NewRebaseMintPercent")
	if err != nil {
		return nil, err
	}
	return &EthRebaserNewRebaseMintPercentIterator{contract: _EthRebaser.contract, event: "NewRebaseMintPercent", logs: logs, sub: sub}, nil
}

// WatchNewRebaseMintPercent is a free log subscription operation binding the contract event 0x59b3ffce759ec92c629beee27554d8fbc2ca1a05020fa0cf500c890c172094be.
//
// Solidity: event NewRebaseMintPercent(uint256 oldRebaseMintPerc, uint256 newRebaseMintPerc)
func (_EthRebaser *EthRebaserFilterer) WatchNewRebaseMintPercent(opts *bind.WatchOpts, sink chan<- *EthRebaserNewRebaseMintPercent) (event.Subscription, error) {

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "NewRebaseMintPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserNewRebaseMintPercent)
				if err := _EthRebaser.contract.UnpackLog(event, "NewRebaseMintPercent", log); err != nil {
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

// ParseNewRebaseMintPercent is a log parse operation binding the contract event 0x59b3ffce759ec92c629beee27554d8fbc2ca1a05020fa0cf500c890c172094be.
//
// Solidity: event NewRebaseMintPercent(uint256 oldRebaseMintPerc, uint256 newRebaseMintPerc)
func (_EthRebaser *EthRebaserFilterer) ParseNewRebaseMintPercent(log types.Log) (*EthRebaserNewRebaseMintPercent, error) {
	event := new(EthRebaserNewRebaseMintPercent)
	if err := _EthRebaser.contract.UnpackLog(event, "NewRebaseMintPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthRebaserNewReserveContractIterator is returned from FilterNewReserveContract and is used to iterate over the raw logs and unpacked data for NewReserveContract events raised by the EthRebaser contract.
type EthRebaserNewReserveContractIterator struct {
	Event *EthRebaserNewReserveContract // Event containing the contract specifics and raw log

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
func (it *EthRebaserNewReserveContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserNewReserveContract)
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
		it.Event = new(EthRebaserNewReserveContract)
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
func (it *EthRebaserNewReserveContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserNewReserveContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserNewReserveContract represents a NewReserveContract event raised by the EthRebaser contract.
type EthRebaserNewReserveContract struct {
	OldReserveContract common.Address
	NewReserveContract common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewReserveContract is a free log retrieval operation binding the contract event 0xce840d2205f08f33375689943da5da9fdfde146fcbb5553b17910a60c8284a20.
//
// Solidity: event NewReserveContract(address oldReserveContract, address newReserveContract)
func (_EthRebaser *EthRebaserFilterer) FilterNewReserveContract(opts *bind.FilterOpts) (*EthRebaserNewReserveContractIterator, error) {

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "NewReserveContract")
	if err != nil {
		return nil, err
	}
	return &EthRebaserNewReserveContractIterator{contract: _EthRebaser.contract, event: "NewReserveContract", logs: logs, sub: sub}, nil
}

// WatchNewReserveContract is a free log subscription operation binding the contract event 0xce840d2205f08f33375689943da5da9fdfde146fcbb5553b17910a60c8284a20.
//
// Solidity: event NewReserveContract(address oldReserveContract, address newReserveContract)
func (_EthRebaser *EthRebaserFilterer) WatchNewReserveContract(opts *bind.WatchOpts, sink chan<- *EthRebaserNewReserveContract) (event.Subscription, error) {

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "NewReserveContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserNewReserveContract)
				if err := _EthRebaser.contract.UnpackLog(event, "NewReserveContract", log); err != nil {
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

// ParseNewReserveContract is a log parse operation binding the contract event 0xce840d2205f08f33375689943da5da9fdfde146fcbb5553b17910a60c8284a20.
//
// Solidity: event NewReserveContract(address oldReserveContract, address newReserveContract)
func (_EthRebaser *EthRebaserFilterer) ParseNewReserveContract(log types.Log) (*EthRebaserNewReserveContract, error) {
	event := new(EthRebaserNewReserveContract)
	if err := _EthRebaser.contract.UnpackLog(event, "NewReserveContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthRebaserTransactionFailedIterator is returned from FilterTransactionFailed and is used to iterate over the raw logs and unpacked data for TransactionFailed events raised by the EthRebaser contract.
type EthRebaserTransactionFailedIterator struct {
	Event *EthRebaserTransactionFailed // Event containing the contract specifics and raw log

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
func (it *EthRebaserTransactionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserTransactionFailed)
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
		it.Event = new(EthRebaserTransactionFailed)
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
func (it *EthRebaserTransactionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserTransactionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserTransactionFailed represents a TransactionFailed event raised by the EthRebaser contract.
type EthRebaserTransactionFailed struct {
	Destination common.Address
	Index       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransactionFailed is a free log retrieval operation binding the contract event 0x8091ecaaa54ebb82e02d36c2c336528e0fcb9b3430fc1291ac88295032b9c263.
//
// Solidity: event TransactionFailed(address indexed destination, uint256 index, bytes data)
func (_EthRebaser *EthRebaserFilterer) FilterTransactionFailed(opts *bind.FilterOpts, destination []common.Address) (*EthRebaserTransactionFailedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "TransactionFailed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &EthRebaserTransactionFailedIterator{contract: _EthRebaser.contract, event: "TransactionFailed", logs: logs, sub: sub}, nil
}

// WatchTransactionFailed is a free log subscription operation binding the contract event 0x8091ecaaa54ebb82e02d36c2c336528e0fcb9b3430fc1291ac88295032b9c263.
//
// Solidity: event TransactionFailed(address indexed destination, uint256 index, bytes data)
func (_EthRebaser *EthRebaserFilterer) WatchTransactionFailed(opts *bind.WatchOpts, sink chan<- *EthRebaserTransactionFailed, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "TransactionFailed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserTransactionFailed)
				if err := _EthRebaser.contract.UnpackLog(event, "TransactionFailed", log); err != nil {
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

// ParseTransactionFailed is a log parse operation binding the contract event 0x8091ecaaa54ebb82e02d36c2c336528e0fcb9b3430fc1291ac88295032b9c263.
//
// Solidity: event TransactionFailed(address indexed destination, uint256 index, bytes data)
func (_EthRebaser *EthRebaserFilterer) ParseTransactionFailed(log types.Log) (*EthRebaserTransactionFailed, error) {
	event := new(EthRebaserTransactionFailed)
	if err := _EthRebaser.contract.UnpackLog(event, "TransactionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthRebaserTreasuryIncreasedIterator is returned from FilterTreasuryIncreased and is used to iterate over the raw logs and unpacked data for TreasuryIncreased events raised by the EthRebaser contract.
type EthRebaserTreasuryIncreasedIterator struct {
	Event *EthRebaserTreasuryIncreased // Event containing the contract specifics and raw log

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
func (it *EthRebaserTreasuryIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthRebaserTreasuryIncreased)
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
		it.Event = new(EthRebaserTreasuryIncreased)
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
func (it *EthRebaserTreasuryIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthRebaserTreasuryIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthRebaserTreasuryIncreased represents a TreasuryIncreased event raised by the EthRebaser contract.
type EthRebaserTreasuryIncreased struct {
	ReservesAdded    *big.Int
	YamsSold         *big.Int
	YamsFromReserves *big.Int
	YamsToReserves   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTreasuryIncreased is a free log retrieval operation binding the contract event 0xb335015c214ae37ed112cc5eb042235c0ea40a7617987e1bd847839143872350.
//
// Solidity: event TreasuryIncreased(uint256 reservesAdded, uint256 yamsSold, uint256 yamsFromReserves, uint256 yamsToReserves)
func (_EthRebaser *EthRebaserFilterer) FilterTreasuryIncreased(opts *bind.FilterOpts) (*EthRebaserTreasuryIncreasedIterator, error) {

	logs, sub, err := _EthRebaser.contract.FilterLogs(opts, "TreasuryIncreased")
	if err != nil {
		return nil, err
	}
	return &EthRebaserTreasuryIncreasedIterator{contract: _EthRebaser.contract, event: "TreasuryIncreased", logs: logs, sub: sub}, nil
}

// WatchTreasuryIncreased is a free log subscription operation binding the contract event 0xb335015c214ae37ed112cc5eb042235c0ea40a7617987e1bd847839143872350.
//
// Solidity: event TreasuryIncreased(uint256 reservesAdded, uint256 yamsSold, uint256 yamsFromReserves, uint256 yamsToReserves)
func (_EthRebaser *EthRebaserFilterer) WatchTreasuryIncreased(opts *bind.WatchOpts, sink chan<- *EthRebaserTreasuryIncreased) (event.Subscription, error) {

	logs, sub, err := _EthRebaser.contract.WatchLogs(opts, "TreasuryIncreased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthRebaserTreasuryIncreased)
				if err := _EthRebaser.contract.UnpackLog(event, "TreasuryIncreased", log); err != nil {
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

// ParseTreasuryIncreased is a log parse operation binding the contract event 0xb335015c214ae37ed112cc5eb042235c0ea40a7617987e1bd847839143872350.
//
// Solidity: event TreasuryIncreased(uint256 reservesAdded, uint256 yamsSold, uint256 yamsFromReserves, uint256 yamsToReserves)
func (_EthRebaser *EthRebaserFilterer) ParseTreasuryIncreased(log types.Log) (*EthRebaserTreasuryIncreased, error) {
	event := new(EthRebaserTreasuryIncreased)
	if err := _EthRebaser.contract.UnpackLog(event, "TreasuryIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
