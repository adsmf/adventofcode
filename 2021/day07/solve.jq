def a(v): if v<0 then -v else v end;
.|split(",")|map(tonumber)|
[
  (sort|.[length/2]) as $md|(map(a($md-.)) |add),
  (add/length|floor) as $mn|(map(a(.-$mn)|(.*(.+1)/2))|add)
]
