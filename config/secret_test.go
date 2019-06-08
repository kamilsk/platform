package config_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"

	. "github.com/kamilsk/platform/config"
)

func TestSecret_Printing(t *testing.T) {
	type password struct {
		XMLName struct{} `json:"-"        xml:"password"   yaml:"-"`
		Value   Secret   `json:"password" xml:"value,attr" yaml:"password"`
	}
	secret := password{Value: "secret"}

	tests := []struct {
		name    string
		marshal func(password) ([]byte, error)
	}{
		{
			"print by `%#v`",
			func(pass password) ([]byte, error) {
				str := fmt.Sprintf("%#v", pass)
				return []byte(str), nil
			},
		},
		{
			"print by `%s`",
			func(pass password) ([]byte, error) {
				str := fmt.Sprintf("%s", pass.Value) //nolint:gosimple
				return []byte(str), nil
			},
		},
		{
			"json marshal",
			func(pass password) ([]byte, error) {
				return json.Marshal(pass)
			},
		},
		{
			"xml marshal",
			func(pass password) ([]byte, error) {
				return xml.Marshal(pass)
			},
		},
		{
			"yaml marshal",
			func(pass password) ([]byte, error) {
				return yaml.Marshal(pass)
			},
		},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			raw, err := tc.marshal(secret)
			assert.NoError(t, err)
			assert.False(t, strings.Contains(string(raw), string(secret.Value)))
		})
	}
}
