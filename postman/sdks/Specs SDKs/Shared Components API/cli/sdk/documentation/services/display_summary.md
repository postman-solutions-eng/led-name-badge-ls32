# DisplaySummary

A list of all methods in the `DisplaySummary` service. Click on the method name to view detailed information about that method.

| Methods                                       | Description                                                                                                                                                                                                                                                                                                                                                                                         |
| :-------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [CreateDisplaySummary](#createdisplaysummary) | Displays a summary message on the LED badge. Without an API key (`X-API-Key` header or `apiKey` in the body), shows the built-in demo summary text. When a key is provided, the endpoint fetches services from the Postman API Catalog and renders a compact health summary on the badge. Note: The API does not validate the `type` parameter and will return success for any type value provided. |

## CreateDisplaySummary

Displays a summary message on the LED badge. Without an API key (`X-API-Key` header or `apiKey` in the body), shows the built-in demo summary text. When a key is provided, the endpoint fetches services from the Postman API Catalog and renders a compact health summary on the badge. Note: The API does not validate the `type` parameter and will return success for any type value provided.

- HTTP Method: `POST`
- Endpoint: `/display-summary`

**Parameters**

| Name                        | Type                        | Required | Description                 |
| :-------------------------- | :-------------------------- | :------- | :-------------------------- |
| ctx                         | Context                     | ✅       | Default go language context |
| createDisplaySummaryRequest | CreateDisplaySummaryRequest | ✅       |                             |

**Return Type**

`CreateDisplaySummaryOkResponse`

**Example Usage Code Snippet**

```go
led-display-api display-summary create-display-summary --body '{"type":"type","customText":"customText","apiKey":"apiKey","systemEnvironmentId":"e8f94f60-f018-425a-afdd-dfbec894def8"}'
```
