package plugin

import (
	"context"
	"testing"

	"github.com/corazawaf/coraza/v3/seclang"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/operators"
)

func TestPlugin(t *testing.T) {
	waf := coraza.NewWAF()
	tx := waf.NewTransaction(context.Background())
	op, err := operators.Get("rx")
	if err != nil {
		t.Error(err)
	}
	if err := op.Init(coraza.RuleOperatorOptions{Arguments: "^foo.*$"}); err != nil {
		t.Error(err)
	}
	if !op.Evaluate(tx, "foo") {
		t.Error("failed to match regex")
	}
}

func TestRxMacro(t *testing.T) {
	waf := coraza.NewWAF()
	rules := `
SecAction "id:100,setvar:'tx.macros=some'"
`
	parser := seclang.NewParser(waf)
	err := parser.FromString(rules)
	if err != nil {
		t.Error(err)
	}
	tx := waf.NewTransaction(context.Background())

	op, err := operators.Get("rx")
	if err != nil {
		t.Error(err)
	}
	if err := op.Init(coraza.RuleOperatorOptions{Arguments: "%{tx.macros}"}); err != nil {
		t.Error(err)
	}
	if op.Evaluate(tx, "somedata") {
		t.Error("error test case for rx")
	}
}

func TestSomePayloads(t *testing.T) {
	waf := coraza.NewWAF()
	tx := waf.NewTransaction(context.Background())
	op, err := operators.Get("rx")
	if err != nil {
		t.Error(err)
	}
	if err := op.Init(coraza.RuleOperatorOptions{Arguments: `(?i:(?:(?:n(?:and|ot)|(?:x?x)?or|between|\|\||like|and|div|&&)[\s(]+\w+[\s)]*?[!=+]+[\s\d]*?[\"'=()]|\d(?:\s*?(?:between|like|x?or|and|div)\s*?\d+\s*?[\-+]|\s+group\s+by.+\()|/\w+;?\s+(?:between|having|select|like|x?or|and|div)\W|--\s*?(?:(?:insert|update)\s*?\w{2,}|alter|drop)|#\s*?(?:(?:insert|update)\s*?\w{2,}|alter|drop)|;\s*?(?:(?:insert|update)\s*?\w{2,}|alter|drop)|\@.+=\s*?\(\s*?select|[^\w]SET\s*?\@\w+))`}); err != nil {
		t.Error(err)
	}
	if !op.Evaluate(tx, "var= @.= ( SELECT\"") {
		t.Error("failed to match regex")
	}
}
