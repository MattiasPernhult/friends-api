# Friends API
REST API for fetching data about the tv show Friends

## Endpoints
### Overview
|Method       |Endpoint         |Description                                                                   |
|:------------|:----------------|:-----------------------------------------------------------------------------|
|GET          |/persons         |Returns persons from Friends, both regulars and recurring characters          |

### GET /friends

##### Query Parameters
|Parameter|Description|Required|Type|Default|Valid values|
|:--------|:----------|:-------|:---|:------|:-----------|
|limit|Limit the returned results|No|Number|25|Number greater than 0|
|orderBy|Sort the result on number of episodes|No|String|"-numberOfEpisodes"|"-numberOfEpisodes" and "numberOfEpisodes"|
