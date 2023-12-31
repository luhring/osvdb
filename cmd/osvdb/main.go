package main

import (
	"fmt"
	"github.com/luhring/osvdb"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "osvdb",
		Short: "osvdb is a command line tool for building and querying a database of Open Source Vulnerability (OSV) data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(buildCmd())

	return cmd
}

func buildCmd() *cobra.Command {
	cfg := buildCmdConfig{}

	cmd := &cobra.Command{
		Use:           "build",
		Short:         "build the OSV database",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			paths := args

			var inputs []osvdb.Input
			for _, path := range paths {
				fi, err := os.Stat(path)
				if err != nil {
					return fmt.Errorf("failed to stat path: %w", err)
				}

				if fi.IsDir() {
					input := osvdb.NewFSInput(os.DirFS(path), cfg.recursive)
					inputs = append(inputs, input)
					continue
				}

				input := osvdb.NewFileInput(path)
				inputs = append(inputs, input)
			}

			build := osvdb.Build{
				OutputDatabaseLocation: cfg.outputDatabaseLocation,
				OverwriteDatabase:      cfg.overwriteDatabase,
			}

			err := build.Do(cmd.Context(), inputs...)
			if err != nil {
				return fmt.Errorf("failed to build database: %w", err)
			}

			return nil
		},
	}

	cfg.addToCommand(cmd)

	return cmd
}

type buildCmdConfig struct {
	outputDatabaseLocation string
	overwriteDatabase      bool
	recursive              bool
}

func (cfg *buildCmdConfig) addToCommand(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&cfg.outputDatabaseLocation, "output", "o", "", "output database location")
	cmd.Flags().BoolVar(&cfg.overwriteDatabase, "overwrite", false, "overwrite database if it already exists")
	cmd.Flags().BoolVarP(&cfg.recursive, "recursive", "R", false, "recursive into directories to find JSON files")
}
