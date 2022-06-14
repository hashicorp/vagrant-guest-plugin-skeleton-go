package main

import (
	"os"

	"github.com/hashicorp/vagrant-guest-plugin-skeleton/internal/guest"
	sdk "github.com/hashicorp/vagrant-plugin-sdk"
)

// Options are the SDK options to use for instantiation.
var ComponentOptions = []sdk.Option{
	sdk.WithComponents(
		&guest.MySpecialGuest{},
	),
	sdk.WithName("myspecialguest"),
}

func main() {
	sdk.Main(ComponentOptions...)
	os.Exit(0)
}
