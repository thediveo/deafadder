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
	"net"
	"time"

	"github.com/knadh/koanf/v2"
	"github.com/spf13/pflag"
)

// DeafAdder (“anguis fragilis sensu stricto”, better known as “slowworm”) wraps
// a koanf.Koanf configuration data object. On top, DeafAdder objects provide
// convenience accessor functions for [pflag]-likeo  configuration data types.
//
// [pflag]: https://github.com/spf13/pflag
type DeafAdder struct {
	*koanf.Koanf
}

// New returns a new DeafAdder object, wrapping the passed koanf.Koanf
// configuration data object.
func New(k *koanf.Koanf) *DeafAdder {
	d := &DeafAdder{
		k,
	}
	return d
}

// GetBool returns the bool value of a configuration setting with the given
// name.
func (d *DeafAdder) GetBool(path string) (v bool, err error) {
	return as(d, path, (*pflag.FlagSet).Bool, false)
}

// GetBytesBase64 returns the []byte value of a configuration setting with the
// given name.
func (d *DeafAdder) GetBytesBase64(path string) (v []byte, err error) {
	return as(d, path, (*pflag.FlagSet).BytesBase64, true)
}

// GetBytesHex returns the []byte value of a configuration setting with the
// given name.
func (d *DeafAdder) GetBytesHex(path string) (v []byte, err error) {
	return as(d, path, (*pflag.FlagSet).BytesHex, true)
}

// GetCount returns the int value of a configuration setting with the given
// name.
func (d *DeafAdder) GetCount(path string) (v int, err error) {
	return as(d, path, counterConstructor, false)
}

func counterConstructor(fs *pflag.FlagSet, name string, _ int, usage string) *int {
	return fs.Count(name, usage)
}

// GetDuration returns the duration value of a configuration setting with the
// given name.
func (d *DeafAdder) GetDuration(path string) (v time.Duration, err error) {
	return as(d, path, (*pflag.FlagSet).Duration, false)
}

// GetDurationSlice returns the []time.Duration value of a configuration setting
// with the given name.
func (d *DeafAdder) GetDurationSlice(path string) (v []time.Duration, err error) {
	return as(d, path, (*pflag.FlagSet).DurationSlice, false)
}

// GetFloat32 returns the float32 value of a configuration setting with the
// given name.
func (d *DeafAdder) GetFloat32(path string) (v float32, err error) {
	return as(d, path, (*pflag.FlagSet).Float32, false)
}

// GetFloat32Slice returns the []float32 value of a configuration setting with
// the given name.
func (d *DeafAdder) GetFloat32Slice(path string) (v []float32, err error) {
	return as(d, path, (*pflag.FlagSet).Float32Slice, false)
}

// GetFloat64 returns the float64 value of a configuration setting with the
// given name.
func (d *DeafAdder) GetFloat64(path string) (v float64, err error) {
	return as(d, path, (*pflag.FlagSet).Float64, false)
}

// GetFloat64Slice returns the []float64 value of a configuration setting with
// the given name.
func (d *DeafAdder) GetFloat64Slice(path string) (v []float64, err error) {
	return as(d, path, (*pflag.FlagSet).Float64Slice, false)
}

// GetInt returns the int value of a configuration setting with the given name.
func (d *DeafAdder) GetInt(path string) (v int, err error) {
	return as(d, path, (*pflag.FlagSet).Int, false)
}

// GetIntSlice returns the []int value of a configuration setting with the given
// name.
func (d *DeafAdder) GetIntSlice(path string) (v []int, err error) {
	return as(d, path, (*pflag.FlagSet).IntSlice, false)
}

// GetInt8 returns the int8 value of a configuration setting with the given
// name.
func (d *DeafAdder) GetInt8(path string) (v int8, err error) {
	return as(d, path, (*pflag.FlagSet).Int8, false)
}

// GetInt16 returns the int16 value of a configuration setting with the given
// name.
func (d *DeafAdder) GetInt16(path string) (v int16, err error) {
	return as(d, path, (*pflag.FlagSet).Int16, false)
}

// GetInt32 returns the int32 value of a configuration setting with the given
// name.
func (d *DeafAdder) GetInt32(path string) (v int32, err error) {
	return as(d, path, (*pflag.FlagSet).Int32, false)
}

// GetInt32Slice returns the []int32 value of a configuration setting with the
// given name.
func (d *DeafAdder) GetInt32Slice(path string) (v []int32, err error) {
	return as(d, path, (*pflag.FlagSet).Int32Slice, false)
}

// GetInt64 returns the int64 value of a configuration setting with the given
// name.
func (d *DeafAdder) GetInt64(path string) (v int64, err error) {
	return as(d, path, (*pflag.FlagSet).Int64, false)
}

// GetInt64Slice returns the []int64 value of a configuration setting with the
// given name.
func (d *DeafAdder) GetInt64Slice(path string) (v []int64, err error) {
	return as(d, path, (*pflag.FlagSet).Int64Slice, false)
}

// GetIP returns the net.IP value of a configuration setting with the given
// name.
func (d *DeafAdder) GetIP(path string) (v net.IP, err error) {
	return as(d, path, (*pflag.FlagSet).IP, true)
}

// GetIPSlice returns the []net.IP value of a configuration setting with the
// given name.
func (d *DeafAdder) GetIPSlice(path string) (v []net.IP, err error) {
	return as(d, path, (*pflag.FlagSet).IPSlice, false)
}

// GetIPNet returns the net.IPNet value of a configuration setting with the
// given name.
func (d *DeafAdder) GetIPNet(path string) (v net.IPNet, err error) {
	return as(d, path, (*pflag.FlagSet).IPNet, false)
}

// GetIPNetSlice returns the []net.IPNet value of a configuration setting with
// the given name.
func (d *DeafAdder) GetIPNetSlice(path string) (v []net.IPNet, err error) {
	return as(d, path, (*pflag.FlagSet).IPNetSlice, false)
}

// GetString returns the string value of a configuration setting with the given
// name.
func (d *DeafAdder) GetString(path string) (v string, err error) {
	return as(d, path, (*pflag.FlagSet).String, false)
}

// GetStringSlice returns the []string value of a configuration setting with the
// given name.
func (d *DeafAdder) GetStringSlice(path string) (v []string, err error) {
	return as(d, path, (*pflag.FlagSet).StringSlice, false)
}

// GetStringArray returns the []string value of a configuration setting with the
// given name.
func (d *DeafAdder) GetStringArray(path string) (v []string, err error) {
	return as(d, path, (*pflag.FlagSet).StringArray, false)
}

// GetUint returns the uint value of a configuration setting with the given
// name.
func (d *DeafAdder) GetUint(path string) (v uint, err error) {
	return as(d, path, (*pflag.FlagSet).Uint, false)
}

// GetUintSlice returns the []uint value of a configuration setting with the
// given name.
func (d *DeafAdder) GetUintSlice(path string) (v []uint, err error) {
	return as(d, path, (*pflag.FlagSet).UintSlice, false)
}

// GetUint8 returns the uint8 value of a configuration setting with the given
// name.
func (d *DeafAdder) GetUint8(path string) (v uint8, err error) {
	return as(d, path, (*pflag.FlagSet).Uint8, false)
}

// GetUint16 returns the uint16 value of a configuration setting with the given
// name.
func (d *DeafAdder) GetUint16(path string) (v uint16, err error) {
	return as(d, path, (*pflag.FlagSet).Uint16, false)
}

// GetUint32 returns the uint32 value of a configuration setting with the given
// name.
func (d *DeafAdder) GetUint32(path string) (v uint32, err error) {
	return as(d, path, (*pflag.FlagSet).Uint32, false)
}

// GetUint64 returns the uint64 value of a configuration setting with the given
// name.
func (d *DeafAdder) GetUint64(path string) (v uint64, err error) {
	return as(d, path, (*pflag.FlagSet).Uint64, false)
}
