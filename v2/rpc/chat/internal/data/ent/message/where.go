// Code generated by ent, DO NOT EDIT.

package message

import (
	"chat/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Updated applies equality check predicate on the "updated" field. It's identical to UpdatedEQ.
func Updated(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdated), v))
	})
}

// SessionKey applies equality check predicate on the "sessionKey" field. It's identical to SessionKeyEQ.
func SessionKey(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionKey), v))
	})
}

// From applies equality check predicate on the "from" field. It's identical to FromEQ.
func From(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFrom), v))
	})
}

// To applies equality check predicate on the "to" field. It's identical to ToEQ.
func To(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTo), v))
	})
}

// SessionType applies equality check predicate on the "session_type" field. It's identical to SessionTypeEQ.
func SessionType(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionType), v))
	})
}

// ClientMsgID applies equality check predicate on the "client_msg_id" field. It's identical to ClientMsgIDEQ.
func ClientMsgID(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientMsgID), v))
	})
}

// ServerMsgSeq applies equality check predicate on the "server_msg_seq" field. It's identical to ServerMsgSeqEQ.
func ServerMsgSeq(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServerMsgSeq), v))
	})
}

// MsgType applies equality check predicate on the "msg_type" field. It's identical to MsgTypeEQ.
func MsgType(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgType), v))
	})
}

// MsgData applies equality check predicate on the "msg_data" field. It's identical to MsgDataEQ.
func MsgData(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgData), v))
	})
}

// MsgResCode applies equality check predicate on the "msg_res_code" field. It's identical to MsgResCodeEQ.
func MsgResCode(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgResCode), v))
	})
}

// MsgFeature applies equality check predicate on the "msg_feature" field. It's identical to MsgFeatureEQ.
func MsgFeature(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgFeature), v))
	})
}

// MsgStatus applies equality check predicate on the "msg_status" field. It's identical to MsgStatusEQ.
func MsgStatus(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgStatus), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// UpdatedEQ applies the EQ predicate on the "updated" field.
func UpdatedEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdated), v))
	})
}

// UpdatedNEQ applies the NEQ predicate on the "updated" field.
func UpdatedNEQ(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdated), v))
	})
}

// UpdatedIn applies the In predicate on the "updated" field.
func UpdatedIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdated), v...))
	})
}

// UpdatedNotIn applies the NotIn predicate on the "updated" field.
func UpdatedNotIn(vs ...time.Time) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdated), v...))
	})
}

// UpdatedGT applies the GT predicate on the "updated" field.
func UpdatedGT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdated), v))
	})
}

// UpdatedGTE applies the GTE predicate on the "updated" field.
func UpdatedGTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdated), v))
	})
}

// UpdatedLT applies the LT predicate on the "updated" field.
func UpdatedLT(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdated), v))
	})
}

// UpdatedLTE applies the LTE predicate on the "updated" field.
func UpdatedLTE(v time.Time) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdated), v))
	})
}

// SessionKeyEQ applies the EQ predicate on the "sessionKey" field.
func SessionKeyEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionKey), v))
	})
}

// SessionKeyNEQ applies the NEQ predicate on the "sessionKey" field.
func SessionKeyNEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessionKey), v))
	})
}

// SessionKeyIn applies the In predicate on the "sessionKey" field.
func SessionKeyIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSessionKey), v...))
	})
}

// SessionKeyNotIn applies the NotIn predicate on the "sessionKey" field.
func SessionKeyNotIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSessionKey), v...))
	})
}

// SessionKeyGT applies the GT predicate on the "sessionKey" field.
func SessionKeyGT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessionKey), v))
	})
}

// SessionKeyGTE applies the GTE predicate on the "sessionKey" field.
func SessionKeyGTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessionKey), v))
	})
}

// SessionKeyLT applies the LT predicate on the "sessionKey" field.
func SessionKeyLT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessionKey), v))
	})
}

// SessionKeyLTE applies the LTE predicate on the "sessionKey" field.
func SessionKeyLTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessionKey), v))
	})
}

// SessionKeyContains applies the Contains predicate on the "sessionKey" field.
func SessionKeyContains(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSessionKey), v))
	})
}

// SessionKeyHasPrefix applies the HasPrefix predicate on the "sessionKey" field.
func SessionKeyHasPrefix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSessionKey), v))
	})
}

// SessionKeyHasSuffix applies the HasSuffix predicate on the "sessionKey" field.
func SessionKeyHasSuffix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSessionKey), v))
	})
}

// SessionKeyEqualFold applies the EqualFold predicate on the "sessionKey" field.
func SessionKeyEqualFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSessionKey), v))
	})
}

// SessionKeyContainsFold applies the ContainsFold predicate on the "sessionKey" field.
func SessionKeyContainsFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSessionKey), v))
	})
}

// FromEQ applies the EQ predicate on the "from" field.
func FromEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFrom), v))
	})
}

// FromNEQ applies the NEQ predicate on the "from" field.
func FromNEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFrom), v))
	})
}

// FromIn applies the In predicate on the "from" field.
func FromIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFrom), v...))
	})
}

// FromNotIn applies the NotIn predicate on the "from" field.
func FromNotIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFrom), v...))
	})
}

// FromGT applies the GT predicate on the "from" field.
func FromGT(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFrom), v))
	})
}

// FromGTE applies the GTE predicate on the "from" field.
func FromGTE(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFrom), v))
	})
}

// FromLT applies the LT predicate on the "from" field.
func FromLT(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFrom), v))
	})
}

// FromLTE applies the LTE predicate on the "from" field.
func FromLTE(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFrom), v))
	})
}

// ToEQ applies the EQ predicate on the "to" field.
func ToEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTo), v))
	})
}

// ToNEQ applies the NEQ predicate on the "to" field.
func ToNEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTo), v))
	})
}

// ToIn applies the In predicate on the "to" field.
func ToIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTo), v...))
	})
}

// ToNotIn applies the NotIn predicate on the "to" field.
func ToNotIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTo), v...))
	})
}

// ToGT applies the GT predicate on the "to" field.
func ToGT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTo), v))
	})
}

// ToGTE applies the GTE predicate on the "to" field.
func ToGTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTo), v))
	})
}

// ToLT applies the LT predicate on the "to" field.
func ToLT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTo), v))
	})
}

// ToLTE applies the LTE predicate on the "to" field.
func ToLTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTo), v))
	})
}

// ToContains applies the Contains predicate on the "to" field.
func ToContains(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTo), v))
	})
}

// ToHasPrefix applies the HasPrefix predicate on the "to" field.
func ToHasPrefix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTo), v))
	})
}

// ToHasSuffix applies the HasSuffix predicate on the "to" field.
func ToHasSuffix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTo), v))
	})
}

// ToEqualFold applies the EqualFold predicate on the "to" field.
func ToEqualFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTo), v))
	})
}

// ToContainsFold applies the ContainsFold predicate on the "to" field.
func ToContainsFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTo), v))
	})
}

// SessionTypeEQ applies the EQ predicate on the "session_type" field.
func SessionTypeEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessionType), v))
	})
}

// SessionTypeNEQ applies the NEQ predicate on the "session_type" field.
func SessionTypeNEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessionType), v))
	})
}

// SessionTypeIn applies the In predicate on the "session_type" field.
func SessionTypeIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSessionType), v...))
	})
}

// SessionTypeNotIn applies the NotIn predicate on the "session_type" field.
func SessionTypeNotIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSessionType), v...))
	})
}

// SessionTypeGT applies the GT predicate on the "session_type" field.
func SessionTypeGT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessionType), v))
	})
}

// SessionTypeGTE applies the GTE predicate on the "session_type" field.
func SessionTypeGTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessionType), v))
	})
}

// SessionTypeLT applies the LT predicate on the "session_type" field.
func SessionTypeLT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessionType), v))
	})
}

// SessionTypeLTE applies the LTE predicate on the "session_type" field.
func SessionTypeLTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessionType), v))
	})
}

// ClientMsgIDEQ applies the EQ predicate on the "client_msg_id" field.
func ClientMsgIDEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDNEQ applies the NEQ predicate on the "client_msg_id" field.
func ClientMsgIDNEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDIn applies the In predicate on the "client_msg_id" field.
func ClientMsgIDIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientMsgID), v...))
	})
}

// ClientMsgIDNotIn applies the NotIn predicate on the "client_msg_id" field.
func ClientMsgIDNotIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientMsgID), v...))
	})
}

// ClientMsgIDGT applies the GT predicate on the "client_msg_id" field.
func ClientMsgIDGT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDGTE applies the GTE predicate on the "client_msg_id" field.
func ClientMsgIDGTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDLT applies the LT predicate on the "client_msg_id" field.
func ClientMsgIDLT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDLTE applies the LTE predicate on the "client_msg_id" field.
func ClientMsgIDLTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDContains applies the Contains predicate on the "client_msg_id" field.
func ClientMsgIDContains(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDHasPrefix applies the HasPrefix predicate on the "client_msg_id" field.
func ClientMsgIDHasPrefix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDHasSuffix applies the HasSuffix predicate on the "client_msg_id" field.
func ClientMsgIDHasSuffix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDEqualFold applies the EqualFold predicate on the "client_msg_id" field.
func ClientMsgIDEqualFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientMsgID), v))
	})
}

// ClientMsgIDContainsFold applies the ContainsFold predicate on the "client_msg_id" field.
func ClientMsgIDContainsFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientMsgID), v))
	})
}

// ServerMsgSeqEQ applies the EQ predicate on the "server_msg_seq" field.
func ServerMsgSeqEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServerMsgSeq), v))
	})
}

// ServerMsgSeqNEQ applies the NEQ predicate on the "server_msg_seq" field.
func ServerMsgSeqNEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldServerMsgSeq), v))
	})
}

// ServerMsgSeqIn applies the In predicate on the "server_msg_seq" field.
func ServerMsgSeqIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldServerMsgSeq), v...))
	})
}

// ServerMsgSeqNotIn applies the NotIn predicate on the "server_msg_seq" field.
func ServerMsgSeqNotIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldServerMsgSeq), v...))
	})
}

// ServerMsgSeqGT applies the GT predicate on the "server_msg_seq" field.
func ServerMsgSeqGT(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldServerMsgSeq), v))
	})
}

// ServerMsgSeqGTE applies the GTE predicate on the "server_msg_seq" field.
func ServerMsgSeqGTE(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldServerMsgSeq), v))
	})
}

// ServerMsgSeqLT applies the LT predicate on the "server_msg_seq" field.
func ServerMsgSeqLT(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldServerMsgSeq), v))
	})
}

// ServerMsgSeqLTE applies the LTE predicate on the "server_msg_seq" field.
func ServerMsgSeqLTE(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldServerMsgSeq), v))
	})
}

// MsgTypeEQ applies the EQ predicate on the "msg_type" field.
func MsgTypeEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgType), v))
	})
}

// MsgTypeNEQ applies the NEQ predicate on the "msg_type" field.
func MsgTypeNEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMsgType), v))
	})
}

// MsgTypeIn applies the In predicate on the "msg_type" field.
func MsgTypeIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMsgType), v...))
	})
}

// MsgTypeNotIn applies the NotIn predicate on the "msg_type" field.
func MsgTypeNotIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMsgType), v...))
	})
}

// MsgTypeGT applies the GT predicate on the "msg_type" field.
func MsgTypeGT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMsgType), v))
	})
}

// MsgTypeGTE applies the GTE predicate on the "msg_type" field.
func MsgTypeGTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMsgType), v))
	})
}

// MsgTypeLT applies the LT predicate on the "msg_type" field.
func MsgTypeLT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMsgType), v))
	})
}

// MsgTypeLTE applies the LTE predicate on the "msg_type" field.
func MsgTypeLTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMsgType), v))
	})
}

// MsgDataEQ applies the EQ predicate on the "msg_data" field.
func MsgDataEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgData), v))
	})
}

// MsgDataNEQ applies the NEQ predicate on the "msg_data" field.
func MsgDataNEQ(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMsgData), v))
	})
}

// MsgDataIn applies the In predicate on the "msg_data" field.
func MsgDataIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMsgData), v...))
	})
}

// MsgDataNotIn applies the NotIn predicate on the "msg_data" field.
func MsgDataNotIn(vs ...string) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMsgData), v...))
	})
}

// MsgDataGT applies the GT predicate on the "msg_data" field.
func MsgDataGT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMsgData), v))
	})
}

// MsgDataGTE applies the GTE predicate on the "msg_data" field.
func MsgDataGTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMsgData), v))
	})
}

// MsgDataLT applies the LT predicate on the "msg_data" field.
func MsgDataLT(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMsgData), v))
	})
}

// MsgDataLTE applies the LTE predicate on the "msg_data" field.
func MsgDataLTE(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMsgData), v))
	})
}

// MsgDataContains applies the Contains predicate on the "msg_data" field.
func MsgDataContains(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMsgData), v))
	})
}

// MsgDataHasPrefix applies the HasPrefix predicate on the "msg_data" field.
func MsgDataHasPrefix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMsgData), v))
	})
}

// MsgDataHasSuffix applies the HasSuffix predicate on the "msg_data" field.
func MsgDataHasSuffix(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMsgData), v))
	})
}

// MsgDataEqualFold applies the EqualFold predicate on the "msg_data" field.
func MsgDataEqualFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMsgData), v))
	})
}

// MsgDataContainsFold applies the ContainsFold predicate on the "msg_data" field.
func MsgDataContainsFold(v string) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMsgData), v))
	})
}

// MsgResCodeEQ applies the EQ predicate on the "msg_res_code" field.
func MsgResCodeEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgResCode), v))
	})
}

// MsgResCodeNEQ applies the NEQ predicate on the "msg_res_code" field.
func MsgResCodeNEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMsgResCode), v))
	})
}

// MsgResCodeIn applies the In predicate on the "msg_res_code" field.
func MsgResCodeIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMsgResCode), v...))
	})
}

// MsgResCodeNotIn applies the NotIn predicate on the "msg_res_code" field.
func MsgResCodeNotIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMsgResCode), v...))
	})
}

// MsgResCodeGT applies the GT predicate on the "msg_res_code" field.
func MsgResCodeGT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMsgResCode), v))
	})
}

// MsgResCodeGTE applies the GTE predicate on the "msg_res_code" field.
func MsgResCodeGTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMsgResCode), v))
	})
}

// MsgResCodeLT applies the LT predicate on the "msg_res_code" field.
func MsgResCodeLT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMsgResCode), v))
	})
}

// MsgResCodeLTE applies the LTE predicate on the "msg_res_code" field.
func MsgResCodeLTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMsgResCode), v))
	})
}

// MsgFeatureEQ applies the EQ predicate on the "msg_feature" field.
func MsgFeatureEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgFeature), v))
	})
}

// MsgFeatureNEQ applies the NEQ predicate on the "msg_feature" field.
func MsgFeatureNEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMsgFeature), v))
	})
}

// MsgFeatureIn applies the In predicate on the "msg_feature" field.
func MsgFeatureIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMsgFeature), v...))
	})
}

// MsgFeatureNotIn applies the NotIn predicate on the "msg_feature" field.
func MsgFeatureNotIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMsgFeature), v...))
	})
}

// MsgFeatureGT applies the GT predicate on the "msg_feature" field.
func MsgFeatureGT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMsgFeature), v))
	})
}

// MsgFeatureGTE applies the GTE predicate on the "msg_feature" field.
func MsgFeatureGTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMsgFeature), v))
	})
}

// MsgFeatureLT applies the LT predicate on the "msg_feature" field.
func MsgFeatureLT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMsgFeature), v))
	})
}

// MsgFeatureLTE applies the LTE predicate on the "msg_feature" field.
func MsgFeatureLTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMsgFeature), v))
	})
}

// MsgStatusEQ applies the EQ predicate on the "msg_status" field.
func MsgStatusEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMsgStatus), v))
	})
}

// MsgStatusNEQ applies the NEQ predicate on the "msg_status" field.
func MsgStatusNEQ(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMsgStatus), v))
	})
}

// MsgStatusIn applies the In predicate on the "msg_status" field.
func MsgStatusIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMsgStatus), v...))
	})
}

// MsgStatusNotIn applies the NotIn predicate on the "msg_status" field.
func MsgStatusNotIn(vs ...int8) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMsgStatus), v...))
	})
}

// MsgStatusGT applies the GT predicate on the "msg_status" field.
func MsgStatusGT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMsgStatus), v))
	})
}

// MsgStatusGTE applies the GTE predicate on the "msg_status" field.
func MsgStatusGTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMsgStatus), v))
	})
}

// MsgStatusLT applies the LT predicate on the "msg_status" field.
func MsgStatusLT(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMsgStatus), v))
	})
}

// MsgStatusLTE applies the LTE predicate on the "msg_status" field.
func MsgStatusLTE(v int8) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMsgStatus), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.Message {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		p(s.Not())
	})
}
