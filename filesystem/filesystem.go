package filesystem

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func MakeDirIfNotExists(filePath string) (err error) {
	dir := filepath.Dir(filePath)
	dirInfo, err := os.Stat(dir)
	if dirInfo == nil || os.IsNotExist(err) {
		logrus.WithField("_func_", "MakeDirIfNotExists").WithField("filePath", filePath).Infoln()
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "MakeDirIfNotExists failed "+filePath)
		}
	}

	return nil
}

func NewFile(filePath string, data []byte) (err error) {
	logrus.WithField("_func_", "NewFile").WithField("filePath", filePath).Infoln()
	err = MakeDirIfNotExists(filePath)
	if err != nil {
		return errors.Wrap(err, "NewFile failed "+filePath)
	}

	err = ioutil.WriteFile(filePath, data, os.ModePerm|os.ModeExclusive)
	if err != nil {
		return errors.Wrap(err, "NewFile failed "+filePath)
	}

	return nil
}

func AppendFile(filePath string, data []byte) (err error) {
	logrus.WithField("_func_", "AppendFile").WithField("filePath", filePath).Infoln()
	err = MakeDirIfNotExists(filePath)
	if err != nil {
		return errors.Wrap(err, "AppendFile failed "+filePath)
	}

	err = ioutil.WriteFile(filePath, data, os.ModeAppend|os.ModePerm|os.ModeExclusive)
	if err != nil {
		return errors.Wrap(err, "AppendFile failed "+filePath)
	}

	return nil
}
