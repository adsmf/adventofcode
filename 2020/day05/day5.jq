# jq -r -s --raw-input -f day5.jq input.txt
[
  split("\n")[:-1][]
  | split("")
  | reduce .[] as $char (0; (.*2)+(if ($char=="R" or $char=="B") then 1 else 0 end))
]
| "Part 1: \(.|max)\nPart 2: \(([range(.|min;.|max)]-.)[0])" 
