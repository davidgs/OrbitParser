# OrbitParser

Parse output from the Orbit JSON into a `.csv`

It doesn't pull every field from the JSON, but you can add other fields to the `struct` to have them parsed as well.

## Usage:

`go run Parser.go` will read from the Orbit API -- You will have filled in your organization and API_KEY first -- until all records are
exhausted.

### Runtime Options:
 `-file=FILE.json` will read from a local JSON file rather than the Orbit API
 `-out=FILE.csv` will write the output to a named file.

### Important

You **must** Edit the file to add your Orbit API_KEY and Organization
