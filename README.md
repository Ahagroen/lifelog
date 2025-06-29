# Lifelog
A small go CLI for logging your life like a program, with as little friction as possible.

## Installation

`go install github.com/ahagroen/lifelog` or via the pre-built binaries in the releases section. Lifelog will by default place the .txt and config .json files in the `~/.lifelog` directory (or on windows `C:users\$USER\.lifelog` by default). It is recommended to alias the `log` command for minimum friction in logging.

## Usage
`lifelog log [message]` can be used to add a new log message to the file. Message can include inline tags beginning with `@`, which can then be filtered on later.

`lifelog show [filter]` can be used to output the current log file (can be piped to less/grep etc.). Filter is an optional filter for tags
