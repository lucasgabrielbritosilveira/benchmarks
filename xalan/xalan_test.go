package xalan

import (
	"testing"
)

var files = []string{
	"files/acks.xml",
	"files/binding.xml",
	"files/changes.xml",
	"files/concepts.xml",
	"files/controls.xml",
	"files/datatypes.xml",
	"files/expr.xml",
	"files/index.xml",
	"files/intro.xml",
	"files/model.xml",
	"files/prod-notes.xml",
	"files/references.xml",
	"files/rpm.xml",
	"files/schema.xml",
	"files/structure.xml",
	"files/template.xml",
	"files/terms.xml",
	"files/ui.xml",
}

func BenchmarkXalan(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
