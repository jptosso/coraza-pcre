package main

import (
	"testing"

	"github.com/jptosso/coraza-waf/v2"
	"github.com/jptosso/coraza-waf/v2/operators"
)

func TestPlugin(t *testing.T) {
	waf := coraza.NewWaf()
	tx := waf.NewTransaction()
	op, err := operators.GetOperator("rx")
	if err != nil {
		t.Error(err)
	}
	if err := op.Init("^foo.*$"); err != nil {
		t.Error(err)
	}
	if !op.Evaluate(tx, "foo") {
		t.Error("failed to match regex")
	}
}
