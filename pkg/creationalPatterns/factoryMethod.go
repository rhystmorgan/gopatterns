// factoryMethod.go
package creationalPatterns

import "fmt"

// Factory Method Pattern - UI Component Library Example
//
// Imagine you're building a component library that needs to support
// different themes (Material, Bootstrap, etc). The factory method
// lets each theme decide how to create its components.

// =============================================================================
// Product Interface - What all buttons can do
// =============================================================================

type Button interface {
	Render() string
	OnClick(handler func())
}

// =============================================================================
// Concrete Products - Theme-specific buttons
// =============================================================================

// MaterialButton - Google Material Design style
type MaterialButton struct {
	label     string
	onClick   func()
	ripple    bool
	elevation int
}

func (b *MaterialButton) Render() string {
	return fmt.Sprintf(
		`<button class="mdc-button mdc-button--raised" style="elevation: %d">
                        <span class="mdc-button__ripple"></span>
                        <span class="mdc-button__label">%s</span>
                </button>`, b.elevation, b.label)
}

func (b *MaterialButton) OnClick(handler func()) {
	b.onClick = handler
}

// BootstrapButton - Bootstrap style
type BootstrapButton struct {
	label   string
	onClick func()
	variant string // "primary", "secondary", "danger"
	size    string // "sm", "md", "lg"
}

func (b *BootstrapButton) Render() string {
	return fmt.Sprintf(
		`<button class="btn btn-%s btn-%s">%s</button>`,
		b.variant, b.size, b.label)
}

func (b *BootstrapButton) OnClick(handler func()) {
	b.onClick = handler
}

// TailwindButton - Tailwind CSS style
type TailwindButton struct {
	label   string
	onClick func()
	classes string
}

func (b *TailwindButton) Render() string {
	return fmt.Sprintf(
		`<button class="%s">%s</button>`,
		b.classes, b.label)
}

func (b *TailwindButton) OnClick(handler func()) {
	b.onClick = handler
}

// =============================================================================
// Creator Interface - The Factory
// =============================================================================

type ComponentFactory interface {
	CreateButton(label string) Button
	// Could expand: CreateInput(), CreateModal(), CreateCard(), etc.
}

// =============================================================================
// Concrete Creators - Theme-specific factories
// =============================================================================

type MaterialFactory struct{}

func (f *MaterialFactory) CreateButton(label string) Button {
	return &MaterialButton{
		label:     label,
		ripple:    true,
		elevation: 2,
	}
}

type BootstrapFactory struct {
	defaultVariant string
	defaultSize    string
}

func NewBootstrapFactory() *BootstrapFactory {
	return &BootstrapFactory{
		defaultVariant: "primary",
		defaultSize:    "md",
	}
}

func (f *BootstrapFactory) CreateButton(label string) Button {
	return &BootstrapButton{
		label:   label,
		variant: f.defaultVariant,
		size:    f.defaultSize,
	}
}

type TailwindFactory struct{}

func (f *TailwindFactory) CreateButton(label string) Button {
	return &TailwindButton{
		label:   label,
		classes: "px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600",
	}
}

// =============================================================================
// Client Code - App doesn't know which theme it's using
// =============================================================================

// RenderLoginForm works with ANY theme - doesn't know or care which one
func RenderLoginForm(factory ComponentFactory) {
	submitBtn := factory.CreateButton("Sign In")
	cancelBtn := factory.CreateButton("Cancel")

	fmt.Println("Login Form:")
	fmt.Println(submitBtn.Render())
	fmt.Println(cancelBtn.Render())
	fmt.Println()
}

// GetFactoryFromConfig simulates reading theme from app config
func GetFactoryFromConfig(theme string) ComponentFactory {
	switch theme {
	case "material":
		return &MaterialFactory{}
	case "bootstrap":
		return NewBootstrapFactory()
	case "tailwind":
		return &TailwindFactory{}
	default:
		return &BootstrapFactory{} // fallback
	}
}

func ExampleFactoryMethod() {
	// App config determines the theme - only place that knows the concrete type
	theme := "material" // Could come from config file, env var, etc.
	factory := GetFactoryFromConfig(theme)

	// All components use the factory - completely decoupled from theme
	RenderLoginForm(factory)

	// Switch themes easily
	fmt.Println("--- Switching to Tailwind ---")
	factory = GetFactoryFromConfig("tailwind")
	RenderLoginForm(factory)
}

//  Output:
//  Login Form:
//  <button class="mdc-button mdc-button--raised" style="elevation: 2">
//      <span class="mdc-button__ripple"></span>
//      <span class="mdc-button__label">Sign In</span>
//  </button>
//  <button class="mdc-button mdc-button--raised" style="elevation: 2">
//      <span class="mdc-button__ripple"></span>
//      <span class="mdc-button__label">Cancel</span>
//  </button>
//
//  --- Switching to Tailwind ---
//  Login Form:
//  <button class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">Sign
//  In</button>
//  <button class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600">Cancel</button>
