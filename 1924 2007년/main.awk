{
    month = $1
    date  = $2

    doms[ 1] = 31
    doms[ 2] = 28
    doms[ 3] = 31
    doms[ 4] = 30
    doms[ 5] = 31
    doms[ 6] = 30
    doms[ 7] = 31
    doms[ 8] = 31
    doms[ 9] = 30
    doms[10] = 31
    doms[11] = 30
    doms[12] = 31

    for (i = 1; i < month; i++)
    {
        date += doms[i]
    }

    days[0] = "SUN"
    days[1] = "MON"
    days[2] = "TUE"
    days[3] = "WED"
    days[4] = "THU"
    days[5] = "FRI"
    days[6] = "SAT"

    print days[date % 7]
}
