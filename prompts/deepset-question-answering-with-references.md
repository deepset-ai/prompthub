This prompt is designed to answer a `query` given a list of `documents`. It will also reference which document the final answer came from. For example, if the answer was in the first document, the first document, the restult will indicate that the answer was 'stated in Document[1]'

## How to use in Haystack

To use this prompt with Haystack, we recommend defining an `output_parser` with `AnswerParser(reference_pattern=r"Document\[(\d+)\]")` in the `PromptTemplate`. This way, the result of the `PromptNode` will return Haystack `Answer` objects.

```python
import os

from haystack.nodes import AnswerParser, PromptNode, PromptTemplate

question_answering_with_references = PromptTemplate("deepset/question-answering-with-references", output_parser=AnswerParser(reference_pattern=r"Document\[(\d+)\]"))

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=question_answering_with_references, query="YOUR_QUERY", documents="YOUR_DOCUMENTS")
```