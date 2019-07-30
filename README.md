# Telegraf Example Plugin

This is an example repo containing an external, runtime loaded plugin using
Go's plugin package.  This feature requires Telegraf 1.12 and is considered
experimental.

### Building
```
make
```

### Loading
```
telegraf --config telegraf.conf --plugin-directory ./path/to/shared-objects
```
