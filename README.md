## color

[ANSI EscapeCodes](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors)


## Install

```bash
go get github.com/sweetpotato0/color
```


## Usage

```go
package main
     
import "github.com/sweetpotato0/color"

func main() {
    color.Common("Common, %s", "JackieZhang")
    color.Red("Red, %s", "JackieZhang")
    color.Blue("Blue, , %s", "JackieZhang")
    color.Yellow("Yellow, %s", "JackieZhang")
    color.Green("Green%s", "JackieZhang")
    color.Black("Black, %s", "JackieZhang")
    color.Purple("Purple, %s", "JackieZhang")
    color.HiGreen("HiGreen, %s", "JackieZhang")
    color.White("White, %s", "JackieZhang")
}
```