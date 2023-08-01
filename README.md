# Go+SvelteKit Template

This is a template application. I hope to create a functioning CLI tool for bootstrapping these types of applications. The idea is that you'll be able to optionally add in TailwindCSS or multiple SPA's served by the Go backend using Echo v4. The Go backend can contain any required database logic and scheduled task execution (business logic). The goal here is to create a reliable starter app providing a middle ground between performance and developer experience.

```bash
.
├── client                  # SvelteKit SPA
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
