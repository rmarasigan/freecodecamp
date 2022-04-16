# slack-age-bot

Prerequisite:
- You should have a Slack account
- You should have a Workspace on Slack
- Go to api.slack.com/apps

#### Create New App From Scratch
![Slack API](/assets/img/slack-bot-create-new-app.png "create an app")

#### Name app & choose workspace
![Slack API Name app & choose workspace](/assets/img/slack-bot-app-workspace.png "name app & choose workspace")

#### Socket Mode > Enable Socket Mode
![Slack API Socket Mode](/assets/img/slack-bot-enable-socket.png "enable socket mode")

##### socket-token
![Slack API Generated Socket Token](/assets/img/slack-bot-generated-socket-token.png "generated socket token")

#### Event Subscriptions > Enable Events > Subscribe to bot events
![Slack API Subscribe to bot events](/assets/img/slack-bot-subscribe-bot-events.png "subscribe to bot events")

#### OAuth & Permissions > Scopes > Add an OAuth Scope
![Slack API OAuth Scope](/assets/img/slack-bot-oauth-scope.png "oauth scope")

#### OAuth & Permissions Install to Workspace
![Slack API OAuth Install to Workspace](/assets/img/slack-bot-install-worksapace.png "install to workspace")

```bash
dev@dev:~/go/src/github.com/development/slack-bot-age$ go build
dev@dev:~/go/src/github.com/development/slack-bot-age$ go run main.go
```
![Slack Age Bot](/assets/img/slack-age-bot.png "age-bot")