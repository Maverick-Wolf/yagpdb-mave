{{/*Does not really require an explaination y'all know what this command is.

	Usuage: 
    
    -pp [user] where user is optional 
    
    Recommended trigger: Regex trigger with trigger `\A(-|<@!?204255221017214977>\s*)(pp|penis)(\s+|\z)`
    NOTE : For any further doubts or bugs you can contact me on discord "Maverick Wolf#1010"
*/}}
{{$args := parseArgs 0 "" (carg "member" "the pp user lol")}}
{{$user := .User}}
{{with $x := ($args.Get 0)}}
	{{$user = $x.User}}
{{end}}
{{/*The variables below declare the range for the small,medium and big pp size to show some comments accordingly*/}}
{{$small := 6}}
{{$medium := 10}}
{{$big := 13}}{{/*also the size of the biggest pp,change it if you want*/}}
{{$smallpp := cslice "Hope you don’t get turned down by ladies with that baby carrot you got down there."
                     "Your dick is so small,you could make love to an ant."
                     "Your dick is so small,it makes my pinky finger look like a giant."
                     "Your dick is so small,the only thing you hear in bed is, ‘is it in?"
  		               "I bet your condoms look like the thumb of a latex glove"
   		               "Wow, I've Never Been Able To Fit A Whole One In My Mouth Before! ;)"
		                 "Legit feel sorry for you man"}}
{{$mediumpp := cslice "Aight Aight, the nerd got some meat"
  					  "Under-construction third leg, I see."
  					  "You surprisingly have got a nice piece of meat there"
  					  "Average pp better than nothing"
 					    "Don't Worry, I've Seen Smaller."
					    "Pretty decent !"}}
{{$bigpp := cslice "She is gonna enjoy today"
 				 "That’s one handsome penis."
 				 "Just wanna you to know, you’re ENORMOUS…"
 				 "Can’t believe all of this is on you!"
 				 "You gonna be good at your job ;)"
 				 "Oh oh oh!"
				 "Aight Aight no need to flex that"}}
{{$size := randInt (add $big 1)}}
{{$comment := ""}}
{{if and (ge $size 0) (le $size $small)}}
	{{$comment = index (shuffle $smallpp) (randInt (len $smallpp))}} 
{{else if and (ge $size (add $small 1)) (le $size $medium)}}
	{{$comment = index (shuffle $mediumpp) (randInt (len $mediumpp))}}
{{else}}
	{{$comment = index (shuffle $bigpp) (randInt (len $bigpp))}}
{{end}}
{{$pp := "8"}}
{{range seq 0 $size}}
	{{- $pp = (print $pp "=") -}}
{{- end -}}
{{$pp = (print $pp "D")}}
{{sendMessage nil (cembed "author" (sdict "name" $user.Username "icon_url" ($user.AvatarURL "512")) "description" (print $comment "\n" $pp) "color" 0x00ffd9)}}
