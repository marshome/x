package jsonhelper

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func SimpleJson(o interface{}) string {
	if o == nil {
		return ""
	}

	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		logrus.Warnln(errors.Wrap(err, "SimpleJson"))
	}

	return string(data)
}
