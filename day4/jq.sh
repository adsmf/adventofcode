#!/bin/bash

input="278384-824795"
echo "Input: $input"
echo "Expected 921, 603 ..."

time jq '
def repetitions:
    reduce
        (split(""))[] as $digit
        ( [];
            if . == [] then
                [ [ $digit, 1] ]
            else
                .[length-1] as $last
                | if $last[0] == $digit then
                    .[length-1] = [$digit, $last[1] + 1]
                else
                    . + [[$digit, 1]]
                end
            end
        ) |
    .[][1];

def valid:
    tostring|
    select(length==6)|
    (split("")) as $digits |
    ($digits|sort) as $sorted |
    select($digits==$sorted) |
    repetitions | select(. > 1);

def valid_part1:
    [valid | select(.>=2)]|length>0|select(.==true);

def valid_part2:
    [valid | select(.==2)]|length>0|select(.==true);


split("-") |
    ([range(.[0]|tonumber;.[1]|tonumber)|valid_part1]|length) as $d1 |
    ([range(.[0]|tonumber;.[1]|tonumber)|valid_part2]|length) as $d2 |
{"part1": $d1,"part2": $d2}
' -s --raw-input <<<"$input"
