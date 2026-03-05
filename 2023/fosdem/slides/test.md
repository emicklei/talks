---
marp: true
theme: default
paginate: true
color: "#EEE"
backgroundColor: "#222"
footer: <h3>melrōse.org</h3>
header: <h4>FOSDEM Brussels, February 2023, Melrōse - music programming</h4>

_head.md_ 1

---
# Melrōse, program and play music

#### Ernest Micklei, February 2023

![height:100px center](/img/emicklei_hackers_logo.png)

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
  background-color: transparent;
}
</style>
<script src="slides/play.js"></script>

_title.md_ 2

---
# Motivation

- Explore patterns in music expressed by programming constructs

- Challenge to find the right abstraction and functions

- Offer quick audio feedback

_intro.md_ 3

---
# What is Melrōse ?

- Language to create music, as programs

- Runtime environment to play programs

![domeka](img/domeka.png)

_whatisit.md_ 4

---
![center](img/overview.png)

_overview.md_ 5

---
# Language Bits

 

_language_bits.md_ 6

---
![note_spec center](img/note.png)

_note.md_ 7

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

_note_examples.md_ 8

---
![seq_spec center](img/sequence.png)

_sequence.md_ 9

---
# Examples: sequence

    sequence('c d e f g a b 1c5 8c5 8a 8b 8g 8e 8f 8g (1c 1e 1g)')

    sequence('= 8a3 8b3 2c-  = 8c 8d 2e-  = 8d 8e 2f-   = 8c 8b3 2a3-')

_sequence_examples.md_ 10

---
# Other music primitives

    chord('c#/m')

    scale(2,'8g#3')

    progression('2c#', 'I VI II V')

    chordsequence('a3 b3 c#/m')

_other_creates.md_ 11

---
# Composition

![height:400px center](img/composition.png)

_composition.md_ 12

---
# Example: composition

    y = sequence('f#2 c#3 f#3 a3 c# f# = ')
    p = resequence('3 4 2 5 1 6 2 5', y)
    f = fraction(8, p)
    
    jf = join(repeat(2,f),
        repeat(2,pitch(1,f)), 
        repeat(2,pitch(-2,f)), 
        repeat(2,pitch(3,f)))

_composition_examples.md_ 13

---
# Drum patterns

`notemap` can create a sequence using `dots and bangs`.    
    
    kick = note('16c2')

    channel(10, notemap('!...!...!...!...', kick))

_drum.md_ 14

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

_drum_merge.md_ 15

---
# More composition functions

- at, duration, dynamic, dynamicmap, export, `fraction`, fractionmap, group, if, import, index, interval, iterator, join, joinmap, listen, `merge`, midi_send, next, notemap, octave, octavemap, onbar, pitch, pitchmap, print, random, `repeat`, replace, resequence, `reverse`, stretch, track, `transpose` (pitch), transposemap (pitchmap), undynamic, ungroup, value, velocitymap

_more_composition.md_ 16

---
# Tool bits

_play_bits.md_ 17

---
# No sound

Melrōse does not produce any sound directly.

The tool sends `MIDI`.

_nosound.md_ 18

---
# MIDI

- Musical Instrument Digital Interface is an industry standard music technology protocol
- Binary protocol


![height:250px](img/beatbars.jpeg)

_midi.md_ 19

---
# Two runtimes

## melrose binary 

- Command Line Interface (cli)
- Read Evaluate Print Loop (REPL)
- HTTP interface
    - Visual Studio Code plugin

## Web

- https://play.melrōse.org

_cli_http.md_ 20

---
# MIDI communication

![height:400px center ](img/melrose-port-daw.png)

_midi_com.md_ 21

---
# Playing

## Audio functions

- bpm(...)
- play(...)
- sync(...)
- loop(...)

## MIDI output

- channel(10, ...)
- device(1, ...)

_playing.md_ 22

---
# Go bits

_go_bits.md_ 23

---
# Program to sound

![height:500px center](img/go_flow.png)

_note_sched.md_ 24

---
![height:500px center](img/timeline_design.png)

_timeline_design.md_ 25

---
# WebAssembly

- play.melrōse.org uses a `wasm` compiled version of melrōse in Go

- `WebMIDI` to send MIDI events

- sharing scripts with other musicians

_assembly.md_ 26

---
# Sound bits ... 

_sound_bits.md_ 27

---
# Sources

- Open source, MIT licensed

- github.com/emicklei/melrose

- https://melrōse.org

- (music) contributions are welcome

_demo.md_ 28

