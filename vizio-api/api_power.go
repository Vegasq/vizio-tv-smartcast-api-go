package vizio_api

type PowerAPI struct {
	URL   string
	TOKEN string
}

func (a *PowerAPI) PowerOn() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 11, 1, "KEYPRESS")
}
func (a *PowerAPI) PowerOff() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 11, 0, "KEYPRESS")
}
func (a *PowerAPI) PowerToggle() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 11, 2, "KEYPRESS")
}
