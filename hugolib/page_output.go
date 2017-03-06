// Copyright 2017 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"github.com/spf13/hugo/output"
)

// PageOutput represents one of potentially many output formats of a given
// Page.
type PageOutput struct {
	*Page

	outputType output.Type
}

func newPageOutput(p *Page, outputType output.Type) *PageOutput {
	// TODO(bep) output avoid copy of first?
	p = p.copy()
	return &PageOutput{Page: p, outputType: outputType}
}

// copy creates a copy of this PageOutput with the lazy sync.Once vars reset
// so they will be evaluated again, for word count calculations etc.
func (p *PageOutput) copy() *PageOutput {
	c := *p
	c.Page = p.Page.copy()
	return &c
}