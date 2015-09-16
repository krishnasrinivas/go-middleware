package middleware

import "net/http"

type MWserve func(http.ResponseWriter, *http.Request, http.HandlerFunc)

type MW struct {
	mw   []MWserve
	next []http.HandlerFunc
	h    http.Handler
}

func New(h http.Handler, mwarr ...MWserve) *MW {
	len := len(mwarr)
	if len == 0 {
		return nil
	}
	m := new(MW)
	m.mw = make([]MWserve, len)
	m.next = make([]http.HandlerFunc, len)
	m.h = h
	for i := 0; i < len; i++ {
		m.mw[i] = mwarr[i]
		m.next[i] = m.NextFunction(i)
	}
	m.next[len-1] = func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
	return m
}

func (m *MW) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mw[0](w, r, m.next[0])
}

func (m *MW) NextFunction(i int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.mw[i+1](w, r, m.next[i+1])
	}
}


