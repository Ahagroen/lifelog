# Lifelog
A small go CLI for logging your life like a program

## Usage
`lifelog log [message]` can be used to add a new log message to the file. Message can include inline tags beginning with `@`, which can then be filtered on later.

`lifelog show [filter]` can be used to output the current log file (can be piped to less/grep etc.). Filter is an optional filter for tags
