# Friends API
REST API for fetching data about the tv show Friends

## Endpoints
### Overview
|Method       |Endpoint         |Description                                                                   |
|:------------|:----------------|:-----------------------------------------------------------------------------|
|GET          |/persons         |Returns persons from Friends, both regulars and recurring characters          |

### GET /friends

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

Hello
