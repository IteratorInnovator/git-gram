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

The following GitHub webhook events and their action types can be supported. Check the boxes for events currently implemented:

- [ ] `branch_protection_configuration` - Branch protection configuration activity
  - [ ] **Action:** `disabled`
- [ ] `branch_protection_rule` - Branch protection rule activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
- [ ] `check_run` - Check run activity
  - [ ] **Action:** `completed`
  - [ ] **Action:** `created`
  - [ ] **Action:** `requested_action`
  - [ ] **Action:** `rerequested`
- [ ] `check_suite` - Check suite activity
  - [ ] **Action:** `completed`
  - [ ] **Action:** `requested`
  - [ ] **Action:** `rerequested`
- [ ] `code_scanning_alert` - Code scanning alert activity
  - [ ] **Action:** `appeared_in_branch`
  - [ ] **Action:** `closed_by_user`
  - [ ] **Action:** `created`
  - [ ] **Action:** `fixed`
  - [ ] **Action:** `reopened`
  - [ ] **Action:** `reopened_by_user`
- [ ] `commit_comment` - Commit comment activity
  - [ ] **Action:** `created`
- [x] `create` - Branch or tag creation
- [ ] `custom_property` - Custom property activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `updated`
- [ ] `custom_property_values` - Custom property values activity
  - [ ] **Action:** `updated`
- [x] `delete` - Branch or tag deletion
- [ ] `dependabot_alert` - Dependabot alert activity
  - [ ] **Action:** `auto_dismissed`
  - [ ] **Action:** `auto_reopened`
  - [ ] **Action:** `created`
  - [ ] **Action:** `dismissed`
  - [ ] **Action:** `fixed`
  - [ ] **Action:** `reintroduced`
  - [ ] **Action:** `reopened`
- [ ] `deploy_key` - Deploy key activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
- [ ] `deployment` - Deployment activity
  - [ ] **Action:** `created`
- [ ] `deployment_protection_rule` - Deployment protection rule activity
  - [ ] **Action:** `requested`
- [ ] `deployment_review` - Deployment review activity
  - [ ] **Action:** `approved`
  - [ ] **Action:** `rejected`
  - [ ] **Action:** `requested`
- [ ] `deployment_status` - Deployment status activity
  - [ ] **Action:** `created`
- [ ] `discussion` - Discussion activity
  - [ ] **Action:** `answered`
  - [ ] **Action:** `category_changed`
  - [ ] **Action:** `closed`
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `labeled`
  - [ ] **Action:** `locked`
  - [ ] **Action:** `pinned`
  - [ ] **Action:** `reopened`
  - [ ] **Action:** `transferred`
  - [ ] **Action:** `unanswered`
  - [ ] **Action:** `unlabeled`
  - [ ] **Action:** `unlocked`
  - [ ] **Action:** `unpinned`
- [ ] `discussion_comment` - Discussion comment activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
- [ ] `fork` - Repository forked
- [ ] `github_app_authorization` - GitHub App authorization activity
  - [ ] **Action:** `revoked`
- [ ] `gollum` - Wiki page activity
- [ ] `installation` - GitHub App installation activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `new_permissions_accepted`
  - [ ] **Action:** `suspend`
  - [ ] **Action:** `unsuspend`
- [ ] `installation_repositories` - GitHub App installation repositories activity
  - [ ] **Action:** `added`
  - [ ] **Action:** `removed`
- [ ] `installation_target` - GitHub App installation target activity
  - [ ] **Action:** `renamed`
- [ ] `issue_comment` - Issue or PR comment activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
- [ ] `issue_dependencies` - Issue dependencies activity
  - [ ] **Action:** `blocked_by_added`
  - [ ] **Action:** `blocked_by_removed`
  - [ ] **Action:** `blocking_added`
  - [ ] **Action:** `blocking_removed`
- [ ] `issues` - Issue activity
  - [ ] **Action:** `assigned`
  - [ ] **Action:** `closed`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `demilestoned`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `labeled`
  - [ ] **Action:** `locked`
  - [ ] **Action:** `milestoned`
  - [ ] **Action:** `opened`
  - [ ] **Action:** `pinned`
  - [ ] **Action:** `reopened`
  - [ ] **Action:** `transferred`
  - [ ] **Action:** `unassigned`
  - [ ] **Action:** `unlabeled`
  - [ ] **Action:** `unlocked`
  - [ ] **Action:** `unpinned`
- [ ] `label` - Label activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
- [ ] `marketplace_purchase` - GitHub Marketplace purchase activity
  - [ ] **Action:** `cancelled`
  - [ ] **Action:** `changed`
  - [ ] **Action:** `pending_change`
  - [ ] **Action:** `pending_change_cancelled`
  - [ ] **Action:** `purchased`
- [ ] `member` - Repository collaborator activity
  - [ ] **Action:** `added`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `removed`
- [ ] `membership` - Team membership activity
  - [ ] **Action:** `added`
  - [ ] **Action:** `removed`
- [ ] `merge_group` - Merge group activity
  - [ ] **Action:** `checks_requested`
  - [ ] **Action:** `destroyed`
- [ ] `meta` - Webhook activity
  - [ ] **Action:** `deleted`
- [ ] `milestone` - Milestone activity
  - [ ] **Action:** `closed`
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `opened`
- [ ] `org_block` - Organization blocking activity
  - [ ] **Action:** `blocked`
  - [ ] **Action:** `unblocked`
- [ ] `organization` - Organization activity
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `member_added`
  - [ ] **Action:** `member_invited`
  - [ ] **Action:** `member_removed`
  - [ ] **Action:** `renamed`
- [ ] `package` - GitHub Packages activity
  - [ ] **Action:** `published`
  - [ ] **Action:** `updated`
- [ ] `page_build` - GitHub Pages build activity
- [ ] `personal_access_token_request` - Personal access token request activity
  - [ ] **Action:** `approved`
  - [ ] **Action:** `cancelled`
  - [ ] **Action:** `created`
  - [ ] **Action:** `denied`
- [ ] `ping` - Webhook ping
- [ ] `project` - Project (classic) activity
  - [ ] **Action:** `closed`
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `reopened`
- [ ] `project_card` - Project (classic) card activity
  - [ ] **Action:** `converted`
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `moved`
- [ ] `project_column` - Project (classic) column activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `moved`
- [ ] `projects_v2` - Projects activity
  - [ ] **Action:** `closed`
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `reopened`
- [ ] `projects_v2_item` - Projects item activity
  - [ ] **Action:** `archived`
  - [ ] **Action:** `converted`
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `reordered`
  - [ ] **Action:** `restored`
- [ ] `projects_v2_status_update` - Projects status update activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
- [ ] `public` - Repository visibility changed to public
- [ ] `pull_request` - Pull request activity
  - [ ] **Action:** `assigned`
  - [ ] **Action:** `auto_merge_disabled`
  - [ ] **Action:** `auto_merge_enabled`
  - [ ] **Action:** `closed`
  - [ ] **Action:** `converted_to_draft`
  - [ ] **Action:** `demilestoned`
  - [ ] **Action:** `dequeued`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `enqueued`
  - [ ] **Action:** `labeled`
  - [ ] **Action:** `locked`
  - [ ] **Action:** `milestoned`
  - [ ] **Action:** `opened`
  - [ ] **Action:** `ready_for_review`
  - [ ] **Action:** `reopened`
  - [ ] **Action:** `review_request_removed`
  - [ ] **Action:** `review_requested`
  - [ ] **Action:** `synchronize`
  - [ ] **Action:** `unassigned`
  - [ ] **Action:** `unlabeled`
  - [ ] **Action:** `unlocked`
- [ ] `pull_request_review` - Pull request review activity
  - [ ] **Action:** `dismissed`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `submitted`
- [ ] `pull_request_review_comment` - Pull request review comment activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
- [ ] `pull_request_review_thread` - Pull request review thread activity
  - [ ] **Action:** `resolved`
  - [ ] **Action:** `unresolved`
- [x] `push` - Git push to repository
- [ ] `registry_package` - Registry package activity
  - [ ] **Action:** `published`
  - [ ] **Action:** `updated`
- [ ] `release` - Release activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `prereleased`
  - [ ] **Action:** `published`
  - [ ] **Action:** `released`
  - [ ] **Action:** `unpublished`
- [ ] `repository` - Repository activity
  - [x] **Action:** `archived`
  - [x] **Action:** `created`
  - [x] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [x] **Action:** `privatized`
  - [x] **Action:** `publicized`
  - [x] **Action:** `renamed`
  - [ ] **Action:** `transferred`
  - [x] **Action:** `unarchived`
- [ ] `repository_advisory` - Repository security advisory activity
  - [ ] **Action:** `published`
  - [ ] **Action:** `reported`
- [ ] `repository_dispatch` - Repository dispatch event
- [ ] `repository_import` - Repository import activity
- [ ] `repository_ruleset` - Repository ruleset activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
- [ ] `repository_vulnerability_alert` - Repository vulnerability alert activity
  - [ ] **Action:** `create`
  - [ ] **Action:** `dismiss`
  - [ ] **Action:** `resolve`
- [ ] `secret_scanning_alert` - Secret scanning alert activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `reopened`
  - [ ] **Action:** `resolved`
  - [ ] **Action:** `revoked`
  - [ ] **Action:** `validated`
- [ ] `secret_scanning_alert_location` - Secret scanning alert location activity
  - [ ] **Action:** `created`
- [ ] `secret_scanning_scan` - Secret scanning scan activity
  - [ ] **Action:** `completed`
- [ ] `security_advisory` - Security advisory activity
  - [ ] **Action:** `performed`
  - [ ] **Action:** `published`
  - [ ] **Action:** `updated`
  - [ ] **Action:** `withdrawn`
- [ ] `security_and_analysis` - Security and analysis activity
- [ ] `sponsorship` - Sponsorship activity
  - [ ] **Action:** `cancelled`
  - [ ] **Action:** `created`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `pending_cancellation`
  - [ ] **Action:** `pending_tier_change`
  - [ ] **Action:** `tier_changed`
- [ ] `star` - Star activity
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
- [ ] `status` - Commit status updated
- [ ] `sub_issues` - Sub-issues activity
  - [ ] **Action:** `parent_issue_added`
  - [ ] **Action:** `parent_issue_removed`
  - [ ] **Action:** `sub_issue_added`
  - [ ] **Action:** `sub_issue_removed`
- [ ] `team` - Team activity
  - [ ] **Action:** `added_to_repository`
  - [ ] **Action:** `created`
  - [ ] **Action:** `deleted`
  - [ ] **Action:** `edited`
  - [ ] **Action:** `removed_from_repository`
- [ ] `team_add` - Team added to repository
- [ ] `watch` - Watch activity
  - [ ] **Action:** `started`
- [ ] `workflow_dispatch` - Workflow dispatch event
- [ ] `workflow_job` - Workflow job activity
  - [ ] **Action:** `completed`
  - [ ] **Action:** `in_progress`
  - [ ] **Action:** `queued`
  - [ ] **Action:** `waiting`
- [ ] `workflow_run` - Workflow run activity
  - [ ] **Action:** `completed`
  - [ ] **Action:** `in_progress`
  - [ ] **Action:** `requested`

> **Note:** Additional events may be supported. This list will be updated as implementation progresses.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

**Questions or Issues?** Open an issue on GitHub.
