This prompt creates a single question for each document in the provided `documents`

## How to use in Haystack

```python 
import os

from haystack.nodes import PromptNode, PromptTemplate

question_generation_template = PromptTemplate("deepset/question-answering")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=question_generation_template, documents="YOUR_DOCUMENTS")
```

