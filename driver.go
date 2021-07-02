package slf4go_multi

import (
	slog "github.com/go-eden/slf4go"
)

type MultiDriver struct {
	drivers []slog.Driver
}

func NewMultiDriver(drivers ...slog.Driver) *MultiDriver {
	return &MultiDriver{
		drivers: drivers,
	}
}

func (d *MultiDriver) Name() string {
	return "slf4go-multi"
}

func (d *MultiDriver) Print(l *slog.Log) {
	for _, driver := range d.drivers {
		driver.Print(l)
	}
}

func (d *MultiDriver) GetLevel(logger string) (sl slog.Level) {
	l := slog.FatalLevel
	for _, driver := range d.drivers {
		driverLevel := driver.GetLevel(logger)
		if driverLevel < l {
			l = driverLevel
		}
	}
	return l
}
