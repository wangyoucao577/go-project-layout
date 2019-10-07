# Go Project Layout
This is my understanding of how to structure `Go` project layout. The most important part is how to organize your `Go` code into packages.          

## Motivation
A good project layout will make your source code **easier to understand, easier to test, and eaiser to maintain**.     
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
This topic is kind of best practices. I'd like to discuss some of them and my opinion, then maybe deep into more best practices in the future.     


### Top dir `/cmd`
[**My opinion: Strongly recommended**]     
The `cmd` layout pattern is very useful when you need to have more than one application binary.      

- Each binary gets a subdirectory (e.g., `your_project/cmd/your_app`).   
  - The subdirectory is the `main` package for your application.     
  - Don't put a lot of code in the `main` package. Instead, it should only used to initialises and ties everything together.     
- This patterns also helps you keep your project/package **go gettable**. 
  - It means you can use the `go get` command to fetch (and install) your project, its applications and its libraries (e.g., `go get github.com/your_github_username/your_project/cmd/your_app`). 
- The [official Go tool](https://github.com/golang/tools/tree/master/cmd) is one example of the `cmd` layout patter. A number of other well known projects use the same pattern: [Kubernetes](https://github.com/kubernetes/kubernetes/tree/master/cmd), [Docker](https://github.com/moby/moby/tree/master/cmd), [Prometheus](https://github.com/prometheus/prometheus/tree/master/cmd), [Influxdb](https://github.com/influxdata/influxdb/tree/master/cmd).


### Top dir `/pkg`
[**My opinion: NOT good enough, use it if good for you**]    
- People also recommend to set up a `/pkg` top dir in `Go` project to put your public libraries.    
  - These libraries can be used internally by your application.     
  - They can also be used by external projects. It's an informal contract between you and other external users of your code. Other projects will import these libraries expecting them to work.        
- But I think it's NOT a good enough idea because
  - It brings confusion since `Go` workspaces have a directory with the same name and that directory has a different purpose(it’s used to store object files for the packages the `Go` compiler builds).     
  - The meaning of "public libraries" is not clear enough.    
    - Even many well known projects use this pattern ([Kubernetes](https://github.com/kubernetes/kubernetes/tree/master/pkg), [Docker](https://github.com/moby/moby/tree/master/pkg), [Grafana](https://github.com/grafana/grafana/tree/master/pkg), [Influxdb](https://github.com/influxdata/influxdb/tree/master/pkg), [Etcd](https://github.com/coreos/etcd/tree/master/pkg)), [Docker](https://github.com/moby/moby/tree/master/pkg) and [Etcd](https://github.com/coreos/etcd/tree/master/pkg) comment the meaning of "public libraries" again. Both of them limit it only as **utility packages**, and these **utility packages** possible to be moved out into its own repository in the future:          
      - [Docker](https://github.com/moby/moby/tree/master/pkg): `pkg/` is a collection of utility packages used by the [moby](https://github.com/moby/moby/) project without being specific to its internals.    
      - [Etcd](https://github.com/coreos/etcd/tree/master/pkg): `pkg/` is a collection of utility packages used by etcd without being specific to etcd itself. A package belongs here only if it could possibly be moved out into its own repository in the future.


### Group packages by dependency     
[**My opinion: Strongly recommended**]     
- Domain types
  - Domain Types are types that model your business functionality and objects.    
  - [Ben Johnson - Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) recommend to place your into root package. This package only contains simple data types, so it will be extremely simple.      
  - In my opinion, either place them into root package or place them into a individual subpackage are ok. The most important part here is this package should be extermely simple, and not depend on any other package in application!     
- Each package should only has a single purpose.    
- Packages names describe their purpose.    
- When necessary, use a descriptive parent package and several children implementing the functionality.     
  - like the [encoding package](https://github.com/golang/go/tree/master/src/encoding).    
- `main` package ties together dependencies

### More best practices

- Follow conventions
  - Be similar to the standard library and other popular packages 
  - Don't surprise people
  - Be obvious, not clever
- Use [Go modules](https://github.com/golang/go/wiki/Modules)
- Avoid global scope and `init()`

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
- [golang/tools](https://github.com/golang/tools)
- [moby/moby](https://github.com/moby/moby)
- [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)
- [etcd-io/etcd](https://github.com/etcd-io/etcd)
- [influxdata/influxdb](https://github.com/influxdata/influxdb)
- [prometheus/prometheus](https://github.com/prometheus/prometheus)
