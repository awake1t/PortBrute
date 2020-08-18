package plugins_test

import (
	"asset-scan/plugins"
	"testing"
)

func TestScanFtp(t *testing.T) {
	t.Log(plugins.ScanFtp("127.0.0.1", "21", "root", "123456"))
}
