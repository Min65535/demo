#!/bin/sh

echo "-----build start-----"
# shellcheck disable=SC2164
cd /Users/min/codes/vw/vw1/src/ios

# /Users/min/Library/Developer/Xcode/DerivedData/
# curl --location 'http://192.168.0.1:3011/api/v1/file/put_file' --form 'f1=@"/Users/min/hello.ipa"'

# shellcheck disable=SC2034
helloArchivePath='/Users/min/pack/vw/ios/archive/hello.xcarchive'
# shellcheck disable=SC2034
helloExportPath='/Users/min/pack/vw/ios/export/hello'
# shellcheck disable=SC2034
exportPlistFilePath='/Users/min/pack/vw/ios/export/plist'
# shellcheck disable=SC2034
hostAddr="192.168.0.1:3011"

mkdir -p ${helloArchivePath}
mkdir -p ${helloExportPath}
mkdir -p ${exportPlistFilePath}


# shellcheck disable=SC2034
helloArchiveFile=${helloArchivePath}"/hello.xcarchive"
# shellcheck disable=SC2034
exportPlistFile=${exportPlistFilePath}"/exportPlist.plist"


xcodebuild archive \
-workspace helloworkspace.xcworkspace \
-scheme helloapp \
-configuration Debug \
-sdk iphoneos16.2 \
-archivePath ${helloArchiveFile} \
-allowProvisioningUpdates


# shellcheck disable=SC2034
appPath=${helloArchiveFile}"/Products/Applications"

# shellcheck disable=SC2164
cd ${appPath}

dst=$(date "+%Y_%m_%d_%H_%M_%S")
echo "${dst}"
# shellcheck disable=SC2128
# shellcheck disable=SC2027
ipaName="hello_"${dst}".ipa"

echo "${ipaName}"
mkdir -p Payload && cp -rf hello.app Payload && zip -r hello.zip Payload && mv hello.zip "${ipaName}"

# shellcheck disable=SC2034
appFile=$(pwd)"/"${ipaName}
# shellcheck disable=SC2089
forms='f1=@"'${appFile}'"'

# shellcheck disable=SC2027
urlRequest="http://"${hostAddr}"/api/v1/file/put_file"

# shellcheck disable=SC2090
curl --location "${urlRequest}" --form "${forms}"

cp -rf ${ipaName} /Volumes/software/ios/fileserver/static

# shellcheck disable=SC2028
echo "\napplication upload successfully \n"

# shellcheck disable=SC2027
urlIpa="http://"${hostAddr}"/static/"${ipaName}
# shellcheck disable=SC2027
INFORM="iphone application of project xxxxx "${RUN_ENV}" build successfully, and the download link is "${urlIpa}
# shellcheck disable=SC2089
INFORM_STR='{"msg_type":"text","content":{"text":"'${INFORM}'"}}'
# shellcheck disable=SC2090
echo "$INFORM_STR"
curl -X POST -H "Content-Type: application/json" \
  -d '{"msg_type":"text","content":{"text":"'"${INFORM}"'"}}' \
  https://xxxxxxxx.cn/xxxx/bot/v2/hook/11111127ee8

# shellcheck disable=SC2028
echo "\napplication inform successfully \n"



#xcodebuild  -exportArchive \
#            -archivePath ${helloArchiveFile} \
#            -exportPath ${helloExportPath} \
#            -exportOptionsPlist ${exportPlistFile}







echo "-----build finish-----"
