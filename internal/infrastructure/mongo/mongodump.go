// Copyright (C) MongoDB, Inc. 2014-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// This file modified by lsjtop10 on 2023.07.31
// original file: mongo-tools/mongodump/main/mongodump.go

package mongoInf

import (
	"time"

	"github.com/mongodb/mongo-tools/common/log"
	"github.com/mongodb/mongo-tools/common/signals"
	"github.com/mongodb/mongo-tools/common/util"
	"github.com/mongodb/mongo-tools/mongodump"
)

const (
	progressBarLength   = 24
	progressBarWaitTime = time.Second * 3
)

var (
	VersionStr = "built-without-version-string"
	GitCommit  = "build-without-git-commit"
)

func ExecMongodump(options []string) error {
	// initialize command-line opts

	opts, err := mongodump.ParseOptions(options, VersionStr, GitCommit)
	if err != nil {
		log.Logvf(log.Always, "error parsing command line options: %s", err.Error())
		log.Logvf(log.Always, util.ShortUsage("mongodump"))
		return err
	}

	// init logger
	log.SetVerbosity(opts.Verbosity)

	// verify uri options and log them
	opts.URI.LogUnsupportedOptions()

	dump := mongodump.MongoDump{
		ToolOptions:   opts.ToolOptions,
		OutputOptions: opts.OutputOptions,
		InputOptions:  opts.InputOptions,
	}

	finishedChan := signals.HandleWithInterrupt(dump.HandleInterrupt)
	defer close(finishedChan)

	if err = dump.Init(); err != nil {
		log.Logvf(log.Always, "Failed: %v", err)
		return err
	}

	if err = dump.Dump(); err != nil {
		log.Logvf(log.Always, "Failed: %v", err)
		return err
	}

	return nil
}
