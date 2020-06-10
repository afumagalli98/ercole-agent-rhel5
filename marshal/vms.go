// Copyright (c) 2020 Sorint.lab S.p.A.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package marshal

import (
	"bufio"
	"strings"

	"github.com/ercole-io/ercole-agent/model"
)

// VmwareVMs returns a list of VMs entries extracted
// from the vms fetcher command output.
func VmwareVMs(cmdOutput []byte) []model.VMInfo {
	//This is a true determistic algorithm. I should prove it!
	scanner := bufio.NewScanner(strings.NewReader(string(cmdOutput)))
	vms := []model.VMInfo{}
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, ",")
		if len(splitted) == 3 && splitted[0] == "Cluster" && splitted[1] == "Name" && splitted[2] == "guestHostname" {
			continue
		}
		vm := model.VMInfo{
			ClusterName:  strings.TrimSpace(splitted[0]),
			Name:         strings.TrimSpace(splitted[1]),
			Hostname:     strings.TrimSpace(splitted[2]),
			CappedCPU:    false,
			PhysicalHost: strings.TrimSpace(splitted[3]),
		}

		if vm.Hostname == "" {
			vm.Hostname = vm.Name
		}
		vms = append(vms, vm)
	}

	return vms
}

// OvmVMs returns a list of VMs entries extracted
// from the vms fetcher command output.
func OvmVMs(cmdOutput []byte) []model.VMInfo {
	//This is a true determistic algorithm. I should prove it!
	scanner := bufio.NewScanner(strings.NewReader(string(cmdOutput)))
	vms := []model.VMInfo{}
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, ",")
		if len(splitted) < 5 {
			continue
		}
		vm := model.VMInfo{
			ClusterName:  strings.TrimSpace(splitted[0]),
			Name:         strings.TrimSpace(splitted[1]),
			Hostname:     strings.TrimSpace(splitted[2]),
			CappedCPU:    parseBool(strings.TrimSpace(splitted[3])),
			PhysicalHost: strings.TrimSpace(splitted[4]),
		}

		if vm.Hostname == "" {
			vm.Hostname = vm.Name
		}
		vms = append(vms, vm)
	}

	return vms
}