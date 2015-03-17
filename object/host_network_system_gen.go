/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package object

import (
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type HostNetworkSystem struct {
	Common
}

func NewHostNetworkSystem(c *vim25.Client, ref types.ManagedObjectReference) *HostNetworkSystem {
	return &HostNetworkSystem{
		Common: NewCommon(c, ref),
	}
}

// AddPortGroup wraps methods.AddPortGroup
func (o HostNetworkSystem) AddPortGroup(portgrp types.HostPortGroupSpec) error {
	req := types.AddPortGroup{
		This:    o.Reference(),
		Portgrp: portgrp,
	}

	_, err := methods.AddPortGroup(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// AddServiceConsoleVirtualNic wraps methods.AddServiceConsoleVirtualNic
func (o HostNetworkSystem) AddServiceConsoleVirtualNic(portgroup string, nic types.HostVirtualNicSpec) (string, error) {
	req := types.AddServiceConsoleVirtualNic{
		This:      o.Reference(),
		Portgroup: portgroup,
		Nic:       nic,
	}

	res, err := methods.AddServiceConsoleVirtualNic(context.TODO(), o.c, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

// AddVirtualNic wraps methods.AddVirtualNic
func (o HostNetworkSystem) AddVirtualNic(portgroup string, nic types.HostVirtualNicSpec) (string, error) {
	req := types.AddVirtualNic{
		This:      o.Reference(),
		Portgroup: portgroup,
		Nic:       nic,
	}

	res, err := methods.AddVirtualNic(context.TODO(), o.c, &req)
	if err != nil {
		return "", err
	}

	return res.Returnval, nil
}

// AddVirtualSwitch wraps methods.AddVirtualSwitch
func (o HostNetworkSystem) AddVirtualSwitch(vswitchName string, spec *types.HostVirtualSwitchSpec) error {
	req := types.AddVirtualSwitch{
		This:        o.Reference(),
		VswitchName: vswitchName,
		Spec:        spec,
	}

	_, err := methods.AddVirtualSwitch(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// QueryNetworkHint wraps methods.QueryNetworkHint
func (o HostNetworkSystem) QueryNetworkHint(device []string) error {
	req := types.QueryNetworkHint{
		This:   o.Reference(),
		Device: device,
	}

	_, err := methods.QueryNetworkHint(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// RefreshNetworkSystem wraps methods.RefreshNetworkSystem
func (o HostNetworkSystem) RefreshNetworkSystem() error {
	req := types.RefreshNetworkSystem{
		This: o.Reference(),
	}

	_, err := methods.RefreshNetworkSystem(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// RemovePortGroup wraps methods.RemovePortGroup
func (o HostNetworkSystem) RemovePortGroup(pgName string) error {
	req := types.RemovePortGroup{
		This:   o.Reference(),
		PgName: pgName,
	}

	_, err := methods.RemovePortGroup(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// RemoveServiceConsoleVirtualNic wraps methods.RemoveServiceConsoleVirtualNic
func (o HostNetworkSystem) RemoveServiceConsoleVirtualNic(device string) error {
	req := types.RemoveServiceConsoleVirtualNic{
		This:   o.Reference(),
		Device: device,
	}

	_, err := methods.RemoveServiceConsoleVirtualNic(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// RemoveVirtualNic wraps methods.RemoveVirtualNic
func (o HostNetworkSystem) RemoveVirtualNic(device string) error {
	req := types.RemoveVirtualNic{
		This:   o.Reference(),
		Device: device,
	}

	_, err := methods.RemoveVirtualNic(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// RemoveVirtualSwitch wraps methods.RemoveVirtualSwitch
func (o HostNetworkSystem) RemoveVirtualSwitch(vswitchName string) error {
	req := types.RemoveVirtualSwitch{
		This:        o.Reference(),
		VswitchName: vswitchName,
	}

	_, err := methods.RemoveVirtualSwitch(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// RestartServiceConsoleVirtualNic wraps methods.RestartServiceConsoleVirtualNic
func (o HostNetworkSystem) RestartServiceConsoleVirtualNic(device string) error {
	req := types.RestartServiceConsoleVirtualNic{
		This:   o.Reference(),
		Device: device,
	}

	_, err := methods.RestartServiceConsoleVirtualNic(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateConsoleIpRouteConfig wraps methods.UpdateConsoleIpRouteConfig
func (o HostNetworkSystem) UpdateConsoleIpRouteConfig(config types.BaseHostIpRouteConfig) error {
	req := types.UpdateConsoleIpRouteConfig{
		This:   o.Reference(),
		Config: config,
	}

	_, err := methods.UpdateConsoleIpRouteConfig(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateDnsConfig wraps methods.UpdateDnsConfig
func (o HostNetworkSystem) UpdateDnsConfig(config types.BaseHostDnsConfig) error {
	req := types.UpdateDnsConfig{
		This:   o.Reference(),
		Config: config,
	}

	_, err := methods.UpdateDnsConfig(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateIpRouteConfig wraps methods.UpdateIpRouteConfig
func (o HostNetworkSystem) UpdateIpRouteConfig(config types.BaseHostIpRouteConfig) error {
	req := types.UpdateIpRouteConfig{
		This:   o.Reference(),
		Config: config,
	}

	_, err := methods.UpdateIpRouteConfig(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateIpRouteTableConfig wraps methods.UpdateIpRouteTableConfig
func (o HostNetworkSystem) UpdateIpRouteTableConfig(config types.HostIpRouteTableConfig) error {
	req := types.UpdateIpRouteTableConfig{
		This:   o.Reference(),
		Config: config,
	}

	_, err := methods.UpdateIpRouteTableConfig(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateNetworkConfig wraps methods.UpdateNetworkConfig
func (o HostNetworkSystem) UpdateNetworkConfig(config types.HostNetworkConfig, changeMode string) (*types.HostNetworkConfigResult, error) {
	req := types.UpdateNetworkConfig{
		This:       o.Reference(),
		Config:     config,
		ChangeMode: changeMode,
	}

	res, err := methods.UpdateNetworkConfig(context.TODO(), o.c, &req)
	if err != nil {
		return nil, err
	}

	return &res.Returnval, nil
}

// UpdatePhysicalNicLinkSpeed wraps methods.UpdatePhysicalNicLinkSpeed
func (o HostNetworkSystem) UpdatePhysicalNicLinkSpeed(device string, linkSpeed *types.PhysicalNicLinkInfo) error {
	req := types.UpdatePhysicalNicLinkSpeed{
		This:      o.Reference(),
		Device:    device,
		LinkSpeed: linkSpeed,
	}

	_, err := methods.UpdatePhysicalNicLinkSpeed(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePortGroup wraps methods.UpdatePortGroup
func (o HostNetworkSystem) UpdatePortGroup(pgName string, portgrp types.HostPortGroupSpec) error {
	req := types.UpdatePortGroup{
		This:    o.Reference(),
		PgName:  pgName,
		Portgrp: portgrp,
	}

	_, err := methods.UpdatePortGroup(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateServiceConsoleVirtualNic wraps methods.UpdateServiceConsoleVirtualNic
func (o HostNetworkSystem) UpdateServiceConsoleVirtualNic(device string, nic types.HostVirtualNicSpec) error {
	req := types.UpdateServiceConsoleVirtualNic{
		This:   o.Reference(),
		Device: device,
		Nic:    nic,
	}

	_, err := methods.UpdateServiceConsoleVirtualNic(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateVirtualNic wraps methods.UpdateVirtualNic
func (o HostNetworkSystem) UpdateVirtualNic(device string, nic types.HostVirtualNicSpec) error {
	req := types.UpdateVirtualNic{
		This:   o.Reference(),
		Device: device,
		Nic:    nic,
	}

	_, err := methods.UpdateVirtualNic(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}

// UpdateVirtualSwitch wraps methods.UpdateVirtualSwitch
func (o HostNetworkSystem) UpdateVirtualSwitch(vswitchName string, spec types.HostVirtualSwitchSpec) error {
	req := types.UpdateVirtualSwitch{
		This:        o.Reference(),
		VswitchName: vswitchName,
		Spec:        spec,
	}

	_, err := methods.UpdateVirtualSwitch(context.TODO(), o.c, &req)
	if err != nil {
		return err
	}

	return nil
}
