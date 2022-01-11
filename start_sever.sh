nohup go run main.go > main.out 2>&1 &
nohup ./start_swagger.sh > swagger.out 2>&1 &

sleep 3s

Pidofmain=$(pidof main)
Pidofswagger=$(pidof swagger)

if [ $Pidofmain ]; then
  echo "sever started , pid:$Pidofmain"
else
  echo "sever start filed"
fi

if [ $Pidofswagger ]; then
  echo "swagger sever started , pid:$Pidofswagger"
else
  echo "sever start filed"
fi