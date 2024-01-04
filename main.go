/*
 * Copyright 2021-2023 JetBrains s.r.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/JetBrains/qodana-cli/v2023/cmd"
	"github.com/JetBrains/qodana-cli/v2023/core"
	"github.com/JetBrains/qodana-cli/v2023/platform"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	platform.InterruptChannel = make(chan os.Signal, 1)
	signal.Notify(platform.InterruptChannel, os.Interrupt)
	signal.Notify(platform.InterruptChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-platform.InterruptChannel
		platform.WarningMessage("Interrupting Qodana CLI...")
		log.SetOutput(io.Discard)
		core.CheckForUpdates(platform.Version)
		core.ContainerCleanup()
		_ = platform.QodanaSpinner.Stop()
		os.Exit(0)
	}()
	cmd.InitCli()
	cmd.Execute()
}
