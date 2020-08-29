package plugins_test

import (

	"testing"
)

func TestScanSsh(t *testing.T) {
	t.Log(plugins.ScanSsh("127.0.0.1", "22", "root", "123456"))
}
