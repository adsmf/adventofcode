# jq -r -s --raw-input -f day4.jq input.txt
[
  split("\n\n") | .[] | gsub("\n";" ") | split(" ") | map(select(.|contains(":")) | split(":") | {(.[0]): (.[1])}) | add
]
| [
  ([.[]|keys|select(.|contains(["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]))]|length),
  (
    [
      .[]|[
        (.byr|tostring|test("^(19[^01][0-9]|200[012])$";"")),
        (.iyr|tostring|test("^(201[0-9]|2020)$";"")),
        (.eyr|tostring|test("^(202[0-9]|2030)$";"")),
        (.hcl|tostring|test("^#[0-9a-f]{6}$";"")),
        (.pid|tostring|test("^[0-9]{9}$";"")),
        (.ecl|tostring|test("^(amb|blu|brn|gry|grn|hzl|oth)$";"")),
        (.hgt|tostring|test("^(((59|6[0-9]|7[0123456])in)|((1[5678][0-9]|19[0123])cm))$";""))
      ] | select(all)
    ] | length
  )
]| "Part 1: \(.[0])\nPart 2: \(.[1])"
