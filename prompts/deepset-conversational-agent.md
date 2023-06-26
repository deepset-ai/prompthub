This prompt is intended to be used with a Haystack `Agent`, specificallly, a [`ConversationalAgent`](https://docs.haystack.deepset.ai/docs/agent#conversational-agent).

To learn how to use this prompt and the `ConversationalAgent`, check out our [Building a Conversational Chat App Tutorial](https://haystack.deepset.ai/tutorials/24_building_chat_app).

By default, the `ConversationalAgent` in Haystack would use this prompt without having to provide it using `prompt_template`. However, below is an example of how you might manually create a `PromptTemplate` with this prompt, and how you might use it with the ConversationalAgent.

## How to use in Haystack

1. Create a `PromptTemplate` that uses `deepset/conversational-agent`:

```python
import os

from haystack.nodes import PromptNode, PromptTemplate

conversational_agent_from_hub = PromptTemplate("deepset/conversational-agent")

prompt_node = PromptNode(
    "gpt-3.5-turbo",
    api_key=os.environ.get("OPENAI_API_KEY"),
    max_length=256,
    stop_words=["Observation:"],
)
```

2. Create a `ConversationalAgent` with `prompt_template`:

```python
from haystack.agents.conversational import ConversationalAgent

agent = ConversationalAgent(prompt_node=prompt_node,
                            prompt_template=conversational_agent_from_hub,
                            tools=[search_tool])
```