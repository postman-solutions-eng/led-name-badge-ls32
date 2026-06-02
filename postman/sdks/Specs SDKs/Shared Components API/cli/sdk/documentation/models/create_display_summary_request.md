# CreateDisplaySummaryRequest

**Properties**

| Name                | Type   | Required | Description                                                                                                                                                                                       |
| :------------------ | :----- | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Type                | string | ❌       | The type of summary message to display (e.g., welcome, status, alert, info)                                                                                                                       |
| CustomText          | string | ❌       | Optional custom text to append to the summary                                                                                                                                                     |
| APIKey              | string | ❌       | Optional alternative to the `X-API-Key` header for supplying a Postman API key (`PMAK-...`) used to fetch the API catalog                                                                         |
| SystemEnvironmentID | string | ❌       | Optional API catalog system environment ID. When omitted, the first production environment is used, otherwise the first environment with associations, otherwise the first available environment. |
