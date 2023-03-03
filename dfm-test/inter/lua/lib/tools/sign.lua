md5 = require("md5")
local signMD5Utils = {}
function signMD5Utils.SignForMD5(param, key)
    return string.upper(signMD5Utils.GetMD5(signMD5Utils.FormatMapToSignStr(param) .. "&key=" .. key))
end

function signMD5Utils.GetMD5(msg)
    return md5.sumhexa(msg)
end

function signMD5Utils.VerifyForMD5(param, sign, key)
    return string.upper(sign) == signMD5Utils.SignForMD5(param, key)
end

function signMD5Utils.BytesToHex(sum)
    local s = ""
    for i = 1, string.len(sum) do
        local cha = tonumber(string.byte(sum, i, i))
        local str = string.format("%02X", cha)
        s = s .. str
    end
    return s
end

function signMD5Utils.FormatMapToSignStr(param)
    local buff = ""
    local ks = {}
    for key, _ in pairs(param) do
        table.insert(ks, key)
    end
    table.sort(ks)
    local buf = ""
    for _, v in pairs(ks) do
        if string.upper(v) ~= "SIGN" then
            local pp = param[v]
            if type(pp) == "table" then
                buf = buf .. v .. "=" .. signMD5Utils.FormatMapToSignStr(pp)
            elseif type(pp) == "string" then
                buf = buf .. v .. "=" .. pp
            end
        end
    end
    return buff
end

-- local tbl = {"1", "2", "4", "5", "3"}

-- function sortGT(a, b)
--     return a > b
-- end

-- print("before sort.")
-- print(tbl[1], tbl[5])
-- table.sort(tbl)
-- print("after less than sort.")
-- print(tbl[1], tbl[5])
-- table.sort(tbl, sortGT)
-- print("after greater than sort.")
-- print(tbl[1], tbl[5])

--[[ print(md5)
print(package.path)
print(signMD5Utils.GetMD5("1111111")) ]]
return signMD5Utils
