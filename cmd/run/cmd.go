// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package run

import (
	"github.com/spf13/cobra"

	"github.com/ava-labs/avalanchego/vms/rpcchainvm"

	"github.com/ava-labs/xsvm"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Runs an XSVM plugin",
		RunE:  runFunc,
	}
}

func runFunc(*cobra.Command, []string) error {
	rpcchainvm.Serve(&xsvm.VM{})
	return nil
}
