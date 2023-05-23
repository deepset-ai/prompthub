This prompt is designed to accept a `target_language` and translate each document procided in `documents` to this target language.

## How to use in Haystack

```python
import os

from haystack.nodes import PromptNode, PromptTemplate

translation_template = PromptTemplate("deepset/translation")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=translation_template, documents=["YOUR_DOCUMENTS"], target_language="your_target_language")
```

### Example

For example, let's assume we have a set of english documents in `english_documents` that we want to translate to French.

```python
prompt_node.prompt(prompt_template=translation_template, documents=english_documents, target_language="french")
```