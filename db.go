package main

import (
	"io"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	"path/filepath"
	"time"
)

const PREFIX = "godm"

// File stores the file information from the GridFS file collection.
// The struct definition was copied from the mgo driver (type gfsFile
// in gridfs.go). It's not clear why this type was not exported.
type File struct {
	Id          interface{} "_id"
	ChunkSize   int         "chunkSize"
	UploadDate  time.Time   "uploadDate"
	Length      int64       ",minsize"
	MD5         string
	Filename    string    ",omitempty"
	ContentType string    "contentType,omitempty"
	Metadata    *bson.Raw ",omitempty"
}

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
	g := session.DB(d.Database).GridFS(PREFIX)
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
	g := session.DB(d.Database).GridFS(PREFIX)
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

// Delete removes all files with the specified name from the database.
func (d *Db) Delete(file string) error {
	session, err := mgo.Dial(d.Address)
	if err != nil {
		return err
	}
	defer session.Close()
	g := session.DB(d.Database).GridFS(PREFIX)
	err = g.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

// List returns a slice containing all file descriptors matching the provided
// BSON query.
func (d *Db) List(query interface{}) ([]File, error) {
	session, err := mgo.Dial(d.Address)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	g := session.DB(d.Database).GridFS(PREFIX)
	var result []File
	err = g.Find(query).All(&result)
	return result, err
}
