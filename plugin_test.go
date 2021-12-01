package plugin

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

func TestSomePayloads(t *testing.T) {
	waf := coraza.NewWaf()
	tx := waf.NewTransaction()
	op, err := operators.GetOperator("rx")
	if err != nil {
		t.Error(err)
	}
	if err := op.Init(`(?i:(?:(?:n(?:and|ot)|(?:x?x)?or|between|\|\||like|and|div|&&)[\s(]+\w+[\s)]*?[!=+]+[\s\d]*?[\"'=()]|\d(?:\s*?(?:between|like|x?or|and|div)\s*?\d+\s*?[\-+]|\s+group\s+by.+\()|/\w+;?\s+(?:between|having|select|like|x?or|and|div)\W|--\s*?(?:(?:insert|update)\s*?\w{2,}|alter|drop)|#\s*?(?:(?:insert|update)\s*?\w{2,}|alter|drop)|;\s*?(?:(?:insert|update)\s*?\w{2,}|alter|drop)|\@.+=\s*?\(\s*?select|[^\w]SET\s*?\@\w+))`); err != nil {
		t.Error(err)
	}
	if !op.Evaluate(tx, "var= @.= ( SELECT\"") {
		t.Error("failed to match regex")
	}
}
