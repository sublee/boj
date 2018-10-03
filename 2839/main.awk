{
    x = $1

    # Divide by 5

    i  = int(x / 5)
    x %= 5

    if (x == 0)
    {
        print i
        exit
    }

    # Divide by 3

    while (x%3 != 0 && i > 0)
    {
        i--
        x += 5
    }

    i += int(x / 3)
    x %= 3

    if (x != 0)
    {
        print -1
        exit
    }

    print i
}
