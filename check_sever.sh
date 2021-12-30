Pidofmain=$(pidof main)
Pidofswagger=$(pidof swagger)

if [ $Pidofmain ]; then
  echo "sever started , pid:$Pidofmain"
else
  echo "sever never start"
fi

if [ $Pidofswagger ]; then
  echo "swagger sever started , pid:$Pidofswagger"
else
  echo "sever never start"
fi