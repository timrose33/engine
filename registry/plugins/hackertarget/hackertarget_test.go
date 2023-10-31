package main

import (
	"fmt"
	"log"
	"net/netip"
	"os"
	"testing"

	oamConfig "github.com/owasp-amass/config/config"
	"github.com/owasp-amass/engine/scheduler"
	"github.com/owasp-amass/engine/sessions"
	"github.com/owasp-amass/engine/types"
	oam "github.com/owasp-amass/open-asset-model"
	oamDom "github.com/owasp-amass/open-asset-model/domain"
	oamNet "github.com/owasp-amass/open-asset-model/network"
)

func TestLookup(t *testing.T) {
	logger := log.New(os.Stdout, "Test: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Create a new scheduler.
	scheduler := scheduler.NewScheduler(logger, nil)

	manager := sessions.NewStorage(logger)

	// Create a new config
	config := oamConfig.NewConfig()

	config.Scope.Domains = []string{"owasp.org"}

	transformationSub := oamConfig.Transformation{From: "FQDN", To: "ALL"}
	config.Transformations["FQDN->ALL"] = &transformationSub

	transformationIP := oamConfig.Transformation{From: "IPAddress", To: "ALL"}
	config.Transformations["IPAddress->ALL"] = &transformationIP
	// Create a new session.
	session, err := sessions.NewSession(config)
	if err != nil {
		log.Fatalf("Failed to create a new session: %v", err)
	}

	UUID, err := manager.Add(session)
	if err != nil {
		log.Fatalf("Failed to add the session: %v", err)
	}

	// Create an IP event.
	ip, err := netip.ParseAddr("8.8.8.8")
	if err != nil {
		t.Errorf("Failed to parse IP address: %v", err)
	}
	ipAsset := types.AssetData{
		OAMAsset: oamNet.IPAddress{
			Address: ip,
			Type:    "IPv4",
		},
		OAMType: oam.IPAddress,
	}
	ipEvent := types.Event{
		SessionID: UUID,
		Data:      ipAsset,
		Sched:     scheduler,
		Session:   session,
	}

	// Create a FQDN event.
	fqdn := "owasp.org"
	fqdnAsset := types.AssetData{
		OAMAsset: oamDom.FQDN{
			Name: fqdn,
		},
		OAMType: oam.FQDN,
	}
	fqdnEvent := types.Event{
		SessionID: UUID,
		Data:      fqdnAsset,
		Sched:     scheduler,
		Session:   session,
	}

	fmt.Println("ipEvent: ", ipEvent)
	fmt.Println("fqdnEvent: ", fqdnEvent)

	plugin := &HackerTargetPlugin{}

	// Test the ipLookup function.
	err = plugin.ipLookup(&ipEvent)
	if err != nil {
		t.Errorf("ipLookup failed: %v", err)
	}

	err = plugin.lookupDomain(&fqdnEvent)
	if err != nil {
		t.Errorf("LookupDomain failed: %v", err)
	}
}
