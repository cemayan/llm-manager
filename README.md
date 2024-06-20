## llm-manager


### Introduction

llm-manager allows the use of common open source LLM libraries with a common payload/input

There are 3 LLM backends available:
- langchaingo (openai)
- lingoose (openai)
- ollama


### Usage

> If you want to use ollama you need to setup ollama-server. You can download [here](https://ollama.ai/download) 

Before using API you need to set env some variables for backend.

```
export LLM_BACKEND=ollama;
export OPENAI_API_KEY=yourkey
```

---

Example api request: 

```bash
curl  'http://localhost:8996/api/v1/commit' \
--header 'Content-Type: application/json' \
--data '{
    "prompt":"Say hi"
}'

```

Example cli command usage:


```bash
go build -o bin/api cmd/main.go
LLM_BACKEND="ollama" ./bin/arm64/darwin/cmd -p "Say hi" | jq .
```

---

Example response from api/cli : 


```json
{
  "data": {
    "model": "llama2",
    "created_at": "2024-06-20T11:04:02.24971Z",
    "response": "Hi! *wave* How are you today?"
  }
}

```