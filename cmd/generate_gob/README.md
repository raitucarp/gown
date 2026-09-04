# Command generate_gob

`generate_gob` is an offline data processing utility used by developers and maintainers of Gown to compile Open English WordNet (OEWN) XML database dumps into the compressed, embedded GOB binary (`oewn.gob.gz`) packaged with Gown.

## Overview

WordNet distributions in XML format (`wn.xml`) are hundreds of megabytes in size and require significant time to parse. The `generate_gob` command:

1. Reads the raw XML lexical resource file (`./data/wn.xml`).
2. Deserializes the lexical entries, synsets, relations, and syntactic behaviors.
3. Encodes the typed structures using Go's binary `encoding/gob` format.
4. Compresses the output using `gzip` and writes it to `oewn.gob.gz`.

End users of the `gown` library do not need to run this command. Gown automatically embeds the generated binary data into the compiled library via Go's `//go:embed` directive.

## Usage

```bash
cd cmd/generate_gob
go run main.go
```
