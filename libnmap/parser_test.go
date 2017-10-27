package libnmap

import (
	"os"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
)

// TestParse is used to test the parsing of the NMAP output
func TestParseValid(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		logrus.WithError(err).Errorln("Failed to get current working directory")
		t.FailNow()
	}

	n, err := Parse(path.Join(wd, "fixtures/nmap_output.xml"))
	if err != nil {
		logrus.WithError(err).Errorln("Failed to load Nmap output")
		t.FailNow()
	}

	if n.Host.Ports.Port[0].AttrPortid != "80" {
		logrus.Errorln("Port number was not expected")
		t.FailNow()
	}
}

// TestParse is used to test the parsing of the NMAP output
func TestParseInvalidXML(t *testing.T) {
	logrus.SetLevel(logrus.FatalLevel)
	wd, err := os.Getwd()
	logrus.SetLevel(logrus.ErrorLevel)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to get current working directory")
		t.FailNow()
	}

	logrus.SetLevel(logrus.FatalLevel)
	_, err = Parse(path.Join(wd, "fixtures/invalid.xml"))
	logrus.SetLevel(logrus.ErrorLevel)
	if err == nil {
		logrus.WithError(err).Errorln("Loaded invalid XML file successfully")
		t.FailNow()
	}
}

// TestParse is used to test the parsing of the NMAP output
func TestParseFailToLoadFile(t *testing.T) {
	logrus.SetLevel(logrus.FatalLevel)
	_, err := Parse(path.Join("doesntexist"))
	logrus.SetLevel(logrus.ErrorLevel)
	if err == nil {
		logrus.WithError(err).Errorln("Loaded non-existant file successfully")
		t.FailNow()
	}
}
