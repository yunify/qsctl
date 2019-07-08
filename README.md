# qsctl

qsctl is intended to be an advanced command line tool for QingStor, it provides
powerful unix-like commands to let you manage QingStor resources just like files
on local machine. Unix-like commands contains: cp, ls, mb, mv, rm, rb, and sync.
All of them support batch processing.

## Installation

### Binary

Visit <https://github.com/yunify/qsctl/releases> to get latest releases.

## Getting Started

To use qsctl, there must be a configuration file , for example

```yaml
access_key_id: 'ACCESS_KEY_ID_EXAMPLE'
secret_access_key: 'SECRET_ACCESS_KEY_EXAMPLE'
```

The configuration file is `~/.qingstor/config.yaml` by default, it also
can be specified by the option `-c /path/to/config`.

You can also config other option like `host` , `port` and so on, just
add lines below into configuration file, for example

```yaml
host: 'qingstor.com'
port: 443
protocol: 'https'
connection_retries: 3
# Valid levels are 'debug', 'info', 'warn', 'error', and 'fatal'.
log_level: 'debug'
```

## Available Commands

Commands supported by qsctl are listed below:

- `cat`: Cat a remote object into stdout.
- `cp`: Copy local file(s) to QingStor or QingStor key(s) to local.
- `tee`: Tee from stdin to a remote object.


## Examples

Stat key in bucket <mybucket> by running

```bash
:) qsctl stat qs://mybucket/test
Key         : test
Size        : 10GB
Type        : application/octet-stream
Modify      : 2019-07-04 09:50:46 +0000 UTC
StorageClass: STANDARD
```

See the detailed usage and more examples with 'qsctl help' or 'qsctl <command> help'.