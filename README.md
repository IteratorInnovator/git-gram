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

- [ ] `branch_protection_configuration` - Branch protection configuration activity
  - [ ] `disabled`
- [ ] `branch_protection_rule` - Branch protection rule activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
- [ ] `check_run` - Check run activity
  - [ ] `completed`
  - [ ] `created`
  - [ ] `requested_action`
  - [ ] `rerequested`
- [ ] `check_suite` - Check suite activity
  - [ ] `completed`
  - [ ] `requested`
  - [ ] `rerequested`
- [ ] `code_scanning_alert` - Code scanning alert activity
  - [ ] `appeared_in_branch`
  - [ ] `closed_by_user`
  - [ ] `created`
  - [ ] `fixed`
  - [ ] `reopened`
  - [ ] `reopened_by_user`
- [ ] `commit_comment` - Commit comment activity
  - [ ] `created`
- [x] `create` - Branch or tag creation
- [ ] `custom_property` - Custom property activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `updated`
- [ ] `custom_property_values` - Custom property values activity
  - [ ] `updated`
- [ ] `delete` - Branch or tag deletion
- [ ] `dependabot_alert` - Dependabot alert activity
  - [ ] `auto_dismissed`
  - [ ] `auto_reopened`
  - [ ] `created`
  - [ ] `dismissed`
  - [ ] `fixed`
  - [ ] `reintroduced`
  - [ ] `reopened`
- [ ] `deploy_key` - Deploy key activity
  - [ ] `created`
  - [ ] `deleted`
- [ ] `deployment` - Deployment activity
  - [ ] `created`
- [ ] `deployment_protection_rule` - Deployment protection rule activity
  - [ ] `requested`
- [ ] `deployment_review` - Deployment review activity
  - [ ] `approved`
  - [ ] `rejected`
  - [ ] `requested`
- [ ] `deployment_status` - Deployment status activity
  - [ ] `created`
- [ ] `discussion` - Discussion activity
  - [ ] `answered`
  - [ ] `category_changed`
  - [ ] `closed`
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `labeled`
  - [ ] `locked`
  - [ ] `pinned`
  - [ ] `reopened`
  - [ ] `transferred`
  - [ ] `unanswered`
  - [ ] `unlabeled`
  - [ ] `unlocked`
  - [ ] `unpinned`
- [ ] `discussion_comment` - Discussion comment activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
- [ ] `fork` - Repository forked
- [ ] `github_app_authorization` - GitHub App authorization activity
  - [ ] `revoked`
- [ ] `gollum` - Wiki page activity
- [ ] `installation` - GitHub App installation activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `new_permissions_accepted`
  - [ ] `suspend`
  - [ ] `unsuspend`
- [ ] `installation_repositories` - GitHub App installation repositories activity
  - [ ] `added`
  - [ ] `removed`
- [ ] `installation_target` - GitHub App installation target activity
  - [ ] `renamed`
- [ ] `issue_comment` - Issue or PR comment activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
- [ ] `issue_dependencies` - Issue dependencies activity
  - [ ] `blocked_by_added`
  - [ ] `blocked_by_removed`
  - [ ] `blocking_added`
  - [ ] `blocking_removed`
- [ ] `issues` - Issue activity
  - [ ] `assigned`
  - [ ] `closed`
  - [ ] `deleted`
  - [ ] `demilestoned`
  - [ ] `edited`
  - [ ] `labeled`
  - [ ] `locked`
  - [ ] `milestoned`
  - [ ] `opened`
  - [ ] `pinned`
  - [ ] `reopened`
  - [ ] `transferred`
  - [ ] `unassigned`
  - [ ] `unlabeled`
  - [ ] `unlocked`
  - [ ] `unpinned`
- [ ] `label` - Label activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
- [ ] `marketplace_purchase` - GitHub Marketplace purchase activity
  - [ ] `cancelled`
  - [ ] `changed`
  - [ ] `pending_change`
  - [ ] `pending_change_cancelled`
  - [ ] `purchased`
- [ ] `member` - Repository collaborator activity
  - [ ] `added`
  - [ ] `edited`
  - [ ] `removed`
- [ ] `membership` - Team membership activity
  - [ ] `added`
  - [ ] `removed`
- [ ] `merge_group` - Merge group activity
  - [ ] `checks_requested`
  - [ ] `destroyed`
- [ ] `meta` - Webhook activity
  - [ ] `deleted`
- [ ] `milestone` - Milestone activity
  - [ ] `closed`
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `opened`
- [ ] `org_block` - Organization blocking activity
  - [ ] `blocked`
  - [ ] `unblocked`
- [ ] `organization` - Organization activity
  - [ ] `deleted`
  - [ ] `member_added`
  - [ ] `member_invited`
  - [ ] `member_removed`
  - [ ] `renamed`
- [ ] `package` - GitHub Packages activity
  - [ ] `published`
  - [ ] `updated`
- [ ] `page_build` - GitHub Pages build activity
- [ ] `personal_access_token_request` - Personal access token request activity
  - [ ] `approved`
  - [ ] `cancelled`
  - [ ] `created`
  - [ ] `denied`
- [ ] `ping` - Webhook ping
- [ ] `project` - Project (classic) activity
  - [ ] `closed`
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `reopened`
- [ ] `project_card` - Project (classic) card activity
  - [ ] `converted`
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `moved`
- [ ] `project_column` - Project (classic) column activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `moved`
- [ ] `projects_v2` - Projects activity
  - [ ] `closed`
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `reopened`
- [ ] `projects_v2_item` - Projects item activity
  - [ ] `archived`
  - [ ] `converted`
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `reordered`
  - [ ] `restored`
- [ ] `projects_v2_status_update` - Projects status update activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
- [ ] `public` - Repository visibility changed to public
- [ ] `pull_request` - Pull request activity
  - [ ] `assigned`
  - [ ] `auto_merge_disabled`
  - [ ] `auto_merge_enabled`
  - [ ] `closed`
  - [ ] `converted_to_draft`
  - [ ] `demilestoned`
  - [ ] `dequeued`
  - [ ] `edited`
  - [ ] `enqueued`
  - [ ] `labeled`
  - [ ] `locked`
  - [ ] `milestoned`
  - [ ] `opened`
  - [ ] `ready_for_review`
  - [ ] `reopened`
  - [ ] `review_request_removed`
  - [ ] `review_requested`
  - [ ] `synchronize`
  - [ ] `unassigned`
  - [ ] `unlabeled`
  - [ ] `unlocked`
- [ ] `pull_request_review` - Pull request review activity
  - [ ] `dismissed`
  - [ ] `edited`
  - [ ] `submitted`
- [ ] `pull_request_review_comment` - Pull request review comment activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
- [ ] `pull_request_review_thread` - Pull request review thread activity
  - [ ] `resolved`
  - [ ] `unresolved`
- [x] `push` - Git push to repository
- [ ] `registry_package` - Registry package activity
  - [ ] `published`
  - [ ] `updated`
- [ ] `release` - Release activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `prereleased`
  - [ ] `published`
  - [ ] `released`
  - [ ] `unpublished`
- [x] `repository` - Repository activity
  - [x] `archived`
  - [x] `created`
  - [x] `deleted`
  - [ ] `edited`
  - [x] `privatized`
  - [x] `publicized`
  - [x] `renamed`
  - [ ] `transferred`
  - [x] `unarchived`
- [ ] `repository_advisory` - Repository security advisory activity
  - [ ] `published`
  - [ ] `reported`
- [ ] `repository_dispatch` - Repository dispatch event
- [ ] `repository_import` - Repository import activity
- [ ] `repository_ruleset` - Repository ruleset activity
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
- [ ] `repository_vulnerability_alert` - Repository vulnerability alert activity
  - [ ] `create`
  - [ ] `dismiss`
  - [ ] `resolve`
- [ ] `secret_scanning_alert` - Secret scanning alert activity
  - [ ] `created`
  - [ ] `reopened`
  - [ ] `resolved`
  - [ ] `revoked`
  - [ ] `validated`
- [ ] `secret_scanning_alert_location` - Secret scanning alert location activity
  - [ ] `created`
- [ ] `secret_scanning_scan` - Secret scanning scan activity
  - [ ] `completed`
- [ ] `security_advisory` - Security advisory activity
  - [ ] `performed`
  - [ ] `published`
  - [ ] `updated`
  - [ ] `withdrawn`
- [ ] `security_and_analysis` - Security and analysis activity
- [ ] `sponsorship` - Sponsorship activity
  - [ ] `cancelled`
  - [ ] `created`
  - [ ] `edited`
  - [ ] `pending_cancellation`
  - [ ] `pending_tier_change`
  - [ ] `tier_changed`
- [ ] `star` - Star activity
  - [ ] `created`
  - [ ] `deleted`
- [ ] `status` - Commit status updated
- [ ] `sub_issues` - Sub-issues activity
  - [ ] `parent_issue_added`
  - [ ] `parent_issue_removed`
  - [ ] `sub_issue_added`
  - [ ] `sub_issue_removed`
- [ ] `team` - Team activity
  - [ ] `added_to_repository`
  - [ ] `created`
  - [ ] `deleted`
  - [ ] `edited`
  - [ ] `removed_from_repository`
- [ ] `team_add` - Team added to repository
- [ ] `watch` - Watch activity
  - [ ] `started`
- [ ] `workflow_dispatch` - Workflow dispatch event
- [ ] `workflow_job` - Workflow job activity
  - [ ] `completed`
  - [ ] `in_progress`
  - [ ] `queued`
  - [ ] `waiting`
- [ ] `workflow_run` - Workflow run activity
  - [ ] `completed`
  - [ ] `in_progress`
  - [ ] `requested`

> **Note:** Additional events may be supported. This list will be updated as implementation progresses.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

**Questions or Issues?** Open an issue on GitHub.
