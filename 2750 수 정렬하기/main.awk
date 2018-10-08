{
    # Keep numbers except the first N.
    if (NR != 1) a[NR-2] = $1
}

END {
    n = NR-1

    # Bubble Sort
    for (i = 0; i < n; i++) {
        for (j = i; j < n; j++) {
            if (a[i] > a[j]) {
                tmp = a[i]
                a[i] = a[j]
                a[j] = tmp
            }
        }
    }

    # Print sorted numbers.
    for (i = 0; i < n; i++) print a[i]
}
