package macsync

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/docker/engine-api/client"
	"github.com/rancher/go-rancher-metadata/metadata"
)

// Some of the tests can run only when in development,
// remember to disable this before commiting the code.
const inDevelopment = false

func TestDoSync(t *testing.T) {
	if !inDevelopment {
		t.Skip("not in development mode")
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugf("TestDoSync")

	mc, err := metadata.NewClientAndWait("http://169.254.169.250/2016-07-29")
	if err != nil {
		logrus.Errorf("error creating metadata client")
		t.Fail()
	}
	dClient, err := client.NewEnvClient()
	if err != nil {
		logrus.Errorf("err=%v", err)
		t.Fail()
	}
	ms := MACSyncer{
		mc: mc,
		dc: dClient,
	}

	ms.doSync()
}
