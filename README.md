<img src="https://upload.wikimedia.org/wikipedia/commons/6/6f/Go_gopher_mascot_bw.png" width="250" height="250" title="GoGopher" alt="GoGopher">
powered by Google Golang </br>

[![Discord Bots](https://discordbots.org/api/widget/status/565229640646393895.svg)](https://discordbots.org/bot/565229640646393895)

# Install 
Just invite the bot to your server [here](https://discordapp.com/api/oauth2/authorize?client_id=565229640646393895&permissions=0&scope=bot). **The profile visibility by default is set to private, it must be set to PUBLIC in order to retriev data form the Overwatch Web Scraper API!!**.</br>*You should check out the !Config Command first*.</br> The bot is currently running on an [Odroid XU4Q](https://www.hardkernel.com/shop/odroid-xu4q-special-price/)</br>
Type !Help (**Case sensitive** not !help or !HeLp or any other modification) to see a list of currently supported commands.</br>
This list changes frequently as many new features coming out. </br>
Join our discord for support, news, or feature requests! [OverwatchTeamBotDeveloper Discord](https://discord.gg/x6RJhg)
# Commands
![](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/Help1.JPG)
![](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/Help2.JPG)
Currently supported commands:</br>
* !Training: Zeigt aktuelle Trainigszeiten/Show current training days<br/>
* !Training \<value\>: Aktualisiert Trainingszeiten/Update training days Bold or italic style? Check out Discord Markup https://gist.github.com/Almeeida/41a664d8d5f3a8855591c2f1e0e07b19<br/>
e.g. !Training "Trainings: :muscle:\r\n:calendar_spiral:Montag: ab 19:30 (Scrim, Review)\r\n:calendar_spiral:Dienstag: ab 19:30 (Scrim, Review)\r\n:medal:Donnerstag ab 19:30 (Ranked)"  Example uses discord markups and emotes like \: muscle\: for :muscle: and \r\n for a new line.</br>
![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/setTeams.JPG)
![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/GetTeams.JPG)
* !Stats \<battletag\>: Spieler Statistiken/Display player stats (z.B. !Stats Krusher-9911)<br/>
![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/Stats.JPG)
* !Register \<battletag\>: Registriert neuen Spieler/ Register new player, the player stats are then stored in the backend and updated on a daily basis (e.g. !Register Krusher-9911)<br/>
* !Update <battletag>: Aktualisiert Statistik für angegebenen Spieler/Update stored player stats (e.g. !Update Krusher-9911)<br/>
  ![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/Update.JPG)
* !Config platform=<pc|xbl|psn> region=<eu|us|asia> prefix=<...>: Default platform is pc and region is eu. If you are playing on Playstation or Xbox call !Config platform=xbl for Xbox and !Config platform=psn for PlayStation. Note you need region only for platform=pc (!Config platform=pc region=eu/us/asia). You can also change the Default prefix from ! To some custom prefix e.g. !Config prefix=&
* !Poll \<n\>: Starts a ready check poll for n players in the channel the command got posted in.
  ![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/NewPoll.JPG)
* !Poll: Gets the status of the current poll. If everybody is ready a message is created and <br/> the creator of the poll gets tagged.
  ![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/PollStatus.JPG)
* !+: Answeres with ready to a ready poll in this channel.
    ![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/Ready.JPG)
* !- \<reason\>: Answers with not ready to a ready poll in this channel. A reason an be defined. **Multiple Words in \"\"!**.
    ![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/NotReady.JPG)
    ![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/NotReadyWithReason.JPG)
* !DeletePoll: Deletes the current poll in the channel if existing.
    ![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/PollDeleted.JPG)
  </br>**When all poll slots are ready a message is created and the initiator of the poll is tagged.**
    ![Example](https://github.com/LukasMarckmiller/OverwatchDiscordBot/blob/master/img/PollFinished.JPG)

# OverwatchDiscordBot
Websocket based discord bot in GO with [Overwatch stats API](https://ow-api.com/) integration.<br/>
You can also copy an piece of code written for this project, and use it in your own project without any restrictions.<br/>
Im aware that there is [this](https://github.com/bwmarrin/discordgo) pretty good discord bindings api out, which also supports websocket connections. But i wanted to write my own small layer.<br/>
... More Infos coming soon...<br/>
**Caution!** This is a spare time project to dive deeper into Golang. This is not a professional, high scalable discord bot. 

# More Features? A Bug?
Just write an [Issue](https://github.com/LukasMarckmiller/OverwatchDiscordBot/issues) if you miss a crutial command/functionality or encountered a bug.

# Contribute
If you want to contribute in this project, feel free to do so.
First you need a Discord Bot API Token in order to send requests to the discord api endpoint. Feel free to contact me and i will send you the bot token. It should also be possible using other valid bot tokens to authenticate the websocket session (not tested).

MIT License

Copyright (c) 2019 Lukas Marckmiller

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

# LOGO
Go Gopher logo created by Renée French for Google Golang.
