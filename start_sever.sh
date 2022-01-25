nohup go run main.go >> ./log/main.out 2>&1 &
nohup ./tools/swagger/start_swagger.sh >> ./log/swagger/swagger.out 2>&1 &

sleep 3s

Pidofmain=$(pidof main)
Pidofswagger=$(pidof swagger)

if [ $Pidofmain ]; then
  echo "sever started , pid:$Pidofmain"
else
  echo "sever start filed"
fi

if [ $Pidofswagger ]; then
  echo "swagger started , pid:$Pidofswagger"
else
  echo "swagger start filed"
fi