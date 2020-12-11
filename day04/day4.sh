#!/usr/bin/env bash

input_file="input.txt"

if ! [ -e "$input_file" ] ; then
  echo "$input_file does not exist" >&2
  exit 1
fi

FIELDS=("byr" "iyr" "eyr" "hgt" "hcl" "ecl" "pid")

valid_passports=0

fields=()
while IFS= read -r line
do
  if [ -z "$line" ] ; then
    missing=()
    for i in "${FIELDS[@]}"; do
      skip=
      for j in "${fields[@]}"; do
        [[ $i == $j ]] && { skip=1; break; }
      done
      [[ -n $skip ]] || missing+=("$i")
    done
    [[ ${#missing[@]} -eq 0 ]] && ((valid_passports++))
    fields=()
  fi
  IFS=" " read -a vals <<< "$line"
  for val in "${vals[@]}" ; do
    read -d ':' key <<< "$val"
    fields+=("$key")
  done
done <"$input_file"

echo "Valid passports: $valid_passports"