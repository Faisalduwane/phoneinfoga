package remote

import (
	"github.com/sundowndev/phoneinfoga/v2/lib/remote/suppliers"
	"net/http"
)

func InitScanners(remote *Library) {
	numverifySupplier := suppliers.NewNumverifySupplier()
	ovhSupplier := suppliers.NewOVHSupplier()

	remote.AddScanner(NewLocalScanner())
	remote.AddScanner(NewNumverifyScanner(numverifySupplier))
	remote.AddScanner(NewGoogleSearchScanner())
	remote.AddScanner(NewOVHScanner(ovhSupplier))
	remote.AddScanner(NewGoogleCSEScanner(&http.Client{}))

	remote.LoadPlugins()
}
