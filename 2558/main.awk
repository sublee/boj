BEGIN { RS = "" ; FS = "\n" }
{
    print $1 + $2
}
