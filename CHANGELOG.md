Changelog
=========

- Added support for option parsers other than the `flag` package by using an interface instead of directly depending on `flag.FlagSet`. Any type that provides methods such as `StringVar` can now be used. (#1)
- Added the BindTag method, allowing arbitrary struct tag names to be used. (#1)

v0.0.4
------
Oct 14, 2025

- Changed recursive binding behavior: Only struct or struct pointer fields explicitly tagged with `flag:""` are now traversed recursively. This prevents unintended recursion into unrelated struct types.

v0.0.3
------
Oct 13, 2025

- Fix: `panic: reflect.Value.Interface: cannot return value obtained from unexported field or method` when unexported fields exist.

v0.0.2
------
Oct 13, 2025

- Made the binding process recursive for non-nil struct and struct pointer fields.
- Added support for `uint` type.

v0.0.1
------
Oct 12, 2025

- Initial release.
- Supported only `bool`, `int`, and `string` field types.
