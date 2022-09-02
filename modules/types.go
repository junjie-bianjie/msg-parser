package msgs

import (
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/perm"
	iritaslashing "github.com/bianjieai/iritamod/modules/slashing"
	tibctranfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibcclient "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	tibcpacket "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stake "github.com/cosmos/cosmos-sdk/x/staking/types"
	coinswap "github.com/irisnet/irismod/modules/coinswap/types"
	htlc "github.com/irisnet/irismod/modules/htlc/types"
	mt "github.com/irisnet/irismod/modules/mt/types"
	nft "github.com/irisnet/irismod/modules/nft/types"
	oracle "github.com/irisnet/irismod/modules/oracle/types"
	random "github.com/irisnet/irismod/modules/random/types"
	record "github.com/irisnet/irismod/modules/record/types"
	service "github.com/irisnet/irismod/modules/service/types"
	token "github.com/irisnet/irismod/modules/token/types"
	models "github.com/kaifei-bianjie/msg-parser/types"
	//"github.com/CosmWasm/wasmd/x/wasm"
	evm "github.com/tharsis/ethermint/x/evm/types"
)

const (
	MsgTypeSend          = "send"
	MsgTypeMultiSend     = "multisend"
	MsgTypeNFTMint       = "mint_nft"
	MsgTypeNFTEdit       = "edit_nft"
	MsgTypeNFTTransfer   = "transfer_nft"
	MsgTypeTransferDenom = "transfer_denom"
	MsgTypeNFTBurn       = "burn_nft"
	MsgTypeIssueDenom    = "issue_denom"
	MsgTypeRecordCreate  = "create_record"

	MsgTypeMTIssueDenom    = "mt_issue_denom"
	MsgTypeMTTransferDenom = "mt_transfer_denom"
	MsgTypeMintMT          = "mint_mt"
	MsgTypeTransferMT      = "transfer_mt"
	MsgTypeEditMT          = "edit_mt"
	MsgTypeBurnMT          = "burn_mt"

	MsgTypeMintToken          = "mint_token"
	MsgTypeBurnToken          = "burn_token"
	MsgTypeEditToken          = "edit_token"
	MsgTypeIssueToken         = "issue_token"
	MsgTypeTransferTokenOwner = "transfer_token_owner"

	MsgTypeDefineService             = "define_service"               // type for MsgDefineService
	MsgTypeBindService               = "bind_service"                 // type for MsgBindService
	MsgTypeUpdateServiceBinding      = "update_service_binding"       // type for MsgUpdateServiceBinding
	MsgTypeServiceSetWithdrawAddress = "service/set_withdraw_address" // type for MsgSetWithdrawAddress
	MsgTypeDisableServiceBinding     = "disable_service_binding"      // type for MsgDisableServiceBinding
	MsgTypeEnableServiceBinding      = "enable_service_binding"       // type for MsgEnableServiceBinding
	MsgTypeRefundServiceDeposit      = "refund_service_deposit"       // type for MsgRefundServiceDeposit
	MsgTypeCallService               = "call_service"                 // type for MsgCallService
	MsgTypeRespondService            = "respond_service"              // type for MsgRespondService
	MsgTypePauseRequestContext       = "pause_request_context"        // type for MsgPauseRequestContext
	MsgTypeStartRequestContext       = "start_request_context"        // type for MsgStartRequestContext
	MsgTypeKillRequestContext        = "kill_request_context"         // type for MsgKillRequestContext
	MsgTypeUpdateRequestContext      = "update_request_context"       // type for MsgUpdateRequestContext
	MsgTypeWithdrawEarnedFees        = "withdraw_earned_fees"         // type for MsgWithdrawEarnedFees

	MsgTypeGrantAllowance  = "grant_allowance"
	MsgTypeRevokeAllowance = "revoke_allowance"

	MsgTypeStakeCreateValidator           = "create_validator"
	MsgTypeStakeEditValidator             = "edit_validator"
	MsgTypeStakeDelegate                  = "delegate"
	MsgTypeStakeBeginUnbonding            = "begin_unbonding"
	MsgTypeBeginRedelegate                = "begin_redelegate"
	MsgTypeUnjail                         = "unjail"
	MsgTypeUnjailValidator                = "unjail_validator"
	MsgTypeSetWithdrawAddress             = "set_withdraw_address"
	MsgTypeWithdrawDelegatorReward        = "withdraw_delegator_reward"
	MsgTypeMsgFundCommunityPool           = "fund_community_pool"
	MsgTypeMsgWithdrawValidatorCommission = "withdraw_validator_commission"
	MsgTypeSubmitProposal                 = "submit_proposal"
	MsgTypeDeposit                        = "deposit"
	MsgTypeVote                           = "vote"

	MsgTypeCreateHTLC = "create_htlc"
	MsgTypeClaimHTLC  = "claim_htlc"
	MsgTypeRefundHTLC = "refund_htlc"

	MsgTypeAddLiquidity    = "add_liquidity"
	MsgTypeRemoveLiquidity = "remove_liquidity"
	MsgTypeSwapOrder       = "swap_order"

	MsgTypeSubmitEvidence  = "submit_evidence"
	MsgTypeVerifyInvariant = "verify_invariant"

	MsgTypeUpdateIdentity = "update_identity"
	MsgTypeCreateIdentity = "create_identity"

	TxTypeRequestRand = "request_rand"

	TxTypeCreateFeed = "create_feed"
	TxTypeEditFeed   = "edit_feed"
	TxTypePauseFeed  = "pause_feed"
	TxTypeStartFeed  = "start_feed"

	MsgTypeUpdateAdmin         = "update_contract_admin"
	MsgTypeClearAdmin          = "clear_contract_admin"
	MsgTypeExecuteContract     = "execute"
	MsgTypeInstantiateContract = "instantiate"
	MsgTypeMigrateContract     = "migrate"
	MsgTypeStoreCode           = "store_code"

	MsgTypeTIBCNftTransfer     = "tibc_nft_transfer"
	MsgTypeTIBCRecvPacket      = "tibc_recv_packet"
	MsgTypeTIBCUpdateClient    = "tibc_update_client"
	MsgTypeTIBCAcknowledgement = "tibc_acknowledge_packet"
	MsgTypeTIBCCleanPacket     = "clean_packet"
	MsgTypeTIBCRecvCleanPacket = "recv_clean_packet"

	MsgTypeEthereumTx = "ethereum_tx"

	DocTypeAssignRoles   = "assign_roles"
	DocTypeUnassignRoles = "unassign_roles"
	DocTypeBlockAccount  = "block_account"
	DocTypeUnlockAccount = "unblock_account"
)

type (
	MsgDocInfo struct {
		DocTxMsg models.TxMsg
		Addrs    []string
		Signers  []string
	}
	SdkMsg sdk.Msg
	Msg    models.Msg

	Coin models.Coin

	Coins []*Coin

	MsgSend      = bank.MsgSend
	MsgMultiSend = bank.MsgMultiSend

	MsgNFTMint       = nft.MsgMintNFT
	MsgNFTEdit       = nft.MsgEditNFT
	MsgNFTTransfer   = nft.MsgTransferNFT
	MsgNFTBurn       = nft.MsgBurnNFT
	MsgIssueDenom    = nft.MsgIssueDenom
	MsgTransferDenom = nft.MsgTransferDenom

	MsgMTMint          = mt.MsgMintMT
	MsgMTEdit          = mt.MsgEditMT
	MsgMTTransfer      = mt.MsgTransferMT
	MsgMTBurn          = mt.MsgBurnMT
	MsgMTIssueDenom    = mt.MsgIssueDenom
	MsgMTTransferDenom = mt.MsgTransferDenom

	MsgDefineService  = service.MsgDefineService
	MsgBindService    = service.MsgBindService
	MsgCallService    = service.MsgCallService
	MsgRespondService = service.MsgRespondService

	MsgUpdateServiceBinding  = service.MsgUpdateServiceBinding
	MsgSetWithdrawAddress    = service.MsgSetWithdrawAddress
	MsgDisableServiceBinding = service.MsgDisableServiceBinding
	MsgEnableServiceBinding  = service.MsgEnableServiceBinding
	MsgRefundServiceDeposit  = service.MsgRefundServiceDeposit
	MsgPauseRequestContext   = service.MsgPauseRequestContext
	MsgStartRequestContext   = service.MsgStartRequestContext
	MsgKillRequestContext    = service.MsgKillRequestContext
	MsgUpdateRequestContext  = service.MsgUpdateRequestContext
	MsgWithdrawEarnedFees    = service.MsgWithdrawEarnedFees

	MsgRecordCreate = record.MsgCreateRecord

	MsgIssueToken         = token.MsgIssueToken
	MsgEditToken          = token.MsgEditToken
	MsgBurnToken          = token.MsgBurnToken
	MsgMintToken          = token.MsgMintToken
	MsgTransferTokenOwner = token.MsgTransferTokenOwner

	MsgStakeCreate                 = stake.MsgCreateValidator
	MsgStakeEdit                   = stake.MsgEditValidator
	MsgStakeDelegate               = stake.MsgDelegate
	MsgStakeBeginUnbonding         = stake.MsgUndelegate
	MsgBeginRedelegate             = stake.MsgBeginRedelegate
	MsgUnjail                      = slashing.MsgUnjail
	MsgUnjailValidator             = iritaslashing.MsgUnjailValidator
	MsgStakeSetWithdrawAddress     = distribution.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward     = distribution.MsgWithdrawDelegatorReward
	MsgFundCommunityPool           = distribution.MsgFundCommunityPool
	MsgWithdrawValidatorCommission = distribution.MsgWithdrawValidatorCommission
	StakeValidator                 = stake.Validator
	Delegation                     = stake.Delegation
	UnbondingDelegation            = stake.UnbondingDelegation

	MsgGrantAllowance  = feegrant.MsgGrantAllowance
	MsgRevokeAllowance = feegrant.MsgRevokeAllowance

	MsgDeposit        = gov.MsgDeposit
	MsgSubmitProposal = gov.MsgSubmitProposal
	TextProposal      = gov.TextProposal
	MsgVote           = gov.MsgVote
	Proposal          = gov.Proposal
	SdkVote           = gov.Vote
	GovContent        = gov.Content

	MsgSwapOrder       = coinswap.MsgSwapOrder
	MsgAddLiquidity    = coinswap.MsgAddLiquidity
	MsgRemoveLiquidity = coinswap.MsgRemoveLiquidity

	MsgClaimHTLC  = htlc.MsgClaimHTLC
	MsgCreateHTLC = htlc.MsgCreateHTLC

	MsgSubmitEvidence  = evidence.MsgSubmitEvidence
	MsgVerifyInvariant = crisis.MsgVerifyInvariant

	MsgCreateFeed = oracle.MsgCreateFeed
	MsgEditFeed   = oracle.MsgEditFeed
	MsgPauseFeed  = oracle.MsgPauseFeed
	MsgStartFeed  = oracle.MsgStartFeed

	MsgRequestRandom = random.MsgRequestRandom

	MsgCreateIdentity = identity.MsgCreateIdentity
	MsgUpdateIdentity = identity.MsgUpdateIdentity

	//MsgStoreCode           = wasm.MsgStoreCode
	//MsgInstantiateContract = wasm.MsgInstantiateContract
	//MsgExecuteContract     = wasm.MsgExecuteContract
	//MsgMigrateContract     = wasm.MsgMigrateContract
	//MsgUpdateAdmin         = wasm.MsgUpdateAdmin
	//MsgClearAdmin          = wasm.MsgClearAdmin
	TIBCAcknowledgement        = tibcpacket.Acknowledgement
	NonFungibleTokenPacketData = tibctranfer.NonFungibleTokenPacketData
	MsgTIBCNftTransfer         = tibctranfer.MsgNftTransfer
	MsgTIBCUpdateClient        = tibcclient.MsgUpdateClient
	MsgTIBCRecvPacket          = tibcpacket.MsgRecvPacket
	MsgTIBCAcknowledgement     = tibcpacket.MsgAcknowledgement
	MsgCleanPacket             = tibcpacket.MsgCleanPacket
	MsgRecvCleanPacket         = tibcpacket.MsgRecvCleanPacket

	MsgEthereumTx = evm.MsgEthereumTx

	MsgAssignRoles    = perm.MsgAssignRoles
	MsgUnassignRoles  = perm.MsgUnassignRoles
	MsgBlockAccount   = perm.MsgBlockAccount
	MsgUnblockAccount = perm.MsgUnblockAccount
)
