## Go Transcriber

### Introduction

This transcriber produces text from sound records, and it is support websocket.

- ASSEMBLY AI transcriber
- WebSocket

### Usage

To run this project, you need to generate assembly-ai api key. You can generate it from https://www.assemblyai.com/app
After generating api_key, you should past it in env file (<ASSEMBLY_AI_API_KEY>)

```
make run
```

#### Other commands

If you want to check coverage output on browser you can use the code below

```
make test_with_profile
```

### Tests 