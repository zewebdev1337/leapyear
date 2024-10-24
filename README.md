# leapyear

Provides a function to determine if a given year is a leap year.

## Constraints
A year is a leap year if:
- It is divisible by 4 AND
- It is NOT divisible by 100 OR it is divisible by 400

## Install
`go get github.com/zewebdev1337/leapyear`

## Usage

```go
import "leapyear"

year := 2024
isLeap := leapyear.Is(year) // returns true
```

## Examples
```
2024 -> true  (divisible by 4)
2000 -> true  (divisible by 400)
2100 -> false (divisible by 100 but not 400)
2023 -> false (not divisible by 4)
```

## License

MIT
