package persian

// Classic Jalali↔Gregorian conversion used widely in Persian open-source code.
// Verified in tests against known reference dates (e.g. Nowruz 1403 = 2024-03-20).

func gregorianToJalali(gy, gm, gd int) (jy, jm, jd int) {
	gDaysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	gy2 := gy - 1600
	gm2 := gm - 1
	gd2 := gd - 1

	gDayNo := 365*gy2 + (gy2+3)/4 - (gy2+99)/100 + (gy2+399)/400
	for i := 0; i < gm2; i++ {
		gDayNo += gDaysInMonth[i]
	}
	if gm2 > 1 && ((gy%4 == 0 && gy%100 != 0) || (gy%400 == 0)) {
		gDayNo++
	}
	gDayNo += gd2

	jDayNo := gDayNo - 79
	jNp := jDayNo / 12053
	jDayNo %= 12053

	jy = 979 + 33*jNp + 4*(jDayNo/1461)
	jDayNo %= 1461
	if jDayNo >= 366 {
		jy += (jDayNo - 1) / 365
		jDayNo = (jDayNo - 1) % 365
	}

	jDaysInMonth := [12]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	var i int
	for i = 0; i < 11 && jDayNo >= jDaysInMonth[i]; i++ {
		jDayNo -= jDaysInMonth[i]
	}
	jm = i + 1
	jd = jDayNo + 1
	return jy, jm, jd
}

func jalaliToGregorian(jy, jm, jd int) (gy, gm, gd int) {
	jDaysInMonth := [12]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

	jy2 := jy - 979
	jm2 := jm - 1
	jd2 := jd - 1

	jDayNo := 365*jy2 + (jy2/33)*8 + ((jy2%33)+3)/4
	for i := 0; i < jm2; i++ {
		jDayNo += jDaysInMonth[i]
	}
	jDayNo += jd2

	gDayNo := jDayNo + 79

	gy = 1600 + 400*(gDayNo/146097)
	gDayNo = gDayNo % 146097

	leap := 1
	if gDayNo >= 36525 {
		gDayNo--
		gy += 100 * (gDayNo / 36524)
		gDayNo = gDayNo % 36524
		if gDayNo >= 365 {
			gDayNo++
		} else {
			leap = 0
		}
	}

	gy += 4 * (gDayNo / 1461)
	gDayNo %= 1461

	if gDayNo >= 366 {
		leap = 0
		gDayNo--
		gy += gDayNo / 365
		gDayNo = gDayNo % 365
	}

	gDaysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	gd = gDayNo + 1
	for gm = 1; gm <= 12; gm++ {
		v := gDaysInMonth[gm-1]
		if gm == 2 {
			v += leap
		}
		if gd <= v {
			break
		}
		gd -= v
	}
	return gy, gm, gd
}

func isJalaliLeapClassic(jy int) bool {
	// 33-year cycle leap residues used by the classic algorithm.
	switch jy % 33 {
	case 1, 5, 9, 13, 17, 22, 26, 30:
		return true
	default:
		return false
	}
}
