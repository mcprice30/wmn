package data

// SensorDataFromBytes will attempt to unmarshall the given bytes to a
// SensorData instance. If unsuccessful, it will return nil.
func SensorDataFromBytes (in []byte) SensorData {
	switch in[0] {
	case GasDataType:
		return GasDataFromBytes(in)
	case HeartRateDataType:
		return HeartRateDataFromBytes(in)
	case LocationDataType:
		return LocationDataFromBytes(in)
	case OxygenDataType:
		return OxygenDataFromBytes(in)
	default:
		return nil
	}
}
