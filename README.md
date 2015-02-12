# awsip.go
awsip.go is simple package for searching ipadders from AWS ip-ranges.json in Go.

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
Create a file named example1.go, and within it code: 
```go
package main
import (
    "bytes"
    "fmt"
    "github.com/akms/awsip"
)
func main() {
    var awsipadder *awsip.AwsJson = awsip.NewAwsIpadder()
    b := bytes.NewBufferString("")
    b.Write(awsipadder.GetIpadderService("EC2").Bytes())
    fmt.Printf(b.String())
}
```
Create a file named example2.go, and within it code:

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

run example.go
```bash
$ go run example1.go
50.19.0.0/16
75.101.128.0/17
54.194.0.0/15
54.208.0.0/15
・・・
54.228.0.0/16
```

```bash
$ go run example2.go
This is Example
deny 54.231.224.0/21
deny 54.238.0.0/16
deny 54.64.0.0/15
deny 54.250.0.0/16
・・・
deny 54.250.253.192/26
```
