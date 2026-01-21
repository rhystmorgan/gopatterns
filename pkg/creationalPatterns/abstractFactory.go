package creationalPatterns

// abstract factory interface
type DeviceFactory interface {
	CreateSmartPhone() SmartPhone
	CreateTablet() Tablet
}

// concrete factory AppleFactory
type AppleFactory struct{}

func (a *AppleFactory) CreateTablet() Tablet {
	return &AppleTablet{}
}

func (a *AppleFactory) CreateSmartPhone() SmartPhone {
	return &AppleSmartPhone{}
}

type SamsungFactory struct{}

func (s *SamsungFactory) CreateSmartPhone() SmartPhone {
	return &SamsungSmartPhone{}
}

func (s *SamsungFactory) CreateTablet() Tablet {
	return &SamsungTablet{}
}

type SmartPhone interface {
	SwitchOn() bool
	Ring()
}

type Tablet interface {
	SwitchOn() bool
}

func (s *SamsungSmartPhone) Ring() {
	fmt.Println("SmartPhone turning on ...")
}

func (s *SamsungSmartPhone) SwitchOn() bool {
	fmt.Println("")

	return true
}

func (a *AppleSmartPhone) SwithcOn() bool {
	fmt.Println("SmartPhone turning on ...")
}

func (a *AppleTablet) Ring() {
	fmt.Println("")

	return true
}

func ClientCode(factory DeviceFactory) {
	smartphone := factory.CreateSmartPhone()
	smartphone.Ring()

	tablet := factory.CreateTablet()
	tablet.SwitchOn()
}

// ClientCode(NewSamsungFactory())
