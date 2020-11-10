# GO .ics creator
You want to create automatic events in .ics format? Then this is the right library for you. You can start ultra fast and have some customization options.


First you have to install the package:

`import "github.com/jojojojonas/ics-creator"`


Here is a small example how to create an ics file over the package:

`create, err := ‌ics.Create(ics.Options{"./termin.ics", ics.Dates{"11.11.2020", "23:15:00"}, ics.Dates{"11.11.2020", "23:45:00"}, "Updates Hilfe bei der Website", "Get some updates done!"})`