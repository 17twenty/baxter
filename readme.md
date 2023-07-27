# Baxter

## Motivation

Building reusable components for web platforms.
I hated the rest.

## How it works

You instantiate your baxter instance and 

```go
// Setup baxter. Choose the backend that makes most
// sense for your use case.
baxter.Init(baxter.InMemory(10)) 

// Subscribe to an event
baxter.Subscribe("event.test", func(event string, meta json.RawMessage) {
    log.Println("Caught event", string(meta))
})

// Start baxters message pump
baxter.Start()

// Publish an event to baxter
baxter.Publish("event.test", []byte("hello"))

// Shutdown baxters message pump
baxter.Stop()
```

## Getting Started

```bash
go build ./cmd/baxter-cli/ && ./baxter-cli 
```

## Todo

- Move to taking an `any` and doing the marshal/unmarshal into JSON automagically
- Threadsafe (use the tools)
- Add the missing magic to safely stop the message pump based on [this article](https://gist.github.com/uudashr/3cf820e3ba902d3c6387abc82c815e66).
- Provide an auth example showing how anyone can incorporate baxter for loosely coupled products.
- Provide another backend
