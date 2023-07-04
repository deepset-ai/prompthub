This prompt is simply designed to answer a `query` given a set of `documents`. There will be 1 answer generated.

## How to use in Haystack
To use this prompt with Haystack, we recommend defining an `output_parser` with `AnswerParser()` in the `PromptTemplate`. This way, the result of the `PromptNode` will return Haystack `Answer` objects

```python
import os

from haystack.nodes import AnswerParser, PromptNode, PromptTemplate

question_answering_template = PromptTemplate("deepset/question-answering", output_parser=AnswerParser())

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=question_answering_template, query="YOUR_QUERY", documents="YOUR_DOCUMENTS")

```
