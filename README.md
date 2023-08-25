# GO+SPA Template

This application is set up to serve any SPA application that compiles to static files. In this example, I'm using SvelteKit. React, Angular, and Vue should all work as well however. 

Future features include:

- Rebuilding of clients when their files change
- Dependencies are automatically installed if they are missing

```bash
.
├── client                  # SvelteKit SPA (Example)
│   ├── package.json
│   ├── package-lock.json
│   ├── playwright.config.ts
│   ├── README.md
│   ├── src
│   │   ├── app.d.ts
│   │   ├── app.html
│   │   ├── index.test.ts
│   │   ├── lib
│   │   │   └── index.ts
│   │   └── routes
│   │       └── +page.svelte
│   ├── static
│   │   └── favicon.png
│   ├── svelte.config.js
│   ├── tests
│   │   └── test.ts
│   ├── tsconfig.json
│   └── vite.config.ts
├── cmd                    # Tools
│   ├── cli                # Used for creating supported modules
│   └── service            # Used for starting the web service
├── controllers            # Contains functions for handling HTTP Requests
│   └── general.go
├── go.mod
├── go.sum
├── initializers           # Used for creating things like database connections at initialization time
├── README.md
└── storage                # Contains functions for fetching data from databases
```

To register your SPA clients, just add them to this slice! This file can be found [here](/cmd/service/main.go)

``` go
// add any SPA clients here
var clients []*server.Client = []*server.Client{
	{
		Name:      "client",
		Prefix:    "/",
		OutputDir: "client/build",
		BuildCmd:  []string{"npm", "run", "build"},
	},
}
```
