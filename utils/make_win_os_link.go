package utils

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// thanks https://stackoverflow.com/questions/32438204/create-a-windows-shortcut-lnk-in-go
func MakeWinOSLink(src string, dst string) error {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)

	oleShellObject, err := oleutil.CreateObject("WScript.Shell")

	if err != nil {
		return err
	}

	defer oleShellObject.Release()

	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)

	if err != nil {
		return err
	}

	defer wshell.Release()

	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)

	if err != nil {
		return err
	}

	idispatch := cs.ToIDispatch()

	oleutil.PutProperty(idispatch, "TargetPath", src)

	oleutil.CallMethod(idispatch, "Save")

	return nil
}
