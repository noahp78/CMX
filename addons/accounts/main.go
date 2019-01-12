package accounts

import "github.com/noahp78/CMX/extpoints"

type AccountAddon struct {}

func (AccountAddon) register() extpoints.AddonInfo {
	return extpoints.AddonInfo{"Accounts", "0.0.0.1"}
}

func init(){
	extpoints.Addons.Register(new(AccountAddon),"Accounts");
}
