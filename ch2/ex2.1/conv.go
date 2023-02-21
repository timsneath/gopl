package tempconv

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CtoK(c Celsius) Kelvin     { return Kelvin(c - 273.15) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FtoK(f Fahrenheit) Kelvin  { return Kelvin(Celsius(f)) }
