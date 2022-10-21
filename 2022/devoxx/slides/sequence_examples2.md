# Example: sequence

Merge, fraction and repeat

    s = sequence(' e  g#  c#5  e5      f#  a  c#5 f#5     g#  b  d#5 g#5   a5 g#5 f#5 e5')
    t = sequence('2e2    2c#2         2f#2   2c#3        2g#2   2d#3')

    repeat(3,fraction(8,merge(s,t)))