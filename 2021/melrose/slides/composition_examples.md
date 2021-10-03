# Example: composition

    y = sequence('f#2 c#3 f#3 a3 c# f# = ')
    p = resequence('3 4 2 5 1 6 2 5', y)
    f = fraction(8, p)
    jf = join(repeat(2,f),
        repeat(2,pitch(1,f)), 
        repeat(2,pitch(-2,f)), 
        repeat(2,pitch(3,f)))