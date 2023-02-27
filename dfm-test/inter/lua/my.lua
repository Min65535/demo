local data = string.char(0x01,0x02,0x00,0x31)
print(data[3])
print(#data)

local c1,c2,s = string.unpack(fmt, s, pos)