BEGIN { total = 0 }
{
    total += ($1<40 ? 40 : $1)
}
END { print total / 5 }
