package buildinfo

import "time"

// TODO: append build info with right key in prometheus metrics
// TODO: append build info in expvar

var info = &buildInfo{
	bin:       "",
	version:   "0.0.1",
	commit:    "NA",
	branch:    "main",
	timestamp: time.Now(),
}

type BuildInfo interface {
	Bin() string
	Version() string
	Commit() string
	Branch() string
	Timestamp() time.Time
}

type buildInfo struct {
	bin       string
	version   string
	commit    string
	branch    string
	timestamp time.Time
}

func (i *buildInfo) Version() string {
	return i.version
}

func (i *buildInfo) Commit() string {
	return i.commit
}

func (i *buildInfo) Branch() string {
	return i.branch
}

func (i *buildInfo) Timestamp() time.Time {
	return i.timestamp
}

func (i *buildInfo) Bin() string {
	return i.bin
}

func Get() BuildInfo {
	return info
}

func SetBin(bin string) {
	info.bin = bin
}

func SetVersion(version string) {
	info.version = version
}

func SetCommit(commit string) {
	info.commit = commit
}

func SetBranch(branch string) {
	info.branch = branch
}

func SetTimestamp(timestamp time.Time) {
	info.timestamp = timestamp
}
