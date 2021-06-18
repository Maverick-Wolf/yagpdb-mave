{{/* As no bot has the inbuilt feature to log images so this custom command allows you to do that.

     Usuage: 
    
    Nothing much, just fill the log channel ID and you good to go, YAGPDB will do the rest.
    
    Recommended trigger:
      Trigger Type: Regex
      Trigger: \A
    NOTE : For any further doubts or bugs you can contact me on discord "Maverick Wolf#6565"
*/}}
{{$log_channel := 773561603659399218}}{{/*Replace with your log channel ID*/}}
{{/*Configration values end here*/}}
{{$extension := ".png" }}{{if reFind "a_" .Guild.Icon}}{{$extension = ".gif" }}{{end}}
{{$serverAvatar := (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon $extension "?size=512")}}{{/*getting the server avatar*/}}
{{$desc := or .Message.Content ""}}
{{$title := (or (and $desc "**Description :**") ""}}
{{$check := true}}{{/*a check for .mp4 and other video files which wont get shown in the embed*/}}
{{range .Message.Attachments}}
	{{if (reFind `(?i)\.(?:mp4|webm|mov)` .Filename)}}
		{{$check = false}}
	{{end}}
	{{if $check}}
		{{$embed := cembed
		"image" (sdict "url" .ProxyURL)
		"description" (print $title "\t" $desc "\n\n **Message sent in** \t " "  <#" $.Channel.ID "> ([Jump to Message](https://discordapp.com/channels/" $.Server.ID "/" $.Channel.ID "/" $.Message.ID "/))")
        "color" (randInt 111111 999999)
		"author" (sdict "name" $.User.String "icon_url" ($.User.AvatarURL "512"))
		"footer" (sdict "text" $.Server.Name "icon_url" $serverAvatar)
		"timestamp" currentTime}}
		{{sendMessage $log_channel  $embed}}
{{end}}{{end}}
