// Copyright 2019 Red Hat
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

package filesystems

import (
	"github.com/coreos/ignition/v2/tests/register"
	"github.com/coreos/ignition/v2/tests/types"
)

func init() {
	register.Register(register.PositiveTest, SkipReadOnlyFilesystems())
}

func SkipReadOnlyFilesystems() types.Test {
	name := "filesystem.skip.readonly"

	mountPathName := "/usr/share/oem1"
	in := types.GetBaseDisk()
	out := types.GetBaseDisk()
	mntDevices := []types.MntDevice{
		{
			Label:        "important-data",
			Substitution: "$DEVICE",
		},
	}
	config := `{
		"ignition": {"version": "$version"},
		"storage": {
			"disks": [{
				"device": "/dev/loop8p1",
				"partitions": [
				{
					"label": "READ-ONLY",
					"number": 1,
					"mountpath": "/usr/share/oem"
				}
				]
			}]
		}
	}`
	configMinVersion := "3.0.0"

	in[0].Partitions.GetPartition("READ-ONLY").FilesystemType = "ext4"
	out[0].Partitions.GetPartition("READ-ONLY").FilesystemType = "ext4"
	in[0].Partitions.GetPartition("READ-ONLY").FilesystemMode = 0444
	out[0].Partitions.GetPartition("READ-ONLY").FilesystemMode = 0777
	out[0].Partitions.GetPartition("READ-ONLY").MountPath = mountPathName

	return types.Test{
		Name:             name,
		In:               in,
		Out:              out,
		MntDevices:       mntDevices,
		Config:           config,
		ConfigMinVersion: configMinVersion,
	}
}