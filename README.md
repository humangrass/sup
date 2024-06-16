[![build](https://github.com/humangrass/sup/actions/workflows/build-test.yml/badge.svg)](https://github.com/humangrass/sup/actions/workflows/build-test.yml)
[![lint](https://github.com/humangrass/sup/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/humangrass/sup/actions/workflows/golangci-lint.yml)

# Sup!

Really simple logger 

## Usage

```go
import "github.com/humangrass/sup"

sup.SetProjectName("Yo!")
sup.Info("Wazzaaaaa!!!") // 2000-07-07 22:33:44 [Yo!] Wazzaaaaa!!!
```
