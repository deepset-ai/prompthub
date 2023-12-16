This prompt is designed to produce a summary for each document provided in `documents`

## How to use in Haystack

```python
import os

from haystack.nodes import PromptNode, PromptTemplate

summarization_template = PromptTemplate("deepset/summarization")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=summarization_template, documents="YOUR_DOCUMENTS")
```
