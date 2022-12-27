// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"os"

	"github.com/ava-labs/avalanchego/utils/formatting"
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
	config, err := ParseFlags(flags, args)
	if err != nil {
		return err
	}

	genesisBytes, err := genesis.Codec.Marshal(genesis.Version, config.Genesis)
	if err != nil {
		return err
	}

	if config.Encoding == binaryEncoding {
		_, err = os.Stdout.Write(genesisBytes)
		return err
	}

	// hex encoded
	encoded, err := formatting.Encode(formatting.Hex, genesisBytes)
	if err != nil {
		return err
	}
	_, err = os.Stdout.WriteString(encoded)

	return err
}
