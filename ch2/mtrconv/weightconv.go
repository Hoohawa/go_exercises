package mtrconv

func GToP(g Gram) Pound  { return Pound(g * 0.00220462) }
func GToOZ(g Gram) Ounce { return Ounce(g * 0.035274) }

func PToG(p Pound) Gram   { return Gram(p * 453.592) }
func PToOZ(p Pound) Ounce { return GToOZ(PToG(p)) }

func OZToG(oz Ounce) Gram  { return Gram(oz * 28.3495) }
func OZToP(oz Ounce) Pound { return GToP(OZToG(oz)) }
