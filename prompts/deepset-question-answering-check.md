This prompt is designed to check whether a query can be answered with the provided context or not. The `PromptNode` that uses this prompt will reply with "yes" or "no"

## How to use in Haystack

To use this prompt with Haystack, we recommend defining an `output_parser` with `AnswerParser()` in the `PromptTemplate`. This way, the result of the `PromptNode` will return Haystack `Answer` objects

```python
import os

from haystack.nodes import AnswerParser, PromptNode, PromptTemplate

question_answering_check = PromptTemplate("deepset/question-answering-check", output_parser=AnwerParser())

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"),)

prompt_node.prompt(prompt_template=question_answering_check, query="YOUR_QUERY", documents="YOUR_documents")
```

### Example

With the following `documents`:

```python
from haystack.schema import Document

doc1 = Document("Paris is the capital of France")

doc2 = Document("Haystack is an NLP framework")

doc3 = Document("deepset is the company behind Haystack")

doc4 = Document("Obama was the president of the USA from 2009 to 2017")

documents = [doc1, doc2, doc3, doc4]
```

```python
prompt_node.prompt(prompt_template=question_answering_check, query="What is the capital of France?", documents=documents)
```

```
[<Answer {'answer': 'Yes', 'type': 'generative', 'score': None, 'context': None, 'offsets_in_document': None, 'offsets_in_context': None, 'document_ids': ['6d0bfeb903d67c894289b56dfe046038'], 'meta': {'prompt': 'Does the following context contain the answer to the question? Context: Paris is the capital of France; Question: What is the capital of France?; Please answer yes or no! Answer:'}}>,
 <Answer {'answer': 'No', 'type': 'generative', 'score': None, 'context': None, 'offsets_in_document': None, 'offsets_in_context': None, 'document_ids': ['5d5c5f51a9e47775025fcbb6864b4ca5'], 'meta': {'prompt': 'Does the following context contain the answer to the question? Context: Haystack is an NLP framework; Question: What is the capital of France?; Please answer yes or no! Answer:'}}>,
 <Answer {'answer': 'No', 'type': 'generative', 'score': None, 'context': None, 'offsets_in_document': None, 'offsets_in_context': None, 'document_ids': ['e3095e3756210b55a6678a241b1b2bea'], 'meta': {'prompt': 'Does the following context contain the answer to the question? Context: deepset is the company behind Haystack; Question: What is the capital of France?; Please answer yes or no! Answer:'}}>,
 <Answer {'answer': 'No', 'type': 'generative', 'score': None, 'context': None, 'offsets_in_document': None, 'offsets_in_context': None, 'document_ids': ['35812bb4434ee594b0493d758d4b818e'], 'meta': {'prompt': 'Does the following context contain the answer to the question? Context: Obama was the president of the USA from 2009 to 2017; Question: What is the capital of France?; Please answer yes or no! Answer:'}}>]
```