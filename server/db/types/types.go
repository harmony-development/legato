package types

import (
	"fmt"
	"strings"
)

type PermissionsNode struct {
	Node  string
	Allow bool
}

func (p *PermissionsNode) Deserialize(s string) (ok bool) {
	trimmed := strings.Split(strings.TrimSuffix(strings.TrimPrefix(s, "("), ")"), ",")

	if len(trimmed) != 3 {
		return false
	}

	if trimmed[2] == "t" {
		p.Allow = true
	} else {
		p.Allow = false
	}

	p.Node = trimmed[1]

	return true
}

func (p PermissionsNode) Serialize() string {
	node := "f"
	if p.Allow {
		node = "t"
	}
	return fmt.Sprintf("(%s,%s)", p.Node, node)
}

type PermissionsData struct {
	Roles      map[uint64][]PermissionsNode
	Categories map[uint64][]uint64
	Channels   map[uint64]map[uint64][]PermissionsNode
}
