# Merge patterns of notes

    kick = note('c2')
    snare = note('b2')
    closehi = midi(4,53,72)

    drums = merge(
        notemap('!.!....!.!!!...!',kick),
        notemap('....!.......!...',snare),
        notemap('!.!.!.!.!.!.!.!.',closehi))

    channel(10,repeat(4,fraction(16,drums)))