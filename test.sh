# curl --location --request POST 'http://localhost:8080/douyin/user/register/?username=test2&password=123456' \
# --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'

# curl --location --request POST 'http://localhost:8080/douyin/user/login/?username=test&password=123456' \
# --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'

#!/bin/bash
function get_json(){
  echo "${1//\"/}" | sed "s/.*$2:\([^,}]*\).*/\1/"
}

function deal_json(){
        cmd="curl --location --request POST 'http://localhost:8080/douyin/user/login/?username=test&password=123456' \
                --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'"
        echo ${cmd} 
        api_result=`eval ${cmd}`
        status_code=$(get_json "${api_result}" "status_code") # 从api_result中获取status对应的值
        echo ${status_code} 
        token=$(get_json "${api_result}" "token") # 从api_result中获取status对应的值
        echo ${token} 

        cmd="curl --location --request GET 'http://localhost:8080/ping?token=${token}' \
                --header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'"
        echo ${cmd}
        api_result=`eval ${cmd}`
        echo ${api_result}
}

deal_json