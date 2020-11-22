## GO .ics creator
You want to create automatic events in .ics format? Then this is the right library for you. You can start ultra fast and have some customization options.

## Install
First you have to install the package:

```console
go get github.com/jojojojonas/ics
```

## How to use?
Here is a small example how to create an ics file over the package:

```go
create, err := ‌ics.Create(ics.Options{"./termin.ics", ics.Dates{"11.11.2020", "23:15:00"}, ics.Dates{"11.11.2020", "23:45:00"}, "Updates Hilfe bei der Website", "Get some updates done!"})
If err != nil {
	fmt.Println("Error: ", err)
}
```

## Help
If you have any questions or comments, please contact us by e-mail at [info@hilfebeiderwebsite.de](mailto:info@hilfebeiderwebsite.de)