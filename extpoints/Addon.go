package extpoints

type AddonInfo struct {
	name 	string;
	version string;
}

type Addon interface {
	register() AddonInfo
}