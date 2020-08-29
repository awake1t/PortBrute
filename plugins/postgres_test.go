package plugins_test

import (
	"testing"
)

func TestScanPostgres(t *testing.T) {
	t.Log(plugins.ScanPostgres("127.0.0.1", "5432", "postgres", "123456"))
}
