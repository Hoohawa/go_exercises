package mtrconv

func MToMI(m Meter) Mile { return Mile(m * 0.000621371) }
func MToIN(m Meter) Inch { return Inch(m * 39.3701) }

func MIToM(mi Mile) Meter { return Meter(mi * 1609.34) }
func MIToIN(mi Mile) Inch { return MToIN(MIToM(mi)) }

func INToM(in Inch) Meter { return Meter(in * 0.0254) }
func INToMI(in Inch) Mile { return MToMI(INToM(in)) }
