package yml

import (
	"fmt"
	"testing"

	"github.com/google/cel-policy-templates-go/test"
)

func TestParse(t *testing.T) {
	tr := test.NewReader("../../../test/testdata")
	tests, err := tr.ReadCases("parse")
	if err != nil {
		t.Fatal(err)
	}
	for _, tst := range tests {
		tc := tst
		t.Run(tc.ID, func(tt *testing.T) {
			tmpl, iss := Parse(tc.In)
			if iss.Err() != nil {
				tt.Fatal(iss.Err())
			}
			if tc.Out != "" {
				dbg := Encode(tmpl, RenderDebugIDs)
				if tc.Out != dbg {
					fmt.Println(dbg)
					tt.Errorf("got:\n%s\nwanted:\n%s\n", dbg, tc.Out)
				}
			}
		})
	}
}
