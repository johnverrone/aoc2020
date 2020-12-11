#!/usr/bin/awk -f

function valid(key, val) {
  if (key == "byr") {
    return val ~ /^[[:digit:]]{4}$/ && val >= 1920 && val <= 2002
  }
  if (key == "iyr") {
    return val ~ /^[[:digit:]]{4}$/ && val >= 2010 && val <= 2020
  }
  if (key == "eyr") {
    return val ~ /^[[:digit:]]{4}$/ && val >= 2020 && val <= 2030
  }
  if (key == "hgt") {
    return val ~ /^[[:digit:]]{3}cm$/ && substr(val, 1, 3) >= 150 && substr(val, 1, 3) <= 193 \
           || val ~ /^[[:digit:]]{2}in$/ && substr(val, 1, 2) >= 59 && substr(val, 1, 2) <= 76
  }
  if (key == "hcl") {
    return val ~ /^#[[:xdigit:]]{6}$/
  }
  if (key == "ecl") {
    return val ~ /^(amb|blu|brn|gry|grn|hzl|oth)$/
  }
  if (key == "pid") {
    return val ~ /^[[:digit:]]{9}$/
  }
  if (key == "cid") {
    return 1
  }
  printf "Record %d invalid key: %s\n", NR, key
  return 0
}

BEGIN { RS = "" }

NF <= 6 { next }
NF == 7 {
  for (i=1; i <= NF; ++i) {
    if (substr($i, 1, 3) == "cid") {
      next
    }
  }
}

{
  for (i = 1; i <= NF; ++i) {
    split($i, a, /:/)
    if (! valid(a[1], a[2])) {
      next
    }
  }
  ++count
}

END { print count }