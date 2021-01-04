{{/*This command allows you to see details of a perticular case using the case number.

     Usuage: 

    -viewcase <case_number>

    Recommended trigger: Regex trigger with trigger `\A(-|<@!?204255221017214977>\s*)(viewcase|seecase|vc|sc|case)(\s+|\z)`
*/}}
{{$args := parseArgs 1 "`-viewcase case_number`" (carg "int" "case number")}}
{{$a := ""}}{{$name := ""}}{{$iconurl := "" }}{{$warnname := "" }}{{$reason := "" }}{{$userid := ""}}{{$action := ""}}{{$link := ""}}{{$msgid := ""}}{{$channel := ""}}{{$server := .Guild.ID}}
{{$b := ($args.Get 0)}}
{{with (dbGet $b "viewcase")}}	
{{$a = sdict .Value}}
	{{$name = $a.name}}
	{{$iconurl = $a.avatar}}
	{{$warnname = $a.warnname}}
	{{$reason = $a.reason}}
	{{$action = $a.action}}
    	{{$userid = $a.userid}}
	{{$msgid = $a.msgid}}
	{{$channel = $a.channel}}
{{$embed := cembed "author" (sdict "name" $name "url" "" "icon_url" $iconurl) "description" (print "[**Jump to Warn**](https://discordapp.com/channels/" $server "/" $channel "/" $msgid "/)\n**User** : `" $warnname "`\t`(" $userid ")`\n **Action** : `" $action "`\n **Reason** : `" $reason "`") "color" 0x00ff00 "footer" (sdict "text" (print "Case#" $b))}}
{{sendMessage nil $embed}} 
{{else}}
Could not find the case specified,Please make sure the case number is correct or the case has not been deleted
{{end}}
