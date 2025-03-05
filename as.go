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

package deafadder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/pflag"
)

// as looks up the value for the specified path, and if successful, returns the
// value as of type T and using pflag conversion rules. If the value does not
// exist or cannot be converted into a value of type T, an error is returned
// instead.
func as[T any](d *DeafAdder, path string, fn func(*pflag.FlagSet, string, T, string) *T, treatAsScalar bool) (v T, err error) {
	// Let's see if we can get a configValue for the specified element; if not, we're
	// done, nothing we can do about it.
	configValue := d.Get(path)
	if configValue == nil {
		return v, fmt.Errorf("no such configuration setting %s", path)
	}
	// We got a value, so let's now create/construct a suitable throw-away flag
	// as part of a throw-away flag set. Notice how creating the flag will give
	// us the pointer to where the flag stores its particular value (note that
	// this is "value" not "Value").
	fs := pflag.NewFlagSet("deafadder-dummy-flagset", pflag.ContinueOnError)
	fnCtor := reflect.ValueOf(fn)
	flagValueT := fnCtor.Type().In(2)
	var pFlagValue *T = fnCtor.Call([]reflect.Value{
		reflect.ValueOf(fs),
		rvName,
		reflect.Zero(flagValueT),
		rvUsage,
	})[0].Interface().(*T)
	// Now we want to set the flag's value, tickling its conversion and checking
	// logic. Here, we need to differentiate between scalar and slice flag
	// values. Note that scalar flags can have struct types, such as net.IP.
	flag := fs.Lookup(flagName)
	if !treatAsScalar && flagValueT.Kind() == reflect.Slice {
		if reflect.TypeOf(configValue).Kind() != reflect.Slice {
			return v, fmt.Errorf("value for configuration setting %s must be slice", path)
		}
		if fsv, ok := flag.Value.(pflag.SliceValue); ok {
			cr := reflect.ValueOf(configValue)
			cl := make([]string, cr.Len())
			for idx := range cl {
				cl[idx] = fmt.Sprintf("%v", cr.Index(idx))
			}
			if err := fsv.Replace(cl); err != nil {
				return v, err
			}
			return *pFlagValue, nil
		}
		cr := reflect.ValueOf(configValue)
		sl := make([]string, cr.Len())
		for idx := range sl {
			sl[idx] = fmt.Sprintf("%v", cr.Index(idx))
		}
		if err := flag.Value.Set(strings.Join(sl, ",")); err != nil {
			return v, err
		}
		return *pFlagValue, nil
	}
	if err := flag.Value.Set(fmt.Sprintf("%v", configValue)); err != nil {
		return v, err
	}
	return *pFlagValue, nil
}

const flagName = "flag"

var (
	rvName  = reflect.ValueOf(flagName)
	rvUsage = reflect.ValueOf("not a trace of hint")
)
