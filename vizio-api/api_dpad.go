package vizio_api

type DPadAPI struct {
	URL   string
	TOKEN string
}

func (a *DPadAPI) Up() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 3, 8, "KEYPRESS")
}
func (a *DPadAPI) Down() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 3, 0, "KEYPRESS")
}
func (a *DPadAPI) Left() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 3, 1, "KEYPRESS")
}
func (a *DPadAPI) Right() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 3, 7, "KEYPRESS")
}
func (a *DPadAPI) Select() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 3, 2, "KEYPRESS")
}
func (a *DPadAPI) Back() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 4, 0, "KEYPRESS")
}
func (a *DPadAPI) Back2() {
	KeyCommand(a.URL, a.TOKEN, "/key_command/", 4, 3, "KEYPRESS")
}
