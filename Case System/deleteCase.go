{{/*This command allows you to delete a perticular case using the case number.

     Usuage: 

    -delcase <case_number>

    Recommended trigger: Regex trigger with trigger `\A(-|<@!?204255221017214977>\s*)(delcase|deletecase|clearcase|dc|clc)(\s+|\z)`
*/}}
{{$args := parseArgs 1 "correct usuage is -delcase case_number" (carg "int" "case number")}}
{{$caseno := ($args.Get 0)}}
{{$id := (toInt (dbGet $caseno "userid").Value)}}
{{dbDel $caseno "viewcase"}}
{{dbDel $caseno $id}}
{{dbDel $caseno "userid"}}
`Deleted the case`
