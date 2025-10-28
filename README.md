# Pulumi Forgejo provider

This is at a very early stage. Please don't use it. It'll only disappoint you.

## What is implemented so far?

- Nothing
- Nada
- Not a sausage

## What still needs to be implemented

- `*`

## What can it be used for currently?

- Taking up space on your harddrive.
- Consuming inodes

# Generating OpenAPI v3.0 spec from Swagger v2

ForgeJo spec: https://code.forgejo.org/swagger.v1.json

Swagger provide a tool to convert swagger to openapi specs. To generate `openapi.yml` from the published swagger spec:

```bash
curl "https://converter.swagger.io/api/convert?url=https://code.forgejo.org/swagger.v1.json" -H "Accept: application/yaml" -o ./provider/cmd/pulumi-gen-forgejo/openapi.yml
```

Checking the generated file:

```bash
go run github.com/getkin/kin-openapi/cmd/validate@latest -- ./provider/cmd/pulumi-gen-forgejo/openapi.yml
```



