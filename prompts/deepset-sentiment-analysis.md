For each document in the provided `documents`, tells you whether it is 'postivie', 'negative' or 'neutral'.

## How to use in Haystack

```python
import os

from haystack.nodes import PromptNode, PromptTemplate

sentiment_analysis_template = PromptTemplate("deepset/sentiment-analysis")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=sentiment_analysis_template, documents="YOUR_DOCUMENTS")
```

### Example

For example, with the following documents:

```python
from haystack.schema import Document

doc1 = Document("I am super happy")

doc2 = Document("Today sucks")

doc3 = Document("There is nothing wrong with what's going on")

documents = [doc1, doc2, doc3]
```

```python
prompt_node.prompt(prompt_template=sentiment_analysis_template, documents=documents)

```

Result:
```
['Positive', 'Negative.', 'Neutral']

```