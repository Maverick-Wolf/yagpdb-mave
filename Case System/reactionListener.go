{{/*
 	This command manages the pagination of the cases command.

	Recommended trigger: Reaction trigger on Reaction Added only.
	NOTE : For any further doubts or bugs you can contact me on discord "Maverick Wolf#1010"
*/}}
{{$action := .Reaction.Emoji.Name}} 
{{$validEmojis := cslice "▶️" "◀️"}} 
{{$isValid := false}} 
{{$page := 0}} 
{{$id := 0}}
{{$display := ""}}
{{with and (eq .ReactionMessage.Author.ID 204255221017214977) .ReactionMessage.Embeds}} 
	{{$embed := index . 0}} 
	{{if and (eq $embed.Title "Cases") $embed.Footer}} 
		{{$page = reFind `\d+` $embed.Footer.Text}} 
		{{$id = reFind `\d{17,19}` $embed.Author.Name}}
		{{$isValid = true}} 
	{{end}}
{{end}}
{{if and (in $validEmojis $action) $isValid $page}} 
	{{deleteMessageReaction nil .ReactionMessage.ID .User.ID $action}}
	{{if eq $action "▶️" }} {{$page = add $page 1}} 
	{{else}} {{$page = sub $page 1}} {{end}}
	{{if ge $page 1}} 
		{{$skip := mult (sub $page 1) 10}} 
		{{$cases := dbTopEntries $id 10 $skip}}
{{if (len $cases)}}
{{- range $cases -}}
{{$value := str .Value}} 
		{{$display = joinStr "" $display "\n" $value}} 
{{- end -}}
{{$user := userArg $id}}
{{editMessage nil .ReactionMessage.ID (cembed "author" (sdict "name" (print $user.Username " (" $user.ID ")") "icon_url" ($user.AvatarURL "512")) "title" "Cases" "description" $display "footer" (sdict "text" (print "Page " $page)) "color" 0x00ffd9)}}
{{end}}{{end}}{{end}}
