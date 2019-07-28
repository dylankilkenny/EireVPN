package test

import (
	"eirevpn/api/router"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDB()
	r = router.SetupRouter(false)
	code := m.Run()
	os.Exit(code)
}
