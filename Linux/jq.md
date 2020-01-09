# JQ

<https://github.com/stedolan/jq>

Command-line JSON processor

## error handling

```sh
# Include All Errors in output as a string
cat example.log | jq -R '. as $raw | try fromjson catch $raw'

# Skip all errors and keep going
cat example.log | jq -R 'fromjson?'
```
