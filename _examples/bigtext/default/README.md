# bigtext/default

![Animation](animation.svg)

```go
package main

import (
	"github.com/fujie-xiyou/pterm"
	"github.com/fujie-xiyou/pterm/putils"
)

func main() {
	pterm.DefaultBigText.WithLetters(putils.LettersFromString("PTerm")).Render()
}

```
