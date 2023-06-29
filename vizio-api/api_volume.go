package vizio_api

type VolumeAPI struct {
	URL   string
	TOKEN string
}

func (a *VolumeAPI) VolumeUp() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 5, 1, "KEYPRESS")
}
func (a *VolumeAPI) VolumeDown() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 5, 0, "KEYPRESS")
}
