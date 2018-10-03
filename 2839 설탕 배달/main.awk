{
    n = $1

    if (n >= 10)
    {
        # Pattern above 10: 23434 34545 45656...
        print int(n/5) + substr("01212", n%5 +1, 1)
        #                                    └─ substr() uses 1-based indexing.
    }
    else
    {
        # Only 4 answers: 3:1, 5:1, 6:2, 9:3
        print (n%3 * n%5 == 0) ? int(n/3) : -1
    }
}
