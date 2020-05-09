# CSV2YAML

This is a tool to convert CSV files into YAML.

## Things Need To Know

For now it treat first column as list head and other columns as its child.

## Parameters

This tool has two parameters, one is `-s` for source csv file & `-d` for destination yaml file. With `-h` or `--help` you can find details.


```cassandraql
$ csv2yaml --help

This tool is used to convert csv file to yaml

Usage: csv2yaml -s [csvfile] -d [yaml file]

Usage:
  csv2yaml [flags]

Flags:
  -d, --destination string   Yaml File Name or Path To Save File.
  -h, --help                 help for csv2yaml
  -s, --source string        Source CSV Filename or File Path
Please add '.yml' or '.yaml' to destination file

```

## How To Install

Just download the binary according to your platform and make it executable to run it.


[Mac](https://github.com/mohsinzaheer25/csv2yaml/releases/download/V1.0/csv2yaml)

[Linux](https://github.com/mohsinzaheer25/csv2yaml/releases/download/V1.0/csv2yaml_linux)

[Windows](https://github.com/mohsinzaheer25/csv2yaml/releases/download/V1.0/csv2yaml.exe)

