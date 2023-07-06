This prompt is intended to be used in a `PromptNode` which receives documents with corresponding scores. For example, a pipeline with a `Retriever` which returnes scored documents is a good candidate to use in combination with this prompt.

## How to use in Haystack

For example, given a query, the `WebRetriever` will return documents from the web, with scores. This prompt might be used in the following pipeline:

```python
import os

from haystack.nodes import PromptNode, PromptTemplate, WebRetriever

question_anwering_with_scores = PromptTemplate("deepset/question-answering-with-document-scores")

prompt_node =  PromptNode("text-davinci-003",
        api_key=os.environ.get("OPENAI_API_KEY"),
        max_length=256,
        default_prompt_template=question_anwering_with_scores)

retriever = WebRetriever(api_key=os.environ.get("SERPER_API_KEY"))

```

```python
from haystack import Pipeline

pipe = Pipeline()

pipe.add_node(component=retriever, name="WebRetriever", inputs=['Query'])
pipe.add_node(component=prompt_node, name="QAwithScoresPrompt", inputs=['WebRetriever'])
```

```python
pipe.run(query="What years was Obama president?")
```
