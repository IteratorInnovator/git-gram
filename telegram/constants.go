package telegram

const HelpMessage string = `*Git Gram – GitHub Notifications Bot*

Receive GitHub activity updates as Telegram messages directly in this chat\.

*Getting Started*

1\. Use ` + "`/start`" + ` to get your GitHub app installation link
2\. Install the GitGram app on your GitHub account or organization
3\. Return here and run ` + "`/status`" + ` to confirm the connection

*Commands*

` + "`/start`" + ` – Get GitHub app installation link and welcome message
` + "`/help`" + ` – Display this help message
` + "`/status`" + ` – View current GitHub installation and mute status
` + "`/mute`" + ` – Pause GitHub notifications in this chat
` + "`/unmute`" + ` – Resume GitHub notifications in this chat
` + "`/unlink`" + ` – Disconnect the GitHub installation from this chat

*Tips*

💡 Use ` + "`/mute`" + ` during meetings or focus time, then ` + "`/unmute`" + ` when ready for updates
🔄 To switch accounts or reinstall: run ` + "`/unlink`" + ` first, then ` + "`/start`" + ` again
📊 Check ` + "`/status`" + ` anytime to verify your connection and notification settings`


var InvalidCommandMessage string = `❓ *Command not recognized*

I didn't understand that command\. Try one of these:

` + "`/start`" + ` – Get started with Git Gram
` + "`/help`" + ` – View all available commands
` + "`/status`" + ` – Check your connection status

Need help? Use ` + "`/help`" + ` to see the full command list\.`

var InstallationMessage string = `Install GitHub app`
