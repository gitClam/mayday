Pidofmain=$(pidof main)

if [ $Pidofmain ]; then
  echo "sever started , pid:$Pidofmain"
else
  echo "sever never start"
fi
