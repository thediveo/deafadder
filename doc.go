/*
Package deafadder wraps [koanf.Koanf] configuration data objects, adding
convenience accessor functions for [pflag] configuration data types.

Suppose you have developed a service or CLI tool using [cobra] and [pflag] to
parse and handle CLI flags. Sweet and easy. Next, you want to support further
configuration data sources, such as configuration files, opting for the [koanf]
module(s). This assumes that in your service/CLI tool code base you are using
the [pflag.FlagSet] flag value accessor functions, and not directly accessing
the underlying flag values directly.

Here, the deafadder module comes in: in order to make this upgrade as simple and
painless as possible without having to throw away all the pflag-related code,
you basically just swap in the pflag value-compatible accessor code from the
[DeafAdder] object, which is a simple wrapper for [koanf.Koanf] objects.

	# before
	var cmd *cobra.Command
	addr, err := cmd.PersistentFlags().GetIP("addr")

	# after
	var k *koadnf.Koanf
	d := deadadder.New(k)

	addr, err := d.GetIP("addr")

# What's a Deaf Adder?

The name “deafadder” (“anguis fragilis sensu stricto”, better known as
“slowworm”) is a pun on the Go [cobra] and [viper] modules: it refers to a
[species of legless lizard] native to western Eurasia.

# Technical Note

Under its hood (or rather, skin) this package leverages the conversion functions
implemented in the [pflag] package, inheriting its behavior but also some
quirks. For instance, [DeafAdder.GetIP] currently does not report any errors in
case of invalid textual IP addresses.

[pflag]: https://github.com/spf13/pflag
[cobra]: https://github.com/spf13/cobra
[viper]: https://github.com/spf13/viper
[species of legless lizard]: https://en.wikipedia.org/wiki/Common_slow_worm
*/
package deafadder
