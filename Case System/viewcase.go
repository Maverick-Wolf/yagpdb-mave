{{/*This command allows you to see details of a perticular case using the case number.

     Usuage: 
    
    -viewcase <case_number>
    
    Recommended trigger: Regex trigger with trigger `\A(-|<@!?204255221017214977>\s*)(viewcase|seecase|vc|sc|case)(\s+|\z)`
    NOTE : For any further doubts or bugs you can contact me on discord "Maverick Wolf#1010"
*/}}
{{$args := parseArgs 1 "`-viewcase case_number`" (carg "int" "case number")}}
{{$a := ""}}{{$name := ""}}{{$iconurl := "" }}{{$warnname := "" }}{{$reason := "" }}{{$userid := ""}}{{$action := ""}}{{$link := ""}}{{$msgid := ""}}{{$channel := ""}}{{$server := .Guild.ID}}
{{$b := ($args.Get 0)}}
{{if (dbGet $b "viewcase")}}
{{with (dbGet $b "viewcase")}}	
{{$a = sdict .Value}}
	{{$name = $a.Get "name"}}
	{{$iconurl = $a.Get "avatar"}}
	{{$warnname = $a.Get "warnname"}}
	{{$reason = $a.Get "reason"}}
	{{$action = $a.Get "action"}}
    	{{$userid = $a.Get "userid"}}
	{{$msgid = $a.Get "msgid"}}
	{{$channel = $a.Get "channel"}}
{{end}}
{{$embed := cembed "author" (sdict "name" $name "url" "" "icon_url" $iconurl) "description" (print "[**Jump to Warn**](https://discordapp.com/channels/" $server "/" $channel "/" $msgid "/)\n**User** : `" $warnname "`\t`(" $userid ")`\n **Action** : `" $action "`\n **Reason** : `" $reason "`") "color" 0x00ff00 "footer" (sdict "text" (print "Case#" $b))}}
{{sendMessage nil $embed}} 
{{else}}
{{sendMessage nil "Could not find the case specified,Please make sure the case number is correct or the case has not been deleted"}}
{{end}}