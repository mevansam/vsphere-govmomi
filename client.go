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

package govmomi

import (
	"net/url"

	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

type Client struct {
	*vim25.Client

	SessionManager *session.Manager
}

// NewClientFromClient creates and returns a new client structure from a
// soap.Client instance. The remote ServiceContent object is retrieved and
// populated in the Client structure before returning.
func NewClientFromClient(vimClient *vim25.Client) (*Client, error) {
	c := Client{
		Client:         vimClient,
		SessionManager: session.NewManager(vimClient),
	}

	return &c, nil
}

// NewClient creates a new client from a URL. The client authenticates with the
// server before returning if the URL contains user information.
func NewClient(u *url.URL, insecure bool) (*Client, error) {
	soapClient := soap.NewClient(u, insecure)
	vimClient, err := vim25.NewClient(context.TODO(), soapClient)
	if err != nil {
		return nil, err
	}

	c, err := NewClientFromClient(vimClient)
	if err != nil {
		return nil, err
	}

	// Only login if the URL contains user information.
	if u.User != nil {
		err = c.SessionManager.Login(context.TODO(), u.User)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// convience method for logout via SessionManager
func (c *Client) Logout() error {
	err := c.SessionManager.Logout(context.TODO())

	// We've logged out - let's close any idle connections
	c.CloseIdleConnections()

	return err
}

// RoundTrip dispatches to the RoundTripper field.
func (c *Client) RoundTrip(ctx context.Context, req, res soap.HasFault) error {
	return c.RoundTripper.RoundTrip(ctx, req, res)
}

func (c *Client) PropertyCollector() *property.Collector {
	return property.DefaultCollector(c.Client)
}

func (c *Client) Properties(obj types.ManagedObjectReference, p []string, dst interface{}) error {
	return c.PropertyCollector().RetrieveOne(context.TODO(), obj, p, dst)
}

func (c *Client) PropertiesN(objs []types.ManagedObjectReference, p []string, dst interface{}) error {
	return c.PropertyCollector().Retrieve(context.TODO(), objs, p, dst)
}

func (c *Client) WaitForProperties(obj types.ManagedObjectReference, ps []string, f func([]types.PropertyChange) bool) error {
	return property.Wait(context.TODO(), c.PropertyCollector(), obj, ps, f)
}
