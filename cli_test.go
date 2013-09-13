package main

import (
	//"github.com/capitancambio/go-subcommand"
	"github.com/daisy-consortium/pipeline-clientlib-go"
	"testing"
)

var SCRIPT pipeline.Script = pipeline.Script{
	Nicename:    "test-to-test",
	Description: "Mocked script",
	Homepage:    "daisy.org",
	Href:        "daisy.org/test",
	Id:          "test",
	Options: []pipeline.Option{
		pipeline.Option{
			Required:  true,
			Sequence:  false,
			Name:      "test-opt",
			Ordered:   false,
			Mediatype: "xml",
			Desc:      "I'm a test option",
			Type:      "anyFileURI",
		},
		pipeline.Option{
			Required:  false,
			Sequence:  false,
			Name:      "another-opt",
			Ordered:   false,
			Mediatype: "xml",
			Desc:      "I'm a test option",
			Type:      "boolean",
		},
	},
	Inputs: []pipeline.Input{
		pipeline.Input{
			Desc:      "input port",
			Mediatype: "application/x-dtbook+xml",
			Name:      "source",
			Sequence:  true,
		},
		pipeline.Input{
			Desc:      "input port not seq",
			Mediatype: "application/x-dtbook+xml",
			Name:      "single",
			Sequence:  false,
		},
	},
}

var in1, in2 = "tmp/dir1/file.xml", "tmp/dir2/file.xml"

func TestCliNonRequiredOptions(t *testing.T) {
        link:=PipelineLink{pipeline: &PipelineTest{false, 0}}
        cli, err := NewCli("testprog", link)
        if err != nil {
                t.Error("Unexpected error")
        }
        cli.AddScripts([]pipeline.Script{SCRIPT},link,false)
        //parser.Parse([]string{"test","--i-source","value"})
        err = cli.Run([]string{"test", "--i-source", "./tmp/file", "--i-single", "./tmp/file2", "--x-test-opt", "./myfile.xml"})
        if err != nil {
                t.Errorf("Non required option threw an error")
        }
}

