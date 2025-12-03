package service

import "github.com/google/wire"

// ProviderSet is gateway providers.
var ProviderSet = wire.NewSet(NewAgentService, NewProviderService, NewModalService, NewAgentWorkspaceService)
