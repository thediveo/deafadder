// Copyright 2025 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package sub

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("submerging key-value maps", func() {

	It("creates the necessary sub-maps", func() {
		src := map[string]any{}
		dst := map[string]any{}
		Merge([]string{"inside", "job"})(src, dst)
		Expect(dst).To(Equal(map[string]any{
			"inside": map[string]any{
				"job": map[string]any{},
			},
		}))
	})

	It("keeps existing sub-maps", func() {
		src := map[string]any{}
		dst := map[string]any{
			"inside": map[string]any{
				"fool": "bar",
			},
		}
		Merge([]string{"inside", "job"})(src, dst)
		Expect(dst).To(Equal(map[string]any{
			"inside": map[string]any{
				"fool": "bar",
				"job":  map[string]any{},
			},
		}))
	})

	It("busts unstable elements", func() {
		src := map[string]any{}
		dst := map[string]any{
			"inside": 42,
		}
		Merge([]string{"inside", "job"})(src, dst)
		Expect(dst).To(Equal(map[string]any{
			"inside": map[string]any{
				"job": map[string]any{},
			},
		}))
	})

	It("merges at a deeper level", func() {
		src := map[string]any{
			"truth": "false",
		}
		dst := map[string]any{
			"config": map[string]any{
				"fool": "bar",
			},
		}
		Merge([]string{"config"})(src, dst)
		Expect(dst).To(Equal(map[string]any{
			"config": map[string]any{
				"truth": "false",
				"fool":  "bar",
			},
		}))
	})

})
