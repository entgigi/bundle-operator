package bundles

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type ComponentType string

const (
	ManifestComponentType ComponentType = "MANIFEST"
	PluginComponentType   ComponentType = "PLUGIN"
)

type Component struct {
	Name string        `json:"name,omitempty"`
	Type ComponentType `json:"type,omitempty"`
	Spec interface{}   `yaml:"-"`
}

type Plugin struct {
	Repository      string `json:"repository,omitempty"`
	Tag             string `json:"tag,omitempty"`
	Digest          string `json:"digest,omitempty"`
	HealthCheckPath string `json:"healthCheckPath,omitempty"`
	Port            int    `json:"port,omitempty"`
	IngressName     string `json:"ingressName,omitempty"`
	IngressHost     string `json:"ingressHost,omitempty"`
	IngressPath     string `json:"ingressPath,omitempty"`
}

type Manifest struct {
	FilePath string `json:"filePath,omitempty"`
}

type BundleDescriptor struct {
	Version      string      `json:"version"`
	Name         string      `json:"name"`
	Descriptor   string      `json:"descriptor"`
	Dependencies []string    `json:"dependencies"`
	Components   []Component `json:"components"`
}

func (s *Component) UnmarshalYAML(n *yaml.Node) error {
	type S Component
	type T struct {
		*S   `yaml:",inline"`
		Spec yaml.Node `yaml:"spec"`
	}

	obj := &T{S: (*S)(s)}
	if err := n.Decode(obj); err != nil {
		return err
	}

	switch s.Type {
	case ManifestComponentType:
		s.Spec = new(Manifest)
	case PluginComponentType:
		s.Spec = new(Plugin)
	default:
		panic(fmt.Sprintf("kind unknown %s", s.Type))
	}
	return obj.Spec.Decode(s.Spec)
}
