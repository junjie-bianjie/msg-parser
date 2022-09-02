package perm

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type DocMsgBlockAccount struct {
	Address  string `bson:"address" json:"address"`
	Operator string `bson:"operator" json:"operator"`
}

func (m *DocMsgBlockAccount) GetType() string {
	return DocTypeBlockAccount
}

func (m *DocMsgBlockAccount) BuildMsg(v interface{}) {
	msg := v.(*MsgBlockAccount)
	m.Address = msg.Address
	m.Operator = msg.Operator
}

func (m *DocMsgBlockAccount) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgBlockAccount)
	addrs = append(addrs, msg.Address, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
