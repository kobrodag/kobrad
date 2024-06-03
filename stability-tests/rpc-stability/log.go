package main

import (
	"github.com/kobradag/kobrad/infrastructure/logger"
	"github.com/kobradag/kobrad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
