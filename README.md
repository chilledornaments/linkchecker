# linkchecker

## Building

`go build`!

## Running

`./linkchecker https://example.com`

### Batches

Simple solution with Bash:

```shell
for site in https://example.com https://example.com/test; do ./linkchecker $site; done
```

## Outputs

### OK

Checking link https://example.com

Link ok for https://example.com

### 404

Checking link https://example.com/test123456

404 found for https://example.com/test123456