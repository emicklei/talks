---
marp: true
theme: default
paginate: true
color: "#EEE"
backgroundColor: "#222"
footer: <h3>melrōse.org</h3>
header: <h4>Golang Meetup Amsterdam, October 2021</h4>

---
# Melrōse, program and play music

#### Ernest Micklei, October 2021

<style>
pre,code {
  background: #eee;
  color: black;
}
</style>
<script src="slides/play.js"></script>

---
# Motivation

- Explore patterns in music expressed by programming constructs

- Challenge to find the right abstraction and functions

- Offer quick audio feedback

---
# What is Melrōse ?

- Language to create music, as programs

- Runtime environment to play programs

![domeka](img/domeka.png)

---
# Language Bits

 

---
# Note
 
![note_spec](img/note.png)

---
# Examples: note

      note('c') 
      note('c5') // octave 5
      note('8c') // duration 1/8
      note('e#') // sharp
      note('b_') // flat

      note('.1a#3++')

Raw

      midi(2,37,72) // 1/2 duration, MIDI nr 37, velocity 72  

---
# Sequence

![seq_spec](img/sequence.png)

---
# Examples: sequence

    f1 = sequence('c e g c')

---
# Other music primitives

    chord('c#/m')

    scale(2,'8g#3')

    progression('2c#', 'I VI II V')

    chordsequence('a3 b3 c#/m')

---
# Composition

![comp](img/composition.png)

---
# Example: composition

    y = sequence('f#2 c#3 f#3 a3 c# f# = ')
    p = resequence('3 4 2 5 1 6 2 5', y)
    f = fraction(8, p)
    jf = join(repeat(2,f),
        repeat(2,pitch(1,f)), 
        repeat(2,pitch(-2,f)), 
        repeat(2,pitch(3,f)))

---
# More composition functions

- at, duration, dynamic, dynamicmap, export, fraction, fractionmap, group, if, import, index, interval, iterator, join, joinmap, listen, merge, midi_send, next, notemap, octave, octavemap, onbar, pitch, pitchmap, print, random, repeat, replace, resequence, reverse, stretch, track, transpose (pitch), transposemap (pitchmap), undynamic, ungroup, value, velocitymap

---
# Tool bits

---
# Go bits

---
# Sound bits

