package utli

import (
	"fmt"
	"path"

	foUtli "github.com/wei0831/fo/utli"
)

// Move TODO
func Move(dir, to, fanhao, number, exclude string, wet bool) {
	// Not provide? Same as dir
	if to == "" {
		to = dir
	}

	toFind := fmt.Sprintf(`.*%s[-|_]*(%s).*(\..*)`, regexCaseIns(fanhao), number)

	foUtli.Movematches(dir, path.Join(to, fanhao), toFind, exclude, foUtli.FileAndFolder, wet)
}
