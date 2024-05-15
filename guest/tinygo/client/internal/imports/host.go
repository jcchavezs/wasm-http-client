package imports

func HTTPDo(
	methodPtr uintptr, methodLen uint32,
	urlPtr uintptr, urlLen uint32,
	headersPtr uintptr, headersLen uint32,
	bodyPtr uintptr, bodyLen uint32,
) uint32 {
	return httpDo(methodPtr, methodLen, urlPtr, urlLen, headersPtr, headersLen, bodyPtr, bodyLen)
}
