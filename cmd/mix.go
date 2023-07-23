package cmd

import (
	"github.com/samber/lo"
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
	Use:   "mix",
	Short: "make a cocktail.",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 1 && args[0] == "Vodka" {
			return []string{"Orange Juice"}, cobra.ShellCompDirectiveNoFileComp
		}
		_, unusedIngredients := lo.Difference(args, availableIngredients)
		return unusedIngredients, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mixCmd)
}
