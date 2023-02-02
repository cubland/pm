package main

import (
	"bytes"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Args:  matchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Use:   "pm [package] [version]",
	Short: "Cub Package Manager",
	Long: `pm is the Cub Package Manager prototype. It tries to mimic the behavior
of existing JS package managers such as npm, yarn and bun's built-in package manager`,
	Run: func(cmd *cobra.Command, args []string) {
		resp := fetchPackage(args[0], args[1])
		buffer := bytes.NewBuffer(resp)
		ExtractTarGz(buffer, args[0])
	},
}

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "Install all dependencies listed inside package.json",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}
