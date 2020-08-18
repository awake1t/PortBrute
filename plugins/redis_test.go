package plugins_test

import (
	"asset-scan/plugins"

	"testing"
)

func TestScanRedis(t *testing.T) {
	t.Log(plugins.ScanRedis("127.0.0.1", "6379", "", ""))
}
