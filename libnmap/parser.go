package libnmap

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

// Parse is used to parse and return a structured object of Nmap output
func Parse(fp string) (Nmaprun, error) {
	result := Nmaprun{}

	data, err := ioutil.ReadFile(fp)
	if err != nil {
		logrus.WithError(err).Errorln(fmt.Sprintf("Failed to open file at path %s", fp))
		return result, err
	}

	err = xml.Unmarshal(data, &result)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to parse Nmap XML")
		return result, err
	}
	return result, nil
}
