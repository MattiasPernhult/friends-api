# Friends API
REST API for fetching data about the tv show Friends

## Endpoints
### Overview
|Method       |Endpoint         |Description                                                                   |
|:------------|:----------------|:-----------------------------------------------------------------------------|
|GET|[/persons](#get-persons)|Returns persons from Friends, both regulars and recurring characters|
|GET|[/episodes](#get-episodes)|Return episodes from Friends|

### GET persons

##### JSON Schema

|Name|Type|Shown|
|:---|:---|:----|
|**firstAppearance**|String|Always|
|**lastAppearance**|String|Always|
|**numberOfEpisodes**|Number|Always|
|**name**|String|Always|
|**potrayedBy**|[String]|With parameter `inlcude`|
|**nicknames**|[String]|With parameter `inlcude`|
|**gender**|String|With parameter `include`|
|**dateOfBirth**|String ISO Date|With parameter `include`|
|**occupations**|[String]|With parameter `inlcude`|
|**relatives**|[name: String, relationship: String]|With parameter `inlcude`|

##### Query Parameters

|Parameter|Default|Valid values|
|:--------|:------|:-----------|
|limit|25|Number greater than 0|
|orderBy|`-numberOfEpisodes`|`-numberOfEpisodes` `numberOfEpisodes`|
|include|*See JSON Schema*|*See JSON Schema*|

Parameter `orderBy` will default to **-numberOfEpisodes**, the dash means it will sort by descending and without the dash it will sort by ascending.

Parameter `include` will include more information in the response. To include several use comma between each include:

``` bash
/persons?include=nicknames,occupations
```

##### Example Response

When doing a GET Request to the endpoint and adding query parameter limit to 2, it will return the following response.


```JSON
{
  "people": [
    {
      "firstAppearance": "The Pilot",
      "lastAppearance": "The Last One, Part 2",
      "numberOfEpisodes": 236,
      "name": "Monica E. Geller-Bing"
    },
    {
      "firstAppearance": "The Pilot",
      "lastAppearance": "The Last One, Part 2",
      "numberOfEpisodes": 236,
      "name": "Rachel Green"
    }
  ]
}
```

### GET episodes

##### JSON Schema

|Name|Type|Shown|
|:---|:---|:----|
|**Season**|String|Always|
|**Episode**|String|Always|
|**SeasonNumber**|Number|Always|
|**EpisodeNumber**|Number|Always|
|**Title**|String|Always|
|**Airdate**|String ISO Date|Always|
|**conversations**|[said: String, person: String]|Always|
|**plot**|String|Always|

##### Query Parameters

|Parameter|Default|Valid values|
|:--------|:------|:-----------|
|limit|25|Number greater than 0|
|orderBy|`seasonNumber`|`seasonNumber` `-seasonNumber` `-episodeNumber` `-episodeNumber`|

Parameter `orderBy` will default to **seasonNumber**.

##### Example Response

``` bash
/episodes?limit=1
```


```JSON
{
  "episodes": [
    {
      "season": "season1",
      "episode": "episode01",
      "seasonNumber": 1,
      "episodeNumber": 1,
      "title": "Pilot",
      "airdate": "1992-09-22T22:00:00",
      "conversations": [
        {
          "said": "There's nothing to tell. He's just some guy I work with.",
          "person": "Monica Geller"
        },
        {
          "said": "C'mon, you're going out with the guy. There's gotta be something wrong with him.",
          "person": "Joey Tribianni"
        },
        {
          "said": "So does he have a hump? A hump and a hairpiece?",
          "person": "Chandler Bing"
        },
        {
          "said": "Wait, does he eat chalk? [The others stare, bemused] It's just that I don't want her to go through what I went through with Carl - ohh!",
          "person": "Phoebe Buffay"
        },
        {
          "said": "Okay, everybody relax. This is not a date. It's just two people going out to dinner, and not having sex.",
          "person": "Monica Geller"
        },
        {
          "said": "Sounds like a date to me.",
          "person": "Chandler Bing"
        },
        {
          "said": "You know what the scariest part is? What if there's only one woman for everybody, you know? What if you get one woman, and that's it? Unfortunately, in my case, there was only one woman for her.",
          "person": "Ross Geller"
        },
        {
          "said": "What are you talking about, one woman? That's like saying there's only one flavor of ice cream for you. Let me tell you something, Ross, there's lots of flavors out there. There's Rocky Road and Cookie Dough and Bing Cherry Vanilla. You could get 'em with jimmies or nuts or whip cream. This is the best thing that ever happened to you. You got married, you were like what, eight? Welcome back to the world! Grab a spoon!",
          "person": "Joey Tribianni"
        },
        {
          "said": "I honestly don't know if I'm hungry or horny.",
          "person": "Ross Geller"
        },
        {
          "said": "Then stay the hell out of my freezer!",
          "person": "Chandler Bing"
        }
      ],
      "plot": "Rachel leaves Barry at the alter and moves in with Monica. Monica goes on a date with Paul the wine guy, who turns out to be less than sincere. Ross is depressed about his failed marriage. Joey compares women to ice cream. Everyone watches Spanish soaps. Ross reveals his high school crush on Rachel. "
    }
  ]
}
```
