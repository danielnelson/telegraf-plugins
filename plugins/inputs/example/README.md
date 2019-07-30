# Example Input Plugin

The most basic plugin that could exist.

### Configuration

```toml
[[inputs.example]]
  value = 123
```

### Metrics

- example
  - tags:
    - source
  - fields:
    - value (integer)

### Example Output

```
example,host=example.org,source=example.org value=123i 1564463260000000000
```
