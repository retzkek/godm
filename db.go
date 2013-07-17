package godm

import (
	"io"
	"labix.org/v2/mgo"
	"os"
	"path/filepath"
)

type Db struct {
	Address  string
	Database string
}

// Store puts the specified file into the database.
func (d *Db) Store(file string) error {
	session, err := mgo.Dial(d.Address)
	if err != nil {
		return err
	}
	defer session.Close()
	g := session.DB(d.Database).GridFS("godm")
	f, err := g.Create(filepath.Base(file))
	if err != nil {
		return err
	}
	source, err := os.Open(file)
	if err != nil {
		return err
	}
	defer source.Close()
	_, err = io.Copy(f, source)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves the specified file from the database and saves it in the
// current working directory.
func (d *Db) Get(file string) error {
	session, err := mgo.Dial(d.Address)
	if err != nil {
		return err
	}
	defer session.Close()
	g := session.DB(d.Database).GridFS("godm")
	f, err := g.Open(file)
	if err != nil {
		return err
	}
	dest, err := os.Create(file)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, f)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
