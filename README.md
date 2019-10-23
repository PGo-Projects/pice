# pice

An offline diceware passphrase generator that offers more flexibility when incorporating
punctunations and numbers.

## Usage

On Unix systems, run as:

```bash
pice [flags]
```

On Windows systems, run as:

```bash
pice.exe [flags]
```

Available flags:

- `-c, --caps`: whether to use capitialization or not (defaults to false)
- `-h, --help`: get help on pice
- `-n, --numbers`: whether to use numbers or not (defaults to false)
- `-p, --punctunations`: whether to use punctunations or not (defaults to false)
- `-l, --wordList string`: specify the path of the word list to use (defaults to
  "wordlists/eff_large.wordlist")
- `-w, --words int`: number of words to use when generating the passphrase (defaults to 6)


Examples:

```bash
# generate a 10 word passphrase with capitialization and numbers
pice -c -n -w 10

# generates a 6 word passphrase with wordlists/eff_short.wordlist
# (note you have to add the wordlist yourself)
pice -l "wordlists/eff_short.wordlist"
```
