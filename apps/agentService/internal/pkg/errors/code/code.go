package code

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrAgentUnKnown         = errors.New(300000, "", "internal error")
	ErrAgentNoAuth          = errors.New(300001, "", "modal no auth")
	ErrAgentNotFound        = errors.New(300002, "", "modal not found")
	ErrAgentVersionExists   = errors.New(300003, "", "agent version already exists")
	ErrAgentVersionNotFound = errors.New(300004, "", "agent version not found")

	ErrProviderUnknown  = errors.New(300100, "", "provider internal error")
	ErrProviderNoAuth   = errors.New(300101, "", "provider no auth")
	ErrProviderNotFound = errors.New(300102, "", "provider not found")
	ErrProviderNoActive = errors.New(300103, "", "provider no active")

	ErrModalUnknown  = errors.New(300200, "", "modal internal error")
	ErrModalNoAuth   = errors.New(300201, "", "modal no auth")
	ErrModalNotFound = errors.New(300202, "", "modal not found")
	ErrModalNoActive = errors.New(300203, "", "modal no active")

	ErrWorkspaceUnknown = errors.New(300300, "", "workspace internal error")
	ErrWorkspaceNoAuth  = errors.New(300301, "", "workspace no auth")
)
