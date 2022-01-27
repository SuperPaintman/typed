# typed [🇹]

`typed` is a typed Go stdlib (almost).

> It is a set of experimental packages that heavily use features of
> [Go 1.18](https://tip.golang.org/doc/go1.18).

---

## Examples

```go
package main

import (
  "fmt"
  "github.com/SuperPaintman/typed/mathx"
)

func main() {
  v := mathx.Max(1, 20)

  fmt.Printf("v: %d\n", v)
  // Output: v: 20
}
```

---

## FAQ

> **Question**:
>
> Should I use this package in production?
>
> **Answer**:
>
> At the moment, I think not. As long as
> [Go 1.18](https://tip.golang.org/doc/go1.18) is in beta, it is better to
> taste Go generics, but it is questionable to use them in real projects.
>
> Later, probably yes.

> **Question**:
>
> Why every package has `x` suffix?
>
> **Answer**:
>
> To avoid collisions with the std Go packages and save idiomatic
> [package naming](https://go.dev/blog/package-names) simultaneously.
>
> Also, `t` (or any other letters) looks awkward and unreadable in package names
> (like `matht`).

---

#### License

[MIT](./LICENSE)

---

With 🫀 by [Aleksandr Krivoshchekov (@SuperPaintman)](https://github.com/SuperPaintman)
