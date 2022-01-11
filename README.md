# fill
is simple tool to fill a go template using the provided environment variables.

Please provide a golang template file (see: https://golangdocs.com/templates-in-golang ), 
using the -t flag ( mandatory )

You can refer the variables by a "." (dot) prefix

 e.g. the following will print out the current content of PATH on your runtime environment.
```
{{.PATH}}
```

Happy templating !!!
