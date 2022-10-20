// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package hoge

import (
	"time"

	"entgo.io/ent/examples/start/ent/util/ulid"
)

const (
	// Label holds the string label denoting the hoge type in the database.
	Label = "hoge"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeHogeAdministrators holds the string denoting the hoge_administrators edge name in mutations.
	EdgeHogeAdministrators = "hoge_administrators"
	// Table holds the table name of the hoge in the database.
	Table = "hoges"
	// HogeAdministratorsTable is the table that holds the hoge_administrators relation/edge.
	HogeAdministratorsTable = "hoge_administrators"
	// HogeAdministratorsInverseTable is the table name for the HogeAdministrator entity.
	// It exists in this package in order to avoid circular dependency with the "hogeadministrator" package.
	HogeAdministratorsInverseTable = "hoge_administrators"
	// HogeAdministratorsColumn is the table column denoting the hoge_administrators relation/edge.
	HogeAdministratorsColumn = "hoge_id"
)

// Columns holds all SQL columns for hoge fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)