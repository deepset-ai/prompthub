Given a set of `document` and a single `query`, this prompt is designed to answer the query once per document.

## How to use in Haystack

To use this prompt with Haystack, we recommend defining an `output_parser` with `AnswerParser()` in the `PromptTemplate`. This way, the result of the `PromptNode` will return Haystack `Answer` objects

```python
import os

from haystack.nodes import AnswerParser, PromptNode, PromptTemplate

question_answering_per_doc = PromptTemplate("deepset/question-answering-per-document", output_parser=AnwerParser())

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=question_answering_per_doc, query="YOUR_QUERY", documents="YOUR_documents")
```

### Example

With the following `documents`:

```python
from haystack.schema import Document

doc1 = Document("The main suspect is John")

doc2 = Document("The main suspect is Jane")

documents = [doc1, doc2]
```

```python
prompt_node.prompt(prompt_template=question_answering_per_doc, query="Who is the main suspect?", documents=documents)
```

```
[<Answer {'answer': 'John is the main suspect.', 'type': 'generative', 'score': None, 'context': None, 'offsets_in_document': None, 'offsets_in_context': None, 'document_ids': ['8bb5a67252fea7718402141b6d256512'], 'meta': {'prompt': 'Given the context please answer the question. Context: The main suspect is John; Question: Who is the main suspect?; Answer:'}}>,
 <Answer {'answer': 'Jane is the main suspect.', 'type': 'generative', 'score': None, 'context': None, 'offsets_in_document': None, 'offsets_in_context': None, 'document_ids': ['736881f380410324ca181100caefc35f'], 'meta': {'prompt': 'Given the context please answer the question. Context: The main suspect is Jane; Question: Who is the main suspect?; Answer:'}}>]
```