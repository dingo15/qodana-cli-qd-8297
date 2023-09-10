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

// note: this file is generated by ../scripts/generate_bots.go and uses data from:
// https://github.com/JetBrains/qodana-cli/blob/main/bots.json

package cloud

var (
	GitHubBotSuffix = "[bot]@users.noreply.github.com"
	CommonGitBots   = []string{
		"cla-bot@users.noreply.github.com",
		"codecov-io@users.noreply.github.com",
		"snyk-bot@snyk.io",
		"gitlab-bot@gitlab.com",
		"bot@swc.rs",
	}
)