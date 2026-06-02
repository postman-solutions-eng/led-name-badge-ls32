# PredefinedIcons

A list of all methods in the `PredefinedIcons` service. Click on the method name to view detailed information about that method.

| Methods                                   | Description                                                                                                                                         |
| :---------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------- |
| [GetPredefinedIcons](#getpredefinedicons) | Returns a list of all available icon codes that can be used in display text. Icons are returned as simple string codes in the `:icon_name:` format. |

## GetPredefinedIcons

Returns a list of all available icon codes that can be used in display text. Icons are returned as simple string codes in the `:icon_name:` format.

- HTTP Method: `GET`
- Endpoint: `/predefined-icons`

**Parameters**

| Name | Type    | Required | Description                 |
| :--- | :------ | :------- | :-------------------------- |
| ctx  | Context | ✅       | Default go language context |

**Return Type**

`GetPredefinedIconsOkResponse`

**Example Usage Code Snippet**

```go
led-display-api predefined-icons get-predefined-icons
```
