package cmd

import (
	"fmt"
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
	wordListPath    string
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
	piceCmd.PersistentFlags().StringVarP(&wordListPath, "wordList", "l", "wordlists/eff_large.wordlist", "specify the path of the word list to use")
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
	if err := picegen.LoadWordList(wordListPath); err != nil {
		output.Errorln(err)
		os.Exit(1)
	}
	result, err := picegen.Generate(wordCount, options)
	if err != nil {
		output.Errorln(err)
		os.Exit(1)
	}
	output.Success(fmt.Sprintf("Your passphrase is: %s", result))
}

func Execute() {
	if err := piceCmd.Execute(); err != nil {
		output.Errorln(err)
		os.Exit(1)
	}
}
