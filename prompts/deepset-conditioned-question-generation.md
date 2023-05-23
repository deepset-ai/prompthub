This prompt aims to generate a question, given `documents` and an `answers`.

## How to use in Haystack

```python
from haystack.nodes import PromptNode, PromptTemplate

documents = [MY_DOCUMENTS]
answers = [MY_ANSWERS]

conditioned_question_generation_from_hub = PromptTemplate("deepset/conditioned-question-generation")

prompt_node = PromptNode(model_name_or_path="text-davinci-003", api_key='YOUR_API_KEY')

prompt_node.prompt(prompt_template=conditioned_question_generation_from_hub, documents=documents, answers=answers)
```

### Example

With the following news snippets as documents and answers:

```python
from haystack.schema import Document

# https://www.theguardian.com/business/2023/feb/12/inflation-may-have-peaked-but-the-cost-of-living-pain-is-far-from-over
news_economics = Document(
    """At long last, Britain’s annual inflation rate is on the way down. After hitting the highest level since the 1980s, heaping pressure on millions of households as living costs soared, official figures this week could bring some rare good news.
City economists expect UK inflation to have cooled for a third month running in January – the exact number is announced on Wednesday – helped by falling petrol prices and a broader decline in the global price of oil and gas in recent months. The hope now is for a sustained decline in the months ahead, continuing a steady drop from the peak of 11.1% seen in October.
The message from the Bank of England has been clear. Inflation is on track for a “rapid” decline over the coming months, raising hopes that the worst of Britain’s cost of living crisis is now in the rearview mirror.
There are two good reasons for this. Energy costs are moving in the right direction, while the initial rise in wholesale oil and gas prices that followed Russia’s invasion of Ukraine in February last year will soon drop from the calculation of the annual inflation rate."""
)

# https://www.theguardian.com/science/2023/feb/13/starwatch-orions-belt-and-sirius-lead-way-to-hydras-head
news_science = Document(
    """On northern winter nights, it is so easy to be beguiled by the gloriously bright constellations of Orion, the hunter, and Taurus, the bull, that one can overlook the fainter constellations.
So this week, find the three stars of Orion’s belt, follow them down to Sirius, the brightest star in the night sky, and then look eastward until you find the faint ring of stars that makes up the head of Hydra, the water snake. The chart shows the view looking south-east from London at 8pm GMT on Monday, but the view will be similar every night this week.
Hydra is the largest of the 88 modern constellations covering an area of 1,303 square degrees. To compare, nearby Orion only covers 594 square degrees. Hydra accounts for most of its area by its length, crossing more than 100 degrees of the sky (the full moon spans half a degree).
As evening becomes night and into the early hours, the rotation of Earth causes Hydra to slither its way across the southern meridian until dawn washes it from the sky. From the southern hemisphere, the constellation is easily visible in the eastern sky by mid-evening."""
)

# https://www.theguardian.com/music/2023/jan/30/salisbury-cathedral-pipe-organ-new-life-holst-the-planets
news_culture = Document(
    """A unique performance of Gustav Holst’s masterwork The Planets – played on a magnificent pipe organ rather than by an orchestra and punctuated by poems inspired by children’s responses to the music – is to be staged in the suitably vast Salisbury Cathedral.
The idea of the community music project is to introduce more people, young and old, to the 140-year-old “Father” Willis organ, one of the treasures of the cathedral.
It is also intended to get the children who took part and the adults who will watch and listen thinking afresh about the themes Holst’s suite tackles – war, peace, joy and mysticism – which seem as relevant now as when he wrote the work a century ago.
John Challenger, the cathedral’s principal organist, said: “We have a fantastic pipe organ largely as it was when built. It’s a thrilling thing. I view it as my purpose in life to share it with as many people as possible.”
The Planets is written for a large orchestra. “Holst calls for huge instrumental forces and an unseen distant choir of sopranos and altos,” said Challenger. But he has transposed the suite for the organ, not copying the effect of the orchestral instruments but finding a new version of the suite."""
)

# https://www.theguardian.com/sport/blog/2023/feb/14/multi-million-dollar-wpl-auction-signals-huge-step-forward-for-womens-sport
news_sport = Document(
    """It was only a few days ago that members of the Australian women’s cricket team were contemplating how best to navigate the impending “distraction” of the inaugural Women’s Premier League auction, scheduled during the first week of the T20 World Cup. “It’s a little bit awkward,” captain Meg Lanning said in South Africa last week. “But it’s just trying to embrace that and understanding it’s actually a really exciting time and you actually don’t have a lot of control over most of it, so you’ve just got to wait and see.”
What a pleasant distraction it turned out to be. Lanning herself will be $192,000 richer for three weeks’ work with the Delhi Capitals. Her teammate, Ash Gardner, will earn three times that playing for the Gujarat Giants. The allrounder’s figure of $558,000 is more than Sam Kerr pockets in a season with Chelsea and more than the WNBA’s top earner, Jackie Young.
If that sounds like a watershed moment, it’s perhaps because it is. And it is not the only one this past week. The NRLW made its own wage-related headlines on Tuesday, to the effect that the next (agreed in principle) collective bargaining agreement will bring with it a $1.5m salary cap in 2027, at an average salary of $62,500. Women’s rugby, too, is making moves, with news on the weekend that Rugby Australia will begin contracting the Wallaroos."""
)

news = [news_economics, news_science, news_culture, news_sport]
answers = ["1980s", "Orion", "The Planets", "the first week of the T20 World Cup"]
```

```python
prompt_node.prompt(prompt_template="deepset/conditioned-question-generation", documents=news, answers=answers)
```

```
['What was the last time UK inflation rate hit the highest level?',
 'What constellation is easy to spot in the eastern sky by mid-evening in the southern hemisphere?',
 "How has John Challenger transposed Gustav Holst's suite for the pipe organ?",
 "What event was Meg Lanning referring to as a possible distraction for the Australian women's cricket team?"]
```