# ExtFieldGen
ExtFieldGen is a gqlgen plugin which enables you to specify extra fields with directives.

## Usage
```go
package main

import (
    "fmt"
    "os"
    "github.com/99designs/gqlgen/api"
    "github.com/99designs/gqlgen/codegen/config"
    "github.com/Warashi/extfieldgen"
)

func main() {
    cfg, err := config.LoadConfigFromDefaultLocations()
    if err != nil {
        fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
        os.Exit(2)
    }
    if err := api.Generate(cfg, api.PrependPlugin(extfieldgen.New())); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(3)
    }
}
```

## Schema Example
```graphql
type Foo {
  bar: String!
}

type Bar 
  @extraField(name: "FooBar", type: "*Foo")
  @extraField(name: "FooBarBaz", type: "[]*Foo")
  @extraField(name: "RecursiveBar", type: "*Bar")
  @extraField(name: "SomeFieldWithDescription", type: "*Bar", description: "SomeFieldWithDescription is example field")
  @extraField(name: "SomeIntField", type: "int")
  @extraField(name: "OtherPackageType", type: "github.com/you/pkg/model.User")
{
  baz: String!
}
```
