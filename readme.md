# Bitset

Go port of [std::bitset](https://en.cppreference.com/w/cpp/utility/bitset/set).

Currently, supports 4 variants:
* Bitset8 (stores 8 bits on [0:7])
* Bitset16 (stores 16 bits on [0:15])
* Bitset32 (stores 32 bits on [0:31])
* Bitset64 (stores 64 bits on [0:63])

### Usage

```go
var b Bitset
b.SetBit(7, true)
println(b.CheckBit(7)) // true
```
