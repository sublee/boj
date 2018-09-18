#!/usr/bin/env bash

function solve() {
  local x="$1"
  local stick=64
  local count=0

  while [[ "$stick" -gt 0 ]]; do
    if [[ "$stick" -le "$x" ]]; then
      let 'x-=stick'
      let 'count++'
    fi

    let 'stick=stick/2'
  done

  echo "$count"
}

read x
solve "$x"
