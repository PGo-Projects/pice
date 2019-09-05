package picegen

import (
	"bufio"
	"crypto/rand"
	"errors"
	"math/big"
	"os"
	"regexp"
	"strings"
)

type Options struct {
	UseCaps          bool
	UseNumbers       bool
	UsePunctunations bool
}

var (
	dicewareNumberAndWordRegex = regexp.MustCompile(`(\d+)\s+(.+)`)

	wordList      []string
	punctunations = []string{"~", "!", "#", "$", "%", "^", "&", "*", "(", ")", "-", "=", "+", "[", "]", "\\", "{", "}", ":", ";", "\"", "'", "<", ">", "?", "/"}

	ErrPunctunationGen = errors.New("Was unable to generate a punctunation")
	ErrWordGen         = errors.New("Was unable to generate a word")
	ErrPassphraseGen   = errors.New("Was unable to generate a passphrase")
	ErrWordListFormat  = errors.New("The word list is improperly formatted")
)

func LoadWordList(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	return nil
}

func Generate(wordCount uint64, options Options) (string, string, error) {
	var components []string
	shouldCapitalize := false
	for i := uint64(0); i < wordCount; i++ {
		word, err := GenerateWord()
		if err != nil {
			return "", "", err
		}

		if shouldCapitalize {
			components = append(components, strings.Title(word))
			shouldCapitalize = false
		} else {
			components = append(components, word)
		}

		n, err := rand.Int(rand.Reader, big.NewInt(6))
		if err != nil {
			return "", "", err
		}

		switch n.Int64() {
		case 0:
			if options.UseNumbers {
				number, err := GenerateNumber()
				if err != nil {
					return "", "", err
				}
				components = append(components, number)
			}
		case 1:
			if options.UsePunctunations {
				punctunation, err := GeneratePunctunation()
				if err != nil {
					return "", "", err
				}
				components = append(components, punctunation)
			}
		case 2:
			shouldCapitalize = options.UseCaps
		default:
			if !n.IsInt64() {
				return "", "", ErrPassphraseGen
			}
		}
	}

	componentsString := strings.Join(components, " ")
	passphraseString := strings.Join(components, "")

	return componentsString, passphraseString, nil
}

func GenerateWord() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(7776))
	if err != nil {
		return "", err
	}
	if n.IsInt64() {
		_, word, err := getDicewareNumAndWord(wordList[n.Int64()])
		return word, err
	}
	return "", ErrWordGen
}

func GenerateNumber() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		return "", err
	}
	return n.String(), nil
}

func GeneratePunctunation() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(punctunations))))
	if err != nil {
		return "", err
	}
	if n.IsInt64() {
		return punctunations[n.Int64()], nil
	}
	return "", ErrPunctunationGen
}

func getDicewareNumAndWord(entry string) (string, string, error) {
	matches := dicewareNumberAndWordRegex.FindStringSubmatch(entry)
	if len(matches) != 3 {
		return "", "", ErrWordListFormat
	}
	return matches[1], matches[2], nil
}
