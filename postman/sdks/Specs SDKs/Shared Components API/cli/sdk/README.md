# LedDisplayAPI Go SDK 1.0.0

Welcome to the LedDisplayAPI SDK documentation. This guide will help you get started with integrating and using the LedDisplayAPI SDK in your project.

## Versions

- SDK version: `1.0.0`

## About the API

# API for controlling LED name badge displays

This API allows you to update the text and visual content displayed on LED badges, including support for alphanumeric text and predefined icon codes.
<img src="https://content.pstmn.io/05b1ef2c-9fd3-4f0e-8b2f-76998fb3f6e5/aW1hZ2UucG5n" width="240">

## Supported Icons

The following icon codes can be embedded in text using the `:icon_name:` syntax:

- `:ball:` - Filled circle
- `:happy:` - Simple smiley face
- `:happy2:` - Larger smiley (2 columns wide)
- `:heart:` - Outline heart
- `:HEART:` - Filled heart
- `:heart2:` - Larger outline heart (2 columns wide)
- `:HEART2:` - Larger filled heart (2 columns wide)
- `:fablab:` - FabLab logo
- `:bicycle:` - Bicycle icon (3 columns wide)
- `:bicycle_r:` - Bicycle facing right (3 columns wide)
- `:owncloud:` - OwnCloud logo (3 columns wide)
- `:octocat:` - GitHub Octocat
- `:smile:` - Smile emoji
- `:star:` - Star icon
- `:sun:` - Sun icon

## Character Restrictions

**Supported characters:**

- Letters: A-Z, a-z
- Numbers: 0-9
- Special characters: `^ !"$%&/()=?` °}\]\[{@ \~ |<>,;.:-_#'+_\`
- German umlauts: äöüÄÖÜß
- French/European accents: àäòöùüèéêëôöûîïÿçÀÅÄÉÈÊËÖÔÜÛÙŸŠ

**NOT supported:**

- Unicode emoji (e.g., 🌍) - will return a 400 error

**Escape sequences:**

- `::` - Escapes to a single literal colon

**Custom images:**

- Use `:path/to/image.png:` syntax to reference custom images
- Images must be exactly 11 pixels high

## Additional Capabilities (not exposed in API yet)

- **Brightness control:** 25%, 50%, 75%, 100%
- **Animation modes:** 0-8 (various scroll and display effects)
- **Speed control:** 1-8 (animation speed)
- **Blink effect:** Enable/disable blinking
- **Ant effect:** Marching border animation

## How the spec was created

This spec was created by having analyzed the backend Python code with Postman Agent mode with the following prompt:

> The spec and the Postman collection both seem very incomplete, not containing all possible icons and other character restrictions. Can you scan the backend API code to find out what additional endpoints and capabilities regarding icons and character restrictions there are and update the spec accordingly?

## Table of Contents

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
- [Authentication](#authentication)
  - [API Key Authentication](#api-key-authentication)
- [Setting a Custom Timeout](#setting-a-custom-timeout)
- [Sample Usage](#sample-usage)
- [Services](#services)
  - [Response Wrappers](#response-wrappers)
- [Models](#models)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### API Key Authentication

The led-display-api API uses API keys as a form of authentication. An API key is a unique identifier used to authenticate a user, developer, or a program that is calling the API.

#### Setting the API key

When you initialize the SDK, you can set the API key as follows:

```go
export LED_DISPLAY_API_API_KEY="YOUR-API-KEY"
led-display-api <command>
```

If you need to set or update the API key after initializing the SDK, you can use:

```go
led-display-api config set api_key "YOUR-API-KEY"
```

## Setting a Custom Timeout

You can set a custom timeout for the SDK's HTTP requests as follows:

```go

```

# Sample Usage

Below is a comprehensive example demonstrating how to authenticate and call a simple endpoint:

```go
led-display-api predefined-icons get-predefined-icons

```

## Services

The SDK provides various services to interact with the API.

<details>
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                          |
| :------------------------------------------------------------ |
| [DisplayText](documentation/services/display_text.md)         |
| [PredefinedIcons](documentation/services/predefined_icons.md) |
| [DisplaySummary](documentation/services/display_summary.md)   |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `LedDisplayAPIResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                            | Description                                 |
| :------- | :------------------------------ | :------------------------------------------ |
| Data     | `T`                             | The body of the API response                |
| Metadata | `LedDisplayAPIResponseMetadata` | Status code and headers returned by the API |

#### `LedDisplayAPIError[T]`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                         | Description                                                       |
| :------- | :--------------------------- | :---------------------------------------------------------------- |
| Err      | `error`                      | The error that occurred                                           |
| Data     | `*T`                         | The deserialized error response data (nil if unmarshaling failed) |
| Body     | `[]byte`                     | The raw body of the API response                                  |
| Metadata | `LedDisplayAPIErrorMetadata` | Status code and headers returned by the API                       |

#### `LedDisplayAPIResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details>
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                         | Description |
| :------------------------------------------------------------------------------------------- | :---------- |
| [CreateDisplayTextOkResponse](documentation/models/create_display_text_ok_response.md)       |             |
| [CreateDisplayTextRequest](documentation/models/create_display_text_request.md)              |             |
| [GetPredefinedIconsOkResponse](documentation/models/get_predefined_icons_ok_response.md)     |             |
| [CreateDisplaySummaryOkResponse](documentation/models/create_display_summary_ok_response.md) |             |
| [CreateDisplaySummaryRequest](documentation/models/create_display_summary_request.md)        |             |

</details>
