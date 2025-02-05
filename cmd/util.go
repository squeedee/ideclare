package cmd

import (
	"strings"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type platformsValue struct {
	platforms []v1.Platform
}

func (ps *platformsValue) Set(platform string) error {
	if ps.platforms == nil {
		ps.platforms = []v1.Platform{}
	}
	p, err := parsePlatform(platform)
	if err != nil {
		return err
	}
	pv := platformValue{p}
	ps.platforms = append(ps.platforms, *pv.platform)
	return nil
}

func (ps *platformsValue) String() string {
	ss := make([]string, 0, len(ps.platforms))
	for _, p := range ps.platforms {
		ss = append(ss, p.String())
	}
	return strings.Join(ss, ",")
}

func (ps *platformsValue) Type() string {
	return "platform(s)"
}

type platformValue struct {
	platform *v1.Platform
}

func (pv *platformValue) Set(platform string) error {
	p, err := parsePlatform(platform)
	if err != nil {
		return err
	}
	pv.platform = p
	return nil
}

func (pv *platformValue) String() string {
	return platformToString(pv.platform)
}

func (pv *platformValue) Type() string {
	return "platform"
}

func platformToString(p *v1.Platform) string {
	if p == nil {
		return "all"
	}
	return p.String()
}

func parsePlatform(platform string) (*v1.Platform, error) {
	if platform == "all" {
		return nil, nil
	}

	return v1.ParsePlatform(platform)
}
