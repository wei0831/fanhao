package utli

import (
	"fmt"

	foUtli "github.com/wei0831/fo/utli"
)

// Rename TODO
func Rename(dir, to, fanhao, sperator, number, exclude string, wet bool) {
	toFind := fmt.Sprintf(`.*%s[-|_]*(%s).*(\..*)`, regexCaseIns(fanhao), number)
	toReplace := fmt.Sprintf("%s%s$1$2", fanhao, sperator)

	foUtli.ReplaceName(dir, to, toFind, toReplace, exclude, foUtli.FileOnly, wet)
}
