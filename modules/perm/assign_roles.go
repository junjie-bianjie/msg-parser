package perm

import (
	"github.com/bianjieai/iritamod/modules/perm/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type DocMsgAssignRoles struct {
	Address  string   `bson:"address" json:"address"`
	Roles    []string `bson:"roles" json:"roles"`
	Operator string   `bson:"operator" json:"operator"`
}

func (m *DocMsgAssignRoles) GetType() string {
	return DocTypeAssignRoles
}

func (m *DocMsgAssignRoles) BuildMsg(v interface{}) {
	msg := v.(*MsgAssignRoles)
	m.Address = msg.Address

	for _, r := range msg.Roles {
		name := types.Role_name[int32(r)]
		m.Roles = append(m.Roles, name)
	}
	m.Operator = msg.Operator
}

func (m *DocMsgAssignRoles) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgAssignRoles)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
