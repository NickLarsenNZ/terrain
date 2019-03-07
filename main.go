package main

import (
	"fmt"
	gohcl2 "github.com/hashicorp/hcl2/gohcl"
	hcl2 "github.com/hashicorp/hcl2/hcl"
	hcl2parse "github.com/hashicorp/hcl2/hclparse"
)

type managedResource struct {
	Type   string    `hcl:"type,label"`
	Name   string    `hcl:"name,label"`
	Config hcl2.Body `hcl:",remain"`
}

type topLevel struct {
	/*	Atlas     *atlas            `hcl:"atlas,block"`
		Datas     []dataResource    `hcl:"data,block"`
		Modules   []module          `hcl:"module,block"`
		Outputs   []output          `hcl:"output,block"`
		Providers []provider        `hcl:"provider,block"`*/
	Resources []managedResource `hcl:"resource,block"`
	/*	Terraform *terraform        `hcl:"terraform,block"`
		Variables []variable        `hcl:"variable,block"`
		Locals    []*locals         `hcl:"locals,block"`*/
}

func main() {

	filename := "test.hcl"
	var resources = []managedResource{}
	var raw topLevel

	parser := hcl2parse.NewParser()
	f, diags := parser.ParseHCLFile(filename)
	if diags.HasErrors() {
		panic("diags has errors on loading file")
	}

	diags = gohcl2.DecodeBody(f.Body, nil, &raw)
	if diags.HasErrors() {
		panic("diags has errors on decoding body")
	}

	for _, rawR := range raw.Resources {

		fmt.Printf("Resource: %s.%s\n", rawR.Type, rawR.Name)
		resources = append(resources, rawR)

	}

	fmt.Printf("count: %d", len(resources))

}
