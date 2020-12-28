{{/* This command allows you to see all the cases that a user has.

     Usuage: 
    
    -cases <user> [page] where page is optional 
    
    Recommended trigger: Regex trigger with trigger `\A(-|<@!?204255221017214977>\s*)(cases|allcase)(\s+|\z)`
    NOTE : For any further doubts or bugs you can contact me on discord "Maverick Wolf#1010"
*/}}
{{$args := parseArgs 1 "correct usuage is -cases @user" (carg "member" "target user") (carg "int" "page number")}}
{{$user := ($args.Get 0).User}}
{{$id := ($args.Get 0).User.ID}}
{{$display := ""}}
{{$page := 1}} {{/* Default page to start at */}}
{{if ($args.IsSet 1)}}
{{if ge ($args.Get 1) 1}}
{{$page = ($args.Get 1)}}
{{else}}
{{$page = 1}}
{{end}}{{end}}
{{$skip := mult (sub $page 1) 10}}
{{$cases := dbTopEntries $id 10 $skip}}
{{if not $cases}}
{{$display = "`No Cases exists on this page`"}}
{{else}}
{{- range $cases -}}
		{{ $value := str .Value }} 
		{{ $display = joinStr "" $display "\n" $value}} 
{{- end -}}{{- end -}}
{{$id := sendMessageRetID nil (cembed "author" (sdict "name" (print $user.Username " (" $user.ID ")") "icon_url" ($user.AvatarURL "512")) "title" "Cases" "description" $display "footer" (sdict "text" (print "Page " $page)) "color" 0x00ffd9)}}
{{addMessageReactions nil $id "◀️" "▶️"}}