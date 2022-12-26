// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"os"

	"github.com/ava-labs/xsvm/genesis"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	c := &cobra.Command{
		Use:   "genesis",
		Short: "Creates a chain's genesis and prints it to stdout",
		RunE:  genesisFunc,
	}
	flags := c.Flags()
	AddFlags(flags)
	return c
}

func genesisFunc(c *cobra.Command, args []string) error {
	flags := c.Flags()
	g, err := ParseFlags(flags, args)
	if err != nil {
		return err
	}

	genesisBytes, err := genesis.Codec.Marshal(genesis.Version, g)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(genesisBytes)
	return err
}
