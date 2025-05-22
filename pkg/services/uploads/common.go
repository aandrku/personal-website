package uploads

import (
	"io"
	"log"
	"os"
)

const (
	UploadsDir = "./data/uploads/"
)

func init() {
	err := os.Mkdir(UploadsDir, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Unable to create %s directory %v", UploadsDir, err)
	}
}

func Get() ([]os.FileInfo, error) {
	var infos []os.FileInfo

	files, err := os.ReadDir(UploadsDir)
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
	f, err := os.Create(UploadsDir + name)
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
	err := os.Remove(UploadsDir + name)
	if err != nil {
		return err
	}

	return nil
}
