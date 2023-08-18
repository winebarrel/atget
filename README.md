# atget

CLI to get airtable records.

## Usage

```
Usage: atget --token=STRING --base=STRING --table=STRING <filter>

Arguments:
  <filter>    A formula used to filter records. see https://support.airtable.com/docs/formula-field-reference

Flags:
  -h, --help            Show context-sensitive help.
      --token=STRING    Airtable API token ($AIRTABLE_TOKEN).
  -b, --base=STRING     Base name.
  -t, --table=STRING    Table name.
      --version
```

```sh
$ atget -b company -t employees '{email}="scott@example.com"'
{"first_name":"scott","last_name":"tiger","email":"scott@example.com"}
```
