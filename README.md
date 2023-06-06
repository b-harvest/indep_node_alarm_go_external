# INDEP_NODE_ALARM_GO_EXTERNAL

INDEP_NODE_ALARM_GO_EXTERNAL is a Go application that continuously checks the TCP dial status of a list of servers. If a connection cannot be established to a server, INDEP_NODE_ALARM_GO_EXTERNAL sends alerts via Telegram and PagerDuty.

## Installation

First, clone the repository:

```bash
git clone https://github.com/b-harvest/indep_node_alarm_go_external.git
cd indep_node_alarm_go_external
```

Then, build and install the application with:

This will build the application and move the resulting executable to your GOPATH/bin directory.

```
make install
```

## Configuration

INDEP_NODE_ALARM_GO_EXTERNAL uses a TOML file for configuration. Here's an example config.toml file:

```
BotToken = "YOUR_BOT_TOKEN"
ChatID = YOUR_CHAT_ID
PagerDutyToken = "YOUR_PAGERDUTY_API_TOKEN"
PagerDutyUserID = "YOUR_PAGERDUTY_USER_ID"

[Servers]
"Server1-name" = "server1:port1"
"Server2-name" = "server2:port2"
```
Replace "server1:port1", "server2:port2", "YOUR_BOT_TOKEN", "YOUR_PAGERDUTY_API_TOKEN", "YOUR_PAGERDUTY_USER_ID", and YOUR_CHAT_ID with your actual server addresses, API tokens, user ID, and chat ID.

## Usage

You can run INDEP_NODE_ALARM_GO_EXTERNAL with:

```
indep_ext --config path/to/config.toml
```

Replace path/to/config.toml with the path to your config.toml file.

