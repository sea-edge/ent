// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"entgo.io/ent/entc/integration/privacy/ent/schema"
	"entgo.io/ent/entc/integration/privacy/ent/task"
	"entgo.io/ent/entc/integration/privacy/ent/team"
	"entgo.io/ent/entc/integration/privacy/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	taskMixin := schema.Task{}.Mixin()
	task.Policy = privacy.NewPolicies(taskMixin[0], taskMixin[1], schema.Task{})
	task.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := task.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	taskHooks := schema.Task{}.Hooks()

	task.Hooks[1] = taskHooks[0]
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescTitle is the schema descriptor for title field.
	taskDescTitle := taskFields[0].Descriptor()
	// task.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	task.TitleValidator = taskDescTitle.Validators[0].(func(string) error)
	teamMixin := schema.Team{}.Mixin()
	team.Policy = privacy.NewPolicies(teamMixin[0], schema.Team{})
	team.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := team.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescName is the schema descriptor for name field.
	teamDescName := teamFields[0].Descriptor()
	// team.NameValidator is a validator for the "name" field. It is called by the builders before save.
	team.NameValidator = teamDescName.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	user.Policy = privacy.NewPolicies(userMixin[0], userMixin[1], schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
}
