- 再帰的にたどる対象は `flag:""` というタグがついた構造体/ポインターのみとした

v0.0.3
======
Oct 13, 2025

- エクスポートされていないフィールドがあると、`panic: reflect.Value.Interface: cannot return value obtained from unexported field or method` が発生する修正

v0.0.2
======
Oct 13, 2025

- 非nil な構造体型フィールド下のタグも再帰的に参照するようにした
- uint型もサポート

v0.0.1
======
Oct 12, 2025

- 初版
- bool, int, string 型フィールドのみをサポート
