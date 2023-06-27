This prompt is intended to be used with a Haystack `Agent`, specifically, a [`ConversationalAgent`](https://docs.haystack.deepset.ai/docs/agent#conversational-agent).

To learn how to use this prompt and the `ConversationalAgent`, check out our [Building a Conversational Chat App Tutorial](https://haystack.deepset.ai/tutorials/24_building_chat_app).

By default, the `ConversationalAgent` without any tools would use this prompt. If you want to more flexibility, below is an example of how you might use this prompt with the `Agent`.

## How to use in Haystack

```python
import os

from haystack.nodes import PromptNode
from haystack.agents.utils import agent_without_tools_parameter_resolver
from haystack.agents.memory import ConversationMemory
from haystack.agents import Agent

prompt_node = PromptNode("gpt-3.5-turbo", api_key=os.environ.get("OPENAI_API_KEY"), max_length=256)
memory = ConversationMemory()

conversational_agent_without_tools = Agent(
    prompt_node=prompt_node,
    memory=memory,
    prompt_template="deepset/conversational-agent-without-tools",
    prompt_parameters_resolver=agent_without_tools_parameter_resolver,
    final_answer_pattern=r"^([\s\S]+)$",
)
```