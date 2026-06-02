# CreateDisplaySummaryOkResponse

**Properties**

| Name              | Type                             | Required | Description                                                                                                |
| :---------------- | :------------------------------- | :------- | :--------------------------------------------------------------------------------------------------------- |
| Status            | string                           | ❌       | Status of the operation                                                                                    |
| Source            | displaysummary.Source            | ❌       | Present when a Postman API key was supplied. Indicates the summary was built from the Postman API catalog. |
| Text              | string                           | ❌       | The text rendered on the LED badge                                                                         |
| SystemEnvironment | displaysummary.SystemEnvironment | ❌       | The API catalog system environment used to list services                                                   |
| ServiceCount      | int64                            | ❌       | Number of services returned for the summary                                                                |
| HasMore           | bool                             | ❌       | True when additional catalog services exist beyond the fetched page                                        |
| Services          | []displaysummary.Services        | ❌       | Catalog services included in the summary                                                                   |

# Source

Present when a Postman API key was supplied. Indicates the summary was built from the Postman API catalog.

**Properties**

| Name              | Type   | Required | Description           |
| :---------------- | :----- | :------- | :-------------------- |
| PostmanAPICatalog | string | ✅       | "postman-api-catalog" |

# SystemEnvironment

The API catalog system environment used to list services

**Properties**

| Name | Type   | Required | Description |
| :--- | :----- | :------- | :---------- |
| ID   | string | ❌       |             |
| Name | string | ❌       |             |

# Services

**Properties**

| Name   | Type                  | Required | Description |
| :----- | :-------------------- | :------- | :---------- |
| ID     | string                | ❌       |             |
| Name   | string                | ❌       |             |
| Status | displaysummary.Status | ❌       |             |

# Status

**Properties**

| Name     | Type   | Required | Description |
| :------- | :----- | :------- | :---------- |
| Healthy  | string | ✅       | "healthy"   |
| Warning  | string | ✅       | "warning"   |
| Critical | string | ✅       | "critical"  |
| Inactive | string | ✅       | "inactive"  |
