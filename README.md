# Git Gram

A Telegram bot that delivers GitHub repository activity notifications directly to your Telegram chats—no YAML configuration required.

## Key Features

- **Zero Configuration**: No workflow files or manual webhook setup needed
- **Direct Notifications**: Receive GitHub events instantly in Telegram
- **GitHub App Integration**: Install once on your account or selected repositories
- **Simple Commands**: Manage notifications through intuitive slash commands
- **Flexible Control**: Mute/unmute notifications per chat as needed
- **Secure**: Webhook signature verification and secure credential handling

## Architecture Overview

```
┌─────────────┐
│ Telegram    │
│ User        │
└──────┬──────┘
       │
       │ /start, /status, etc.
       ▼
┌─────────────────┐
│ Telegram Bot    │
│ (Bot API)       │
└────────┬────────┘
         │
         │ commands
         ▼
┌─────────────────────────────────────┐
│ Backend API                         │
│ • Process commands                  │
│ • Store mappings (chat ↔ install)  │
│ • Receive GitHub webhooks           │
│ • Send Telegram messages            │
└──────┬──────────────────────┬───────┘
       │                      ▲
       │ provide link         │ webhook events
       ▼                      │
┌─────────────────┐    ┌──────┴──────┐
│ GitHub App      │    │  GitHub     │
│ Installation    │◄───┤  Webhooks   │
└─────────────────┘    └─────────────┘
```

**Flow:**
1. User chats with Telegram bot
2. Bot provides GitHub App installation link
3. User installs GitHub App on their account/repos
4. GitHub sends webhook events to backend
5. Backend maps installation to Telegram chat(s)
6. Backend sends formatted notifications to Telegram

## Supported Commands

| Command | Description |
|---------|-------------|
| `/start` | Begin onboarding and get the GitHub App installation link |
| `/help` | Display available commands and usage instructions |
| `/status` | Show current notification status and linked repositories |
| `/mute` | Temporarily disable notifications for this chat |
| `/unmute` | Re-enable notifications for this chat |
| `/unlink` | Disconnect GitHub installation from this chat |

## Supported GitHub Webhook Events

The following GitHub webhook events can be supported. Check the boxes for events currently implemented:

- [ ] `pull_request` - Pull request opened, closed, merged, etc.
- [ ] `pull_request_review` - Pull request reviews submitted
- [ ] `pull_request_review_comment` - Comments on PR reviews
- [ ] `issues` - Issues opened, closed, reopened, etc.
- [ ] `issue_comment` - Comments on issues and pull requests
- [ ] `push` - Commits pushed to repository
- [ ] `create` - Branch or tag created
- [ ] `delete` - Branch or tag deleted
- [ ] `release` - Release published, updated, deleted
- [ ] `fork` - Repository forked
- [ ] `star` - Repository starred
- [ ] `watch` - Repository watched
- [ ] `commit_comment` - Comment on a commit
- [ ] `deployment` - Deployment created
- [ ] `deployment_status` - Deployment status updated
- [ ] `status` - Commit status updated
- [ ] `check_run` - Check run created, updated, completed
- [ ] `check_suite` - Check suite completed
- [ ] `workflow_run` - GitHub Actions workflow run completed

> **Note:** Additional events may be supported. Update this list as implementation progresses.

## Setup and Configuration

### Prerequisites

- Go `<VERSION>` or higher (e.g., `1.21+`)
- A GitHub account to create a GitHub App
- A Telegram bot token from [@BotFather](https://t.me/botfather)
- A publicly accessible HTTPS URL for receiving webhooks
- Database (PostgreSQL, MySQL, SQLite, etc.) - `<SPECIFY_DATABASE>`

### Environment Variables

Create a `.env` file or set the following environment variables:

```bash
# Telegram Bot Configuration
TELEGRAM_BOT_TOKEN=<YOUR_TELEGRAM_BOT_TOKEN>

# GitHub App Configuration
GITHUB_APP_ID=<YOUR_GITHUB_APP_ID>
GITHUB_APP_PRIVATE_KEY=<YOUR_GITHUB_APP_PRIVATE_KEY_PATH_OR_CONTENT>
GITHUB_WEBHOOK_SECRET=<YOUR_WEBHOOK_SECRET>

# Backend Configuration
PUBLIC_BASE_URL=<YOUR_PUBLIC_HTTPS_URL>
PORT=<SERVER_PORT>

# Database Configuration
DATABASE_URL=<YOUR_DATABASE_CONNECTION_STRING>
# Example: postgresql://user:password@localhost:5432/gitgram
# Example: sqlite://./gitgram.db
```

## GitHub App Configuration

### Creating the GitHub App

1. Go to **Settings** > **Developer settings** > **GitHub Apps** > **New GitHub App**
2. Fill in the required fields:
   - **GitHub App name**: `<YOUR_APP_NAME>` (e.g., "Git Gram Bot")
   - **Homepage URL**: `<YOUR_PROJECT_URL>`
   - **Webhook URL**: `<YOUR_PUBLIC_BASE_URL>/webhooks/github`
   - **Webhook secret**: Generate a secure random string and save it as `GITHUB_WEBHOOK_SECRET`

### Required Permissions

Configure the following permissions (adjust based on events you want to support):

**Repository permissions:**
- **Contents**: Read-only (for push events, branch/tag creation)
- **Issues**: Read-only (for issue events)
- **Metadata**: Read-only (required)
- **Pull requests**: Read-only (for PR events)
- **Commit statuses**: Read-only (for status events)
- **Checks**: Read-only (for check run events)

**Organization permissions:**
- **Members**: Read-only (optional, for organization events)

### Subscribe to Events

Select the webhook events you want to receive (match with supported events above):

- `[ ]` Pull requests
- `[ ]` Issues
- `[ ]` Issue comments
- `[ ]` Pushes
- `[ ]` Branch or tag creation
- `[ ]` Branch or tag deletion
- `[ ]` Releases
- `[ ]` Stars
- `[ ]` <OTHER_EVENTS>

After creating the app:
1. Generate a private key and download it
2. Note your App ID
3. Install the app on your account or organization

## Telegram Bot Setup

### Creating the Bot

1. Open Telegram and search for [@BotFather](https://t.me/botfather)
2. Send `/newbot` and follow the prompts:
   - Choose a display name (e.g., "Git Gram")
   - Choose a username (e.g., "gitgram_bot")
3. BotFather will provide your bot token - save it as `TELEGRAM_BOT_TOKEN`
4. (Optional) Set bot commands with `/setcommands`:

```
start - Begin setup and get GitHub App link
help - Show available commands
status - View notification status
mute - Disable notifications
unmute - Enable notifications
unlink - Disconnect GitHub integration
```

### Bot Link

Your bot will be accessible at:
```
https://t.me/<YOUR_BOT_USERNAME>
```

## Running Locally

### 1. Clone the Repository

```bash
git clone <REPOSITORY_URL>
cd gitgram
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Database

```bash
# Run migrations (example)
<DATABASE_MIGRATION_COMMAND>
```

### 4. Run the Backend

```bash
# Option 1: Using go run
go run cmd/server/main.go

# Option 2: Build and run
go build -o gitgram cmd/server/main.go
./gitgram
```

### 5. Expose Local Server (for Development)

If testing locally, use a tool like [ngrok](https://ngrok.com/) to expose your local server:

```bash
ngrok http <YOUR_LOCAL_PORT>
# Update PUBLIC_BASE_URL with the ngrok HTTPS URL
```

## Deployment

### Building the Container

```bash
# Build Docker image
docker build -t gitgram:latest .

# Run container
docker run -d \
  --name gitgram \
  -p <PORT>:<PORT> \
  --env-file .env \
  gitgram:latest
```

### Container Registry

```bash
# Tag and push to registry
docker tag gitgram:latest <REGISTRY_URL>/gitgram:latest
docker push <REGISTRY_URL>/gitgram:latest
```

### Deployment Notes

- Ensure your webhook endpoint `<PUBLIC_BASE_URL>/webhooks/github` is publicly accessible via HTTPS
- Use a reverse proxy (nginx, Caddy) or cloud load balancer for SSL termination
- Set up health check endpoints for monitoring
- Configure auto-restart policies for production
- Consider horizontal scaling for high-traffic scenarios

**Deployment Options:**
- `<CLOUD_PROVIDER>` (e.g., AWS, GCP, Azure, DigitalOcean)
- `<CONTAINER_ORCHESTRATION>` (e.g., Kubernetes, Docker Swarm)
- `<PAAS>` (e.g., Heroku, Railway, Render)

## Security Notes

### Webhook Verification

- **Always verify GitHub webhook signatures** using `GITHUB_WEBHOOK_SECRET`
- Reject requests with invalid or missing signatures
- Implement signature verification before parsing payload

### Credential Management

- **Never log or expose sensitive values**: bot tokens, private keys, webhook secrets
- Use environment variables or secure secret management systems
- Rotate credentials periodically
- Use `.gitignore` to prevent committing `.env` files

### Rate Limiting

- Implement rate limiting on the webhook endpoint to prevent abuse
- Consider GitHub's webhook delivery retry behavior
- Add request throttling for Telegram API calls to avoid hitting rate limits

### Additional Recommendations

- Use HTTPS only for all communications
- Implement input validation and sanitization
- Keep dependencies up to date
- Monitor for suspicious webhook patterns
- Set up logging and alerting for security events

## Roadmap

- [ ] Support for GitHub Discussions events
- [ ] Custom notification templates per event type
- [ ] Multi-language support for bot messages
- [ ] Group chat support with permission management
- [ ] Notification filtering by repository, branch, or user
- [ ] Integration with GitHub Actions workflow status
- [ ] Dashboard for managing multiple installations
- [ ] Notification scheduling (quiet hours)

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

`<SPECIFY_LICENSE>` (e.g., MIT, Apache 2.0, GPL-3.0)

---

**Questions or Issues?** Open an issue on GitHub or reach out via `<CONTACT_METHOD>`.
