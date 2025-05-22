package uploads

import (
	"io"
	"log"
	"os"
)

const (
	uploadsDir = "./data/uploads/"
)

func init() {
	err := os.Mkdir(uploadsDir, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Unable to create %s directory %v", uploadsDir, err)
	}
}

func Get() ([]os.FileInfo, error) {
	var infos []os.FileInfo

	files, err := os.ReadDir(uploadsDir)
	if err != nil {
		return infos, err
	}

	for _, f := range files {
		info, err := f.Info()
		if err != nil {
			return infos, err
		}

		infos = append(infos, info)
	}

	return infos, nil
}

func New(r io.Reader, name string) error {
	f, err := os.Create(uploadsDir + name)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, r)
	if err != nil {
		return err
	}
	return nil

}

func Remove(name string) error {
	err := os.Remove(uploadsDir + name)
	if err != nil {
		return err
	}

	return nil
}
