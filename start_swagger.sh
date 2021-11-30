echo "生成JSON中"
swagger generate spec -o ./swagger.json
echo "生成完毕"
echo "开启服务器"
swagger serve -F=swagger swagger.json -p=8081 0.0.0.0 apis --no-open