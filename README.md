# Friends API
REST API for fetching data about the tv show Friends

## Endpoints
### Overview
|Method       |Endpoint         |Description                                                                   |
|:------------|:----------------|:-----------------------------------------------------------------------------|
|GET          |/persons         |Returns persons from Friends, both regulars and recurring characters          |

### GET /friends

##### Query Parameters
|Parameter|Required|Type|Default|Valid values|
|:--------|:-------|:---|:------|:-----------|
|limit|No|Number|25|Number greater than 0|
|orderBy|No|String|`-numberOfEpisodes`|`-numberOfEpisodes` `numberOfEpisodes`|
