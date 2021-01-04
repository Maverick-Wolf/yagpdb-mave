{{/*
    This is the main custom command which generates the major information such as the case number and the details about the same and sets it to a database to be used by other cc's

    Add this code in the Moderation DM template space of "Mute,Ban,Kick and Warn" under `Control Panel->Tools and Utilities->Moderation->Then the respective Moderation`
*/}}

{{$dm := 1}}{{/* change to 0 if you don't wanna DM the offender about the moderation action*/}}
{{if $dm}}
You have been {{.ModAction}}{{/*You can edit this if you want to chnage the DM message*/}}
{{if .Reason}}**Reason:** {{.Reason}}{{end}}{{end}}
{{$case_number := (toInt (dbIncr 77 "cv" 1))}}
{{$action := .ModAction.Prefix}}
{{$a := 0}}
{{if eq $action "Muted" "Unmuted"}}
{{$a = (sub (len .ModAction.Prefix) 1)}}
{{else}}
{{$a = (sub (len .ModAction.Prefix) 2)}}
{{end}}
{{$title := (slice .ModAction.Prefix 0 $a)}}
{{$id := .User.ID}}
{{$channel := .Channel.ID}}
{{$x := sendMessageRetID nil (cembed "title" (print $action " | case# " $case_number) "description" (print "**Offender :** `" .User.Username "`\t`(" .User.ID ")`\n**Reason :** `" .Reason "`\n**Moderator :** `" .Author.String "`") "timestamp" currentTime "color" 0xfc0606)}}
{{dbSet $case_number "viewcase" (sdict "name" .Author.Username "warnname" .User.Username "avatar" (.Author.AvatarURL "512") "reason" .Reason "userid" $id "action" (.ModAction.Prefix) "channel" $channel "msgid" $x)}}{{/*for viewcase*/}}
{{dbSet $case_number $id (print "Case# **" $case_number "**\t\t**| " $title " Reason:** `" .Reason "`")}}{{/*for per user case viewing*/}}
{{dbSet $case_number "userid" (str $id)}}{{/* for delete case*/}}
