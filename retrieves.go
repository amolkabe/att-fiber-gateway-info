package main

import (
	"edgan/att-fiber-gateway-info/internal/logging"
)

func (rc *gatewayClient) retrieveAction(
	action string, actionPages map[string]string, configs configs, flags *flags, model string, returnFact string,
) (string, error) {
	fact := empty

	// Get the specified page based on action
	page := returnActionPage(action, actionPages)

	// login is not required for most pages
	loginRequired := false

	// pages that require login
	loginPages := []string{"ipalloc", "nattable", "reset"}

	for _, loginPage := range loginPages {
		if page == loginPage {
			if configs.Password == empty {
				logging.LogFatal("Password is required")
			}
			loginRequired = true
		}
	}

	if loginRequired {
		logging.DebugLog(*flags.Debug, "LoginRequired true")
		performLogin(rc, configs, flags)
	}

	fact, err := rc.getPage(action, configs, flags, model, page, returnFact)

	return fact, err
}
