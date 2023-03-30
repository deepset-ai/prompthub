## How to use in Haystack
Here's an example of how this prompt is intended to be used with Haystack. This `PromptTemplate` is best used alongside a `Shaper` in an `output_shapers` variable.

```python
from haystack.nodes import PromptTemplate

question_answering_template = PromptTemaplte(name="deepset/question-answering", output_shapers=[Shaper(func="strings_to_answers", outputs=["answers"], inputs={"strings": "results"})])

```