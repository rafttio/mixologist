package cmd

import (
	"github.com/spf13/cobra"
)

var availableIngredients = []string{
	"Vodka",
	"Gin",
	"Orange Juice",
	"Triple Sec",
	"Tequila",
	"simple syrup",
	"white rum",
	"Kahlua coffee liqueur",
}

var mixCmd = &cobra.Command{
	Use:       "mix",
	Short:     "make a cocktail.",
	ValidArgs: availableIngredients,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mixCmd)
}
