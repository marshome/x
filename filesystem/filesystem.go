package filesystem

import (
	"path/filepath"
	"os"
	"io/ioutil"
)

func MakeDirIfNotExists(filePath string)(err error) {
	dir := filepath.Dir(filePath)
	dirInfo, err := os.Stat(dir)
	if dirInfo == nil || os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewFile(filePath string, data []byte) (err error) {
	err = MakeDirIfNotExists(filePath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, os.ModePerm | os.ModeExclusive)
	if err != nil {
		return err
	}

	return nil
}

func AppendFile(filePath string,data[] byte)(err error) {
	err = MakeDirIfNotExists(filePath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, data, os.ModeAppend | os.ModePerm | os.ModeExclusive)
	if err != nil {
		return err
	}

	return nil
}