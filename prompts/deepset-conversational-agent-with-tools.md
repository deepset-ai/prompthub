This prompt is intended to be used with a Haystack `Agent`, specificallly, a (`ConversationalAgentWithTools`)[]

By default, the `ConversationalAgentWithTools` in Haystack would use this prompt without you having to provide it via `prompt_template`. However, below is an example of how you might manually create a `PromptTemplate` with this prompt, and how you might use it with the Agent.

## How to use in Haystack]

1. Create a `PromptNode` that uses `deepset/conversational-agent-with-tools` as the `default_prompt_template`.

```python
import os

from haystack.agent import ConversationalAgentWithTools
from haystack.nodes import PromptNode, PromptTemplate

conversational_agent_with_tools_from_hub = PromptTemplate("deepset/conversational-agent-with-tools")

prompt_node = PromptNode(
    "gpt-3.5-turbo",
    api_key=os.environ.get("OPENAI_API_KEY"),
    max_length=256,
    stop_words=["Observation:"],
)
```

2. Create a `ConversationalAgentWithTools`

```python
from haystack.agents import ToolsManager

agent = ConversationalAgentWithTools(prompt_node=prompt_node, prompt_template=conversational_agent_with_tools_from_hub, tools_manager=ToolsManager(tools=[YOUR_TOOLS]))
```