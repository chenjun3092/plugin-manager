package events

import (
	docker "github.com/fsouza/go-dockerclient"
	"github.com/rancher/log"
	"github.com/rancher/plugin-manager/network"
)

type NetworkManagerHandler struct {
	nm *network.Manager
}

func (h *NetworkManagerHandler) Handle(event *docker.APIEvents) error {
	if err := h.nm.Evaluate(event.ID); err != nil {
		log.Errorf("Failed to evaluate network state for %s: %v", event.ID, err)
		return err
	}
	return nil
}
