// Package ent provides an interface for interacting with
// ent (ent codegen) as a package rather than an executable.
// # This manifest was generated by ymir. DO NOT EDIT.
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration --target ./ ../schema
