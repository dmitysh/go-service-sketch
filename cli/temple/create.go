package temple

import (
	"os"
	"path"
	"text/template"
)

func createFile(folderPath, filename, templatePath string, templateData map[string]string) error {
	err := os.MkdirAll(folderPath, defaultFolderPerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	f, err := os.Create(path.Join(folderPath, filename))
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl, err := template.ParseFS(templatesFolder, templatePath)
	if err != nil {
		return err
	}

	if err = tmpl.Execute(f, templateData); err != nil {
		return err
	}

	return nil
}
