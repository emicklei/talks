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
# What is Melrōse ?

- Expression language to write music, as programs

- Runtime to play and interact

![domeka](img/domeka.png)

---
![note_spec center](img/note.png)

---
![seq_spec center](img/sequence.png)

---
# Example: sequence

    sequence('c d e f g a b 1c5 8c5 8a 8b 8g 8e 8f 8g (1c 1e 1g)')

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
# Example: sequence

Merge, fraction and repeat

    s = sequence(' e  g#  c#5  e5      f#  a  c#5 f#5     g#  b  d#5 g#5   a5 g#5 f#5 e5')
    t = sequence('2e2    2c#2         2f#2   2c#3        2g#2   2d#3')

    repeat(3,fraction(8,merge(s,t)))

---
# More composition functions

- at, duration, dynamic, dynamicmap, export, `fraction`, fractionmap, group, if, import, index, interval, iterator, join, joinmap, listen, `merge`, midi_send, next, notemap, octave, octavemap, onbar, pitch, pitchmap, print, random, `repeat`, replace, resequence, `reverse`, stretch, track, `transpose` (pitch), transposemap (pitchmap), undynamic, ungroup, value, velocitymap

---
# Drum patterns

`notemap` can create a sequence using `dots and bangs`.

&nbsp;

    kick = note('16c2')

    channel(10, notemap('!...!!..!!!...!', kick))

---
# No sound

Melrōse does not produce any sound directly.

The tool sends `MIDI`

---
# MIDI communication

![height:400px center ](img/melrose-port-daw.png)

---
# Show me and let's hear it!

---
# Thank you for listening !

- Open source, MIT licensed

- github.com/emicklei/melrose

- https://melrōse.org

- (music) contributions are welcome

