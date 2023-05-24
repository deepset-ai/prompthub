This prompt is designed to categorize each document provided in `documents` based on the categories provided in `options`.

## How to use in Haystack

```python 
import os

from haystack.nodes import PromptNode, PromptTemplate

topic_classifier_template = PromptTemplate("deepset/topic-classification")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=topic_classifier_template, documents="YOUR_DOCUMENTS", options=["A LIST OF TOPICS"])
```

### Example

For example, with a list of news articles defined in `news`, we may use this template as follows:

```python
prompt_node.prompt(prompt_template=topic_classifier_template, documents=news, options=["economics", "politics", "culture"])
```