This prompt is designed to answer a `query` from a set of given `options`. 

## How to use in Haystack

To use this prompt with Haystack, we recommend defining an `output_parser` with `AnswerParser()` in the `PromptTemplate`. This way, the result of the `PromptNode` will return Haystack `Answer` objects

```python
import os

from haystack.nodes import AnswerParser
from haystack.nodes import PromptNode, PromptTemplate

multiple_choice = PromptTemplate("deepset/multiple-choice-question-answering", output_parser=AnswerParser())

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=multiple_choice, query="YOUR_QUERY", options="YOUR_OPTIONS")
```

### Example

```python
prompt_node.prompt(prompt_template=multiple_choice, query="What is the capital of France?", options=["Paris", "London", "Istanbul"])
```