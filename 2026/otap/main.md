---
marp: true
theme: default
paginate: true 
color: "#EEE"
backgroundColor: "#222"
footer: <h3>BBTG</h3>
header: <h4>Platform Engineering, September 2026, State of OTAP</h4>

---
# State of OTAP
September 2026, Kickoff Platform Engineering

### Ernest Micklei

- Software Artist
- Platform Engineer
- Cloud Architect
- Cloud Architect

---
# Who am i

![height:80px center](./img/emicklei_hackers_logo.png)

- chief platform engineering (day) 
- open-source developer (night)
    - **go-restful** - REST-ful framework (part of k8s)
    - **melrōse** - Music programming
    - **proto** - ProtocolBuffers parser
    - **dot** - Graphviz DOT writer
    - **pgtalk** - Postgres access code generator

---
# Why OTAP ?


---
# Why OTAP?

![height:400px center](./img/rabo_otap.avif)

---
# Typical Multi-tier Web Application

![width:70% center](./img/OTAP-Slate-1.png)

---
# Connecting Deployed Components

![width:70% center](./img/OTAP-Slate-2.png)

---
# Naive Approach

![width:70% center](./img/OTAP-Slate-3.png)

---
# Stable, production-like backends

![width:70% center](./img/OTAP-Slate-4.png)

---
# Better isolation for development

![width:70% center](./img/OTAP-Slate-5.png)

---
# Mock their APIs

![width:70% center](./img/OTAP-Slate-6.png)

---
# What if you have Uber.com scale ?

![width:70% center](./img/OTAP-Slate-7.png)

---
# SLATE (Uber)

Short-Lived Application Testing Environments

- **on-demand, ephemeral**
- **production fidelity**
- **tenancy-based routing**
- **data isolation** 
    - test accounts, tenancy-aware Kafka, isolated resources keep test traffic from touching real users

---
# Typical Multi-tier Web Application

![width:70% center](./img/OTAP-Slate-8.png)

---
# thank you


Slide deck (Creative Commons)

**github.com/emicklei/talks**

