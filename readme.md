# Baxter

## Motivation

Building reusable components for web platforms.
I hated the rest.

## How it works

All you need is the main event loop in your code

```pseudocode
function main
    initialize()
    while message != quit
        message := get_next_message()
        process_message(message)
    end while
end function
```

## Getting Started

```bash
go build ./
```
