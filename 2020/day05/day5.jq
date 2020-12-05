# jq -r -s --raw-input -f day5.jq input.txt
[
  split("\n")[:-1]
  | .[]
  | split("")
  | reduce .[] as $char
      (0; (.*2) as $acc | $acc+(if ($char=="R" or $char=="B") then 1 else 0 end))
]
|sort
| first as $min
| last as $max
|[
  .|max,
  ([range($min;$max)]-.)[0]
] |"Part 1: \(.[0])\nPart 2: \(.[1])" 
