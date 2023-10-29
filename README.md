<h1 align="center">Websentry</h1>

<p align="center">
  <img src="./WebSentry.png" alt="Websentry Logo" width="200">
</p>

<p align="center">
  <strong>Check the health of your websites and receive instant status reports with Websentry!</strong>
</p>

[![Website Health Check](https://github.com/theritikchoure/websentry/actions/workflows/website_health.yml/badge.svg)](https://github.com/theritikchoure/websentry/actions/workflows/website_health.yml)

---

## Table of Contents

- [About Websentry](#about-websentry)
- [Key Features](#key-features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Inputs](#inputs)
- [Example Workflow](#example-workflow)
- [Contributing](#contributing)
- [License](#license)

---

## About Websentry

Websentry is a powerful GitHub Action that simplifies website health monitoring. With Websentry, you can effortlessly check the availability of your websites and receive real-time reports on their status. Whether you're a website owner, developer, or operations team, Websentry provides a reliable solution for ensuring your online presence remains uninterrupted.

<!-- ![Websentry Demo](https://yourwebsite.com/demo.gif) -->

---

## Key Features

- **Website Health Monitoring**: Easily check the health of your websites.
- **Configurable Retries**: Define the number of retries and time intervals for health checks.
- **Timeout Control**: Set the timeout for HTTP requests.
- **GitHub Action**: Seamlessly integrate Websentry into your GitHub Actions workflows.
- **User-Friendly**: Websentry is designed to be user-friendly and easily customizable.

---

## Getting Started

### Prerequisites

- A GitHub repository where you want to monitor website health.
- Basic knowledge of GitHub Actions.

### Installation

To start using Websentry, add it as a step in your GitHub Actions workflow. For example, in your `.github/workflows/health_check.yml` file:

```yaml
name: Website Health Check

on:
  schedule:
    - cron: '0 * * * *'  # Run every hour

jobs:
  health-check:
    runs-on: ubuntu-latest
    steps:
    - name: Check Website Health
      uses: theritikchoure/websentry@v1
      with:
        website-url: 'https://yourwebsite.com'
        max-retries: 3
        retry-interval: 5
        request-timeout: 10
```

## Usage
Websentry is designed to be simple and straightforward to use. Once integrated into your workflow, it will continuously monitor the health of your specified website and provide instant status reports.

## Configuration

You can configure Websentry by specifying the following input parameters in your GitHub Actions workflow:

- **website-url**: The URL of the website you want to monitor.
- **max-retries**: The maximum number of retries for checking the website's health.
- **retry-interval**: The time interval between retry attempts.
- **request-timeout**: The timeout for HTTP requests.

## Inputs
| Parameter      | Description                                      | Required | Default   |
| -------------- | ----------------------------------------------- | -------- | --------- |
| website-url    | The URL of the website to check (e.g., example.com) | true     | -         |
| max-retries    | Maximum number of retries (default: 3)          | false    | '3'       |
| retry-interval | Time to wait between retries in seconds (default: 5) | false    | '5'       |
| request-timeout | Request timeout in seconds (default: 10)      | false    | '10'      |


## Example Workflow
Here's an example GitHub Actions workflow that uses Websentry to check the health of a website:

```yaml
name: Website Health Check

on:
  schedule:
    - cron: '0 * * * *'  # Run every hour

jobs:
  health-check:
    runs-on: ubuntu-latest
    steps:
    - name: Check Website Health
      uses: theritikchoure/websentry@v1
      with:
        website-url: 'https://yourwebsite.com'
        max-retries: 3
        retry-interval: 5
        request-timeout: 10
```

## Contributing
We welcome contributions from the community! If you'd like to contribute to Websentry, please follow our Contribution Guidelines.

## License
Distributed under the MIT License. See LICENSE for more information.