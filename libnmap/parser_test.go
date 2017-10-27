package libnmap

import (
	"os"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
)

// TestParse is used to test the parsing of the NMAP output
func TestParse(t *testing.T) {
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
