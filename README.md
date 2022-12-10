# Go Data Examples
A walk through of setting up your environment and doing data science projects in [Go](https://go.dev/)

For community input and resources please checkout the [#data-science](https://gophers.slack.com/?redir=%2Fmessages%2Fgeneral%2F) channel in Gophers Slack.

[![Actions Status](https://github.com/soypete/{}/workflows/build/badge.svg)](https://github.com/soypete/{}/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/soypete/{}/branch/master/graph/badge.svg)](https://codecov.io/gh/soypete/{})

## Intro

Go is a great tool for creating reliable, repeatabile, and stable scripts and applications. Sometimes we want to use go to tackle data problems. This doc is a simple walk through to get started using Go for data research projects (ML, analytics, data science, and stats). It is meant for people who are beginner to intermediate level programmers. It will walk through any installations and setup needed to complete this walk through. 

## Set up Go env

1. Download [Go](https://go.dev/dl/)
	- For more detailed instructions visit this [Tutorial]()

2. Create/clone the repo

```bash
git clone https://github.com/Soypete/go-data-examples.git
```

3. test examples

```bash
go run example/example.go
```

4. Add go plug-in to your editor of choice. 
- [VSCode](https://code.visualstudio.com/docs/languages/go)
- [Vim](https://github.com/fatih/vim-go)
- [Emacs](https://github.com/dominikh/go-mode.el)
- [Sublime](https://github.com/DisposaBoy/GoSublime)

Turn on features like compile , `go fmt`, and `go imports` so that they happen everytime you save your program.

## Getting data

### Scraping
Web or API scraping is a common way for data scientists to get public data sets. This method will require a lot of data cleaning before any insights can be made. 

#### Examples: 

[Meetup Web scraper](https://github.com/Soypete/event-web-crawler)

#### Tutorials and Videos

[scraper build live stream](https://youtu.be/g2i3dhEUQLQ)

### Test Data sets

Many public data sets are open for download from sites like Kaggle. These datasets are great for evaluation models, algorithms, and packages against know solutions. Often solutions in common languages like R or Python are published. This is a good indicator on if your model is accurate enough for production. 

## Exploring the Data

The data set we will be using today is already downloaded and in `data/`. 

### Load data into your program

#### Using CSV

In your go program add a file reader to import your csv file. In go, i/o reads and writes can be made by any library that implements the [io.Reader](https://pkg.go.dev/io#Reader) or [io.Writer interfaces](https://pkg.go.dev/io#Writer). For these files the `csv.Reader` type is best suited for our needs. 

Open a `main.go`file in your repo. 

*note*: _if you are using a your repo and not the clone repo from above you will need to run the `go mod init` [command](https://go.dev/doc/tutorial/create-module)_

Add the following code to your `func main(){}`.

```go
	f, err := os.Open("data/player/player_stats.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
```

The `reader.ReadAll()` will process all the csv rows in the file and store them in the memory as a `[][]string` object. You can access the attributes of the csv file my calling the indexes of the attribute, starting with index `0` for both directs. 

examples:

```go
fmt.Println(data[2][0])
```

The above code would print the 3 rows first column of the csv file.

#### Using Dataframes

*NOTE*: _A dataFrame is a 2 dimensional data structure, like a 2 dimensional array, or a table with rows and columns. It was originated in the R programming language and is commonly buy pandas pracitioners to manipulate in-memory datasets_

We can use a dataframe to import and analyize our player data.The package for using dataframe `go get github.com/go-gota/gota` so it can be added to your `go.mod` file. After you have gotten the `gota` module, add the `Dataframe` to your `func main(){}`.

```go
	f, err := os.Open("data/player/player_stats.csv")
	if err != nil {
		log.Fatalf("unable to access csv: %s", err)
		return
	}
	df := dataframe.ReadCSV(f)
```

Now you can run methods on your player stats `Dataframe`.

Here are some examples of methods you can run.

```go
	fmt.Print(df.Describe())
```

```go
	fmt.Print(df.Select([]string{"player", "age", "position"})
```

```go
	df.Arrange(dataframe.RevSort("goals_per90"))
```

Exercise: Use the [Dataframe godocs](https://pkg.go.dev/github.com/go-gota/gota/dataframe)  to explore and implement other data manipulations.

To see a working example run
```bash
go run examples-dataframes/dataframes.go
```

## Machine Learning

There are several packages for running machine learning in go. Below are some of the packages.

Here is an example of an `Decision Tree` for detecting names and twitter handles. [Example](https://github.com/Soypete/golearn/blob/master/examples/fifaText/main.go)
To learn learn more about this example watch this [video](https://www.youtube.com/watch?v=3mivck8WArg&list=PL8Q5PSrFkjswyF90RdoxzVKMCKnuQGlFc&index=5)

## Additional Resources
- [Gopher Notes](https://github.com/gopherdata/gophernotes) is the Go Kernel for a jupyter notebooks
- [Machine Learning with Go](https://www.packtpub.com/product/machine-learning-with-go/9781785882104)
- [KNN with Go](https://youtu.be/DulSFt37284)
- [Gopher Data](http://gopherdata.io/)
