# CSV2YAML

This is a tool to convert CSV files into YAML.

## Things Need To Know

For now it treat first column as list head and other columns as its child.

## Parameters

This tool has two parameter one is `-s` source csv file & `-d` destination yaml file. With `-h` or `--help` you can find details.


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
