package guest

import (
	"fmt"

	"github.com/hashicorp/vagrant-plugin-sdk/component"
	plugincore "github.com/hashicorp/vagrant-plugin-sdk/core"
)

// MySpecialGuest is a Guest implementation.
type MySpecialGuest struct {
}

// DetectFunc implements component.Guest
func (h *MySpecialGuest) GuestDetectFunc() interface{} {
	return h.Detect
}

func (h *MySpecialGuest) Detect(t plugincore.Target) (bool, error) {
	return false, nil
}

// ParentsFunc implements component.Guest
func (h *MySpecialGuest) ParentFunc() interface{} {
	return h.Parent
}

func (h *MySpecialGuest) Parent() string {
	return ""
}

// HasCapabilityFunc implements component.Guest
func (h *MySpecialGuest) HasCapabilityFunc() interface{} {
	return h.CheckCapability
}

func (h *MySpecialGuest) CheckCapability(n *component.NamedCapability) bool {
	return false
}

// CapabilityFunc implements component.Guest
func (h *MySpecialGuest) CapabilityFunc(name string) interface{} {
	return fmt.Errorf("requested capability %s not found", name)
}

var (
	_ component.Guest = (*MySpecialGuest)(nil)
)
