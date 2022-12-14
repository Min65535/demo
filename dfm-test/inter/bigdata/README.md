## bigdatacron
```text
使用方法:
在logsvr的/data/my/logs目录
nohup ./bigdatacron -from="/data/my/logs/" -to="/data/bigdata-storage/" -spc="0 15 21 * * ?" >bigdatacron_nohup.out 2>&1 &
```

#### start
```shell
#!/usr/bin/env bash

set -e

nohup ./bigdatacron -from="/data/my/logs/" -to="/data/bigdata-storage/" -spc="0 15 21 * * ?" >bigdatacron_nohup.out 2>&1 &

exit 0
```

#### restart
```shell
#!/usr/bin/env bash

set -e

# shellcheck disable=SC2009
outerPid=$(ps -ef | grep 'bigdatacron' | grep -v 'grep' | awk '{print $2}')
if [ -n "${outerPid}""" ]; then kill -9 $outerPid; fi

nohup ./bigdatacron -from="/data/my/logs/" -to="/data/bigdata-storage/" -spc="0 15 21 * * ?" >bigdatacron_nohup.out 2>&1 &

exit 0
```

#### stop
```shell
#!/usr/bin/env bash

set -e

# shellcheck disable=SC2009
outerPid=$(ps -ef | grep 'bigdatacron' | grep -v 'grep' | awk '{print $2}')
if [ -n "${outerPid}""" ]; then kill -9 $outerPid; fi

exit 0
```