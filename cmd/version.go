// Package cmd is the command surface of learnymir cli tool provided by kubuskotak.
// # This manifest was generated by ymir. DO NOT EDIT.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/irhamsahbana/learn-ymir/pkg/version"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   `version`,
		Short: "Print version info",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Service version: %s\n", version.GetVersion().VersionNumber())
			fmt.Printf("Golang version: %s\n", version.GoVersion)
			fmt.Printf("Git commit hash: %s\n", version.GetVersion().Revision)
			fmt.Printf("Built on: %s\n", version.GetVersion().BuildDate)
			fmt.Printf("Built by: %s\n", version.BuildUser)
			return nil
		},
	}
}
