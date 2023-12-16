
This prompt is designed to detect `actionable items` present in a transcript of a meeting or normal text sequences. 
## How to use in Haystack
To use this prompt with Haystack, we recommend defining an `output_parser` with `AnswerParser()` in the `PromptTemplate`. This way, the result of the `PromptNode` will return Haystack `Answer` objects

```python
import os

from haystack.nodes import AnswerParser, PromptNode, PromptTemplate

action_item_detection_template = PromptTemplate("deepset/actionable-item-detection")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"))

prompt_node.prompt(prompt_template=action_item_detection_template, transcript="YOUR_TRANSCRIPT")
```

> Make sure to replace `"YOUR_TRANSCRIPT"` with the actual meeting transcript or sentence you want to analyze.