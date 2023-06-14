# PromptHub

PromptHub serves a collection of ready-made prompts for the most common NLP tasks. The service is deployed at the
public URL https://api.prompthub.deepset.ai and there's a nice UI to browse the prompts at https://prompthub.deepset.ai.

## Prompt format

A prompt is defined in a yaml file with the following format:

```yaml
name: an-example
text: Your prompt text goes here
description: A brief description of what your prompt is for
tags:
  - translation
meta:
  authors:
    - your name
version: v0.1
```

## PromptHub API

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

### Get a model card by model name

Request:

```sh
curl -i http://api.prompthub.deepset.ai/cards/prompt-name
```

Response:

```
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Vary: Origin
Date: Wed, 07 Jun 2023 10:46:53 GMT
Content-Length: 762

This prompt is simply designed to answer a `query` given a set of `documents`. There will be 1 answer generated.

...
```

## API clients

You can consume the PromptHub API natively from one of the supported languages:
- Python: https://github.com/deepset-ai/prompthub-py
- Typescript: https://github.com/deepset-ai/prompthub-ts

## Run the PromptHub API by yourself

The easiest way to run the PromptHub API on your premises is to use the official Docker image:

```sh
$ docker run -p80:80 deepset/prompthub
INFO Prompthub running at 0.0.0.0:80
```

You can also serve your very own set of prompts by overriding the default `prompts` folder in the container:

```sh
$ docker run -p80:80 --volume $PWD/path/to/your/prompts:/prompts deepset/prompthub
INFO Prompthub running at 0.0.0.0:80
```

## Development

You can build and run the PromptHub API service locally following these steps:

- Ensure you have a recent version of [Go](https://go.dev) installed:

```sh
$ go version
go version go1.20.2 darwin/arm64
```

    If you get an error, follow the install instructions from the [official documentation](https://go.dev/doc/install) and
    try again.

- From the root of the repo, build the binary with:

```sh
$ go build
```

- Run the server with:

```sh
$ ./prompthub
```

- Run the tests with:

```sh
$ go test ./...
```
