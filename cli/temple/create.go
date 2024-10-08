package temple

import (
	"go.uber.org/multierr"
	"os"
	"path"
	"text/template"
)

func createFolders(folderCreators []func() error) error {
	var err error

	for _, createFn := range folderCreators {
		err = multierr.Append(err, createFn())
	}

	return err
}

func createFiles(fileCreators []func() error) error {
	var err error

	for _, createFn := range fileCreators {
		err = multierr.Append(err, createFn())
	}

	return err
}

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
