//args: -Emisspell
//config_path: testdata/configs/misspell.yml
package testdata

func Misspell() {
	// comment with incorrect spelling: occurred // ERROR "`occurred` is a misspelling of `occurred`"
}

// the word language should be ignored here: it's set in config
