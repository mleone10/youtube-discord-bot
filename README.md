# youtube-discord-bot
> Bespoke Discord bot designed to poll a list of YouTube channels for new videos, then post them to a Discord channel.

## Setup
1. Create a Discord Bot and invite it to your server
Discord bots join channels when guild admins click on invite links like these.  The client ID comes from the bot management page, the permissions are a computed value based on what we actually need (in this case, permission to send messages and embed links).

```
https://discord.com/oauth2/authorize?client_id=<clientId>&scope=bot&permissions=18432
```
