# ZeroTears
Golang ZeroTier Controller client.
Tears as in rain, not torn buttholes.

## Design notes
- Pass auth secret and hostname in config file
- Operation name as subcommand, options then apply to operation

`zerotears info` < to call /controller and list info

- Output should be listed mostly as tables, colourise these
- Input as command line args as much as possible
- Function to dump controller config so it can be reuploaded somewhere else?
- Option for JSON output instead of tables

## Implementation
- Generic functions for prepping request and returning response
- Operation specific function calls generic 'request' function
- Create structs for response types so they can be decoded
    - Structs can have methods to format tables for them
    - Output methods should have verbose mode implemented as a different method?
