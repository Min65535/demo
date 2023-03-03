--local handler = {appId = "", secretKey = "", url = "", reloadTimeSec = 0}
local handler = {}
function handler.new(appId, secretKey, url, reloadTimeSec)
    local o = {appId = "", secretKey = "", url = "", reloadTimeSec = 0}
    o.appId = appId
    o.secretKey = secretKey
    o.url = url
    o.reloadTimeSec = reloadTimeSec
    return o
end

hd = handler.new("111", "222", "www.baidu.com", 10)

print(hd.appId)


handlerInterface = {
    loadCfg = function (hd)

    end
}
