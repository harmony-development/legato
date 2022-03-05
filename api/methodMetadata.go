package api

import harmonytypesv1 "github.com/harmony-development/legato/gen/harmonytypes/v1"

func meta(auth bool, local bool, owner bool, permissionNode string) *harmonytypesv1.HarmonyMethodMetadata {
	return &harmonytypesv1.HarmonyMethodMetadata{
		RequiresAuthentication: auth,
		RequiresLocal:          local,
		RequiresOwner:          owner,
		RequiresPermissionNode: permissionNode,
	}
}

var MethodMetadata = map[string]*harmonytypesv1.HarmonyMethodMetadata{
	"protocol.auth.v1.AuthService/Federate":       meta(false, false, false, ""),
	"protocol.auth.v1.AuthService/LoginFederated": meta(false, false, false, ""),
	"protocol.auth.v1.AuthService/Key":            meta(false, false, false, ""),
	"protocol.auth.v1.AuthService/BeginAuth":      meta(false, false, false, ""),
	"protocol.auth.v1.AuthService/NextStep":       meta(false, false, false, ""),
	"protocol.auth.v1.AuthService/StepBack":       meta(false, false, false, ""),
	"protocol.auth.v1.AuthService/StreamSteps":    meta(false, false, false, ""),
	"protocol.auth.v1.AuthService/CheckLoggedIn":  meta(true, false, false, ""),
}
