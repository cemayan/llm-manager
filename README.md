## git observer


### Introduction

git-observer allows editing commit messages on Git via LLM models and git hooks when developer wants to commit.

There are 3 LLM backends available:
- langchaingo (openai)
- lingoose (openai)
- ollama


### Usage

> If you want to use ollama you need to setup ollama-server. You can download [here](https://ollama.ai/download) 

> You can find the prepare-commit-msg [here](script/prepare-commit-msg). Default git hook location is .git/hooks. In order to use pre hook you need to copy this file to under .git/hooks folder.

Before using API you need to set env some variables for backend.

```
export LLM_BACKEND=ollama;
export OPENAI_API_KEY=yourkey
```

Example request: 

```bash
curl  'http://localhost:8996/api/v1/commit' \
--header 'Content-Type: application/json' \
--data '{
    "prompt":"You are a developer and your responsible is fix the commit message. You can check the guidline here https://wiki.openstack.org/wiki/GitCommitMessages#Information_in_commit_messages. Can you fix '"$(cat "$1")""
}'

```

For cli


```bash
go build -o bin/api cmd/main.go
LLM_BACKEND="ollama" ./bin/api -p "Say hi" | jq .
```
