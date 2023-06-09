// Package shared is library function for whole system.
// # This manifest was generated by ymir. DO NOT EDIT.
package shared

import "entgo.io/ent/dialect"

// Dialect is the database dialect converter.
func Dialect(d string) string {
	switch d {
	case dialect.Postgres:
		return dialect.Postgres
	case dialect.MySQL:
		return dialect.MySQL
	default:
		return dialect.SQLite
	}
}
