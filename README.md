# Git Gram

A Telegram bot that delivers GitHub repository activity notifications directly to your Telegram chats—no YAML configuration required.

[Bot Link](https://t.me/gitgram_67bot)

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
│ Telegram    │◄─────────────────────┐
│ User        │                      │
└──────┬──────┘                      │
       │                             │
       │ /start, /status, etc.       │ send notifications
       ▼                             │
┌─────────────────┐                  │
│ Telegram Bot    │──────────────────┘
│ (Bot API)       │◄────────┐
└────────┬────────┘         |         
         │                  |  format & call Bot API  
         │ commands         |         
         ▼                  |         
┌─────────────────────────────────────────┐
│ Backend API                             │
│ • Process commands                      │
│ • Store mappings (chat ↔ install)       │
│ • Receive GitHub webhooks               │
│ • Send Telegram messages                │
└──────┬──────────────────────┬───────────┘
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
3. User installs Git Gram App on their GitHub account
4. GitHub sends webhook events to backend
5. Backend maps installation to Telegram chat
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

> **Note:** Additional events may be supported. This list will be updated as implementation progresses.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

**Questions or Issues?** Open an issue on GitHub.
