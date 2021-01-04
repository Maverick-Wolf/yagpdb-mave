{{/*A role info command.
	Usuage: 
    
    -roleinfo <role ID/role Name/role Mention> 
    
    Recommended trigger: Regex trigger with trigger `\A(-|<@!?204255221017214977>\s*)(roleinfo|ri)(\s+|\z)`
    NOTE : For any further doubts or bugs you can contact me on discord "Maverick Wolf#1010"
*/}}
{{$role := ""}}
{{$input := .StrippedMsg}}
{{if reFind `\d{17,19}` $input}}
{{$input = reFind `\d{17,19}` $input}}
{{end}}
{{range .Guild.Roles}}{{if eq (str $input|lower) (str .ID) (lower .Name)}}{{$role = .}}{{end}}{{end}}
{{$all_perms := ""}}
{{$ex := "png" }}{{if reFind "a_" .Guild.Icon}}{{$ex = "gif" }}{{end}}{{$icon := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=512" }}
{{if $role}}
  {{$perm := $role.Permissions}}
  {{$perms_list := cslice "CreateInstantInvite" "KickMembers" "BanMembers" "Administrator" "ManageChannels" "ManageServer" "AddReactions" "ViewAuditLogs" "PrioritySpeaker" "Stream" "ViewChannel" "SendMessages" "SendTTSMessages" "ManageMessages" "EmbedLinks" "AttachFiles" "ReadMessageHistory" "MentionEveryone" "UseExternalEmojis" "ViewGuildInsigths" "VoiceConnect" "VoiceSpeak" "VoiceMuteMembers" "VoiceDeafenMembers" "VoiceMoveMembers" "VoiceUseVAD" "ChangeNickname" "ManageNicknames" "ManageRoles" "ManageWebhooks" "ManageEmojis"}}
  {{range seq 0 (len $perms_list)}}
     {{if (mod $perm 2)}}
	{{if $all_perms}}
	  {{$all_perms = (print $all_perms ", `" (index $perms_list .) "`")}}
	{{else}}
	  {{$all_perms = (print "`" (index $perms_list .) "`")}}
	{{end}}
     {{end}}
        {{$perm = (div $perm 2)}}
  {{end}}
{{$embed := (cembed "fields" (cslice (sdict "name" "Role Name" "value" (print $role.Name) "inline" true) 
				     (sdict "name" "Role Color" "value" (print $role.Color) "inline" true)
                                     (sdict "name" "Role ID" "value" (print $role.ID) "inline" true) 
                                     (sdict "name" "Position" "value" (print (sub (len .Guild.Roles) $role.Position)) "inline" true)
                                     (sdict "name" "Hoisted" "value" (print $role.Hoist) "inline" true)
  			             (sdict "name" "Managed" "value" (print $role.Managed) "inline" true)
  			             (sdict "name" "Permissions" "value" (print $all_perms))) "color" $role.Color
		  "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "512"))
		  "thumbnail" (sdict "url" $icon)
		  "footer" (sdict "text" .Server.Name) "timestamp" currentTime)}}
{{sendMessage nil $embed}}
{{else}}
`Please enter the correct role ID/role Name/role Metion`
{{end}}
