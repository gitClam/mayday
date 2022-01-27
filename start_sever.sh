nohup go run main.go >> ./log/main.out 2>&1 &

sleep 3s

Pidofmain=$(pidof main)

if [ $Pidofmain ]; then
  echo "sever started , pid:$Pidofmain"
else
  echo "sever start filed"
fi
