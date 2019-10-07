# Go Project Layout
This is my understanding of how to structure `Go` project layout. The most important part is how to organize your `Go` code into packages.          

## Motivation
A good project layout will make your source code easier to understand, easier to test, and eaiser to maintain.     
[Kat Zien](https://github.com/katzien/) has a very good summary([Slides](https://github.com/katzien/talks/tree/master/how-do-you-structure-your-apps/gopherconuk-2018-08-03)) for this, so I'd like to refer it directly as below:        

### Questions, decisions 

- Should I put everything in the main package?
- Should I start with one package and extract other packages over time? 
- How do I decide if something should be in its own package?
- Should I just use a framework?
- What's the programming paradigm for Go?
- Microservices or monolith?
- How much should be shared between packages?
       
### Why should we care?
**Because if Go is going to be a language that companies invest in for the long term, the maintenance of Go programs, the ease of which they can change, 
will be a key factor in their decision.** - Dave Cheney, Golang UK 2016 keynote

### Good structure goals
- Consistent.
- Easy to understand, navigate and reason about. ("makes sense") Easy to change, loosely-coupled.
- Easy to test.
- "As simple as possible, but no simpler."
- Design reflects exactly how the software works.
- Structure reflects the design exactly.     

## Note

## References
- [GopherCon UK 2018: Kat Zien - How do you structure your Go apps?](https://www.youtube.com/watch?v=VQym87o91f8&t=3s)     
  - This's an amazing talk for the `Go` project layout I think.    
  - [Slides](https://github.com/katzien/talks/tree/master/how-do-you-structure-your-apps/gopherconuk-2018-08-03)    
  - Code: [katzien/go-structure-examples](https://github.com/katzien/go-structure-examples)    
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)    
  - [No, it's NOT a golang standard!!!](https://github.com/golang-standards/project-layout/issues/38) See more in [Delete this repo #36](https://github.com/golang-standards/project-layout/issues/36) and [This repository has multiple serious issues #41](https://github.com/golang-standards/project-layout/issues/41).     
  - It's just an organization name "golang-standards" used by some one to manage these repos. It's harmful if you really think it's a standard and try to always follow it.     
  - On the other hand, I think this repo still have some benifits: collects many resources for the "golang-project-layout" question, gives suggestion for how to do it, etc.     
  - Try to read it as a post, think about what's valueable for you.             
  - [More resources metioned in #36](https://github.com/golang-standards/project-layout/issues/36#issue-479438391)
- [Ben Johnson - Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
- [Ben Johnson - Building WTF Dial](https://medium.com/wtf-dial/wtf-dial-domain-model-9655cd523182)
- [Kyle C. Quest - Go Project Layout](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)
- [GopherCon Russia 2018: Ashley McNamara + Brian Ketelsen - Go best practices](https://www.youtube.com/watch?v=MzTcsI6tn-0)
  - [slides](https://www.brianketelsen.com/slides/gcru18-best/)
- [Golang UK Conference 2017 | Mat Ryer - Writing Beautiful Packages in Go](https://www.youtube.com/watch?v=cAWlv2SeQus)
- [CodeBear801 - Golang Package layout](https://github.com/CodeBear801/tech_summary/blob/master/tech-summary/language/go/go-package-layout.md)
- [golang/go](https://github.com/golang/go)
- [moby/moby](https://github.com/moby/moby)
- [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)
- [etcd-io/etcd](https://github.com/etcd-io/etcd)

