package temple

import (
	"os"
)

type DBCreator struct {
	fileCreators   []func() error
	folderCreators []func() error
}

func NewDBCreator() *DBCreator {
	d := &DBCreator{}

	d.fileCreators = []func() error{
		d.createPGXFile,
		d.createTransactionFile,
	}

	d.folderCreators = []func() error{
		d.createMigrationsFolder,
	}

	return d
}

func (d *DBCreator) Create() error {
	if err := createFolders(d.folderCreators); err != nil {
		return err
	}

	if err := createFiles(d.fileCreators); err != nil {
		return err
	}

	return nil
}

func (d *DBCreator) createMigrationsFolder() error {
	err := os.MkdirAll("migrations", defaultFolderPerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

func (d *DBCreator) createPGXFile() error {
	return createFile(houstonPath("dobby"), "pgx.go", houstonDobbyPGXPath,
		map[string]string{},
	)
}

func (d *DBCreator) createTransactionFile() error {
	return createFile(houstonPath("dobby"), "transaction.go", houstonDobbyTransactionPath,
		map[string]string{},
	)
}
