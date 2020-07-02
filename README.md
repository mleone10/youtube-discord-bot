# youtube-discord-bot
> Bespoke Discord bot designed to poll a list of YouTube channels for new videos, then post them to a Discord channel.

## Setup
1. Create a Discord Bot and invite it to your server

    Visit the Discord Developer Portal, create a new application, and add a bot.  Take note of the application's **client ID** and the bot's **API token**.
    
    Discord bots join channels when guild admins click on invite links like these.  Add your client ID to the below URL and navigate to it in a browser.  The permission parameter specifies which accesses the bot will be granted (in this case, permission to send messages and embed links).

    ```
    https://discord.com/oauth2/authorize?client_id=<clientId>&scope=bot&permissions=18432
    ```

1. Generate a YouTube API Key

    Follow YouTube's instructions [here](https://developers.google.com/youtube/v3/getting-started).  Create new credentials and note the **API token**.

1. Test your credentials

    Create a `local.env` file with the following:
   
    ```bash
    export YT_API_KEY=<YouTube API Key>
    export YT_DELTA_MINUTES=<variable delta minutes to check, i.e. "30" for last 30 minutes>
    export YT_CHANNELS=<comma-delimited list of YouTube channel names>
    export DISCORD_BOT_TOKEN=<Discord bot API token
    export DISCORD_CHANNEL_ID=<Target Discord channel ID>
    ```
    
    Run `source local.env && go run ./...`.  Videos posted by the listed YouTube channels in the last `<delta>` minutes should be posted to the target Discord channel.

1. Prepare your deploy environment

    TBD

1. Deploy the bot to AWS

    TBD
