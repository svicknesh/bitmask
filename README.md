# Bitmask

Golang helper library for bitmask operations.

## Usage

### Creating new bitmask instance

```go
const (
    SrvCapRSA Bit = 1 << iota // fastest way to initialize capabilities
    SrvCapEC
    SrvCapED
    SrvCapReserved1
    SrvCapReserved2
    SrvCapCrvPub
)

b := New(3)
fmt.Println("new\t", b)

b.Set(SrvCapRSA, SrvCapEC, SrvCapED, SrvCapCrvPub)
//b.Set(SrvCapRSA, SrvCapED)
fmt.Println("set\t", b)

fmt.Println("has RSA\t", b.Has(SrvCapRSA))

b.Remove(SrvCapRSA)
//b.Set(SrvCapRSA, SrvCapED)
fmt.Println("remove\t", b)

fmt.Println("has RSA\t", b.Has(SrvCapRSA))

b.SetAll()
fmt.Println("setall\t", b)

b.Clear()
fmt.Println("clear\t", b)
```

### Decoding from bit string

```go
b, err := NewFromStr("100111")
if nil != err {
    log.Println(err)
    os.Exit(1)
}
fmt.Println("new\t", b)

fmt.Println("has RSA\t", b.Has(SrvCapRSA))

fmt.Println(b.Uint64())
```
