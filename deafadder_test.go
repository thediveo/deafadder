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

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
	"github.com/spf13/pflag"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/thediveo/success"
)

var _ = Describe("retrieving values of all types", func() {

	var d *DeafAdder

	BeforeEach(func() {
		d = New(koanf.New("."))
		Expect(d).NotTo(BeNil())
		DeferCleanup(func() {
			d = nil
		})
	})

	Context("problems, problems, but we've got problems", func() {

		It("returns a lookup error", func() {
			Expect(d.GetInt8("foo.bar")).Error().To(
				MatchError(ContainSubstring("no such configuration setting foo.bar")))
		})

		It("reports a conversion error", func() {
			s := `
fool:
  bar: 42nd
`
			Expect(d.Load(rawbytes.Provider([]byte(s)), yaml.Parser())).To(Succeed())
			Expect(as(d, "fool.bar", (*pflag.FlagSet).Int, false)).Error().To(
				MatchError(ContainSubstring(`"42nd": invalid syntax`)))
		})

		It("reports trying to set a scalar to a sliced flag", func() {
			s := `
fool: 'bar'
`
			Expect(d.Load(rawbytes.Provider([]byte(s)), yaml.Parser())).To(Succeed())
			Expect(as(d, "fool", (*pflag.FlagSet).StringSlice, false)).Error().To(
				MatchError(ContainSubstring(`value for configuration setting fool must be slice`)))
		})

		It("reports trying to set invalid slice values (Replace)", func() {
			s := `
numbers:
  - 1234
  - abc
`
			Expect(d.Load(rawbytes.Provider([]byte(s)), yaml.Parser())).To(Succeed())
			Expect(as(d, "numbers", (*pflag.FlagSet).IntSlice, false)).Error().To(
				MatchError(ContainSubstring(`parsing "abc": invalid syntax`)))
		})

		It("reports trying to set invalid slice values (Set)", func() {
			s := `
subnets:
  - 127.0.0.1/8
  - abc/666
`
			Expect(d.Load(rawbytes.Provider([]byte(s)), yaml.Parser())).To(Succeed())
			Expect(as(d, "subnets", (*pflag.FlagSet).IPNetSlice, false)).Error().To(
				MatchError(ContainSubstring(`invalid string being converted to CIDR:`)))
		})

	})

	It("returns correct values", func() {
		s := `
config:
  bool: 'true'

  base64: 'QmFzZTY0'
  hex: deadbeadcafe

  count: 666

  duration: '25µs'
  duration-slice:
    - 6µs  # mind the tabs
    - 66ms
    - 666s

  float32: '42.666'
  float32-slice:
    - 42
    - 6.66
  float64: '42.666'
  float64-slice:
    - 42
    - 6.66
  
  int: '1234567890'
  int-slice:
    - 1
    - 2
    - 3
  int8: '42'
  int16: '32767'
  int32: '65536'
  int32-slice:
    - 1
    - 2
    - 65536
  int64: '123'
  int64-slice:
    - 1
    - 2
    - 3

  ip: '127.0.0.11'
  ip-slice:
    - 127.0.0.1
    - 192.168.0.1
  ip-net: '127.0.0.1/8'
  ip-nets:
    - 127.0.0.1/8
    - 192.168.0.1/24

  string: hellorld
  string-slice:
    - hello
    - world

  uint: 123
  uint-slice:
    - 123
  uint8: 42
  uint16: 43
  uint32: 44
  uint64: 45
`
		Expect(d.Load(rawbytes.Provider([]byte(s)), yaml.Parser())).To(Succeed())

		Expect(d.GetBool("config.bool")).To(BeTrue())

		Expect(d.GetBytesBase64("config.base64")).To(Equal([]byte("Base64")))
		Expect(d.GetBytesHex("config.hex")).To(Equal([]byte{0xde, 0xad, 0xbe, 0xad, 0xca, 0xfe}))

		Expect(d.GetCount("config.count")).To(Equal(666))

		Expect(d.GetDuration("config.duration")).To(Equal(time.Duration(25 * time.Microsecond)))
		Expect(d.GetDurationSlice("config.duration-slice")).To(Equal(
			[]time.Duration{6 * time.Microsecond, 66 * time.Millisecond, 666 * time.Second}))

		Expect(d.GetFloat32("config.float32")).To(Equal(float32(42.666)))
		Expect(d.GetFloat32Slice("config.float32-slice")).To(Equal([]float32{42, 6.66}))
		Expect(d.GetFloat64("config.float64")).To(Equal(float64(42.666)))
		Expect(d.GetFloat64Slice("config.float64-slice")).To(Equal([]float64{42, 6.66}))

		Expect(d.GetInt("config.int")).To(Equal(int(1234567890)))
		Expect(d.GetIntSlice("config.int-slice")).To(Equal([]int{1, 2, 3}))
		Expect(d.GetInt8("config.int8")).To(Equal(int8(42)))
		Expect(d.GetInt16("config.int16")).To(Equal(int16(32767)))
		Expect(d.GetInt32("config.int32")).To(Equal(int32(65536)))
		Expect(d.GetInt32Slice("config.int32-slice")).To(Equal([]int32{1, 2, 65536}))
		Expect(d.GetInt64("config.int64")).To(Equal(int64(123)))
		Expect(d.GetInt64Slice("config.int64-slice")).To(Equal([]int64{1, 2, 3}))

		Expect(d.GetIP("config.ip")).To(Equal(net.ParseIP("127.0.0.11")))
		Expect(d.GetIPSlice("config.ip-slice")).To(Equal(
			[]net.IP{net.ParseIP("127.0.0.1"), net.ParseIP("192.168.0.1")}))
		_, ipnet := Successful2R(net.ParseCIDR("127.0.0.1/8"))
		Expect(d.GetIPNet("config.ip-net")).To(Equal(*ipnet))
		_, ipnet2 := Successful2R(net.ParseCIDR("192.168.0.1/24"))
		Expect(d.GetIPNetSlice("config.ip-nets")).To(Equal([]net.IPNet{*ipnet, *ipnet2}))

		Expect(d.GetString("config.string")).To(Equal("hellorld"))
		Expect(d.GetStringSlice("config.string-slice")).To(Equal([]string{"hello", "world"}))
		Expect(d.GetStringArray("config.string-slice")).To(Equal([]string{"hello", "world"}))

		Expect(d.GetUint("config.uint")).To(Equal(uint(123)))
		Expect(d.GetUintSlice("config.uint-slice")).To(Equal([]uint{123}))
		Expect(d.GetUint8("config.uint8")).To(Equal(uint8(42)))
		Expect(d.GetUint16("config.uint16")).To(Equal(uint16(43)))
		Expect(d.GetUint32("config.uint32")).To(Equal(uint32(44)))
		Expect(d.GetUint64("config.uint64")).To(Equal(uint64(45)))
	})

})
