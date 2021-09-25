# 7GUIs in Golang-fyne

This is an implementation of 7GUIs in Golang with fyne as the UI toolkit There is one folder for each task with one or more go files. Each folder has a main go file so you can see the resulting GUI application by running the corresponding file.

### Launch tasks

```bash
go run task1/main.go
go run task2/main.go
go run task3/main.go
go run task4/main.go
# ...
```

### BigSure-Apple Silicon issue
Fyne is basically broken, not due to Fyne itself but one of its dependency.
Chasing this up is a rabbit hole, so it will take a while.
