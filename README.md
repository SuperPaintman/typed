# typed [🇹]

`typed` is a typed Go stdlib (almost) for generics.

> It is a set of experimental packages that heavily use features of
> [Go 1.18](https://tip.golang.org/doc/go1.18) generics.

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

```go
package main

import (
  "fmt"
  "github.com/SuperPaintman/typed/containerx/stackx"
)

func main() {
  var stack stackx.Stack[int]

  stack = append(stack, 7331)
  stack = append(stack, 1337)

  v1, ok := stack.Pop()
  fmt.Printf("v1: %d, ok: %v\n", v1, ok)
  // Output: v1: 1337, ok: true

  v2, ok := stack.Pop()
  fmt.Printf("v2: %d, ok: %v\n", v2, ok)
  // Output: v2: 7331, ok: true

  v3, ok := stack.Pop()
  fmt.Printf("v3: %d, ok: %v\n", v3, ok)
  // Output: v3: 0, ok: false
}
```

```go
package main

import (
  "encoding/json"
  "fmt"
  "github.com/SuperPaintman/typed/nullx"
)

func main() {
  var post struct {
    Description nullx.Nullable[string]
    Likes       nullx.Nullable[int]
    Tags        nullx.NonNullableSlice[string]
  }

  if err := json.Unmarshal([]byte(`{ "likes": 0 }`), &post); err != nil {
    panic(err)
  }

  fmt.Printf("%#v\n", post)
  // Output: struct { Description nullx.Nullable[string]; Likes nullx.Nullable[int]; Tags nullx.NonNullableSlice[string] }{Description:nullx.Nullable[string]{Value:"", Valid:false}, Likes:nullx.Nullable[int]{Value:0, Valid:true}, Tags:nullx.NonNullableSlice[string](nil)}

  b, err := json.Marshal(post)
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s\n", b)
  // Output: {"Description":null,"Likes":0,"Tags":[]}
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
