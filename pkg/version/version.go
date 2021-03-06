package version

// Copyright (c) Simon Pelczer 2019. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

var (
	//Version release version of the provider
	Version string
	//GitCommit SHA of the last git commit
	GitCommit string
	//DEVVersion string for the development version
	DEVVersion = "dev"
	// DEVCommit string for the development version
	DEVCommit = "local"
)

// GetReleaseInfo returns commit hash and version tag (if available).
func GetReleaseInfo() (sha, release string) {
	sha = DEVCommit
	release = DEVVersion

	if len(GitCommit) > 0 {
		sha = GitCommit
	}

	if len(Version) > 0 {
		release = Version
	}

	return sha, release
}
