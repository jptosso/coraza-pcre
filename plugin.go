// Copyright 2021 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	pcre "github.com/gijsbers/go-pcre"
	"github.com/jptosso/coraza-waf/v2"
	"github.com/jptosso/coraza-waf/v2/operators"
)

type rx struct {
	re pcre.Regexp
}

func (o *rx) Init(data string) error {
	re, err := pcre.Compile(data, 0)
	o.re = re
	return err
}

func (o *rx) Evaluate(tx *coraza.Transaction, value string) bool {
	m := o.re.MatcherString(value, 0)
	for i := 0; i < m.Groups()+1; i++ {
		if i == 10 {
			return true
		}
		tx.CaptureField(i, m.GroupString(i))
	}
	return m.Matches()
}

func init() {
	operators.RegisterOperator("rx", func() coraza.RuleOperator { return new(rx) })
}

var _ coraza.RuleOperator = &rx{}
