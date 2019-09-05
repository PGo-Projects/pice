package cmd

import (
	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/pice/internal/picegen"
	"github.com/spf13/cobra"
	"os"
)

var (
	wordCount       uint64
	usePunctunation bool
	useNumbers      bool
	useCaps         bool
)

var piceCmd = &cobra.Command{
	Use:   "pice",
	Short: "Pice is an offline diceware passphrase generator",
	Long: `An offline diceware passphrase generator that offers more flexibility
    when incorporating punctunations and numbers.`,
	Run: pice,
}

func init() {
	piceCmd.PersistentFlags().Uint64VarP(&wordCount, "words", "w", 6, "number of words to use when generating the passphrase")
	piceCmd.PersistentFlags().BoolVarP(&usePunctunation, "punctunations", "p", false, "whether to use punctunations or not (default false)")
	piceCmd.PersistentFlags().BoolVarP(&useNumbers, "numbers", "n", false, "whether to use numbers or not (default false)")
	piceCmd.PersistentFlags().BoolVarP(&useCaps, "caps", "c", false, "whether to use capitialization or not (default false)")
}

func pice(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		if err := cmd.Help(); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}
	options := picegen.Options{
		UseCaps:          useCaps,
		UseNumbers:       useNumbers,
		UsePunctunations: usePunctunation,
	}
	if err := picegen.Generate(wordCount, options); err != nil {
		output.Errorln(err)
		os.Exit(1)
	}
}

func Execute() {
	if err := piceCmd.Execute(); err != nil {
		output.Errorln(err)
		os.Exit(1)
	}
}
