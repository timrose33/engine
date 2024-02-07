// Copyright © by Jeff Foley 2023. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package scrape

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/owasp-amass/engine/net/http"
	"github.com/owasp-amass/engine/plugins/support"
	et "github.com/owasp-amass/engine/types"
	oam "github.com/owasp-amass/open-asset-model"
	"github.com/owasp-amass/open-asset-model/domain"
)

type bing struct {
	fmtstr string
}

func NewBing() et.Plugin {
	return &bing{fmtstr: "https://www.bing.com"}
}

func (b *bing) Start(r et.Registry) error {
	name := "Bing-Handler"

	if err := r.RegisterHandler(&et.Handler{
		Name:       name,
		Transforms: []string{"fqdn"},
		EventType:  oam.FQDN,
		Callback:   b.check,
	}); err != nil {
		r.Log().Printf("Failed to register the %s: %v", name, err)
		return err
	}
	return nil
}

func (b *bing) Stop() {}

func (b *bing) check(e *et.Event) error {
	fqdn, ok := e.Asset.Asset.(*domain.FQDN)
	if !ok {
		return errors.New("failed to extract the FQDN asset")
	}

	domlt := strings.ToLower(strings.TrimSpace(fqdn.Name))
	if e.Session.Config().WhichDomain(domlt) != domlt {
		return nil
	}

	matches, err := e.Session.Config().CheckTransformations("fqdn", "fqdn", "bing")
	if err != nil {
		return err
	}
	if !matches.IsMatch("fqdn") {
		return nil
	}

	for i := 1; i < 10; i++ {
		if body, err := b.query(domlt, i); err == nil {
			b.process(e, body)
		}
	}
	return nil
}

func (b *bing) query(name string, itemnum int) (string, error) {
	req := &http.Request{URL: fmt.Sprintf(b.fmtstr, itemnum, name, name)}

	resp, err := http.RequestWebPage(context.TODO(), req)
	if err != nil {
		return "", fmt.Errorf("error fetching URL: %w", err)
	}

	return resp.Body, nil
}

func (b *bing) process(e *et.Event, body string) {
	for _, name := range support.ScrapeSubdomainNames(body) {
		n := strings.ToLower(strings.TrimSpace(name))
		// if the subdomain is not in scope, skip it
		if n != "" && e.Session.Config().IsDomainInScope(n) {
			b.submitCallback(e, n)
		}
	}
}

func (b *bing) submitCallback(e *et.Event, name string) {
	support.AppendToDBQueue(func() {
		if a, err := e.Session.DB().Create(nil, "", &domain.FQDN{Name: name}); err == nil && a != nil {
			_ = e.Dispatcher.DispatchEvent(&et.Event{
				Name:    name,
				Asset:   a,
				Session: e.Session,
			})
		}
	})
}
