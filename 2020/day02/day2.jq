# jq -s --raw-input -f day2.jq input.txt
[
  split("\n") | .[] | select(.!="")
  | split(" ")|{
        policy: (.[0]|split("-")|{
          lower: .[0]|tonumber,
          upper: .[1]|tonumber
        }),
        char: .[1][0:1],
        pass: .[2]
    } as $input
    | $input | {
      inchar: $input.char,
      pass: $input.pass,
      p1: (
        ($input.pass|split("")|map(select(.==$input.char))|length) as $matches
        | ($input.policy.lower)<=$matches and $matches <= ($input.policy.upper)
      ),
      p2: (
        $input.pass[$input.policy.lower-1:$input.policy.lower] as $char1
        | $input.pass[$input.policy.upper-1:$input.policy.upper] as $char2
        | ($char1!=$char2 and ($char1==$input.char or $char2==$input.char))
      )
    }
] as $results
| {
  part1: [$results[] | select(.p1==true)]|length,
  part2: [$results[] | select(.p2==true)]|length,
}
