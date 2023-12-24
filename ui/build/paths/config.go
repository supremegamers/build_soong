// Copyright 2018 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package paths

import "runtime"

type PathConfig struct {
	// Whether to create the symlink in the new PATH for this tool.
	Symlink bool

	// Whether to log about usages of this tool to the soong.log
	Log bool

	// Whether to exit with an error instead of invoking the underlying tool.
	Error bool

	// Whether we use a linux-specific prebuilt for this tool. On Darwin,
	// we'll allow the host executable instead.
	LinuxOnlyPrebuilt bool
}

var Allowed = PathConfig{
	Symlink: true,
	Log:     false,
	Error:   false,
}

var Forbidden = PathConfig{
	Symlink: false,
	Log:     true,
	Error:   true,
}

var Log = PathConfig{
	Symlink: true,
	Log:     true,
	Error:   false,
}

// The configuration used if the tool is not listed in the config below.
// Currently this will create the symlink, but log and error when it's used. In
// the future, I expect the symlink to be removed, and this will be equivalent
// to Forbidden.
var Missing = PathConfig{
	Symlink: true,
	Log:     true,
	Error:   true,
}

var LinuxOnlyPrebuilt = PathConfig{
	Symlink:           false,
	Log:               true,
	Error:             true,
	LinuxOnlyPrebuilt: true,
}

func GetConfig(name string) PathConfig {
	if config, ok := Configuration[name]; ok {
		return config
	}
	return Missing
}

var Configuration = map[string]PathConfig{
	"bash":    Allowed,
	"bindgen": Allowed,
	"dd":      Allowed,
	"diff":    Allowed,
	"dlv":     Allowed,
	"expr":    Allowed,
	"fuser":   Allowed,
	"getopt":  Allowed,
	"git":     Allowed,
	"hexdump": Allowed,
	"install": Allowed,
	"jar":     Allowed,
	"java":    Allowed,
	"javap":   Allowed,
	"lsof":    Allowed,
	"openssl": Allowed,
	"pahole":  Allowed,
	"perl":    Allowed,
	"pstree":  Allowed,
    "realpath": Allowed,
	"rsync":   Allowed,
	"sh":      Allowed,
    "rustc":   Allowed,
	"stubby":  Allowed,
	"tr":      Allowed,
	"unzip":   Allowed,
	"zip":     Allowed,

	"x86_64-linux-android-addr2line":    Allowed,
	"x86_64-linux-android-ar":           Allowed,
	"x86_64-linux-android-as":           Allowed,
	"x86_64-linux-android-c++filt":      Allowed,
	"x86_64-linux-android-dwp":          Allowed,
	"x86_64-linux-android-elfedit":      Allowed,
	"x86_64-linux-android-gcc":          Allowed,
	"x86_64-linux-android-gcc-ar":       Allowed,
	"x86_64-linux-android-gcc-nm":       Allowed,
	"x86_64-linux-android-gcc-ranlib":   Allowed,
	"x86_64-linux-android-gcov":         Allowed,
	"x86_64-linux-android-gcov-tool":    Allowed,
	"x86_64-linux-android-gprof":        Allowed,
	"x86_64-linux-android-ld":           Allowed,
	"x86_64-linux-android-ld.bfd":       Allowed,
	"x86_64-linux-android-ld.gold":      Allowed,
	"x86_64-linux-android-nm":           Allowed,
	"x86_64-linux-android-objcopy":      Allowed,
	"x86_64-linux-android-objdump":      Allowed,
	"x86_64-linux-android-ranlib":       Allowed,
	"x86_64-linux-android-readelf":      Allowed,
	"x86_64-linux-android-size":         Allowed,
	"x86_64-linux-android-strings":      Allowed,
	"x86_64-linux-android-strip":        Allowed,

	// Host toolchain is removed. In-tree toolchain should be used instead.
	// GCC also can't find cc1 with this implementation.
	"ar":         Forbidden,
	"as":         Forbidden,
	"cc":         Forbidden,
	"clang":      Forbidden,
	"clang++":    Forbidden,
	"gcc":        Forbidden,
	"g++":        Forbidden,
	"ld":         Forbidden,
	"ld.bfd":     Forbidden,
	"ld.gold":    Forbidden,
	"pkg-config": Forbidden,

	// These are toybox tools that only work on Linux.
	"pgrep": LinuxOnlyPrebuilt,
	"pkill": LinuxOnlyPrebuilt,
	"ps":    LinuxOnlyPrebuilt,
}

func init() {
	if runtime.GOOS == "darwin" {
		Configuration["sw_vers"] = Allowed
		Configuration["xcrun"] = Allowed

		// We don't have darwin prebuilts for some tools,
		// so allow the host versions.
		for name, config := range Configuration {
			if config.LinuxOnlyPrebuilt {
				Configuration[name] = Allowed
			}
		}
	}
}
