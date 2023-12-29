This prompt will summarize a large document using rolling summarization (or refine) technique. 

## How to use in Haystack

### Split Documents into Chunks:
Before using the prompt, ensure that your documents are split into multiple chunks. Each chunk will undergo the rolling summarization process to generate a refined summary.

Note: The chunk size should be chosen carefully, considering both the content structure and cost implications. Avoid excessively small chunks for efficiency.

### Inputs to the Prompt:

The prompt takes three inputs:
1. Length: Specify the expected word count for the generated summary. This helps control the length of the output.
2. Context: Provide the rolling summary from the previous step. This ensures continuity in the summarization process.
3. Document: Input the text chunk of the current run. This is the content that will be considered in the latest summarization.