package telegram

const HelpMessage string = `*Git Gram – GitHub Notifications Bot*

Receive GitHub activity updates as Telegram messages directly in this chat\.

*Getting Started*

1\. Use /start to get your GitHub app installation link
2\. Install the Git Gram app on your GitHub account or organization
3\. Return here and run /status to confirm the connection

*Commands*

/start – Get GitHub app installation link and welcome message
/help – Display this help message
/status – View current GitHub installation and mute status
/mute – Pause GitHub notifications in this chat
/unmute – Resume GitHub notifications in this chat
/unlink – Disconnect the GitHub installation from this chat

*Tips*

💡 Use /mute during meetings or focus time, then /unmute when ready for updates
🔄 To switch accounts or reinstall: run /unlink first, then /start again
📊 Check /status anytime to verify your connection and notification settings`


const InvalidCommandMessage string = `❓ *Command not recognized*

I didn't understand that command\. Try one of these:

` + "`/start`" + ` – Get started with Git Gram
` + "`/help`" + ` – View all available commands
` + "`/status`" + ` – Check your connection status

Need help? Use ` + "`/help`" + ` to see the full command list\.`


const InstallationMessage string = `🚀 *Connect GitHub with GitGram*

To start receiving notifications:

1\. Tap the *Install GitHub App* button below\.
2\. Choose the repositories you want to connect\.
3\. Return here and send /status to confirm your setup\.`


const SuccessfulInstallationMessage string = `✅ *Git Gram App installed successfully*

You can now receive GitHub notifications in this chat\.

Send /status to check your connection status\.`


const MuteSuccessMessage = `🔕 *Notifications muted*

You will no longer receive GitHub updates in this chat\. Use /unmute to turn notifications back on\.`


const MuteBeforeStartErrorMessage = `⚠️ *Setup required*

You have not started GitGram in this chat yet\.  
Send /start first to link your GitHub installation, then use /mute\.`


const UnmuteSuccessMessage = `🔔 *Notifications unmuted*

You will now receive GitHub updates in this chat again\. Use /mute to turn notifications off\.`


const UnmuteBeforeStartErrorMessage = `⚠️ *Setup required*

You have not started GitGram in this chat yet\.  
Send /start first to link your GitHub installation, then use /unmute\.`


const StatusDocNotFoundMessage string = `📋 *Status*

This chat is not registered with Git Gram yet\.

Send /start to set up your GitHub installation, then run /status again\.`


const StatusNoInstallationMessage string = `📋 *Status*

You started setup, but the Git Gram app is not installed on GitHub yet\.

Open the previous installation link or send /start to get a new one, then install the app and run /status again\.`


const StatusInstalledTemplateMessage string = `📋 *Status*

This chat is linked to the GitHub account: *[%[1]s](https://github.com/%[1]s)*\.

Notifications are currently %s\.

%s\.
Use /unlink to disconnect this chat from GitHub\.`


const UnlinkSuccessMessage string = `✅ *Unlinked*

Git Gram has been disconnected from your GitHub account\.
This chat is no longer linked\.
Use /start to reinstall Git Gram and link again\.`


const UnlinkNotInstalledMessage string = `⚠️ *Nothing to unlink*

This chat is not linked to any GitHub account\.
Use /start to install GitGram and link your GitHub account first\.`


const UnlinkFailedMessage string = `❌ *Unlink failed*

I couldn't unlink your GitHub account\. Please try again\.
If this keeps happening, use /status to confirm your current link state\.`


const DefaultErrorMessage = `⚠️ *Something went wrong*

Failed to process your request\. Please try again later\.`
