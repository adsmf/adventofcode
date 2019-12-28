# https://oeis.org/A006715

$s = 1321131112;
for (2..shift @ARGV) {
    $s =~ s/(.)\1*/(length $&).$1/eg;
}
print "$s\n";
