Discord bots join channels when guild admins click on invite links like these.  The client ID comes from the bot management page, the permissions are a computed value based on what we actually need (in this case, permission to send messages and embed links).

```
https://discord.com/oauth2/authorize?client_id=728013471886606449&scope=bot&permissions=18432
```
