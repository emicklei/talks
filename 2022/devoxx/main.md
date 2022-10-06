---
marp: true
theme: default
paginate: true
color: "#EEE"
backgroundColor: "#222"
footer: <h3>melrōse.org</h3>
header: <h4>Devoxx Antwerp, October 2022, Melrōse - music programming</h4>

---
# Melrōse, program and play music

#### Ernest Micklei, October 2022

<style>
pre,code {
  background: #00527A;
  color: yellow;
}
a {
  color: cyan;
}
img[alt~="center"] {
  display: block;
  margin: 0 auto;
}
</style>
<script src="slides/play.js"></script>

---
# Motivation

- Explore patterns in music expressed by programming constructs

- Create a simple but expressive language

- Offer quick audio feedback

---
# What is Melrōse ?

- Language to create music, as programs

- Runtime environment to play programs

![domeka](img/domeka.png)

---
![note_spec center](img/note.png)

---
![seq_spec center](img/sequence.png)

---
# Other music primitives

    chord('c#/m')

    scale(2,'8g#3')

    progression('2c#', 'I VI II V')

    chordsequence('a3 b3 c#/m')

---
# Composition

![height:400px center](img/composition.png)

---
# Drum patterns

`notemap` can create a sequence using `dots and bangs`.
    
    kick = note('16c2')

    channel(10, notemap('!...!...!...!...', kick))

---
# Merge

    kick = note('c2')
    snare = note('e2')
    closehi = midi(4,42,72)

    drum14 = merge(
        notemap('!.!....!.!!!...!',kick),
        notemap('....!.......!...',snare),
        notemap('!.!.!.!.!.!.!.!.',closehi))

    bpm(120)
    channel(10,repeat(4,fraction(16,drum14)))

---
# No sound

Melrōse does not produce any sound directly.

The tool sends `MIDI`.

---
# MIDI communication

![height:400px center ](img/melrose-port-daw.png)

---
# Show me and let's hear it!

---
# Source bits

- Open source, MIT licensed

- github.com/emicklei/melrose

- https://melrōse.org

- (music) contributions are welcome

