# Go Crash Course

A hands-on, lesson-by-lesson walkthrough of Go, written as runnable, heavily-commented source files.
Each lesson is a single `.go` file you can read top to bottom — the comments *are* the notes.

## Run it

```bash
go run .
```

Every lesson exposes an entry function (`lesson1()`, `lesson2()`, ...). Call the one you want from
`main()` in [main.go](main.go) to see its output.

```bash
go vet ./...   # static checks
gofmt -l .     # formatting check (no output = clean)
```

## Lessons

| # | File | Topics |
|---|------|--------|
| 1 | [lesson1.go](lesson1.go) | Variables & the three declaration forms, basic data types (`int`/`uint` widths, `float`, `bool`, `string`, `byte`, `rune`, `error`, `complex`), arrays, slices, maps, structs, functions and multiple return values |
| 2 | [lesson2.go](lesson2.go) | `if`/`else if`/`else`, `if` with initialization, error-check idiom, `for` (the only loop), `for range`, while-style loops, infinite loops, `break`/`continue`, `switch` (simple, multi-case, expressionless), block scope and shadowing |

## Notes on the code

- Everything lives in `package main` so lessons can call each other freely.
- `lesson2()` calls `InfiniteLoop()` — comment that line out before running it, or you'll be there a while.
- Comments are deliberately verbose. This is a learning repo, not production code.

## Requirements

Go 1.24.5 or newer.
