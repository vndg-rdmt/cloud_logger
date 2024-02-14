# go.cloud_logger [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

## Install

```bash
go get -u github.com/vndg-rdmt/go.cloud_logger
```

## Config

### env (oriented)
```ini
CLOUDLOGGER_LEVEL=
CLOUDLOGGER_FIELD_MESSAGE=
CLOUDLOGGER_FIELD_NAME=
CLOUDLOGGER_FIELD_LEVEL=
CLOUDLOGGER_FIELD_TIMESTAMP=
CLOUDLOGGER_ENDLINE=
CLOUDLOGGER_STDOUT=
CLOUDLOGGER_SKIP_INIT_FAILURE=
CLOUDLOGGER_GRAYLOG_ADDR=
CLOUDLOGGER_ENV_FIELDS=
```

### json
```json
{
    "level": "",
    "field_message": "",
    "field_name": "",
    "field_level": "",
    "field_timestamp": "",
    "endine": "",
    "static_fields": [
        {
            "key": "",
            "value": ""
        }
    ],
    "stdout": false,
    "writers": "",
    "skip_init_failure": "",
    "grayog_addr": "",
    "env_fields": "",
    "key": "",
    "value": "",
    "key": "",
    "value": "",
    "level": "",
    "level": "",
    "path": "",
    "yaml": ""
}
```

### yml/yaml
```yml
level:
field_message:
field_name:
field_level:
field_timestamp:
endine:
static_fields:
stdout:
writers:
skip_init_failure:
grayog_addr:
env_fields:
key:
value:
key:
value:
level:
level:
path:
yaml:
```


# Featured

- Currently disabled file writers.

[ci-img]: https://github.com/vndg-rdmt/go.cloud_logger/actions/workflows/ci.yml/badge.svg

[ci]: https://github.com/vndg-rdmt/go.cloud_logger/actions/workflows/ci.yml

[cov-img]: https://codecov.io/gh/vndg-rdmt/go.cloud_logger/branch/main/graph/badge.svg?token=P0ZUmJlZCV

[cov]: https://codecov.io/gh/vndg-rdmt/go.cloud_logger#go.cloud_logger