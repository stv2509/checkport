package checkport

import (
	i "checkport/init"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"io/ioutil"
	"log"
	"strconv"
)

type Config struct {
	ConnTimeout string  `hcl:"ConnectTimeout"`
	RepTimeout  string  `hcl:"RepeatTimeout"`
	Hosts       []*Host `hcl:"host,block"`
}

type Host struct {
	Name  string   `hcl:",label"`
	Proto string   `hcl:"proto"`
	Ports []string `hcl:"ports"`
	//	Options hcl.Body `hcl:",remain"`
}

// Parse will parse file content into valid config.
func ParseHcl(src []byte, filename string) (c *Config, err error) {
	var diags hcl.Diagnostics

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, fmt.Errorf("config parse: %w", diags)
	}

	c = &Config{}

	diags = gohcl.DecodeBody(file.Body, nil, c)
	if diags.HasErrors() {
		return nil, fmt.Errorf("config parse: %w", diags)
	}

	return c, nil
}

func ShowConfig() (h []*Host, s, r int) {

	filename := i.ConfigPath
	content, err := ioutil.ReadFile(filename)
	if err != nil { // Handle errors reading the config file
		log.Fatalf("\n %s \n", err)
	}
	conf, err := ParseHcl(content, filename)
	if err != nil { // Handle errors reading the config file
		log.Fatalf("\n %s \n", err)
	}

	s, err = strconv.Atoi(conf.ConnTimeout)
	if err != nil { // Handle errors reading the config file
		log.Fatalf("\n %s \n %s\n", err, "Value \"RepeatTimeout\" in config is not type(int)")
	}

	r, err = strconv.Atoi(conf.RepTimeout)
	if err != nil { // Handle errors reading the config file
		log.Fatalf("\n %s \n %s\n", err, "Value \"ConnectTimeout\" in config is not type(int)")
	}

	return conf.Hosts, s, r

}
