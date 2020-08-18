package plugins_test

import (
	"asset-scan/plugins"
	"testing"
)

func TestScanMongodb(t *testing.T) {
	t.Log(plugins.ScanMongodb("127.0.0.1", "27017", "root", "123456"))
}
