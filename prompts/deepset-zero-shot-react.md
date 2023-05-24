This prompt is intended to be used with a Haystack `Agent`. It's the backbone of an LLM powered `Agent` and expects to be provided `Tools`.

To learn how to use this prompt and an `Agent` check out our [Answering Multihop Questions with Agents Tutorial](https://haystack.deepset.ai/tutorials/23_answering_multihop_questions_with_agents).

## How to use in Haystack

To use this prompt from the PromptHub in an Agent:

```python
import os

from haystack.agents import Agent
from haystack.nodes import PromptNode, PromptTemplate

zero_shot_agent_template = PromptTemplate("deepset/zero-shot-react")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"), stop_words=["Observation:"])
agent = Agent(prompt_node=prompt_node, prompt_template=zero_shot_agent_template)

agent.run("Your query")
```