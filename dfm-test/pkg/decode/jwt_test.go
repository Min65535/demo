package decode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testJwtPvtKey = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDv5SaTtEgIt6bf
/xdotCGi9uY2FCobQlKNWKOp3wVcx7Qvep019GCFsXwz2mVy8yDNIwUkCpSQzDe6
2y7/l7zaaZ5Kwnu5RAUbZLIokz7CYgYdEDl8SG+sRK0htzd6XXBMl57IVRhJjEjF
AetY21VMZbarzyGLWmUoSBgIrKT76zhkhD+0LIXaCcPzKpmYe0PrKg3CfPk+167i
9z+lll3JececeQANBVa2fFYX5hOukukycBi0bFIytxXWfwMm3SX35Jl5MeN6DXPD
5Zqf0nMp0O9zciWHsBIq4H1MNZE8TQEnOZYJOxyTmDb4u1CD9/sywEiH3oOqCYNN
DrlR6o6LAgMBAAECggEAGtPUgZYarCblmY0scAAc9t2HlyqgHtZG5xmvi9KuBdcr
dlfO9vaySKE1k5Lr33QDCg3NiF9Kh/rejJ3wXpL4grBnDFM9hNVvgMW9Cr6UuSY/
Ksbittxe6LPxbKDRqnSCl34pOpwRkEAI70csq44ztsx1JjeOt8fPIjcVwPqVpLpM
7wHsi5yu8IuDTr006JWopKnHR7Jg5hJGAboybcnVaCvK8PCWJuY42s9i8uYwHsnX
sW4Au0XklGdBTG3a6YaJhGWiRPX4KeTMfzzxhZy5pwdPuluQEfuO3/eRoT6Wqo9a
6ZYMR5GjeMv3p12X9k2WtEVgTp8krLCPxoe2DT3q6QKBgQD8tCRz5+OVGtoi+iVR
iy4N28UmKJS2Phf9conzS550qFCvR/QmwNe39EhVWqnw0cuI9NhzoS2aLgX3Dmlm
hetiQ4118Orbin3uoqnxnUfbGluUL5lWCs5H1sPpZNrwI6f0VXHmWkjXLfTT+wzv
RiJYxadbBCmPHGLsN2JShJEPZwKBgQDzBjyIXWjzo1TMl1bR22bySQNz5bd6oVmF
W5oJq8yyHXN03kxdPsCLDp73VUd815xKb1fJoxBWWC4BIyQPkbMHeG49L1aJE2Zb
8kWokkTCnQemlcgrQ5w3HcfITgowcVSgFdMT+uLvfVcVnCYSlVljoJ2s7fKGq5PI
H2/QJuAlPQKBgQCIPIY6hpXHWQapPLrJz4MwrX8IJ3ClH6zHUuzUYbw3oFZ/aQdT
rTufTO+CNLLRxgk4+OeIzyhKqu1EWFyyjRhtjYXCQ/QaD/v8n22HeQe4M+mTZmYA
YSr8x+gu99ShClgN/dK+IYaLm5cWgY5joSf7O/QRuZi+MsuSFfnICvg5mQKBgQDu
dpkhKZ9ZKlSEKKAdbc83QKOtrqP6JabE75xXjwddYv6ul46RFIk0KdWcVka+XobS
lyyqA2J9hyslHIk+6bWlR+vMB84+1RVkdXcasOPdnt4p/OEcasw3XPZzOvhOjnX+
W0vyWAUkAHZaEw1cwMWkt67gR9/peySYgxhhRDQ+9QKBgQDKxgvqjr/83pX132hj
B3tJI6xvshdeijtugUwCkSNdSNUsojqXTge4mL+XbhNv/KLxGYFFcbNM5Bs/74CQ
SfRxv2JbWw6i6uY2Ags3+JOwW9D2lojnsauTA0CAHDjZcoYjxmYJsxu9Ka2yiPbk
7D7gKBUI51iFTkzUfQ8mWG761Q==
-----END PRIVATE KEY-----`
	testJwtPvtKey2 = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCsDKAb+i4qJWut
iLVpdQen2Rjmd4IS4FINUUMzJKaFrAkp0LP+gWwiNk51g4+7dqVYapeo8+U0w4HM
0eQ6nyOkoiOdUDCOaok3GnfpEZ++45iHCR7dNw42yzZQB24UZ/VNgrjqAjsHmtUc
zHQzW+RfYaA7K/pBK8IbqLmAw+YJsP14g94H8mCAlmxKeGodEiStu5Y5Y2vQDtBz
hu4KlsWtsAw6NaXv9hdvc3ifEAmRmChO9gDXZ7LGjF2+WSa+m1R5f2L43FMJ41At
IeXlCI4YXZC4TAG+KcvKiDq646e73GmsAfbewyleDwKc0Z/uTyWpKGdppzpm9fZp
32AJCigDAgMBAAECggEAcSewmhX2DNQluLFkFyLYUSVwwgsxqnaMoKOkndtqBhp6
wFy6NRYr52huO822V6YV40vfyUf3pLc9BFe7Um7eA2Z8I/fcp54b9vjaipiDxwlr
hMyT1wxhtdn7M7FfTd59sGfcdFPZPQwxQ8qYVjsGvVC5EJHA+uHrvUNAFhdBkXcY
nffnKxQeoACzHyFGi8vxemAIMl3Wjmq2eDHz8Qtn77muOysnSAY8e+wORxvLl0hX
GG9MbdmX2Oa7mHjk+rDT3LUpGxB7Czld2zOOmTdYepn01WPT0sTZExywASuNWOpP
/xsi1Sh1pYpxgnaVRtpVTRB9GrXsPn5DnrYxoyuKKQKBgQDjSB/79rLFfReozhLb
hATwmHtGiqHhwMvoL00AIIiQI+3pw/ty2xN5gvpyvwOKzNdrfcaVATTCGr+Fgy7R
YdLIXYjce6qjWsS9ABjmVG2z12gokVHq+F80YeUr6pAe91JzKcKaPk8lkyUC/LxK
tYEu85ywsSvdPmK51hBHFekhJQKBgQDByeYwjgldL2oBUhI4HyzN/DbVsZtNg2Vl
ssVxGcgGZb0ueM8C0W4RVygQTbX92z7KS25D75VQXI4a85mskjJBxrRqoe/TPPz2
0kxwlyHpxfpJTTpPFc0WPGie6z7kaehl0GI2ugISPKZ/n3cQo35rEyNM720D0VGE
nWVjbo1ABwKBgQCI1UNsTnoSq90yo952imTu9N5C+fO8Fnasss2I5g1Ruk/iDTu1
Sm/PGCvwKU/YoLqQ3IhR7Qf2VGQ53WCyblKYjrd3Bn0VG/CWWRikku/49hafVd4b
uKyYvNdcOTvLaNsaummOszSzSoNd6Qrzb5L20XPkbMYbzRNjDp1+LpLMgQKBgGKx
kbNO00QLHsC3bKH6dpYdikvA3WhXr+9gYZ/dUnq3m+asDjnQMW9RZQ0QlsxHua3L
RsgAn5nC2Xiucahq+H95VG8uM/bwC6Ekr1t0DQiDFJn6Y+TdJIrbjyjIhEpOCda4
jxxyFRX2n5FFxJzLt1mO8J6BHZMhKpQQmQ9PwU5HAoGALqZgnTEZYyyKSFrdVpse
XpBcyEOMK5hwA9xjVsDXt+5tkRJQ9LoF+RIvCF2r4LSg5rNTTxuljpsNuffvRCFQ
04r2J24P6Vxgem2DG0PEN16UQJNgHGHJ9rfRMpsA4MxQOZGjA/HJIkRAJeWbSgFo
XBmEXsQKVN08B91tNOsoYBg=
-----END PRIVATE KEY-----`
)

func TestJwtToken(t *testing.T) {
	j, err := NewJwtToken(testJwtPvtKey)
	assert.NoError(t, err)
	token, err := j.Encode(1, 2)
	if !assert.NoError(t, err) {
		return
	}

	uid, rid, err := j.Decode(token)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, uint(1), uid)
	assert.Equal(t, uint(2), rid)
}

func TestJwts_Get(t *testing.T) {
	j := make(JwtSets)
	jwtt, err := NewJwtToken(testJwtPvtKey)
	assert.NoError(t, err)
	jwtt2, err := NewJwtToken(testJwtPvtKey2)
	assert.NoError(t, err)
	j["test-env"] = jwtt
	j["test2"] = jwtt2

	token, err := jwtt.Encode(1, 2)
	if !assert.NoError(t, err) {
		return
	}
	uid, rid, err := j.Get("test-env").Decode(token)
	if assert.NoError(t, err) {
		assert.Equal(t, uint(1), uid)
		assert.Equal(t, uint(2), rid)
	}

	_, _, err = j.Get("test2").Decode(token)
	assert.Error(t, err)

	_, _, err = j.Get("nothing").Decode(token)
	assert.EqualError(t, err, "不存在的校验")

	token, err = j.Get("test2").Encode(3, 1)
	if !assert.NoError(t, err) {
		return
	}
	uid, rid, err = jwtt2.Decode(token)
	if assert.NoError(t, err) {
		assert.Equal(t, uint(3), uid)
		assert.Equal(t, uint(1), rid)
	}
}
