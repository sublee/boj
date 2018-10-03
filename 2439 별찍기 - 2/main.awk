{
    for (i = 1; i <= $1; i++)
    {
        for (j = i; j < $1; j++) printf " "
        for (j = 0; j <  i; j++) printf "*"
        print ""
    }
}
