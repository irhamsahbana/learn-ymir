// Code generated by ent, DO NOT EDIT.

package ymir

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the ymir type in the database.
	Label = "ymir"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// Table holds the table name of the ymir in the database.
	Table = "ymirs"
)

// Columns holds all SQL columns for ymir fields.
var Columns = []string{
	FieldID,
	FieldVersion,
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
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion string
)

// OrderOption defines the ordering options for the Ymir queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}
