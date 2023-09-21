This prompt is designed to produce a summary for each document provided in `documents`.
It uses a technique called "chain of density" (from https://arxiv.org/pdf/2309.04269.pdf), that iteratively refines the summary in each step.
The prompt yields 5 summaries of the article with increasing density. Typically, users prefer summary 4 or 5.

## How to use in Haystack

```python
from haystack.nodes import PromptNode, PromptTemplate

prompt_node = PromptNode(model_name_or_path="gpt-4", max_length=3000, api_key="<YOUR_API_KEY>", model_kwargs={"stream":True})
summarization_template = PromptTemplate("deepset/summarization-chain-of-density")

res = prompt_node.prompt(prompt_template=summarization_template, documents=["YOUR_DOCUMENT"])
print(res)
```