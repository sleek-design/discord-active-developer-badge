# Active Developer Badge

Earn the "Active Developer" badge from Discord using this simple self-contained bot!

## Getting Started

1. [Download the bot](https://github.com/benjammin4dayz/discord-active-developer-badge/releases/latest) for your system and extract the file(s) anywhere you prefer

2. [Create a Discord application](https://discord.com/developers/applications) on the account you want the badge on

3. Create a Discord server and [enable Community](https://support.discord.com/hc/en-us/articles/360047132851-Enabling-Your-Community-Server)

## Instructions

1. On the Discord application page, go to the `Bot` tab and make a bot account for the application

2. Click the `Reset Token` button to generate a token (if you have 2fa enabled, you'll need your code)

3. Copy the token to your clipboard

4. Start the bot by running the program extracted in the previous step

5. Paste (<kbd>CTRL</kbd><kbd>SHIFT</kbd><kbd>V</kbd>) the token in the terminal when requested

   - This can also be set via the `BOT_TOKEN` environment variable (either shell or `.env` file)

6. Click the invite link that appears in the terminal to invite the bot to your server

7. In your server, go to a channel and use the `/ping` command

8. Wait ~24 hours, then go to the [Active Developer](https://discord.com/developers/active-developer) page and claim your badge

> [!Important]
> In order to detect command usage, you or at least one person on the team that owns the app needs to have "Use data to improve Discord" enabled within User Settings > Data & Privacy. At least 24 hours need to pass after detecting a command, so be sure to wait at least 24 hours after enabling this setting before trying again.

## Acknowledgement

This project wouldn't exist without inspiration from [hackermondev & contributors](https://github.com/hackermondev/discord-active-developer-badge)
