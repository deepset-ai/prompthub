This prompt is designed to detect the language of given documents. 

## How to use in Haystack

1. Create a `PromptNode`. There are many models to pick from that can be used with `PromptNode`. Here is an example of one that uses `text-davinci-003`.

```python
from haystack.nodes import PromptTemplate, PromptNode

language_detection = PromptTemplate("deepset/language-detection")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key=os.environ.get("OPENAI_API_KEY"),)
```

2. Prompt is with some documents for context

```python
document = [MY_DOCUMENTS]

prompt_node.prompt(prompt_template=language_detection, documents=documents)
```

### Example

With the following `documents`:

```python
from haystack.schema import Document

en_doc = Document("This is a document with English content")

tr_doc = Document("İstanbul'da hava yağışlı")

fr_doc = Document("Bonjour, comment allez-vous?")

de_doc = Document("Ich werde bald in Berlin sein.")

documents = [en_doc, tr_doc, fr_doc, de_doc]
```

```python
prompt_node.prompt(prompt_template="language-detection", documents=documents)
```

```
['English', 'Turkish', 'French', 'German']
```