# Learn Golang 🚀

A concise, hands-on journal of small Go examples and exercises — updated daily while learning Go.

## Progress Table

| Day   | Activity                                | Folder Reference                                                                                                                                                                                                          |
| ----- | --------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Day 1 | Created first Go program (Hello World)  | [`1_hello-world`](./1_hello-world/)                                                                                                                                                                                       |
| Day 2 | Simple value types and examples         | [`2_simple-values`](./2_simple-values/), [`3_variables`](./3_variables/), [`4_constants`](./4_constants/), [`5_for`](./5_for/), [`6_if-else`](./6_if-else/)                                                               |
| Day 3 | Control flow, arrays/slices/maps, funcs | [`7_switch`](./7_switch/), [`8_arrays`](./8_arrays/), [`9_slices`](./9_slices/), [`10_maps`](./10_maps/), [`11_range`](./11_range/), [`12_functions`](./12_functions/), [`13_Variadic-function`](./13_Variadic-function/) |

---

## Day 3 — Quick folder overview

Below are the folders added for Day 3 (7 → 13). Each line links to the folder and the main example file with a short description.

- **7_switch** — [`switch.go`](./7_switch/switch.go)
  - Demonstrates basic `switch`, multiple conditions in a case, and a `type` switch helper function.
- **8_arrays** — [`array.go`](./8_arrays/array.go)
  - Array declaration, length, and a 2D array (matrix) example.
- **9_slices** — [`main.go`](./9_slices/main.go)
  - Slices primer: make, literals, copy, slicing operations, 2D slices, and `slices.Equal` (Go 1.21+).
- **10_maps** — [`map.go`](./10_maps/map.go)
  - Map creation, add/delete, existence check (comma-ok), clearing a map, and `maps.Equal` (Go 1.21+).
- **11_range** — [`range.go`](./11_range/range.go)
  - `range` examples for slices, maps, strings and using `_` to ignore values.
- **12_functions** — [`function.go`](./12_functions/function.go)
  - Functions: single/multiple returns, ignoring return values, and returning a function (closures / higher-order functions).
- **13_Variadic-function** — [`main.go`](./13_Variadic-function/main.go)
  - Variadic function example (`add(nums ...int)`) and unpacking a slice into variadic args.

## Try it (PowerShell)

Run any example directly using `go run` from the repository root. Examples:

```powershell
go run .\7_switch\switch.go
go run .\8_arrays\array.go
go run .\9_slices\main.go
go run .\10_maps\map.go
go run .\11_range\range.go
go run .\12_functions\function.go
go run .\13_Variadic-function\main.go
```

## Notes

- These exercises target Go 1.21+ for a few helper packages (`slices`, `maps`) used in examples. If you're on an older Go version, remove those parts or update Go.
- Contributions welcome! Open a PR with improvements or more examples.

---

This repository will be updated daily with new Go learning exercises and notes.
