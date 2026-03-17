package dlg

import (
	"runtime"

	"github.com/thingscompton/gormt/public/myi18n"
)

// WinMain windows main loop
func WinMain() {
	if runtime.GOOS == "windows" {
		myi18n.SetLocalLG("en")
	}
	OnInitDialog()
}
