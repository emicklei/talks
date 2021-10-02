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
---
# Who is Ernest Micklei

![width:100px](img/emicklei_hackers_logo.png)

- using Go since 2011
- co-organizer Golang Meetup Amsterdam
- 100+ Github repos: *go-restful*, *proto*, *dot*, *forest*, *zazkia*, *pgtalk*, ...
- Google Developer Expert (Go & GCP)


<style>
pre,code {
  background: #eee;
  color: black;
}
</style>
<script src="slides/play.js"></script>
---
# Language Bits
---
# Note

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

    sequence('c e g = (c5 e5 g5)')

- note
- rest `=`
- group `(...)`
---
# Examples: sequence

    f1 = sequence('c e g c')
---
# Play bits
---
# Go bits
---
# Demo
