package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google/google"
)

var seps = "	"

func co(r *schema.Resource, prefix string) {
	for k, e := range r.Schema {
		req := ""
		if e.Required {
			req = req + "R"
		}
		if e.Computed {
			req = req + "C"
		}
		if e.Optional {
			req = req + "O"
		}
		fmt.Printf("%s%s%s%s\n", req, seps, prefix, k)
		if e.Elem != nil {
			t, ok := e.Elem.(*schema.Resource)
			if ok {
				co(t, prefix+k+seps)
			}
		}
	}
}

func main() {
	pf := flag.NewFlagSet("terraform-google-property-exporter", flag.ExitOnError)
	var (
		s = pf.String("separator", "	", "separator char")
		c = pf.Bool("csv", false, "output by csv (overwrite separator option)")
	)
	pf.Parse(os.Args[1:])
	if *c {
		*s = ","
	}
	seps = *s
	rm := google.ResourceMap()
	for k, r := range rm {
		co(r, k+seps)
	}
}
