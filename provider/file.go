package provider

import (
	"log"
	"os"

	"github.com/eliothedeman/img/util"
)

// File data proiver for a file
type File struct {
	Path   string
	ID     string
	Format string
	file   *os.File
}

// Read reads from the file
func (f *File) Read(p []byte) (int, error) {
	if f.file == nil {
		err := f.open()
		if err != nil {
			return 0, err
		}
	}
	n, err := f.Read(p)
	return n, err

}

// Write write to the file
func (f *File) Write(p []byte) (int, error) {
	if f.file == nil {
		err := f.open()
		if err != nil {
			return 0, err
		}
	}
	n, err := f.file.Write(p)
	log.Println(err)
	return n, err

}

// Close close the connection to the file
func (f *File) Close() error {
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}

// open opens a file on disk
func (f *File) open() error {
	var err error
	if f.file == nil {
		// check if the file exists
		// if it is not on disk, create a new one
		if _, err = os.Stat(f.Location()); os.IsNotExist(err) {
			err = f.create()
			if err != nil {
				log.Println(err.Error())
				return err
			}
		} else {
			// open the file
			f.file, err = os.Open(f.Location())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Location returns the path to fetch this data provider at on disk
func (f *File) Location() string {
	return f.Path + "/" + f.ID + "." + f.Format
}

func (f *File) Create(id, codec, size, prefix string) {
	f.ID = id
	f.Path = util.GenPath(id, codec, size, prefix)
	f.Format = codec
}

// create creates a new file on disk
func (f *File) create() error {
	var err error
	// check if there if the path exists, if not create it
	if _, err = os.Stat(f.Path); os.IsNotExist(err) {
		err := os.MkdirAll(f.Path, os.FileMode(0777))
		if err != nil {
			return err
		}
		log.Println("Created directory '" + f.Path + "'")
	}

	// if the file already exists, don't try to create it again
	if _, err = os.Stat(f.Path); !os.IsNotExist(err) {
		log.Println("File '" + f.Location() + "' already exists")
		return nil
	}
	f.file, err = os.Create(f.Location())
	log.Println("created " + f.Location())
	if err != nil {
		return err
	}
	log.Println("Created file '" + f.Location() + "'")
	return nil
}
