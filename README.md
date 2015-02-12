# awsip.go

## Installation
To install `awsip.go`
`$ go get github.com/akms/awsip`

## Getting Started
```go
import (
    "github.com/akms/awsip"
)
var awsipadder *awsip.AwsJson = awsip.NewAwsIpadder()
``` 

## Example
Create a file named example.go, and within it: 
```go
package main
import (
    "bytes"
    "fmt"
    "github.com/akms/awsip"
)
func main() {
    var awsipadder *awsip.AwsJson = awsip.NewAwsIpadder()
    awsipadder.SetHeadder("deny")
    b := bytes.NewBufferString("This is Example")
    b.Write(awsipadder.GetIpadderRegion("ap-northeast-1").Bytes())
    fmt.Printf(b.String())
}
```

`go run example.go`

```shell
This is Example
deny 54.231.224.0/21
deny 54.238.0.0/16
deny 54.64.0.0/15
deny 54.250.0.0/16
・・・
deny 54.250.253.192/26
```
