package confx_test

import (
	"fmt"
	"testing"

	"github.com/eztop/go_common/confx"
	"github.com/eztop/go_common/log"
)

type conf struct {
	A  []string `toml:"a"`
	B  *string  `toml:"b"`
	XX struct {
		C string `toml:"c" `
	} `toml:"xx"`
}

func TestParse(t *testing.T) {
	c := &conf{}
	if err := confx.ParseFile("./testdata/conf.toml", c); err != nil {
		log.Error(err)
	}
	fmt.Println(c)
}

func TestRender(t *testing.T) {
	c := &conf{}
	if err := confx.ParseFile("./testdata/conf.toml", c); err != nil {
		log.Error(err)
	}
	fmt.Println(c)
}
