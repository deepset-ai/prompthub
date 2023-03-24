# Prompt Hub

Prompt Hub serves a collection of ready-made prompts for the most common NLP tasks. The service is deployed at the
public URL https://api.prompthub.deepset.ai

## Prompt format

A prompt is defined in a yaml file with the following format:
```yaml
name: an-example
prompt_text: Your prompt text goes here
description: A brief description of what your prompt is for
tags:
  - translation
meta:
  authors:
    - your name
version: v0.1
```

## Prompt Hub API

### Get all the prompts

`GET /prompts`

Request:
```sh
curl -i -H 'Accept: application/json' http://api.prompthub.deepset.ai/prompts
```

Response:
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 24 Mar 2023 07:59:55 GMT
Content-Length: 42

[
    {
        "name":"deepset/an-example-prompt",
        "tags":["question-answering"],
        "meta":{"authors":["Alice","Bob"]},
        "version":"1.0",
        "prompt_text":"My prompt text",
        "description":"Provides a prompt for question answering with references to documents"
    }
]
```

### Get a specific prompt by name

`GET /prompts/prompt-name`

Request:
```sh
curl -i -H 'Accept: application/json' http://api.prompthub.deepset.ai/prompts/prompt-name
```

Response:
```
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Fri, 24 Mar 2023 08:06:19 GMT
Content-Length: 211

{"name":"prompt-name","tags":["translation"],"meta":{"authors":["vblagoje"]},"version":"v0.1.0","prompt_text":"Your prompt text goes here","description":"Prompt to translate text into a target language"}
```
