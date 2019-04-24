package pipeline

import (
	"io"
	"io/ioutil"
	//"os"

	"gopkg.in/yaml.v2"
)

type (
	Schema struct {
		Version string `yaml:"version"`
		Kind    string `yaml:"kind"`
		Name    string `yaml:"name"`
		Services []*Container `yaml:"services,omitempty"`
		Pipeline []*Container `yaml:"steps"`
	}

	Container struct {
		Name          string                    `yaml:"name"`
		Desc          string                    `yaml:"desc,omitempty"`
		Type          string                    `yaml:"type"`
		Settings      map[string]interface{}    `yaml:"settings,omitempty"`
		Vargs         map[string]interface{}    `yaml:",inline"`
	}
)

// Parse parses the configuration from bytes b.
func Parse(r io.Reader) (*Schema, error) {
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseBytes(out)
}

// ParseBytes parses the configuration from bytes b.
func ParseBytes(b []byte) (*Schema, error) {
	out := new(Schema)
	err := yaml.Unmarshal(b, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// ParseString parses the configuration from string s.
func ParseString(s string) (*Schema, error) {
	return ParseBytes(
		[]byte(s),
	)
}

// UnmarshalYAML implements the Unmarshaller interface.
// func (c *Column) UnmarshalYAML(unmarshal func(interface{}) error) error {
//         slice := yaml.MapSlice{}
//         if err := unmarshal(&slice); err != nil {
//                 return err
//         }
//
//         for _, s := range slice {
//                 container := Container{}
//                 out, _ := yaml.Marshal(s.Value)
//
//                 if err := yaml.Unmarshal(out, &container); err != nil {
//                         return err
//                 }
//                 if container.Name == "" {
//                         container.Name = fmt.Sprintf("%v", s.Key)
//                 }
//                 c.Containers = append(c.Containers, &container)
//         }
//         return nil
// }
