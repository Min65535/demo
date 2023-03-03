sc = require("socket.http")

bo = _G.os.execute("curl ip.sb")
print("bo:", bo)

-- vv,ret=sc.request("https://ipv4.ip.sb/addrinfo")
vv, ret = sc.request("https://ip.sb")
--vv, ret = sc.request("https://ipv4.ip.sb/addrinfo")
print("vv:", vv)
print("ret:", ret)
sc = require("socket.http")
fcc = function()
    local httpc = "http://ip.myhostadmin.net/"
    vvs, rets = sc.request(httpc)
    if rets ~= 200 then
        return ""
    end
    return string.match(vvs, "(%d+.%d+.%d+.%d+)")
end

print("ip_: ", fcc())

package.path = package.path..';share/?.lua'

print(package.path)
