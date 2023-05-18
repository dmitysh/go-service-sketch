package tmplcreator

import (
	"embed"
	"os"
)

const (
	defaultFolderPerm os.FileMode = 0666
	defaultFilePerm   os.FileMode = 0666
)

//go:embed templates
var templatesFolder embed.FS

const (
	makefileTemplatePath  = "templates/makefile.txt"
	gitignoreTemplatePath = "templates/gitignore.txt"
)
