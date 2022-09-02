package perm

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type permClient struct {
}

func NewClient() Client {
	return permClient{}
}

func (nft permClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgAssignRoles:
		docMsg := DocMsgAssignRoles{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUnassignRoles:
		docMsg := DocMsgUnassignRoles{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgBlockAccount:
		docMsg := DocMsgBlockAccount{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUnblockAccount:
		docMsg := DocMsgUnblockAccount{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
