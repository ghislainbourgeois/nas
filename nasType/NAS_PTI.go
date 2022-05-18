// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

package nasType

// PTI 9.6
// PTI Row, sBit, len = [0, 0], 8 , 8
type PTI struct {
	Octet uint8
}

func NewPTI() (pTI *PTI) {
	pTI = &PTI{}
	return pTI
}

// PTI 9.6
// PTI Row, sBit, len = [0, 0], 8 , 8
func (a *PTI) GetPTI() (pTI uint8) {
	return a.Octet
}

// PTI 9.6
// PTI Row, sBit, len = [0, 0], 8 , 8
func (a *PTI) SetPTI(pTI uint8) {
	a.Octet = pTI
}