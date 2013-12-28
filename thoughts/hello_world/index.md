# Hello World!!!

![Road](/hello_world/road.jpg)

Finally, my personal blog was official done. It's one day work, but it's so very satisfing. Bellow are some of Tests.

This website is build by the power of [golang](http://golang.org). It's very simple and very small. Write your acticles in a code editor or a markdown editor, then save it, publish it, which means push your changes to a git repository then redeploy your server, then it's done.

## Heads
### Heads
#### Heads

[A link](/about).

A List:

* Item 1
* Item 2
* Item 3
* Item 4

A Code Block.

```
package main

import (
	"github.com/bom-d-van/me/app"
	"github.com/bom-d-van/me/configs"
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", app.GetThoughts)
	m.Get("/me", app.GetMe)
	m.Get("/thoughts", app.GetThoughts)
	m.Get("/thoughts/:artile_name", app.GetArticle)

	println("Serving Me on Port", configs.Port)
	for {
		http.ListenAndServe(configs.Port, m)
	}
}
```

Some Text Holders.

Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Vivamus vitae risus vitae lorem iaculis placerat. Aliquam sit amet felis. Etiam congue. Donec risus risus, pretium ac, tincidunt eu, tempor eu, quam. Morbi blandit mollis magna. Suspendisse eu tortor. Donec vitae felis nec ligula blandit rhoncus.

Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Vivamus vitae risus vitae lorem iaculis placerat. Aliquam sit amet felis. Etiam congue. Donec risus risus, pretium ac, tincidunt eu, tempor eu, quam. Morbi blandit mollis magna. Suspendisse eu tortor. Donec vitae felis nec ligula blandit rhoncus. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Vivamus vitae risus vitae lorem iaculis placerat. Aliquam sit amet felis. Etiam congue. Donec risus risus, pretium ac, tincidunt eu, tempor eu, quam. Morbi blandit mollis magna. Suspendisse eu tortor. Donec vitae felis nec ligula blandit rhoncus. Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Vivamus vitae risus vitae lorem iaculis placerat. Aliquam sit amet felis. Etiam congue. Donec risus risus, pretium ac, tincidunt eu, tempor eu, quam. Morbi blandit mollis magna. Suspendisse eu tortor. Donec vitae felis nec ligula blandit rhoncus.

This is Open Source, if you are instered in this, you can fork it [here](https://github.com/bom-d-van/me).