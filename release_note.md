v0.0.3
======
Oct 13, 2025

- Fix: `panic: reflect.Value.Interface: cannot return value obtained from unexported field or method` when unexported fields exist.

v0.0.2
======
Oct 13, 2025

- Made the binding process recursive for non-nil struct and struct pointer fields.
- Added support for `uint` type.

v0.0.1
======
Oct 12, 2025

- Initial release.
- Supported only `bool`, `int`, and `string` field types.
