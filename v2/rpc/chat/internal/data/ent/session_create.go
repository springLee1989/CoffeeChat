// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat/internal/data/ent/session"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetCreated sets the "created" field.
func (sc *SessionCreate) SetCreated(t time.Time) *SessionCreate {
	sc.mutation.SetCreated(t)
	return sc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (sc *SessionCreate) SetNillableCreated(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetCreated(*t)
	}
	return sc
}

// SetUpdated sets the "updated" field.
func (sc *SessionCreate) SetUpdated(t time.Time) *SessionCreate {
	sc.mutation.SetUpdated(t)
	return sc
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (sc *SessionCreate) SetNillableUpdated(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetUpdated(*t)
	}
	return sc
}

// SetUserID sets the "user_id" field.
func (sc *SessionCreate) SetUserID(s string) *SessionCreate {
	sc.mutation.SetUserID(s)
	return sc
}

// SetPeerID sets the "peer_id" field.
func (sc *SessionCreate) SetPeerID(s string) *SessionCreate {
	sc.mutation.SetPeerID(s)
	return sc
}

// SetSessionType sets the "session_type" field.
func (sc *SessionCreate) SetSessionType(i int8) *SessionCreate {
	sc.mutation.SetSessionType(i)
	return sc
}

// SetSessionStatus sets the "session_status" field.
func (sc *SessionCreate) SetSessionStatus(i int8) *SessionCreate {
	sc.mutation.SetSessionStatus(i)
	return sc
}

// SetID sets the "id" field.
func (sc *SessionCreate) SetID(i int32) *SessionCreate {
	sc.mutation.SetID(i)
	return sc
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	var (
		err  error
		node *Session
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Session)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SessionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() {
	if _, ok := sc.mutation.Created(); !ok {
		v := session.DefaultCreated
		sc.mutation.SetCreated(v)
	}
	if _, ok := sc.mutation.Updated(); !ok {
		v := session.DefaultUpdated
		sc.mutation.SetUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "Session.created"`)}
	}
	if _, ok := sc.mutation.Updated(); !ok {
		return &ValidationError{Name: "updated", err: errors.New(`ent: missing required field "Session.updated"`)}
	}
	if _, ok := sc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Session.user_id"`)}
	}
	if v, ok := sc.mutation.UserID(); ok {
		if err := session.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Session.user_id": %w`, err)}
		}
	}
	if _, ok := sc.mutation.PeerID(); !ok {
		return &ValidationError{Name: "peer_id", err: errors.New(`ent: missing required field "Session.peer_id"`)}
	}
	if v, ok := sc.mutation.PeerID(); ok {
		if err := session.PeerIDValidator(v); err != nil {
			return &ValidationError{Name: "peer_id", err: fmt.Errorf(`ent: validator failed for field "Session.peer_id": %w`, err)}
		}
	}
	if _, ok := sc.mutation.SessionType(); !ok {
		return &ValidationError{Name: "session_type", err: errors.New(`ent: missing required field "Session.session_type"`)}
	}
	if _, ok := sc.mutation.SessionStatus(); !ok {
		return &ValidationError{Name: "session_status", err: errors.New(`ent: missing required field "Session.session_status"`)}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: session.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: session.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := sc.mutation.Updated(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldUpdated,
		})
		_node.Updated = value
	}
	if value, ok := sc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := sc.mutation.PeerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: session.FieldPeerID,
		})
		_node.PeerID = value
	}
	if value, ok := sc.mutation.SessionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: session.FieldSessionType,
		})
		_node.SessionType = value
	}
	if value, ok := sc.mutation.SessionStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: session.FieldSessionStatus,
		})
		_node.SessionStatus = value
	}
	return _node, _spec
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
