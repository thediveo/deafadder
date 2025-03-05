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

import "github.com/knadh/koanf/maps"

// Merge returns a map merge function that merges its src map into its dest map
// at the specified path, mutating the destination map. For instance, Merge
// allows merging in CLI flag configuration settings at any level deeper inside
// your root configuration map.
//
// The returned merge function has the following properties:
//   - non-strict
//   - creates or replaces the values along the src merge root path with
//     key-value maps as necessary.
func Merge(path []string) func(src, dest map[string]any) {
	return func(src, dest map[string]any) {
		root := dest
		for _, key := range path {
			value, ok := root[key]
			if !ok {
				// nada, so we now need to create the remaining elements;
				// immediately, we just create the missing subordinate map, and
				// then carry on with the next element in question, if any,
				// coming back into this branch over and over again until all
				// missing path element values were created properly.
				child := map[string]any{}
				root[key] = child
				root = child
				continue
			}
			subMap, isMap := value.(map[string]any)
			if !isMap {
				// there's some value, but it ain't a key-value map, so with a
				// nod to "Dark Star" blast the existing unstable element and
				// replace it with an empty key-value map.
				child := map[string]any{}
				root[key] = child
				root = child
				continue
			}
			// just descend... (cue in Carmina Burana)
			root = subMap
		}
		// Carry on by simply merging the src map into the new root map inside
		// the destination map.
		maps.Merge(src, root)
	}
}
