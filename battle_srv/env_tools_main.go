package main

import (
	. "server/common"
	"server/env_tools"
	"server/iap"
	"server/storage"
)

func mmain() {
	MustParseConfig()
	MustParseConstants()
	storage.Init()
	iap.InitTrustedCertManager()
	env_tools.LoadPreConf()
	env_tools.LoadMailboxTranscript()
	if Conf.General.ServerEnv == SERVER_ENV_TEST {
		env_tools.MergeTestPlayerAccounts()
	}
}
