package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	var listVersionsFlag bool
	var rootCmd = &cobra.Command{
		Args:  matchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
		Use:   "pm [package] [version]",
		Short: "Cub Package Manager",
		Long: `pm is the Cub Package Manager prototype. It tries to mimic the behavior
of existing JS package managers such as npm, yarn and bun's built-in package manager`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Ordered JSON map
			// so the latest version is at the bottom and colored
			if len(args) <= 1 && listVersionsFlag {
				packageInfo := getVersions(args[0])
				for _, versions := range packageInfo {
					fmt.Println(versions.Version)
				}
			}
			if len(args) >= 2 {
				resp := fetchPackage(args[0], args[1])
				buffer := bytes.NewBuffer(resp)
				ExtractTarGz(buffer, args[0])
			}
		},
	}

	rootCmd.Flags().BoolVarP(&listVersionsFlag, "version", "v", false, "List all available package versions")

	var freshInstallFlag bool
	var installCmd = &cobra.Command{
		Use:     "install",
		Aliases: []string{"i"},
		Short:   "Install all dependencies listed inside package.json",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if freshInstallFlag {
				fmt.Println("Removing node_modules. This might take a while")
				err := os.RemoveAll("node_modules")
				if err != nil {
					log.Fatal(err)
				}
			}
			install()
		},
	}

	installCmd.Flags().BoolVar(&freshInstallFlag, "fresh", false, "Install all node_modules from scratch")

	rootCmd.AddCommand(installCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing pm '%s'", err)
		os.Exit(1)
	}
}
